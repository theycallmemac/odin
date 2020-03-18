package main

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "os"
    "strings"

    "go.mongodb.org/mongo-driver/bson"
    "github.com/go-chi/chi"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/jobs"
)


// create NewJob type to tbe used for accessing job information
type NewJob struct {
    ID string `yaml:"id"`
    UID string `yaml:"uid"`
    GID string `yaml:"gid"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Stats string `yaml:"stats"`
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
            r.Post("/stats", rs.StatsByID)
            r.Put("/", rs.Update)
    })

    r.Route("/logs", func(r chi.Router) {
        r.Post("/", rs.LogByID)
    })

    return r
}

// this function is used to list the current jobs running
func (rs jobsResource) List(w http.ResponseWriter, r *http.Request) {
    jobList := jobs.GetAll(jobs.SetupClient())
    w.Write([]byte(jobs.Format("ID", "NAME", "DESCRIPTION", "LANGUAGE", "SCHEDULE")))
    for _, job := range jobList {
        w.Write([]byte(jobs.Format(job.ID, job.Name, job.Description, job.Language, job.Schedule[:len(job.Schedule)-1])))
    }
}

// this function is used to create a new job
func (rs jobsResource) Create(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    path := jobs.SetupEnvironment(d)
    status := jobs.InsertIntoMongo(jobs.SetupClient(), d, path)
    w.Write([]byte(status))
}

// this function is used to show a job's description
func (rs jobsResource) DescriptionByID(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    job := jobs.GetJobByValue(jobs.SetupClient(), bson.M{"id": string(d)})
    w.Write([]byte(job.Name + " - " + job.Description + "\n"))
}

// this function is used to show a job's stats
func (rs jobsResource) StatsByID(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    statsList := jobs.GetJobStats(jobs.SetupClient(), string(d))
    w.Write([]byte(jobs.Format("ID", "DESCRIPTION", "", "TYPE", "VALUE")))
    for _, stat := range statsList {
        w.Write([]byte(jobs.Format(stat.ID, stat.Description, "", stat.Type, stat.Value)))
    }
}

// this function is used to update a job
func (rs jobsResource) Update(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    data := strings.Split(string(d), ",")
    job := jobs.GetJobByValue(jobs.SetupClient(), bson.M{"id": data[0]})
    if jobs.NotEmpty(data[1]) {
        job.Name = data[1]
    }
    if jobs.NotEmpty(data[2]) {
        job.Description = data[2]
    }
    if jobs.NotEmpty(data[3]) {
        ioutil.WriteFile(".tmp.yml", []byte("provider:\n  name: 'odin'\n  version: '1.0.0'\njob:\n  name: ''\n  description: ''\n  language: ''\n  file: ''\n  schedule: "+data[3] + "\n\n"), 0654)
        resp := jobs.MakePostRequest("http://localhost:3939/schedule", bytes.NewBuffer([]byte(".tmp.yml")))
        os.Remove(".tmp.yml")
        job.Schedule = resp
    }
    _ = jobs.UpdateJobByValue(jobs.SetupClient(), job)
    w.Write([]byte("Updated job " +  data[0] + " successfully\n"))
}

// this function is used to delete a job
func (rs jobsResource) Delete(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    if jobs.DeleteJobByValue(jobs.SetupClient(), bson.M{"id": string(d)}) {
        w.Write([]byte("Job removed!\n"))
    } else {
        w.Write([]byte("Job with that ID does not exist!\n"))
    }
}

// this function is used to retrieve the logs for a job
func (rs jobsResource) LogByID(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    log, _ := ioutil.ReadFile("/etc/odin/logs/" + string(d))
    w.Write([]byte("\n" + string(log) + "\n"))
}

