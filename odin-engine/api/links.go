package api

import (
	"fmt"
	"strings"

	"github.com/theycallmemac/odin/odin-engine/pkg/repository"
	"github.com/valyala/fasthttp"
)

// LinkJobs is used to link two jobs together
func LinkJobs(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	split := strings.Split(string(ctx.PostBody()), "_")
	from, to, uid := split[0], split[1], split[2]
	if err := repo.AddJobLink(ctx, from, to, uid); err != nil {
		fmt.Fprintf(ctx, "[FAILED] Job %s could not be linked to %s: %v\n", from, to, err)
	} else {
		fmt.Fprintf(ctx, "[SUCCESS] Job %s linked to %s\n", from, to)
	}
}

// UnlinkJobs is used to delete a job link
func UnlinkJobs(repo repository.Repository, ctx *fasthttp.RequestCtx) {
	split := strings.Split(string(ctx.PostBody()), "_")
	from, to, uid := split[0], split[1], split[2]
	if err := repo.DeleteJobLink(ctx, from, to, uid); err != nil {
		fmt.Fprintf(ctx, "[FAILED] Link %s could not be unlinked from %s: %v\n", to, from, err)
	} else {
		fmt.Fprintf(ctx, "[SUCCESS] Link %s unlinked from %s\n", to, from)
	}
}
