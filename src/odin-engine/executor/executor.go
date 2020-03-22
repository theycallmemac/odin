package executor

import (
    "encoding/json"
    "io"
    "os"
    "os/exec"
    "os/user"
    "strings"
    "strconv"
    "syscall"

    "github.com/sirupsen/logrus"
)

type JobNode struct {
    ID string
    UID string
    GID string
    Lang string
    File string
    Schedule []int
}

type Data struct {
    output []byte
    error  error
}

// this function is used to run a job like a shell would run a command
// parameters: ch (channel used to return data), uid (uint32 used to execute as a particular user), gid (uint32 used to execute as a particular group), language (string value of execution language), file (string containing the name of the base file), id (a string containing the jobs id)
// returns: nil
func runCommand(ch chan<- Data, uid uint32, gid uint32, language string, file string, id string) {
    cmd := exec.Command(language, file)
    cmd.SysProcAttr = &syscall.SysProcAttr{}
    cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
    data, err := cmd.CombinedOutput()
    go func() {
        var logFile, _ = os.OpenFile("/etc/odin/logs/" + id, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
        logrus.SetOutput(io.MultiWriter(logFile, os.Stdout))
        logrus.SetFormatter(&logrus.TextFormatter{})
        if id != "" {
            logrus.WithFields(logrus.Fields{
                "id": id,
                "uid": uid,
                "gid": gid,
                "language": language,
                "file": file,
            }).Info("executed")
        }
        if err != nil {
            logrus.WithFields(logrus.Fields{
                "id": id,
                "uid": uid,
                "gid": gid,
                "language": language,
                "file": file,
                "error": err,
            }).Warn("failed")
        }
        ch <- Data{
            error:  err,
            output: data,
        }
    }()
}

// this function is used to run a job like straight from the command line tool
// parameters: filename (a string containing the path to the local file to execute)
// returns: boolean (returns true if the file exists and is executed, false otherwise)
func executeYaml(filename string, done chan bool) {
    if exists(filename) {
        singleChannel := make(chan Data)
        path := strings.Split(filename, "/")
        basePath := strings.Join(path[:len(path)-1], "/")
        language, file := getYaml(filename)
        destFile := basePath + "/" + file
        uid, _ := strconv.ParseUint("0", 10, 32)
        group, _ := user.LookupGroup("odin")
        gid, _ := strconv.Atoi(group.Gid)
        go runCommand(singleChannel, uint32(uid), uint32(gid), language, destFile, "")
        res := <-singleChannel
        ReviewError(res.error, "bool")
        done<- true
        return
    } else {
        done<- false
        return
    }
}

// this function is used to execute a file in /etc/odin/$id
// parameters: contentsJSON (byte array containing uid, gid, language and file information)
// returns: boolean (returns true if the file exists and is executed)
func executeLang(contentsJSON []byte, done chan bool) {
    var jobs []JobNode
    json.Unmarshal(contentsJSON, &jobs)
    for _, job := range jobs {
        go func(job JobNode) {
            channel := make(chan Data)
            uid, _ := strconv.ParseUint(job.UID, 10, 32)
            gid, _ := strconv.ParseUint(job.GID, 10, 32)
            go runCommand(channel, uint32(uid), uint32(gid), job.Lang, job.File, job.ID)
            res := <-channel
            ReviewError(res.error, "bool")
        }(job)
    }
    done<- true
    return
}

// this function is used to decide which of the executeLang and exectureYaml functions to use
// parameters: contents (byte array containing uid, gid, language and file information), process (int used to decide the function to use in the code)
// returns: boolean (returns true if one of the functions executes sucessfully, false otherwise)
func Execute(contents []byte, process int) bool {
    done := make(chan bool)
    switch process {
        case 0:
            go executeLang(contents, done)
        case 1:
            go executeYaml(string(contents), done)
    }
    return <-done
}

