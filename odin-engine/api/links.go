package api

import (
        "fmt"
	"strings"

	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
        "github.com/valyala/fasthttp"
)

// LinkJobs is used to link two jobs together
func LinkJobs(ctx *fasthttp.RequestCtx) {
	split := strings.Split(string(ctx.PostBody()), "_")
	from, to, uid := split[0], split[1], split[2]
	client, _ := jobs.SetupClient()
	updated := jobs.AddJobLink(client, from, to, uid)
	if updated == 1 {
		fmt.Fprintf(ctx, "Job " + from + " linked to " + to + "!\n")
	} else {
		fmt.Fprintf(ctx, "Job " + from + " could not be linked to " + to + ".\n")
	}
}

// UnlinkJobs is used to delete a job link
func UnlinkJobs(ctx *fasthttp.RequestCtx) {
	split := strings.Split(string(ctx.PostBody()), "_")
	from, to, uid := split[0], split[1], split[2]
	client, _ := jobs.SetupClient()
	updated := jobs.DeleteJobLink(client, from, to, uid)
	if updated == 1 {
		fmt.Fprintf(ctx, "Job " + to + " unlinked from " + from + "!\n")
	} else {
		fmt.Fprintf(ctx, "Job " + to + " has no links!\n")
	}
}
