package nosql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/hidal-go/hidalgo/legacy/nosql/mongo"

	"github.com/hidal-go/hidalgo/legacy/nosql"
	base "github.com/theycallmemac/odin/odin-engine/pkg/repository"
)

func init() {
	base.AddRegistration("mongodb", &base.Registration{
		OpenFunc: func(address string, options base.Options) (base.Repository, error) {
			db, err := mongo.Dial(address, base.DefaultDatabase, nosql.Options(options))
			if err != nil {
				return nil, err
			}
			repo := &Repository{
				db: db,
			}
			if err := repo.ensureIndex(context.Background()); err != nil {
				return nil, err
			}
			return repo, nil
		},
	})
}

var (
	ErrInvalidArgument = errors.New("Invalid argument")
)

// Repository is implementation of Repository interface
type Repository struct {
	db nosql.Database
}

func (repo *Repository) ensureIndex(ctx context.Context) error {
	if err := repo.db.EnsureIndex(
		ctx,
		base.ObservabilityTable,
		nosql.Index{
			Fields: []string{"id"},
			Type:   nosql.StringExact,
		},
		nil,
	); err != nil {
		return err
	}

	if err := repo.db.EnsureIndex(
		ctx,
		base.JobTable,
		nosql.Index{
			Fields: []string{"id"},
			Type:   nosql.StringExact,
		},
		nil,
	); err != nil {
		return err
	}

	return nil
}

// CreateJob creates a new job for an user
func (repo *Repository) CreateJob(ctx context.Context, data []byte, path string, uid string) (string, error) {
	job := &base.Job{}
	if err := json.Unmarshal(data, job); err != nil {
		return "", err
	}
	if job.ID == "" {
		return "", ErrInvalidArgument
	} else if _, err := repo.GetJobById(ctx, job.ID, uid); err == nil {
		return "", fmt.Errorf("job with id %s exists", job.ID)
	} else if err != nosql.ErrNotFound {
		return "", err
	}
	job.File = path
	job.Runs = 0
	doc := marshalJob(job)
	_, err := repo.db.Insert(ctx, base.JobTable, []string{job.ID}, doc)
	return job.ID, err
}

// GetJobById returns a job by filtering on a certain value pertaining to that job
func (repo *Repository) GetJobById(ctx context.Context, id string, uid string) (*base.Job, error) {
	if id == "" {
		return nil, ErrInvalidArgument
	}
	doc, err := repo.db.Query(base.JobTable).WithFields(
		nosql.FieldFilter{
			Path:   []string{"_id"},
			Filter: nosql.Equal,
			Value:  nosql.String(id),
		},
		nosql.FieldFilter{
			Path:   []string{"uid"},
			Filter: nosql.Equal,
			Value:  nosql.String(uid),
		},
	).One(ctx)

	if err != nil {
		return nil, err
	}

	job := &base.Job{}
	unmarshalJob(doc, job)
	return job, nil
}

// GetUserJobs returns all jobs belonging to an user
func (repo *Repository) GetUserJobs(ctx context.Context, uid string) ([]*base.Job, error) {
	if uid == "" {
		return nil, ErrInvalidArgument
	}
	iter := repo.db.Query(base.JobTable).WithFields(nosql.FieldFilter{
		Path:   []string{"uid"},
		Filter: nosql.Equal,
		Value:  nosql.String(uid),
	}).Iterate()

	if iter.Err() != nil {
		return nil, iter.Err()
	}
	defer iter.Close()

	results := make([]*base.Job, 0)
	for iter.Next(ctx) {
		doc := iter.Doc()
		job := &base.Job{}
		unmarshalJob(doc, job)
		results = append(results, job)
	}
	return results, nil
}

// GetAll returns all jobs
func (repo *Repository) GetAll(ctx context.Context) ([]base.Job, error) {
	iter := repo.db.Query(base.JobTable).Iterate()
	if iter.Err() != nil {
		return nil, iter.Err()
	}
	defer iter.Close()

	results := make([]base.Job, 0)
	for iter.Next(ctx) {
		doc := iter.Doc()
		job := base.Job{}
		unmarshalJob(doc, &job)
		results = append(results, job)
	}
	return results, nil
}

// UpdateJob modifies a job
func (repo *Repository) UpdateJob(ctx context.Context, job *base.Job) error {
	if job.ID == "" {
		return ErrInvalidArgument
	}
	key := []string{job.ID}
	doc, err := repo.db.Query(base.JobTable).WithFields(
		nosql.FieldFilter{
			Path:   []string{"_id"},
			Filter: nosql.Equal,
			Value:  nosql.String(job.ID),
		},
		nosql.FieldFilter{
			Path:   []string{"uid"},
			Filter: nosql.Equal,
			Value:  nosql.String(job.UID),
		},
	).One(ctx)
	if err != nil {
		return err
	}
	doc["name"] = nosql.String(job.Name)
	doc["description"] = nosql.String(job.Description)
	doc["schedule"] = nosql.String(job.Schedule)
	doc["runs"] = nosql.Int(job.Runs)
	return repo.db.Update(base.JobTable, key).Upsert(doc).Do(ctx)
}

// DeleteJob deletes an user's job
func (repo *Repository) DeleteJob(ctx context.Context, id string, uid string) error {
	if id == "" {
		return ErrInvalidArgument
	}
	return repo.db.Delete(base.JobTable).WithFields(
		nosql.FieldFilter{
			Path:   []string{"_id"},
			Filter: nosql.Equal,
			Value:  nosql.String(id),
		},
		nosql.FieldFilter{
			Path:   []string{"uid"},
			Filter: nosql.Equal,
			Value:  nosql.String(uid),
		},
	).Do(ctx)
}

