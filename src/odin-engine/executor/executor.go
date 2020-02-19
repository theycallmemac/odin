package executor

import (
    "io"
    "os"
    "os/exec"
    "strings"

    "github.com/sirupsen/logrus"
)

type Data struct {
    output []byte
    error  error
}

func runCommand(ch chan<- Data, language string, file string, id string) {
    cmd := exec.Command(language, file)
    // data, err output after job is finished running
    data, err := cmd.CombinedOutput()
    var logFile, _ = os.OpenFile("/etc/odin/logs/" + id, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    logrus.SetOutput(io.MultiWriter(logFile, os.Stdout))
    logrus.SetFormatter(&logrus.TextFormatter{})
    if id != "" {
	logrus.WithFields(logrus.Fields{
	    "id": id,
	    "language": language,
	    "file": file,
        }).Info("executed")
    }
    if err != nil {
        logrus.WithFields(logrus.Fields{
	    "id": id,
	    "language": language,
	    "file": file,
            "error": err,
        }).Warn("failed")
    }
    ch <- Data{
        error:  err,
        output: data,
    }
}

func executeYaml(filename string) bool {
    if exists(filename) {
        singleChannel := make(chan Data)
        path := strings.Split(filename, "/")
        basePath := strings.Join(path[:len(path)-1], "/")
        language, file := getYaml(filename)
        destFile := basePath + "/" + file
        go runCommand(singleChannel, language, destFile, "")
        res := <-singleChannel
        ReviewError(res.error, "bool")
        return true
    } else {
        return false
    }
}

func executeLang(contents string) bool {
    contentList := strings.Split(contents, " ")
    language, filename, id := contentList[0], contentList[1], contentList[2]
    if exists(filename) {
        channel := make(chan Data)
        go runCommand(channel, language, filename, id)
        res := <-channel
        ReviewError(res.error, "bool")
        return true
    } else {
        return false
    }
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

