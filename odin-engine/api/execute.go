package api

import (
	"encoding/json"

	"github.com/theycallmemac/odin/odin-engine/pkg/executor"
	"github.com/theycallmemac/odin/odin-engine/pkg/fsm"
	"github.com/theycallmemac/odin/odin-engine/pkg/repository"
	"github.com/valyala/fasthttp"
)

// ExecNode is a type to be used to unmarshal data into after a HTTP request
// consists of Items, a byte array of marshaled json, and a store of node details
type ExecNode struct {
	Items []byte
	Store fsm.Store
}

// Executor is used to execute the item at the head of the job queue
func Executor(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	var en ExecNode
	json.Unmarshal(ctx.PostBody(), &en)
	go executor.Execute(repo, en.Items, 0, HTTPAddr, en.Store)
}

// ExecuteYaml is used to execute a job passed to the command line tool
func ExecuteYaml(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	var en ExecNode
	json.Unmarshal(ctx.PostBody(), &en)
	go executor.Execute(repo, en.Items, 1, HTTPAddr, en.Store)
}