// GetJobStats returns stats of a job given the job id
func (repo *Repository) GetJobStats(ctx context.Context, id string) ([]*base.JobStats, error) {
	if id == "" {
		return nil, ErrInvalidArgument
	}
	iter := repo.db.Query(base.ObservabilityTable).WithFields(nosql.FieldFilter{
		Path:   []string{"_id"},
		Filter: nosql.Equal,
		Value:  nosql.String(id),
	}).Iterate()

	if iter.Err() != nil {
		return nil, iter.Err()
	}
	defer iter.Close()

	results := make([]*base.JobStats, 0)
	for iter.Next(ctx) {
		doc := iter.Doc()
		jobStats := &base.JobStats{}
		unmarshalJobStats(doc, jobStats)
		results = append(results, jobStats)
	}
	return results, nil
}

// CreateJobStats creates a new job for an user
func (repo *Repository) CreateJobStats(ctx context.Context, js *base.JobStats) error {
	if js.ID == "" {
		return ErrInvalidArgument
	} else if _, err := repo.GetJobStats(ctx, js.ID); err == nil {
		return fmt.Errorf("job stats with id %s exists", js.ID)
	} else if err != nosql.ErrNotFound {
		return err
	}
	doc := marshalJobStats(js)
	_, err := repo.db.Insert(ctx, base.JobTable, []string{js.ID}, doc)
	return err
}

// AddJobLink is used to add links the job is associated with
func (repo *Repository) AddJobLink(ctx context.Context, from string, to string, uid string) error {
	if from == "" || to == "" {
		return ErrInvalidArgument
	}
	key := []string{from}
	doc, err := repo.db.Query(base.JobTable).WithFields(
		nosql.FieldFilter{
			Path:   []string{"_id"},
			Filter: nosql.Equal,
			Value:  nosql.String(from),
		},
		nosql.FieldFilter{
			Path:   []string{"uid"},
			Filter: nosql.Equal,
			Value:  nosql.String(uid),
		},
	).One(ctx)
	if err != nil {
		return err
	}
	doc["links"] = nosql.String(string(doc["links"].(nosql.String)) + to + ",")
	return repo.db.Update(base.JobTable, key).Upsert(doc).Do(ctx)
}

// DeleteJobLink is used to delete links the job is associated with
func (repo *Repository) DeleteJobLink(ctx context.Context, from string, to string, uid string) error {
	if from == "" || to == "" {
		return ErrInvalidArgument
	}
	key := []string{from}
	doc, err := repo.db.Query(base.JobTable).WithFields(
		nosql.FieldFilter{
			Path:   []string{"_id"},
			Filter: nosql.Equal,
			Value:  nosql.String(from),
		},
		nosql.FieldFilter{
			Path:   []string{"uid"},
			Filter: nosql.Equal,
			Value:  nosql.String(uid),
		},
	).One(ctx)
	if err != nil {
		return err
	}
	links := strings.Split(string(doc["links"].(nosql.String)), ",")
	newLinks := ""
	for _, link := range links {
		if link != to && link != "" {
			newLinks = newLinks + link + ","
		}
	}
	doc["links"] = nosql.String(newLinks)
	return repo.db.Update(base.JobTable, key).Upsert(doc).Do(ctx)
}

// Close closes db connection
func (repo *Repository) Close() error {
	return repo.db.Close()
}

func unmarshalJob(doc nosql.Document, job *base.Job) {
	job.ID = valueToString(doc["id"])
	job.UID = valueToString(doc["uid"])
	job.GID = valueToString(doc["gid"])
	job.Name = valueToString(doc["name"])
	job.Description = valueToString(doc["description"])
	job.Language = valueToString(doc["language"])
	job.File = valueToString(doc["file"])
	job.Stats = valueToString(doc["stats"])
	job.Schedule = valueToString(doc["schedule"])
	job.Runs = int(doc["runs"].(nosql.Int))
	job.Links = valueToString(doc["links"])
}

func marshalJob(job *base.Job) nosql.Document {
	return nosql.Document{
		"id":          nosql.String(job.ID),
		"uid":         nosql.String(job.UID),
		"gid":         nosql.String(job.GID),
		"name":        nosql.String(job.Name),
		"description": nosql.String(job.Description),
		"language":    nosql.String(job.Language),
		"file":        nosql.String(job.File),
		"stats":       nosql.String(job.Stats),
		"schedule":    nosql.String(job.Schedule),
		"runs":        nosql.Int(job.Runs),
		"links":       nosql.String(job.Links),
	}
}

func unmarshalJobStats(doc nosql.Document, js *base.JobStats) {
	js.ID = valueToString(doc["id"])
	js.Description = valueToString(doc["desc"])
	js.Type = valueToString(doc["type"])
	js.Value = valueToString(doc["value"])
	js.Timestamp = valueToString(doc["timestamp"])
}

func marshalJobStats(js *base.JobStats) nosql.Document {
	return nosql.Document{
		"id":        nosql.String(js.ID),
		"desc":      nosql.String(js.Description),
		"type":      nosql.String(js.Type),
		"value":     nosql.String(js.Value),
		"timestamp": nosql.String(js.Timestamp),
	}
}

func valueToString(val nosql.Value) string {
	if val == nil {
		return ""
	}
	return string(val.(nosql.String))
}
