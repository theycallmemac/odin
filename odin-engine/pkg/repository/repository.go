package repository

import (
	"context"
	"fmt"
)

const (
	DefaultDatabase    = "odin"
	ObservabilityTable = "observability"
	JobTable           = "jobs"
)

var registry = make(map[string]*Registration)

type Options map[string]interface{}

type Registration struct {
	OpenFunc func(address string, options Options) (Repository, error)
}

func AddRegistration(name string, reg *Registration) error {
	if _, ok := registry[name]; ok {
		return fmt.Errorf("%s is already registered", name)
	}
	registry[name] = reg
	return nil
}

func GetRegistration(name string) (*Registration, error) {
	reg, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("cannot find %s repository", name)
	}
	return reg, nil
}

// Job is a type to be used for accessing and storing job information
type Job struct {
	ID          string `yaml:"id"`
	UID         string `yaml:"uid"`
	GID         string `yaml:"gid"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Language    string `yaml:"language"`
	File        string `yaml:"file"`
	Stats       string `yaml:"stats"`
	Schedule    string `yaml:"schedule"`
	Runs        int
	Links       string
}

// JobStats is a type to be used for accessing and storing job stats information
type JobStats struct {
	ID          string
	Description string
	Type        string
	Value       string
	Timestamp   string
}

type Repository interface {
	// CreateJob creates a new job for an user
	CreateJob(ctx context.Context, data []byte, path string, uid string) (string, error)

	// GetJobById returns a job by job id and user id
	GetJobById(ctx context.Context, id string, uid string) (*Job, error)

	// GetUserJobs returns all jobs belonging to an user
	GetUserJobs(ctx context.Context, uid string) ([]*Job, error)

	// GetAll returns all jobs
	GetAll(ctx context.Context) ([]Job, error)

	// UpdateJob modifies a job
	UpdateJob(ctx context.Context, job *Job) error

	// DeleteJob deletes an user's job
	DeleteJob(ctx context.Context, id string, uid string) error

	// AddJobLink is used to add links the job is associated with
	AddJobLink(ctx context.Context, from string, to string, uid string) error

	// DeleteJobLink is used to delete links the job is associated with
	DeleteJobLink(ctx context.Context, from string, to string, uid string) error

	// GetJobStats returns stats of a job given the job id
	GetJobStats(ctx context.Context, id string) ([]*JobStats, error)

	// CreateJobStats create new job stats
	CreateJobStats(ctx context.Context, js *JobStats) error

	// Close closes db connection
	Close() error
}
