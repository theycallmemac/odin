package main

import (
    "./jobs"
    "net/http"
    "io/ioutil"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/go-chi/chi"
)

type jobsResource struct{}

type NewJob struct {
    ID string `yaml:"id"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Status string `yaml:"status"`
    Schedule string `yaml:"schedule"`
}

func (rs jobsResource) Routes() chi.Router {
    r := chi.NewRouter()
    r.Get("/", rs.List)
    r.Post("/", rs.Create)
    r.Put("/", rs.Delete)

    r.Route("/info", func(r chi.Router) {
            r.Post("/description", rs.DescriptionByID)
            r.Post("/status", rs.StatusByID)
            r.Get("/status/all", rs.AllStatus)
            r.Put("/", rs.Update)
            r.Delete("/", rs.Delete)
    })

    return r
}

func (rs jobsResource) List(w http.ResponseWriter, r *http.Request) {
    jobList := jobs.GetAll(jobs.SetupClient())
    w.Write([]byte(jobs.Format("ID", "NAME", "DESCRIPTION", "LANGUAGE", "STATUS", "SCHEDULE")))
    for _, job := range jobList {
        w.Write([]byte(jobs.Format(job.ID, job.Name, job.Description, job.Language ,job.Status, job.Schedule)))
    }
}

func (rs jobsResource) Create(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    inserted := jobs.InsertIntoMongo(jobs.SetupClient(), d)
    b, _ := inserted.([]byte)
    w.Write(b)
}

func (rs jobsResource) DescriptionByID(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    job := jobs.GetJobByValue(jobs.SetupClient(), bson.M{"id": string(d)})
    w.Write([]byte(job.Name + " - " + job.Description + "\n"))
}

func (rs jobsResource) StatusByID(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    job := jobs.GetJobByValue(jobs.SetupClient(), bson.M{"id": string(d)})
    w.Write([]byte(job.Name + " - " + job.Status + "\n"))
}

func (rs jobsResource) AllStatus(w http.ResponseWriter, r *http.Request) {
    jobs := jobs.GetAll(jobs.SetupClient())
    for _, job := range jobs {
	w.Write([]byte(job.Name + " - " + job.Status + "\n"))
    }
}

func (rs jobsResource) Update(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("update a job"))
}

func (rs jobsResource) Delete(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    _ = jobs.DeleteJobByValue(jobs.SetupClient(), bson.M{"id": string(d)})
    w.Write([]byte("Job removed!\n"))
}

