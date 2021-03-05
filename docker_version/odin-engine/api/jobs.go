package api

import (
	"bytes"
        "fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
	"github.com/theycallmemac/odin/odin-engine/pkg/resources"
        "github.com/valyala/fasthttp"

        "go.mongodb.org/mongo-driver/bson"
)

// AddJob is used to create a new job
func AddJob(ctx *fasthttp.RequestCtx) {
        body := ctx.PostBody()
	path := jobs.SetupEnvironment(body)
	client, err := jobs.SetupClient()
	if err != nil {
		fmt.Fprintf(ctx, "MongoDB cannot be accessed at the moment\n")
	} else {
		status := jobs.InsertIntoMongo(client, body, path, "")
		fmt.Fprintf(ctx, status)
	}
}

// DeleteJob is used to delete a job
func DeleteJob(ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), " ")
	id, uid := args[0], args[1]
	client, err := jobs.SetupClient()
	if err != nil {
		fmt.Fprintf(ctx, "MongoDB cannot be accessed at the moment\n")
	} else {
		os.RemoveAll("/etc/odin/jobs/" + id)
		os.RemoveAll("/etc/odin/logs/" + id)
		if jobs.DeleteJobByValue(client, bson.M{"id": id}, uid) {
			fmt.Fprintf(ctx, "Job removed!\n")
		} else {
			fmt.Fprintf(ctx, "Job with that ID does not exist!\n")
		}
	}
}

// UpdateJob is used to update a job
func UpdateJob(ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), "_")
	id, name, description, schedule, uid := args[0], args[1], args[2], args[3], args[4]
	client, err := jobs.SetupClient()
	if err != nil {
		fmt.Fprintf(ctx, "MongoDB cannot be accessed at the moment\n")
	} else {
		job := jobs.GetJobByValue(client, bson.M{"id": id}, uid)
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
		_ = jobs.UpdateJobByValue(client, job)
		fmt.Fprintf(ctx, "Updated job " + id + " successfully\n")
	}
}

// GetJobDescription is used to show a job's description
func GetJobDescription(ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), "_")
	id, uid := args[0], args[1]
	client, err := jobs.SetupClient()
	if err != nil {
		fmt.Fprintf(ctx, "MongoDB cannot be accessed at the moment\n")
	} else {
		job := jobs.GetJobByValue(client, bson.M{"id": id}, uid)
		fmt.Fprintf(ctx, job.Name + " - " + job.Description + "\n")
	}
}

// UpdateJobRuns is used to update a job's run number
func UpdateJobRuns(ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), " ")
	id, runs, uid := args[0], args[1], args[2]
	client, err := jobs.SetupClient()
	if err != nil {
		fmt.Fprintf(ctx, "MongoDB cannot be accessed at the moment\n")
	} else {
		job := jobs.GetJobByValue(client, bson.M{"id": id}, uid)
		inc, _ := strconv.Atoi(runs)
		job.Runs = job.Runs + inc
		_ = jobs.UpdateJobByValue(client, job)
	}
}

// ListJobs is used to list the current jobs running
func ListJobs(ctx *fasthttp.RequestCtx) {
	client, err := jobs.SetupClient()
	if err != nil {
		fmt.Fprintf(ctx, "MongoDB cannot be accessed at the moment\n")
	} else {
		jobList := jobs.GetUserJobs(client, string(ctx.PostBody()))
		fmt.Fprintf(ctx, jobs.SchFormat("ID", "NAME", "DESCRIPTION", "LANGUAGE", "LINKS", "SCHEDULE"))
		for _, job := range jobList {
			linkLen := len(job.Links) - 1
			if linkLen < 0 {
				linkLen = 0
			}
			fmt.Fprintf(ctx, jobs.SchFormat(job.ID, job.Name, job.Description, job.Language, job.Links[:linkLen], job.Schedule[:len(job.Schedule)-1]))
		}
	}
}

// GetJobLogs is used to retrieve the logs for a job
func GetJobLogs(ctx *fasthttp.RequestCtx) {
	log, _ := ioutil.ReadFile("/etc/odin/logs/" + string(ctx.PostBody()))
	fmt.Fprintf(ctx, "\n" + string(log) + "\n")
}
