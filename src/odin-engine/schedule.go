package main

import (
        "odin/src/odin-engine/scheduler"
        "io/ioutil"
        "net/http"
	"github.com/go-chi/chi"
)

type scheduleResource struct{}

func (rs scheduleResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", rs.Parse)
	return r
}

func (rs scheduleResource) Parse(w http.ResponseWriter, r *http.Request) {
        path, _ := ioutil.ReadAll(r.Body)
        go parser.Execute(string(path))
}

