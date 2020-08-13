package api

import (
	"fmt"
	"strings"

	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
	"github.com/theycallmemac/odin/odin-engine/pkg/repository"
	"github.com/valyala/fasthttp"
)

// JobStats is a type to be used for accessing and storing job stats information
type JobStats struct {
	ID          string
	Description string
	Type        string
	Value       string
	Timestamp   string
}

// AddJobStats is used to parse collected metrics
func AddJobStats(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), ",")
	typeOfValue, desc, value, id, timestamp := args[0], args[1], args[2], args[3], args[4]
	js := &repository.JobStats{
		ID:          id,
		Description: desc,
		Type:        typeOfValue,
		Value:       value,
		Timestamp:   timestamp,
	}
	if err := repo.CreateJobStats(ctx, js); err != nil {
		fmt.Fprintf(ctx, "500")
	} else {
		fmt.Fprintf(ctx, "200")
	}
}

// GetJobStats is used to show stats collected by a specified job
func GetJobStats(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	statsList, err := repo.GetJobStats(ctx, string(ctx.PostBody()))
	if err != nil {
		fmt.Fprintf(ctx, "[FAILED] Cannot get job stats: %v\n", err)
	} else {
		for _, stat := range statsList {
			fmt.Fprintf(ctx, jobs.Format(stat.ID, stat.Description, stat.Type, stat.Value))
		}
	}
}
