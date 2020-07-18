package nosql

import (
	"context"

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

type Repository struct {
	db nosql.Database
}

func (repo *Repository) ensureIndex(ctx context.Context) error {
	return repo.db.EnsureIndex(
		ctx,
		base.ObservabilityTable,
		nosql.Index{
			Fields: []string{"_id"},
			Type:   nosql.StringExact,
		},
		[]nosql.Index{
			{
				Fields: []string{"id"},
				Type:   nosql.StringExact,
			},
		},
	)
}

func (repo *Repository) GetJobStats(ctx context.Context, id string) ([]base.JobStats, error) {
	iter := repo.db.Query(base.ObservabilityTable).WithFields(nosql.FieldFilter{
		Path:   []string{"id"},
		Filter: nosql.Equal,
		Value:  nosql.String(id),
	}).Iterate()

	if iter.Err() != nil {
		return nil, iter.Err()
	}
	defer iter.Close()

	results := make([]base.JobStats, 0)
	for iter.Next(ctx) {
		doc := iter.Doc()
		jobStats := base.JobStats{
			ID:          string(doc["id"].(nosql.String)),
			Description: string(doc["desc"].(nosql.String)),
			Type:        string(doc["type"].(nosql.String)),
			Value:       string(doc["value"].(nosql.String)),
		}
		results = append(results, jobStats)
	}
	return results, nil
}

func (repo *Repository) Close() error {
	return repo.db.Close()
}
