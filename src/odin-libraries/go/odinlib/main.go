package odinlib

import (
    "fmt"
    "io/ioutil"
    "os"
    "time"
)

var ENV_CONFIG bool

var TEST = false

var cfg JobConfig

var odin *Odin

func setAsTest(setting bool) {
    TEST = setting
}

type Odin struct {
    ID string
    Timestamp string
}

func newOdin(id string) *Odin {
    return &Odin{ID: id, Timestamp: fmt.Sprint(time.Now().Unix())}
}

func Setup(config string) (*Odin, string) {
    var path string
    if TEST != false {
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
    if ok || TEST != false {
        ENV_CONFIG = true
        return odin, ""
    } else {
        ENV_CONFIG = false
        return odin, "`ODIN_EXEC_ENV` does not exist"
    }
}

func (odin Odin) Condition(description string, expression string) bool {
    if ENV_CONFIG {
        return Log("condition", description, expression, odin.ID, odin.Timestamp)
    }
    return false
}

func (odin Odin) Watch(description string, expression string) bool {
    if ENV_CONFIG {
        return Log("watch", description, expression, odin.ID, odin.Timestamp)
    }
    return false
}

func (odin Odin) Result(description string, expression string) bool {
    if ENV_CONFIG {
        return Log("result", description, expression, odin.ID, odin.Timestamp)
    }
    return false
}
