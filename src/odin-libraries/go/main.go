package odinlib

import (
    "os"
)

var ENV_CONFIG bool
var TEST = false
var ID string

func setAsTest(setting bool) {
    TEST = setting
}

func Setup(config string) (bool, string) {
    var cfg JobConfig
    if ParseYaml(&cfg, ReadFile(config)) {
        ID = cfg.Job.ID
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
        return Log("condition", description, expression, ID)
    }
    return false
}

func Watch(description string, expression string) bool {
    if ENV_CONFIG {
        return Log("watch", description, expression, ID)
    }
    return false
}

func Result(description string, expression string) bool {
    if ENV_CONFIG {
        return Log("result", description, expression, ID)
    }
    return false
}

