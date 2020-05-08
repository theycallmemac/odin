package main

import (
    "io/ioutil"
    "net/http"
    "strings"

    "github.com/go-chi/chi"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/jobs"
)

// create resource type to be used by the router
type linksResource struct{}

func (rs linksResource) Routes() chi.Router {
    // establish new chi router
    r := chi.NewRouter()

    r.Post("/add", rs.Link)
    r.Post("/delete", rs.Unlink)

    return r
}

// this function is used to link two jobs together
func (rs linksResource) Link(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    split := strings.Split(string(body), "_")
    from, to, uid := split[0], split[1], split[2]
    client, _ := jobs.SetupClient()
    updated := jobs.AddJobLink(client, from, to, uid)
    if updated == 1 {
        w.Write([]byte("Job " + from + " linked to " + to + "!\n"))
    } else {
        w.Write([]byte("Job " + from + " could not be linked to " + to + ".\n"))
    }
}

// this function is used to delete a job link
func (rs linksResource) Unlink(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    split := strings.Split(string(body), "_")
    from, to, uid := split[0], split[1], split[2]
    client, _ := jobs.SetupClient()
    updated := jobs.DeleteJobLink(client, from, to, uid)
    if updated == 1 {
        w.Write([]byte("Job " + to + " unlinked from " + from + "!\n"))
    } else {
        w.Write([]byte("Job " + to + " has no links!\n"))
    }
}
