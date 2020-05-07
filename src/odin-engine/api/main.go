package api

import (
    "os"
    "net/http"
    "syscall"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/jobs"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
)


var httpAddr string

// set Odin ENV variables to be used by running jobs via Odin SDK 
func SetOdinEnv(mongoDbUrl string) {
    // tells SDK that job is running within an Odin Environment
    os.Setenv("ODIN_EXEC_ENV", "True")
    syscall.Exec(os.Getenv("zsh"), []string{os.Getenv("zsh")}, syscall.Environ())
    // Is read by Odin SDK to connect to logging DB
    os.Setenv("ODIN_MONGODB", mongoDbUrl)
}

// create service type to be used by the raft consensus protocol
// consists of a base http address and a store in the finite state machine
type Service struct {
    addr  string
    store fsm.Store
}

// this function is used to initialise a new service struct
// parameters: addr (a string of a http address), store (a store of node details)
// returns: *Service (a newly initialized service struct)
func NewService(addr string, store fsm.Store) *Service {
    return &Service{
        addr:  addr,
	store: store,
    }
}

// this function is called on a service and is used to start it
// parameters: nil
// returns: nil
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
    r.Mount("/leave", leaveResource{}.Routes(s))
    r.Mount("/schedule", scheduleResource{}.Routes())

    // start the countdown timer for the execution until the first job
    go jobs.StartTicker(s.store, s.addr)

    httpAddr = s.addr
    // listen and service on the provided host and port in ~/odin-config.yml
    http.ListenAndServe(s.addr, r)
}
