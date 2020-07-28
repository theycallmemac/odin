package nosql

import (
	"context"
	"testing"

	"github.com/hidal-go/hidalgo/legacy/nosql"
	"github.com/hidal-go/hidalgo/legacy/nosql/mongo"
	base "github.com/theycallmemac/odin/odin-engine/pkg/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

// use mongo as default nosql test db
func testRepo(t *testing.T) *Repository {
	// TODO: pass mongo url as env variable instead
	db, err := mongo.Dial("localhost:27017", base.DefaultDatabase, nosql.Options(nil))
	require.NoError(t, err)
	return &Repository{
		db: db,
	}
}

// drop all related collections
// TODO: find a better way to cleanup with hidal-go
func cleanUp(t *testing.T, ctx context.Context) {
	client, err := mongodriver.NewClient(mongooptions.Client().ApplyURI("mongodb://localhost:27017"))
	require.NoError(t, err)
	client.Connect(ctx)
	err = client.Database(base.DefaultDatabase).Collection(base.ObservabilityTable).Drop(ctx)
	require.NoError(t, err)
	require.NoError(t, client.Disconnect(ctx))
}

func TestRepository_GetJobStats(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name     string
		existing []base.JobStats
		args     args
		want     []base.JobStats
		wantErr  error
	}{
		{
			name: "Simple success",
			existing: []base.JobStats{
				{
					ID:          "1",
					Description: "sample stats",
					Type:        "count",
					Value:       "2",
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want: []base.JobStats{
				{
					ID:          "1",
					Description: "sample stats",
					Type:        "count",
					Value:       "2",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			for _, e := range tt.existing {
				_, err := repo.db.Insert(tt.args.ctx, base.ObservabilityTable, nil, nosql.Document{
					"id":    nosql.String(e.ID),
					"desc":  nosql.String(e.Description),
					"type":  nosql.String(e.Type),
					"value": nosql.String(e.Value),
				})
				require.NoError(t, err)
			}
			got, err := repo.GetJobStats(tt.args.ctx, tt.args.id)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
			assert.ElementsMatch(t, tt.want, got)

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}
