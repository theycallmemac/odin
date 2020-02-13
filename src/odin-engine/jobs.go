package main

import (
    "./jobs"

    "io/ioutil"
    "net/http"

    "go.mongodb.org/mongo-driver/bson"
    "github.com/go-chi/chi"
)


// create NewJob type to tbe used for accessing job information
type NewJob struct {
    ID string `yaml:"id"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Status string `yaml:"status"`
    Schedule string `yaml:"schedule"`
}

// create resource type to be used by the router
type jobsResource struct{}

func (rs jobsResource) Routes() chi.Router {
    // establish new chi router
    r := chi.NewRouter()

    // define routes under the jobs endpoint
    r.Get("/", rs.List)
    r.Post("/", rs.Create)
    r.Put("/", rs.Delete)

    // define routes under the jobs/info endpoint
    r.Route("/info", func(r chi.Router) {
            r.Post("/description", rs.DescriptionByID)
            r.Post("/status", rs.StatusByID)
            r.Get("/status/all", rs.AllStatus)
            r.Put("/", rs.Update)
            r.Delete("/", rs.Delete)
    })

    return r
}

// this function is used to list the current jobs running
func (rs jobsResource) List(w http.ResponseWriter, r *http.Request) {
    jobList := jobs.GetAll(jobs.SetupClient())
    w.Write([]byte(jobs.Format("ID", "NAME", "DESCRIPTION", "LANGUAGE", "STATUS", "SCHEDULE")))
    for _, job := range jobList {
        w.Write([]byte(jobs.Format(job.ID, job.Name, job.Description, job.Language ,job.Status, job.Schedule)))
    }
}

// this function is used to create a new job
func (rs jobsResource) Create(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    inserted := jobs.InsertIntoMongo(jobs.SetupClient(), d)
    b, _ := inserted.([]byte)
    w.Write(b)
}

// this function is used to show a job's description
func (rs jobsResource) DescriptionByID(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    job := jobs.GetJobByValue(jobs.SetupClient(), bson.M{"id": string(d)})
    w.Write([]byte(job.Name + " - " + job.Description + "\n"))
}

// this function is used to show a job's status
func (rs jobsResource) StatusByID(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    job := jobs.GetJobByValue(jobs.SetupClient(), bson.M{"id": string(d)})
    w.Write([]byte(job.Name + " - " + job.Status + "\n"))
}

// this function is used to show the status of all jobs
func (rs jobsResource) AllStatus(w http.ResponseWriter, r *http.Request) {
    jobs := jobs.GetAll(jobs.SetupClient())
    for _, job := range jobs {
	w.Write([]byte(job.Name + " - " + job.Status + "\n"))
    }
}

// this function is used to update a job
func (rs jobsResource) Update(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("update a job"))
}

// this function is used to delete a job
func (rs jobsResource) Delete(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    _ = jobs.DeleteJobByValue(jobs.SetupClient(), bson.M{"id": string(d)})
    w.Write([]byte("Job removed!\n"))
}

