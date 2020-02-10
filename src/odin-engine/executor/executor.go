package executor

import (
    "fmt"
    "os/exec"
    "strings"
)

type Data struct {
    output []byte
    error  error
}

func runCommand(ch chan<- Data, language string, file string) {
    cmd := exec.Command(language, file)
    // data, err output after job is finished running
    data, err := cmd.CombinedOutput()
    fmt.Println(string(data))
    ch <- Data{
        error:  err,
        output: data,
    }
}

func Execute(filename string) bool {
    if exists(filename) {
        channel := make(chan Data)
        path := strings.Split(filename, "/")
        basePath := strings.Join(path[:len(path)-1], "/")
        language, file := getYaml(filename)
        destFile := basePath + "/" + file
        go runCommand(channel, language, destFile)
        res := <-channel
        ReviewError(res.error, "bool")
        return true
    } else {
        return false
    }
}

