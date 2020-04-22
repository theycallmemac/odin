package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/go-chi/chi"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/executor"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
)


// create resource type to be used by the router
type linksResource struct{}

func (rs linksResource) Routes() chi.Router {
    // establish new chi router
    r := chi.NewRouter()

    // define routes under the execute endpoint
    r.Post("/add", rs.Link)
	r.Post("/delete", rs.Unlink)
	r.Get("/list", rs.Unlink)

    return r
}

// this function is used to execute the item at the head of the job queue
func (rs executeResource) Link(w http.ResponseWriter, r *http.Request) {
    var en linksResource
    body, err := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &en)
    executor.ReviewError(err, "bool")
    go executor.Link(en.Items, 0, httpAddr, en.Store)
}

// this function is used to delete a job link
func (rs executeResource) Unlink(w http.ResponseWriter, r *http.Request) {
    var en linksResource
    body, err := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &en)
    executor.ReviewError(err, "bool")
    go executor.Unlink(en.Items, 1, httpAddr, en.Store)
}
