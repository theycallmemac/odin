package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
	"github.com/theycallmemac/odin/odin-engine/pkg/resources"

	"go.mongodb.org/mongo-driver/bson"
)

// create resource type to be used by the router
type jobsResource struct{}

func (rs jobsResource) Routes() chi.Router {
	// establish new chi router
	r := chi.NewRouter()

	// define routes under the jobs endpoint
	r.Post("/", rs.Create)
	r.Put("/", rs.Delete)

	// define routes under the jobs/list endpoint
	r.Route("/list", func(r chi.Router) {
		r.Post("/", rs.List)
	})

	// define routes under the jobs/info endpoint
	r.Route("/info", func(r chi.Router) {
		r.Post("/description", rs.DescriptionByID)
		r.Put("/runs", rs.UpdateRuns)
		r.Put("/", rs.Update)
	})

	r.Route("/logs", func(r chi.Router) {
		r.Post("/", rs.LogByID)
	})

	return r
}

// List is used to list the current jobs running
func (rs jobsResource) List(w http.ResponseWriter, r *http.Request) {
	d, _ := ioutil.ReadAll(r.Body)
	client, err := jobs.SetupClient()
	if err != nil {
		w.Write([]byte("MongoDB cannot be accessed at the moment\n"))
	} else {
		jobList := jobs.GetUserJobs(client, string(d))
		w.Write([]byte(jobs.SchFormat("ID", "NAME", "DESCRIPTION", "LANGUAGE", "LINKS", "SCHEDULE")))
		for _, job := range jobList {
			linkLen := len(job.Links) - 1
			if linkLen < 0 {
				linkLen = 0
			}
			w.Write([]byte(jobs.SchFormat(job.ID, job.Name, job.Description, job.Language, job.Links[:linkLen], job.Schedule[:len(job.Schedule)-1])))
		}
	}
}

// Create is used to create a new job
func (rs jobsResource) Create(w http.ResponseWriter, r *http.Request) {
	d, _ := ioutil.ReadAll(r.Body)
	path := jobs.SetupEnvironment(d)
	client, err := jobs.SetupClient()
	if err != nil {
		w.Write([]byte("MongoDB cannot be accessed at the moment\n"))
	} else {
		status := jobs.InsertIntoMongo(client, d, path, "")
		w.Write([]byte(status))
	}
}

// DescriptionByID is used to show a job's description
func (rs jobsResource) DescriptionByID(w http.ResponseWriter, r *http.Request) {
	d, _ := ioutil.ReadAll(r.Body)
	args := strings.Split(string(d), "_")
	id, uid := args[0], args[1]
	client, err := jobs.SetupClient()
	if err != nil {
		w.Write([]byte("MongoDB cannot be accessed at the moment\n"))
	} else {
		job := jobs.GetJobByValue(client, bson.M{"id": id}, uid)
		w.Write([]byte(job.Name + " - " + job.Description + "\n"))
	}
}

// Update is used to update a job
func (rs jobsResource) Update(w http.ResponseWriter, r *http.Request) {
	d, _ := ioutil.ReadAll(r.Body)
	args := strings.Split(string(d), "_")
	id, name, description, schedule, uid := args[0], args[1], args[2], args[3], args[4]
	client, err := jobs.SetupClient()
	if err != nil {
		w.Write([]byte("MongoDB cannot be accessed at the moment\n"))
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
		w.Write([]byte("Updated job " + id + " successfully\n"))
	}
}

// UpdateRuns is used to update a job's run number
func (rs jobsResource) UpdateRuns(w http.ResponseWriter, r *http.Request) {
	d, _ := ioutil.ReadAll(r.Body)
	args := strings.Split(string(d), " ")
	id, runs, uid := args[0], args[1], args[2]
	client, err := jobs.SetupClient()
	if err != nil {
		w.Write([]byte("MongoDB cannot be accessed at the moment\n"))
	} else {
		job := jobs.GetJobByValue(client, bson.M{"id": id}, uid)
		inc, _ := strconv.Atoi(runs)
		job.Runs = job.Runs + inc
		_ = jobs.UpdateJobByValue(client, job)
	}
}

// Delete is used to delete a job
func (rs jobsResource) Delete(w http.ResponseWriter, r *http.Request) {
	d, _ := ioutil.ReadAll(r.Body)
	args := strings.Split(string(d), " ")
	id, uid := args[0], args[1]
	client, err := jobs.SetupClient()
	if err != nil {
		w.Write([]byte("MongoDB cannot be accessed at the moment\n"))
	} else {
		os.RemoveAll("/etc/odin/jobs/" + id)
		os.RemoveAll("/etc/odin/logs/" + id)
		if jobs.DeleteJobByValue(client, bson.M{"id": id}, uid) {
			w.Write([]byte("Job removed!\n"))
		} else {
			w.Write([]byte("Job with that ID does not exist!\n"))
		}
	}
}

// LogByID is used to retrieve the logs for a job
func (rs jobsResource) LogByID(w http.ResponseWriter, r *http.Request) {
	d, _ := ioutil.ReadAll(r.Body)
	log, _ := ioutil.ReadFile("/etc/odin/logs/" + string(d))
	w.Write([]byte("\n" + string(log) + "\n"))
}
