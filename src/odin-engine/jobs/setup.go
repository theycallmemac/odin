package jobs

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "os/user"
    "strconv"
    "strings"
)

func notDirectory(dir string) bool {
    if  _, err := os.Stat(dir); os.IsNotExist(err) {
        return true
    }
    return false
}

func makeDirectory(name string) {
    os.MkdirAll(name, 0654)
}

func ChownR(path string, uid, gid int) {
    s := strings.Split(path, "/")
   for i := len(s) - 1; i > 2; i-- {
        os.Chown(strings.Join(s[:i], "/"), uid, gid)
    }
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
        group, _ := user.LookupGroup("odin")
        gid, _ := strconv.Atoi(group.Gid)
	ChownR(newFilePath, 0, gid)
	ChownR(logsPath + job.ID, 0, gid)
        ioutil.WriteFile(newFilePath, []byte(""), 0654)
        ioutil.WriteFile(logsPath + job.ID, []byte(""), 0766)
    }
    input, err := ioutil.ReadFile(originalFile)
    if err != nil {
        panic(err)
    }
    ioutil.WriteFile(logsPath, []byte(""), 0766)
    ioutil.WriteFile(newFilePath, input, 0654)
    return newFilePath
}
