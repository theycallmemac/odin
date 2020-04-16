package jobs

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "os/user"
    "strconv"
    "strings"
)

// this function is used to check whether or not a directory exists
// parameters: dir (a string of the directory path to check)
// returns: boolean (true if it exists, false otherwise)
func notDirectory(dir string) bool {
    if  _, err := os.Stat(dir); os.IsNotExist(err) {
        return true
    }
    return false
}

// this function is used to create a new directory
// parameters: name (a string of the directory path to create)
// returns: bool
func makeDirectory(name string) bool {
    err := os.MkdirAll(name, 0654)
    if err != nil {
        if notDirectory(name) {
            return false
        }
    }
    return true
}

// this function is used to recursively change the owner of each subdir under /etc
// parameters: path (a string of the directory path to chown), uid (an int used to set the owner uid), gid (an int used to set the owner gid)
// returns: bool
func ChownR(path string, uid, gid int) bool {
    s := strings.Split(path, "/")
    for i := len(s) - 1; i > 2; i-- {
        err := os.Chown(strings.Join(s[:i], "/"), uid, gid)
        if err != nil {
            return false
        }
    }
    return true
}

// this function is used to kickstart the process for setting up the correct directories and files used by odin
// parameters: d (a byte array containing marshaled JSON)
// returns: string (the path to the newly created file)
func SetupEnvironment(d []byte) string {
    var job NewJob
    err := json.Unmarshal(d, &job)
    if err != nil {
        return ""
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
        ioutil.WriteFile(newFilePath, []byte(""), 0744)
        ioutil.WriteFile(logsPath + job.ID, []byte(""), 0766)
    }
    input, err := ioutil.ReadFile(originalFile)
    if err != nil {
        return ""
    }
    ioutil.WriteFile(logsPath, []byte(""), 0766)
    ioutil.WriteFile(newFilePath, input, 0744)
    return newFilePath
}
