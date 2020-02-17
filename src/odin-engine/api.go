package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/jobs"
)

func main() {
    // restablish new chi router
    r := chi.NewRouter()

    // tell router to use some middlewares
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // set the base endpoint to return nothing
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte(""))
    })

    // define current odin-engine endpoints
    r.Mount("/execute", executeResource{}.Routes())
    r.Mount("/jobs", jobsResource{}.Routes())
    r.Mount("/schedule", scheduleResource{}.Routes())

    // start the countdown timer for the execution until the first job
    go jobs.StartTicker()

    // listen and service on localhost:3939
    http.ListenAndServe(":3939", r)
}
