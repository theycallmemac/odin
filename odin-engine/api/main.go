package api

import (
	"os"
	"syscall"

	"github.com/theycallmemac/odin/odin-engine/pkg/fsm"
	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
	"github.com/theycallmemac/odin/odin-engine/pkg/repository"
	"github.com/valyala/fasthttp"
)

var (
	// HTTPAddr contains the port used by this node
	HTTPAddr string
)

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
	repo  repository.Repository
}

// NewService is used to initialize a new service struct
// parameters: addr (a string of a http address), store (a store of node details)
// returns: *Service (a newly initialized service struct)
func NewService(addr string, store fsm.Store, repo repository.Repository) *Service {
	return &Service{
		addr:  addr,
		store: store,
		repo:  repo,
	}
}

// Start is called on a Service and is used to kick off its execution
// parameters: nil
// returns: nil
func (service *Service) Start() {
	routes := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/cluster/join":
			service.JoinCluster(ctx)
		case "/cluster/leave":
			service.LeaveCluster(ctx)
		case "/execute":
			Executor(service.repo, ctx)
		case "/execute/yaml":
			ExecuteYaml(service.repo, ctx)
		case "/jobs/add":
			AddJob(service.repo, ctx)
		case "/jobs/delete":
			DeleteJob(service.repo, ctx)
		case "/jobs/info/update":
			UpdateJob(service.repo, ctx)
		case "/jobs/info/description":
			GetJobDescription(service.repo, ctx)
		case "/jobs/info/runs":
			UpdateJobRuns(service.repo, ctx)
		case "/jobs/list":
			ListJobs(service.repo, ctx)
		case "/jobs/logs":
			GetJobLogs(ctx)
		case "/links/add":
			LinkJobs(service.repo, ctx)
		case "/links/delete":
			UnlinkJobs(service.repo, ctx)
		case "/schedule":
			GetJobSchedule(ctx)
		case "/stats/add":
			AddJobStats(service.repo, ctx)
		case "/stats/get":
			GetJobStats(service.repo, ctx)
		}
	}

	// start the countdown timer for the execution until the first job
	go jobs.StartTicker(service.repo, service.store, service.addr)

	HTTPAddr = service.addr
	fasthttp.ListenAndServe(HTTPAddr, routes)
}
