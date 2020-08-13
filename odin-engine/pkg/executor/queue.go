package executor

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/theycallmemac/odin/odin-engine/pkg/fsm"
	"github.com/theycallmemac/odin/odin-engine/pkg/repository"
)

// BatchRun is called on a queue type and is used to run the batch loop to run all executions
// parameters:  store (a store of node information)
// returns: nil
func (queue Queue) BatchRun(repo repository.Repository, httpAddr string, store fsm.Store) {
	for _, job := range queue {
		go func(job JobNode) {
			channel := make(chan Data)
			go job.runCommand(repo, channel, httpAddr, store)
		}(job)
	}
}

// UpdateRuns is called on a queue type and is used to update the run number for each job
// parameters: httpAddr (an address string for the node)
// returns: nil
func (queue Queue) UpdateRuns(httpAddr string) {
	for _, job := range queue {
		go func(job JobNode) {
			inc := job.Runs + 1
			go makePutRequest("http://localhost"+httpAddr+"/jobs/info/runs", bytes.NewBuffer([]byte(job.ID+" "+strconv.Itoa(inc)+" "+fmt.Sprint(job.UID))))
		}(job)
	}
}
