package api

import (
    "bytes"
    "encoding/json"
    "fmt"

    "github.com/valyala/fasthttp"
)

// JoinCluster is used to join a node to an exisiting cluster
func (service *Service) JoinCluster(ctx *fasthttp.RequestCtx) {
        m := map[string]string{}
        if err := json.NewDecoder(bytes.NewReader(ctx.PostBody())).Decode(&m); err != nil {
                fmt.Fprintf(ctx, "%d\n", fasthttp.StatusBadRequest)
                return
        }
        if len(m) != 2 {
                fmt.Fprintf(ctx, "%d\n", fasthttp.StatusBadRequest)
                return
        }

        remoteAddr, ok := m["addr"]
        if !ok {
                fmt.Fprintf(ctx, "%d\n", fasthttp.StatusBadRequest)
                return
        }

        nodeID, ok := m["id"]
        if !ok {
                fmt.Fprintf(ctx, "%d\n", fasthttp.StatusBadRequest)
                return
        }
        fmt.Println()
        if err := service.store.Join(nodeID, remoteAddr); err != nil {
                fmt.Fprintf(ctx, "%d\n",fasthttp.StatusInternalServerError)
        }
}

// LeaveCluster is used to remove a node from a cluster
func (service *Service) LeaveCluster(ctx *fasthttp.RequestCtx) {
	m := map[string]string{}
	if err := json.NewDecoder(bytes.NewReader(ctx.PostBody())).Decode(&m); err != nil {
		fmt.Fprintf(ctx, "%d\n", fasthttp.StatusBadRequest)
		return
	}
	nodeID, _ := m["id"]
	service.store.Leave(nodeID)
}
