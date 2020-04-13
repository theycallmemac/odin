package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/go-chi/chi"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/executor"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
)


type AltNode struct {
    Items []byte
    Store fsm.Store
}

// create resource type to be used by the router
type executeResource struct{}

func (rs executeResource) Routes() chi.Router {
    // establish new chi router
    r := chi.NewRouter()

    // define routes under the execute endpoint
    r.Post("/", rs.Executor)
    r.Post("/yaml", rs.ExecuteYaml)

    return r
}

// this function is used to execute the item at the head of the job queue
func (rs executeResource) Executor(w http.ResponseWriter, r *http.Request) {
    var an AltNode
    body, err := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &an)
    executor.ReviewError(err, "bool")
    go executor.Execute(an.Items, 0, httpAddr, an.Store)
}

// this function is used to execute a job passed to the command line tool
func (rs executeResource) ExecuteYaml(w http.ResponseWriter, r *http.Request) {
    var an AltNode
    body, err := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &an)
    executor.ReviewError(err, "bool")
    go executor.Execute(an.Items, 1, httpAddr, an.Store)
}
