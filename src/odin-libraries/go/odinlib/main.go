package odinlib

import (
    "fmt"
    "os"
    "time"
)

var ENV_CONFIG bool
var TEST = false

func setAsTest(setting bool) {
    TEST = setting
}

type Odin struct {
    ID string
    Timestamp string
}

var odin Odin

func Setup(config string) (bool, string) {
    var cfg JobConfig
    if ParseYaml(&cfg, ReadFile(config)) {
        odin.ID = cfg.Job.ID
        odin.Timestamp = fmt.Sprint(time.Now().Unix())
    }
    _, ok := os.LookupEnv("ODIN_EXEC_ENV")
    if ok || TEST != false {
        ENV_CONFIG = true
        return true, ""
    } else {
        ENV_CONFIG = false
        return false, "`ODIN_EXEC_ENV` does not exist"
    }
}

func Condition(description string, expression string) bool {
    if ENV_CONFIG {
        return Log("condition", description, expression, odin.ID, odin.Timestamp)
    }
    return false
}

func Watch(description string, expression string) bool {
    if ENV_CONFIG {
        return Log("watch", description, expression, odin.ID, odin.Timestamp)
    }
    return false
}

func Result(description string, expression string) bool {
    if ENV_CONFIG {
        return Log("result", description, expression, odin.ID, odin.Timestamp)
    }
    return false
}

