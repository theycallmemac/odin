package nosql

import (
	"context"
	"errors"
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
	db, err := mongo.Dial("localhost:27017", base.DefaultDatabase, nil)
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
	err = client.Database(base.DefaultDatabase).Collection(base.JobTable).Drop(ctx)
	require.NoError(t, client.Disconnect(ctx))
}

func TestRepository_CreateJob(t *testing.T) {
	type args struct {
		ctx  context.Context
		data []byte
		path string
		uid  string
	}
	tests := []struct {
		name     string
		existing []*base.Job
		args     args
		want     *base.Job
		wantId   string
		wantErr  error
	}{
		{
			name:     "Simple success",
			existing: nil,
			args: args{
				ctx:  context.Background(),
				data: []byte(`{"id":"1", "uid": "2", "name": "sample", "description": "job desc"}`),
				path: "sample.yaml",
				uid:  "2",
			},
			want: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				File:        "sample.yaml",
			},
			wantId:  "1",
			wantErr: nil,
		},
		{
			name: "Exist",
			existing: []*base.Job{
				{
					ID:          "1",
					UID:         "2",
					Name:        "sample",
					Description: "job desc",
					File:        "sample.yaml",
				},
			},
			args: args{
				ctx:  context.Background(),
				data: []byte(`{"id":"1", "uid": "2", "name": "sample", "description": "job desc"}`),
				path: "sample.yaml",
				uid:  "2",
			},
			want:    nil,
			wantId:  "",
			wantErr: errors.New("job with id 1 exists"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp(t, tt.args.ctx)
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			for _, e := range tt.existing {
				_, err := repo.db.Insert(tt.args.ctx, base.JobTable, []string{e.ID}, marshalJob(e))
				require.NoError(t, err)
			}
			id, err := repo.CreateJob(tt.args.ctx, tt.args.data, tt.args.path, tt.args.uid)
			if tt.wantErr == nil {
				require.NoError(t, err)
				doc, err := repo.db.Query(base.JobTable).WithFields(
					nosql.FieldFilter{
						Path:   []string{"_id"},
						Filter: nosql.Equal,
						Value:  nosql.String(tt.want.ID),
					},
					nosql.FieldFilter{
						Path:   []string{"uid"},
						Filter: nosql.Equal,
						Value:  nosql.String(tt.want.UID),
					},
				).One(tt.args.ctx)
				require.NoError(t, err)
				job := &base.Job{}
				unmarshalJob(doc, job)
				assert.Equal(t, tt.want, job)
				assert.Equal(t, tt.wantId, id)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}

func TestRepository_GetJobById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		uid string
	}
	tests := []struct {
		name     string
		existing []*base.Job
		args     args
		want     *base.Job
		wantErr  error
	}{
		{
			name: "Simple success",
			existing: []*base.Job{
				{
					ID:          "1",
					UID:         "2",
					Name:        "sample",
					Description: "job desc",
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
				uid: "2",
			},
			want: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
			},
			wantErr: nil,
		},
		{
			name:     "Not exist",
			existing: nil,
			args: args{
				ctx: context.Background(),
				id:  "1",
				uid: "2",
			},
			want:    nil,
			wantErr: nosql.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp(t, tt.args.ctx)
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			for _, e := range tt.existing {
				_, err := repo.db.Insert(tt.args.ctx, base.JobTable, []string{e.ID}, marshalJob(e))
				require.NoError(t, err)
			}
			job, err := repo.GetJobById(tt.args.ctx, tt.args.id, tt.args.uid)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
			assert.Equal(t, tt.want, job)

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}

func TestRepository_GetUserJobs(t *testing.T) {
	type args struct {
		ctx context.Context
		uid string
	}
	tests := []struct {
		name     string
		existing []*base.Job
		args     args
		want     []*base.Job
		wantErr  error
	}{
		{
			name: "Simple success",
			existing: []*base.Job{
				{
					ID:          "1",
					UID:         "2",
					Name:        "sample",
					Description: "job desc",
				},
				{
					ID:          "2",
					UID:         "2",
					Name:        "sample",
					Description: "job desc",
				},
				{
					ID:          "3",
					UID:         "3",
					Name:        "sample",
					Description: "job desc",
				},
			},
			args: args{
				ctx: context.Background(),
				uid: "2",
			},
			want: []*base.Job{
				{
					ID:          "1",
					UID:         "2",
					Name:        "sample",
					Description: "job desc",
				},
				{
					ID:          "2",
					UID:         "2",
					Name:        "sample",
					Description: "job desc",
				},
			},
			wantErr: nil,
		},
		{
			name:     "Not exist",
			existing: nil,
			args: args{
				ctx: context.Background(),
				uid: "2",
			},
			want:    []*base.Job{},
			wantErr: nil,
		},
		{
			name: "No user jobs",
			existing: []*base.Job{
				{
					ID:          "1",
					UID:         "3",
					Description: "sample",
				},
			},
			args: args{
				ctx: context.Background(),
				uid: "2",
			},
			want:    []*base.Job{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp(t, tt.args.ctx)
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			for _, e := range tt.existing {
				_, err := repo.db.Insert(tt.args.ctx, base.JobTable, []string{e.ID}, marshalJob(e))
				require.NoError(t, err)
			}
			jobs, err := repo.GetUserJobs(tt.args.ctx, tt.args.uid)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
			assert.Equal(t, tt.want, jobs)

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}

func TestRepository_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []base.Job
		wantErr error
	}{
		{
			name: "Simple success 1",
			want: []base.Job{
				{
					ID:          "1",
					UID:         "2",
					Name:        "sample",
					Description: "job desc",
				},
				{
					ID:          "2",
					UID:         "2",
					Name:        "sample",
					Description: "job desc",
				},
				{
					ID:          "3",
					UID:         "3",
					Name:        "sample",
					Description: "job desc",
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: nil,
		},
		{
			name: "Simple success 2",
			want: []base.Job{
				{
					ID:          "1",
					UID:         "3",
					Description: "sample",
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: nil,
		},
		{
			name: "Not exist",
			args: args{
				ctx: context.Background(),
			},
			want:    []base.Job{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp(t, tt.args.ctx)
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			for _, e := range tt.want {
				_, err := repo.db.Insert(tt.args.ctx, base.JobTable, []string{e.ID}, marshalJob(&e))
				require.NoError(t, err)
			}
			jobs, err := repo.GetAll(tt.args.ctx)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
			assert.Equal(t, tt.want, jobs)

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}

func TestRepository_UpdateJob(t *testing.T) {
	type args struct {
		ctx context.Context
		job *base.Job
	}
	tests := []struct {
		name     string
		existing *base.Job
		args     args
		want     *base.Job
		count    int
		wantErr  error
	}{
		{
			name: "Simple success",
			existing: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
			},
			args: args{
				ctx: context.Background(),
				job: &base.Job{
					ID:          "1",
					UID:         "2",
					Name:        "changed",
					Description: "changed",
					Schedule:    "weekly",
					Runs:        1,
				},
			},
			want: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "changed",
				Description: "changed",
				Schedule:    "weekly",
				Runs:        1,
			},
			count:   1,
			wantErr: nil,
		},
		{
			name:     "Not exist",
			existing: nil,
			args: args{
				ctx: context.Background(),
				job: &base.Job{
					ID:          "1",
					UID:         "2",
					Name:        "changed",
					Description: "changed",
					Schedule:    "weekly",
					Runs:        1,
				},
			},
			want:    nil,
			wantErr: nosql.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp(t, tt.args.ctx)
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			if tt.existing != nil {
				_, err := repo.db.Insert(tt.args.ctx, base.JobTable, []string{tt.existing.ID}, marshalJob(tt.existing))
				require.NoError(t, err)
			}
			err := repo.UpdateJob(tt.args.ctx, tt.args.job)
			if tt.wantErr == nil {
				require.NoError(t, err)
				it := repo.db.Query(base.JobTable).Iterate()
				require.NoError(t, it.Err())
				c := 0
				for it.Next(tt.args.ctx) {
					c++
				}
				assert.Equal(t, tt.count, c)
				doc, err := repo.db.Query(base.JobTable).WithFields(
					nosql.FieldFilter{
						Path:   []string{"_id"},
						Filter: nosql.Equal,
						Value:  nosql.String(tt.existing.ID),
					},
					nosql.FieldFilter{
						Path:   []string{"uid"},
						Filter: nosql.Equal,
						Value:  nosql.String(tt.existing.UID),
					},
				).One(tt.args.ctx)
				require.NoError(t, err)
				job := &base.Job{}
				unmarshalJob(doc, job)
				assert.Equal(t, tt.want, job)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}

func TestRepository_DeleteJob(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		uid string
	}
	tests := []struct {
		name     string
		existing *base.Job
		args     args
		count    int
		wantErr  error
	}{
		{
			name: "Simple success",
			existing: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
				uid: "2",
			},
			count:   0,
			wantErr: nil,
		},
		{
			name: "Not matched",
			existing: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
			},
			args: args{
				ctx: context.Background(),
				id:  "2",
				uid: "3",
			},
			count:   1,
			wantErr: nil,
		},
		{
			name:     "Not exist",
			existing: nil,
			args: args{
				ctx: context.Background(),
				id:  "1",
				uid: "2",
			},
			count:   0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp(t, tt.args.ctx)
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			var err error
			if tt.existing != nil {
				_, err = repo.db.Insert(tt.args.ctx, base.JobTable, []string{tt.existing.ID}, marshalJob(tt.existing))
			}
			require.NoError(t, err)
			err = repo.DeleteJob(tt.args.ctx, tt.args.id, tt.args.uid)
			if tt.wantErr == nil {
				require.NoError(t, err)
				it := repo.db.Query(base.JobTable).Iterate()
				require.NoError(t, it.Err())
				c := 0
				for it.Next(tt.args.ctx) {
					c++
				}
				assert.Equal(t, tt.count, c)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}

func TestRepository_AddJobLink(t *testing.T) {
	type args struct {
		ctx  context.Context
		from string
		to   string
		uid  string
	}
	tests := []struct {
		name     string
		existing *base.Job
		args     args
		want     *base.Job
		wantErr  error
	}{
		{
			name: "No existing links",
			existing: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
			},
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "2",
				uid:  "2",
			},
			want: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "2,",
			},
			wantErr: nil,
		},
		{
			name: "Existing links",
			existing: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "2,",
			},
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "3",
				uid:  "2",
			},
			want: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "2,3,",
			},
			wantErr: nil,
		},
		{
			name:     "No job",
			existing: nil,
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "2",
				uid:  "2",
			},
			want:    nil,
			wantErr: nosql.ErrNotFound,
		},
		{
			name:     "Invalid argument",
			existing: nil,
			args: args{
				ctx:  context.Background(),
				from: "",
				to:   "2",
				uid:  "2",
			},
			want:    nil,
			wantErr: ErrInvalidArgument,
		},
		{
			name:     "Invalid argument 2",
			existing: nil,
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "",
				uid:  "2",
			},
			want:    nil,
			wantErr: ErrInvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp(t, tt.args.ctx)
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			if tt.existing != nil {
				_, err := repo.db.Insert(tt.args.ctx, base.JobTable, []string{tt.existing.ID}, marshalJob(tt.existing))
				require.NoError(t, err)
			}
			err := repo.AddJobLink(tt.args.ctx, tt.args.from, tt.args.to, tt.args.uid)
			if tt.wantErr == nil {
				require.NoError(t, err)
				doc, err := repo.db.Query(base.JobTable).WithFields(
					nosql.FieldFilter{
						Path:   []string{"_id"},
						Filter: nosql.Equal,
						Value:  nosql.String(tt.existing.ID),
					},
					nosql.FieldFilter{
						Path:   []string{"uid"},
						Filter: nosql.Equal,
						Value:  nosql.String(tt.existing.UID),
					},
				).One(tt.args.ctx)
				require.NoError(t, err)
				job := &base.Job{}
				unmarshalJob(doc, job)
				assert.Equal(t, tt.want, job)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}

func TestRepository_DeleteJobLink(t *testing.T) {
	type args struct {
		ctx  context.Context
		from string
		to   string
		uid  string
	}
	tests := []struct {
		name     string
		existing *base.Job
		args     args
		want     *base.Job
		wantErr  error
	}{
		{
			name: "No existing links",
			existing: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "",
			},
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "2",
				uid:  "2",
			},
			want: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "",
			},
			wantErr: nil,
		},
		{
			name: "Existing links",
			existing: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "2,",
			},
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "2",
				uid:  "2",
			},
			want: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "",
			},
			wantErr: nil,
		},
		{
			name: "Existing links 2",
			existing: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "2,3,",
			},
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "2",
				uid:  "2",
			},
			want: &base.Job{
				ID:          "1",
				UID:         "2",
				Name:        "sample",
				Description: "job desc",
				Schedule:    "daily",
				Runs:        0,
				Links:       "3,",
			},
			wantErr: nil,
		},
		{
			name:     "No job",
			existing: nil,
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "2",
				uid:  "2",
			},
			want:    nil,
			wantErr: nosql.ErrNotFound,
		},
		{
			name:     "Invalid argument",
			existing: nil,
			args: args{
				ctx:  context.Background(),
				from: "",
				to:   "2",
				uid:  "2",
			},
			want:    nil,
			wantErr: ErrInvalidArgument,
		},
		{
			name:     "Invalid argument 2",
			existing: nil,
			args: args{
				ctx:  context.Background(),
				from: "1",
				to:   "",
				uid:  "2",
			},
			want:    nil,
			wantErr: ErrInvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp(t, tt.args.ctx)
			repo := testRepo(t)
			require.NoError(t, repo.ensureIndex(tt.args.ctx))

			if tt.existing != nil {
				_, err := repo.db.Insert(tt.args.ctx, base.JobTable, []string{tt.existing.ID}, marshalJob(tt.existing))
				require.NoError(t, err)
			}
			err := repo.DeleteJobLink(tt.args.ctx, tt.args.from, tt.args.to, tt.args.uid)
			if tt.wantErr == nil {
				require.NoError(t, err)
				doc, err := repo.db.Query(base.JobTable).WithFields(
					nosql.FieldFilter{
						Path:   []string{"_id"},
						Filter: nosql.Equal,
						Value:  nosql.String(tt.existing.ID),
					},
					nosql.FieldFilter{
						Path:   []string{"uid"},
						Filter: nosql.Equal,
						Value:  nosql.String(tt.existing.UID),
					},
				).One(tt.args.ctx)
				require.NoError(t, err)
				job := &base.Job{}
				unmarshalJob(doc, job)
				assert.Equal(t, tt.want, job)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}

			cleanUp(t, tt.args.ctx)
			require.NoError(t, repo.Close())
		})
	}
}

func TestRepository_GetJobStats(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name     string
		existing []*base.JobStats
		args     args
		want     []*base.JobStats
		wantErr  error
	}{
		{
			name: "Simple success",
			existing: []*base.JobStats{
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
			want: []*base.JobStats{
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
				_, err := repo.db.Insert(tt.args.ctx, base.ObservabilityTable, []string{e.ID}, nosql.Document{
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
