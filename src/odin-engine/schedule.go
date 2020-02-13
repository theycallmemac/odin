package main

import (
    "./scheduler"
    "io/ioutil"
    "net/http"
	"github.com/go-chi/chi"
)


type StringFormat struct {
    Minute string
    Hour string
    Dom string
    Mon string
    Dow string
}

type scheduleResource struct{}

func (rs scheduleResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", rs.Parse)
	return r
}

func (rs scheduleResource) Parse(w http.ResponseWriter, r *http.Request) {
        path, _ := ioutil.ReadAll(r.Body)
        strs := scheduler.Execute(string(path))
        for _, str := range strs {
            w.Write([]byte(str.Minute + " " + str.Hour + " " + str.Dom +  " " + str.Mon + " " + str.Dow + ","))
        }
}

