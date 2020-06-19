package api

import (
        "os"
        "syscall"

	"github.com/valyala/fasthttp"
        "github.com/theycallmemac/odin/odin-engine/pkg/fsm"
        "github.com/theycallmemac/odin/odin-engine/pkg/jobs"
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
func (service *Service) Start() {
	routes := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
                        case "/cluster/join":
                                service.JoinCluster(ctx)
                        case "/cluster/leave":
                                service.LeaveCluster(ctx)
                        case "/execute":
                                Executor(ctx)
                        case "/execute/yaml":
                                ExecuteYaml(ctx)
                        case "/jobs/add":
                                AddJob(ctx)
                        case "/jobs/delete":
                                DeleteJob(ctx)
                        case "/jobs/info/update":
                                UpdateJob(ctx)
                        case "/jobs/info/description":
                                GetJobDescription(ctx)
                        case "/jobs/info/runs":
                                UpdateJobRuns(ctx)
                        case "/jobs/list":
                                ListJobs(ctx)
                        case "/jobs/logs":
                                GetJobLogs(ctx)
                        case "/links/add":
                                LinkJobs(ctx)
                        case "/links/delete":
                                UnlinkJobs(ctx)
		        case "/schedule":
			        GetJobSchedule(ctx)
                        case "/stats/add":
                                AddJobStats(ctx)
                        case "/stats/get":
                                GetJobStats(ctx)
		}
	}

        // start the countdown timer for the execution until the first job
        go jobs.StartTicker(service.store, service.addr)

        HTTPAddr = service.addr
	fasthttp.ListenAndServe(HTTPAddr, routes)
}
