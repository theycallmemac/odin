package executor

import (
    "fmt"
    "os"
)

func exists(name string) bool {
    _, err := os.Stat(name)
    return ProcessError(err, "dir")
}

func ProcessError(err error, errType string) bool {
    switch errType {
        case "bool":
            return err != nil
        case "dir":
            return !os.IsNotExist(err)
    }
    return false
}

func ReviewError(err error, errType string) {
    if ProcessError(err, errType) {
        fmt.Println(err)
    }
}

