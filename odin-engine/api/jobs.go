package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
	"github.com/theycallmemac/odin/odin-engine/pkg/repository"
	"github.com/theycallmemac/odin/odin-engine/pkg/resources"
	"github.com/valyala/fasthttp"
)

// AddJob is used to create a new job
func AddJob(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	path := jobs.SetupEnvironment(body)
	id, err := repo.CreateJob(ctx, body, path, "")
	if err != nil {
		fmt.Fprintf(ctx, "[FAILED] Job failed to deploy: %v\n", err)
	} else {
		fmt.Fprintf(ctx, "[SUCCESS] Job (%s) deployed successfully\n", id)
	}
}

// DeleteJob is used to delete a job
func DeleteJob(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), " ")
	id, uid := args[0], args[1]
	os.RemoveAll("/etc/odin/jobs/" + id)
	os.RemoveAll("/etc/odin/logs/" + id)
	if err := repo.DeleteJob(ctx, id, uid); err != nil {
		fmt.Fprintf(ctx, "[FAILED] Failed to remove job (%s): %v\n", id, err)
	} else {
		fmt.Fprintf(ctx, "[SUCCESS] Job (%s) removed\n", id)
	}
}

// UpdateJob is used to update a job
func UpdateJob(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), "_")
	id, name, description, schedule, uid := args[0], args[1], args[2], args[3], args[4]
	job := &repository.Job{
		ID:  id,
		UID: uid,
	}
	if resources.NotEmpty(name) {
		job.Name = name
	}
	if resources.NotEmpty(description) {
		job.Description = description
	}
	if resources.NotEmpty(schedule) {
		ioutil.WriteFile(".tmp.yml", []byte("provider:\n  name: 'odin'\n  version: '1.0.0'\njob:\n  name: ''\n  description: ''\n  language: ''\n  file: ''\n  schedule: "+schedule+"\n\n"), 0654)
		resp := jobs.MakePostRequest("http://localhost:3939/schedule", bytes.NewBuffer([]byte(".tmp.yml")))
		os.Remove(".tmp.yml")
		job.Schedule = resp
	}
	if err := repo.UpdateJob(ctx, job); err != nil {
		fmt.Fprintf(ctx, "[FAILED] Failed to update job (%s): %v\n", job.ID, err)
	} else {
		fmt.Fprintf(ctx, "[SUCCESS] Job (%s) updated\n", id)
	}
}

// GetJobDescription is used to show a job's description
func GetJobDescription(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), "_")
	id, uid := args[0], args[1]
	job, err := repo.GetJobById(ctx, id, uid)
	if err != nil {
		fmt.Fprintf(ctx, "[FAILED] Failed to get job (%s): %v", id, err)
	} else {
		fmt.Fprintf(ctx, job.Name+" - "+job.Description+"\n")
	}
}

// UpdateJobRuns is used to update a job's run number
// TODO: Get and update job should be done as a transaction
func UpdateJobRuns(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), " ")
	id, runs, uid := args[0], args[1], args[2]
	job, err := repo.GetJobById(ctx, id, uid)
	if err != nil {
		fmt.Fprintf(ctx, "[FAILED] Failed to update job (%s): %v\n", id, err)
		return
	}
	inc, err := strconv.Atoi(runs)
	if err != nil {
		fmt.Fprint(ctx, "Invalid run")
		return
	}
	job.Runs = job.Runs + inc
	if err := repo.UpdateJob(ctx, job); err != nil {
		fmt.Fprintf(ctx, "[FAILED] Failed to update job (%s): %v\n", job.ID, err)
	} else {
		fmt.Fprintf(ctx, "[SUCCESS] Job (%s) updated\n", id)
	}
}

// ListJobs is used to list the current jobs running
func ListJobs(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	uid := string(ctx.PostBody())
	jobList, err := repo.GetUserJobs(ctx, uid)
	if err != nil {
		fmt.Fprintf(ctx, "[FAILED] Failed to get jobs for user %s\n", uid)
		return
	}
	fmt.Fprintf(ctx, jobs.SchFormat("ID", "NAME", "DESCRIPTION", "LANGUAGE", "LINKS", "SCHEDULE"))
	for _, job := range jobList {
		linkLen := len(job.Links) - 1
		if linkLen < 0 {
			linkLen = 0
		}
		fmt.Fprintf(ctx, jobs.SchFormat(job.ID, job.Name, job.Description, job.Language, job.Links[:linkLen], job.Schedule[:len(job.Schedule)-1]))
	}
}

// GetJobLogs is used to retrieve the logs for a job
func GetJobLogs(ctx *fasthttp.RequestCtx) {
	log, _ := ioutil.ReadFile("/etc/odin/logs/" + string(ctx.PostBody()))
	fmt.Fprintf(ctx, "\n"+string(log)+"\n")
}
