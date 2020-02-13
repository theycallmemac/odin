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

func executeYaml(filename string) bool {
    singleChannel := make(chan Data)
    path := strings.Split(filename, "/")
    basePath := strings.Join(path[:len(path)-1], "/")
    language, file := getYaml(filename)
    destFile := basePath + "/" + file
    go runCommand(singleChannel, language, destFile)
    res := <-singleChannel
    ReviewError(res.error, "bool")
    return true
}

func executeLang(contents string) bool {
    contentList := strings.Split(contents, " ")
    language, filename := contentList[0], contentList[1]
    channel := make(chan Data)
    go runCommand(channel, language, filename)
    res := <-channel
    ReviewError(res.error, "bool")
    return true
}

func Execute(contents string, process int) bool {
    switch process {
        case 0:
            return executeLang(contents)
        case 1:
            return executeYaml(contents)
    }
    return false
}

