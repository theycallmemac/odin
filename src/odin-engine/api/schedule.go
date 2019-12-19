package main

import (
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
	w.Write([]byte("parse a schedule string"))
}

