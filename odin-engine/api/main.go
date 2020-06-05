package api

import (
	"net/http"
	"os"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/theycallmemac/odin/odin-engine/pkg/fsm"
	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
)

var httpAddr string

// SetOdinEnv is used to set variables to be used by running jobs via Odin SDK
// parameters: mongoURL (a string of the address for the MongoDB instance)
// returns: nil
func SetOdinEnv(mongoURL string) {
	// tells SDK that job is running within an Odin Environment
	os.Setenv("ODIN_EXEC_ENV", "True")
	syscall.Exec(os.Getenv("zsh"), []string{os.Getenv("zsh")}, syscall.Environ())
	// Is read by Odin SDK to connect to logging DB
	os.Setenv("ODIN_MONGODB", mongoURL)
}

// Service is a type to be used by the raft consensus protocol
// consists of a base http address and a store in the finite state machine
type Service struct {
	addr  string
	store fsm.Store
}

// NewService is used to initialize a new service struct
// parameters: addr (a string of a http address), store (a store of node details)
// returns: *Service (a newly initialized service struct)
func NewService(addr string, store fsm.Store) *Service {
	return &Service{
		addr:  addr,
		store: store,
	}
}

// Start is called on a Service and is used to kick off its execution
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
	r.Mount("/stats", statsResource{}.Routes())
	r.Mount("/links", linksResource{}.Routes())

	// start the countdown timer for the execution until the first job
	go jobs.StartTicker(s.store, s.addr)

	httpAddr = s.addr
	// listen and service on the provided host and port in ~/odin-config.yml
	http.ListenAndServe(s.addr, r)
}
