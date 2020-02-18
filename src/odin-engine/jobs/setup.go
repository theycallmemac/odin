package jobs

import (
    "encoding/json"
    "io/ioutil"
    "os"
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

func SetupEnvironment(d []byte) {
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
    if notDirectory(jobsPath + job.ID) {
        makeDirectory(jobsPath + job.ID)
        ioutil.WriteFile(jobsPath + job.ID + job.File, []byte(""), 0644)
        ioutil.WriteFile(logsPath + job.ID, []byte(""), 0644)
    }
    input, err := ioutil.ReadFile(job.File)
    if err != nil {
        panic(err)
    }
    ioutil.WriteFile(logsPath, []byte(""), 0644)
    ioutil.WriteFile(jobsPath + job.ID + "/hello.py", input, 0644)
}
