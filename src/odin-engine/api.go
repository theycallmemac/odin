package main

import (
    "io/ioutil"
    "log"
    "os"
    "os/user"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/jobs"

    "gopkg.in/yaml.v2"
)

// create OdinConfig type to be used for accessing config information
type OdinConfig struct {
    Odin OdinType `yaml:"odin"`
    Mongo MongoType `yaml:"mongo"`
}

// create ProviderType type to be used for accessing odin information in the config
type OdinType struct {
    Master string `yaml:"master"`
    Port string `yaml:"port"`
}

// create ProviderType type to be used for accessing mongo information in the config
type MongoType struct {
    Address string `yaml:"address"`
}

// this function is used to read a file
// parameters: name (a string containing the path to a file)
// returns: []byte (an array of bytes containing the contents of the file)
func readFile(name string) []byte {
    file, err := os.Open(name)
    if err != nil {
        log.Fatal(err)
    }
    bytes, err := ioutil.ReadAll(file)
    defer file.Close()
    return bytes
}

// this function is used to unmarshal YAML
// parameters: byteArray (an array of bytes representing the contents of a file)
// returns: Config (a struct form of the YAML)
func unmarsharlYaml(byteArray []byte) OdinConfig {
    var cfg OdinConfig
    err := yaml.Unmarshal([]byte(byteArray), &cfg)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return cfg
}

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
    config := unmarsharlYaml(readFile(usr.HomeDir + "/odin-config.yml"))

    // start the countdown timer for the execution until the first job
    go jobs.StartTicker()

    // listen and service on the provided host and port in ~/odin-config.yml
    http.ListenAndServe(config.Odin.Master + ":" + config.Odin.Port, r)
    
    // set Odin ENV variables to be used by running jobs via Odin SDK 
    setOdinEnv(config.Mongo.Address)
}
