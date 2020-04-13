package main

import (
    "os"
    "net/http"
    "syscall"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/jobs"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
)

// set Odin ENV variables to be used by running jobs via Odin SDK 
func setOdinEnv(mongoDbUrl string) {
    // tells SDK that job is running within an Odin Environment
    os.Setenv("ODIN_EXEC_ENV", "True")
    syscall.Exec(os.Getenv("zsh"), []string{os.Getenv("zsh")}, syscall.Environ())
    // Is read by Odin SDK to connect to logging DB
    os.Setenv("ODIN_MONGODB", mongoDbUrl)
}

type Service struct {
    addr  string
    store fsm.Store
}

func newService(addr string, store fsm.Store) *Service {
    return &Service{
        addr:  addr,
	store: store,
    }
}

func (s *Service) Start() {
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
    r.Mount("/join", joinResource{}.Routes(s))
    r.Mount("/schedule", scheduleResource{}.Routes())

    // start the countdown timer for the execution until the first job
    go jobs.StartTicker(s.store, s.addr)

    // listen and service on the provided host and port in ~/odin-config.yml
    http.ListenAndServe(s.addr, r)
}
