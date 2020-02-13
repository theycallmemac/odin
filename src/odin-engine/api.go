package main

import (
	"net/http"
        "odin/src/odin-engine/jobs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Mount("/execute", executeResource{}.Routes())
	r.Mount("/jobs", jobsResource{}.Routes())
	r.Mount("/schedule", scheduleResource{}.Routes())
        go jobs.StartTicker()
	http.ListenAndServe(":3939", r)
}
