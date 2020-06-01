package api

import (
    "io/ioutil"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/theycallmemac/odin/odin-engine/pkg/jobs"
)

// create resource type to be used by the router
type statsResource struct{}

func (rs statsResource) Routes() chi.Router {
    // establish new chi router
    r := chi.NewRouter()

    // define routes under the schedule endpoint
    r.Post("/get", rs.GetJobStats)
    return r
}

func (rs statsResource) GetJobStats(w http.ResponseWriter, r *http.Request) {
    d, _ := ioutil.ReadAll(r.Body)
    client, err := jobs.SetupClient()
    if err != nil {
        w.Write([]byte("MongoDB cannot be accessed at the moment\n"))
    } else {
        statsList := jobs.GetJobStats(client, string(d))
        w.Write([]byte(jobs.Format("ID", "DESCRIPTION", "TYPE", "VALUE")))
        for _, stat := range statsList {
            w.Write([]byte(jobs.Format(stat.ID, stat.Description, stat.Type, stat.Value)))
        }
    }
}

