package main

import (
        "./executor"
        "io/ioutil"
        "net/http"
	"github.com/go-chi/chi"
)

type executeResource struct{}

func (rs executeResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", rs.Executor)
	r.Post("/yaml", rs.ExecuteYaml)
	return r
}

func (rs executeResource) Executor(w http.ResponseWriter, r *http.Request) {
        path, err := ioutil.ReadAll(r.Body)
        executor.ReviewError(err, "bool")
        go executor.Execute(string(path), 0)
}

func (rs executeResource) ExecuteYaml(w http.ResponseWriter, r *http.Request) {
        path, err := ioutil.ReadAll(r.Body)
        executor.ReviewError(err, "bool")
        go executor.Execute(string(path), 1)
}
