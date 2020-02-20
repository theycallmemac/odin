package executor

import (
    "io"
    "os"
    "os/exec"
    "os/user"
    "strings"
    "strconv"
    "syscall"

    "github.com/sirupsen/logrus"
)

type Data struct {
    output []byte
    error  error
}

func runCommand(ch chan<- Data, uid uint32, gid uint32, language string, file string, id string) {
    cmd := exec.Command(language, file)
    cmd.SysProcAttr = &syscall.SysProcAttr{}
    cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
    data, err := cmd.CombinedOutput()
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
}

func executeYaml(filename string) bool {
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
        return true
    } else {
        return false
    }
}

func executeLang(contents string) bool {
    contentList := strings.Split(contents, " ")
    uid, gid, language, filename, id := contentList[0], contentList[1], contentList[2], contentList[3], contentList[4]
    if exists(filename) {
        channel := make(chan Data)
        uid, _ := strconv.ParseUint(uid, 10, 32)
        gid, _ := strconv.ParseUint(gid, 10, 32)
        go runCommand(channel, uint32(uid), uint32(gid), language, filename, id)
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

