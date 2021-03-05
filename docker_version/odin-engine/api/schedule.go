package api

import (
        "fmt"

	"github.com/theycallmemac/odin/odin-engine/pkg/scheduler"
        "github.com/valyala/fasthttp"
)

// GetJobSchedule is used to parse the request and return a cron time format
func GetJobSchedule(ctx *fasthttp.RequestCtx) {
	strs := scheduler.Execute(string(ctx.PostBody()))
	for _, str := range strs {
		fmt.Fprintf(ctx, str.Minute + " " + str.Hour + " " + str.Dom + " " + str.Mon + " " + str.Dow + ",")
	}
}
