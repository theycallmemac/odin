package main

import (
    "net/http"

    "github.com/go-chi/chi"
)

type jobsResource struct{}

func (rs jobsResource) Routes() chi.Router {
    r := chi.NewRouter()

    r.Get("/", rs.List)
    r.Post("/", rs.Create)
    r.Put("/", rs.Delete)

    r.Route("/{id}", func(r chi.Router) {
            r.Put("/", rs.Update)
            r.Delete("/", rs.Delete)
    })

    return r
}

func (rs jobsResource) List(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("get a list of jobs"))
}

func (rs jobsResource) Create(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("create a job"))
}

func (rs jobsResource) Get(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("get a job"))
}

func (rs jobsResource) Update(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("update a job"))
}

func (rs jobsResource) Delete(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("delete a job"))
}
