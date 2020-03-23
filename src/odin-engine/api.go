package main

import (
    "os"
    "os/user"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/jobs"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/resources"
)

// set Odin ENV variables to be used by running jobs via Odin SDK 
func setOdinEnv(mongoDbUrl string) {
    // tells SDK that job is running within an Odin Environment
    os.Setenv("ODIN_EXEC_ENV", "True")
    // Is read by Odin SDK to connect to logging DB
    os.Setenv("ODIN_MONGODB", mongoDbUrl)
}

func main() {
    // restablish new chi router
    r := chi.NewRouter()

    // tell router to use some middlewares
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // set the base endpoint to return nothing
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(""))
    })

    // define current odin-engine endpoints
    r.Mount("/execute", executeResource{}.Routes())
    r.Mount("/jobs", jobsResource{}.Routes())
    r.Mount("/schedule", scheduleResource{}.Routes())

    // load the odin config yaml
    usr, _ := user.Current()
    config := resources.UnmarsharlYaml(resources.ReadFileBytes(usr.HomeDir + "/odin-config.yml"))

    // start the countdown timer for the execution until the first job
    go jobs.StartTicker()

    // listen and service on the provided host and port in ~/odin-config.yml
    http.ListenAndServe(config.OdinVars.Master + ":" + config.OdinVars.Port, r)

    // set Odin ENV variables to be used by running jobs via Odin SDK 
    setOdinEnv(config.Mongo.Address)
}
