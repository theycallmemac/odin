package main

import (
        "odin/src/odin-engine/executor"
        "fmt"
        "io/ioutil"
        "net/http"
	"github.com/go-chi/chi"
)

type executeResource struct{}

func (rs executeResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", rs.Parse)
	return r
}

func (rs executeResource) Parse(w http.ResponseWriter, r *http.Request) {
        path, _ := ioutil.ReadAll(r.Body)
        go executor.Execute(string(path))
        fmt.Println("job executed successfully")
}

