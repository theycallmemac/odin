package repository

import (
	"context"
	"fmt"
)

const (
	DefaultDatabase    = "odin"
	ObservabilityTable = "observability"
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

type JobStats struct {
	ID          string
	Description string
	Type        string
	Value       string
}

type Repository interface {
	GetJobStats(ctx context.Context, id string) ([]JobStats, error)
	Close() error
}
