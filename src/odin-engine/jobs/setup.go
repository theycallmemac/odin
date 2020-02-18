package jobs

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "strings"
)

func notDirectory(dir string) bool {
    if  _, err := os.Stat(dir); os.IsNotExist(err) {
        return true
    }
    return false
}

func makeDirectory(name string) {
    os.MkdirAll(name, 0644)
}

func SetupEnvironment(d []byte) string {
    var job NewJob
    err := json.Unmarshal(d, &job)
    if err != nil {
        panic(err)
    }
    jobsPath := "/etc/odin/jobs/"
    logsPath := "/etc/odin/logs/"
    if notDirectory(jobsPath) && notDirectory(logsPath) {
        makeDirectory(jobsPath)
        makeDirectory(logsPath)
    }
    originalFile := job.File
    fileSlice := strings.Split(job.File, "/")
    job.File = string(fileSlice[len(fileSlice)-1])
    newFilePath := jobsPath + job.ID + "/" + job.File
    if notDirectory(jobsPath + job.ID) {
        makeDirectory(jobsPath + job.ID)
        ioutil.WriteFile(newFilePath, []byte(""), 0644)
        ioutil.WriteFile(logsPath + job.ID, []byte(""), 0644)
    }
    input, err := ioutil.ReadFile(originalFile)
    if err != nil {
        panic(err)
    }
    ioutil.WriteFile(logsPath, []byte(""), 0644)
    ioutil.WriteFile(newFilePath, input, 0644)
    return newFilePath
}
