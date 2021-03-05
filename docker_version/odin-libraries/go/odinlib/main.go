package odinlib

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// EnvConfig is a boolean used to check if the ODIN_EXEC_ENV environment variable has been set
var EnvConfig bool

// Test is a boolean used to check if the current run is a test case
var Test = false

var cfg JobConfig

var odin *Odin

func setAsTest(setting bool) {
	Test = setting
}

// Odin is a type used to consolidate information about the Job ID and the time of stat collection
type Odin struct {
	ID        string
	Timestamp string
}

// newOdin is used to create a new Odin struct
// parameters: id (a string of a Job ID)
// returns: *Odin (a new Odin struct)
func newOdin(id string) *Odin {
	return &Odin{ID: id, Timestamp: fmt.Sprint(time.Now().Unix())}
}

// Setup is used to define the correct metadata surrounding the job
// parameters: config (a string containing the path to the Yaml file)
// RETURNS: *Odin (a new Odin struct), string (a potential error message)
func Setup(config string) (*Odin, string) {
	var path string
	if Test != false {
		path = ""
		files, _ := ioutil.ReadDir("/etc/odin/jobs/")
		for _, f := range files {
			newPath := "/etc/odin/jobs/" + f.Name() + "/" + config
			if _, err := os.Stat(newPath); !os.IsNotExist(err) {
				path = newPath
			}
		}
		if path == "" {
			return nil, "Failed to read yaml config file"
		}
	} else {
		path = config
	}
	if ParseYaml(&cfg, ReadFile(path)) {
		odin = newOdin(cfg.Job.ID)
	}
	_, ok := os.LookupEnv("ODIN_EXEC_ENV")
	if ok || Test != false {
		EnvConfig = true
		return odin, ""
	}
	EnvConfig = false
	return odin, "`ODIN_EXEC_ENV` does not exist"
}

// Condition is used to capture a boolean operation in a job in the form of if statements
func (odin Odin) Condition(description string, expression string) bool {
	if EnvConfig {
		return Log("condition", description, expression, odin.ID, odin.Timestamp)
	}
	return false
}

// Watch is used to capture any arbitrary value in a job
func (odin Odin) Watch(description string, expression string) bool {
	if EnvConfig {
		return Log("watch", description, expression, odin.ID, odin.Timestamp)
	}
	return false
}

// Result is used to capture the output status of a job
func (odin Odin) Result(description string, expression string) bool {
	if EnvConfig {
		return Log("result", description, expression, odin.ID, odin.Timestamp)
	}
	return false
}
