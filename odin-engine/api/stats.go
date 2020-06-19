package api

import (
	"context"
        "fmt"
	"strings"

	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
        "github.com/valyala/fasthttp"

	"go.mongodb.org/mongo-driver/mongo"
)

// JobStats is a type to be used for accessing and storing job stats information
type JobStats struct {
	ID          string
	Description string
	Type        string
	Value       string
	Timestamp   string
}

// AddJobStats is used to parse collected metrics
func AddJobStats(ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.PostBody()), ",")
	typeOfValue, desc, value, id, timestamp := args[0], args[1], args[2], args[3], args[4]
	client, err := jobs.SetupClient()
	if err != nil {
		fmt.Fprintf(ctx, "MongoDB cannot be accessed at the moment\n")
	} else {
		if InsertIntoMongo(client, typeOfValue, desc, value, id, timestamp) {
			fmt.Fprintf(ctx,"200")
		} else {
			fmt.Fprintf(ctx,"500")
		}
	}
}

// InsertIntoMongo is used to add collected metrics to the observability collection
// parameters: client (a *mongo.Client), typeOfValue (a string of the type of value being stored), desc (a string describing the value being stored), value (a string of the value being stored), id (a string of the associated Job ID), timestamp (a string of the unix time at which the operation took place)
// returns: bool (true is successful, false if otherwise)
func InsertIntoMongo(client *mongo.Client, typeOfValue string, desc string, value string, id string, timestamp string) bool {
	var js JobStats
	js.ID = id
	js.Description = desc
	js.Type = typeOfValue
	js.Value = value
	js.Timestamp = timestamp
	collection := client.Database("odin").Collection("observability")
	_, err := collection.InsertOne(context.TODO(), js)
	client.Disconnect(context.TODO())
	if err != nil {
		return false
	}
	return true
}

// GetJobStats is used to show stats collected by a specified job
func GetJobStats(ctx *fasthttp.RequestCtx) {
	client, err := jobs.SetupClient()
	if err != nil {
		fmt.Fprintf(ctx, "MongoDB cannot be accessed at the moment\n")
	} else {
		statsList := jobs.GetJobStats(client, string(ctx.PostBody()))
		fmt.Fprintf(ctx, jobs.Format("ID", "DESCRIPTION", "TYPE", "VALUE"))
		for _, stat := range statsList {
			fmt.Fprintf(ctx, jobs.Format(stat.ID, stat.Description, stat.Type, stat.Value))
		}
	}
}
