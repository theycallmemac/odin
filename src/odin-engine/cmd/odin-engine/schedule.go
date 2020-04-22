package main

import (
    "io/ioutil"
    "net/http"

    "github.com/go-chi/chi"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/scheduler"
)

// create resource type to be used by the router
type scheduleResource struct{}

func (rs scheduleResource) Routes() chi.Router {
    // establish new chi router
    r := chi.NewRouter()

    // define routes under the schedule endpoint
    r.Post("/", rs.Parse)
    return r
}

// this function is used to parse the request and return a cron time format
func (rs scheduleResource) Parse(w http.ResponseWriter, r *http.Request) {
        path, _ := ioutil.ReadAll(r.Body)
        strs := scheduler.Execute(string(path))
        for _, str := range strs {
            w.Write([]byte(str.Minute + " " + str.Hour + " " + str.Dom +  " " + str.Mon + " " + str.Dow + ","))
        }
}

