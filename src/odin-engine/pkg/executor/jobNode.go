package executor

import (
    "fmt"
    "io"
    "os"
    "os/exec"
    "os/user"
    "syscall"

    "github.com/sirupsen/logrus"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/resources"
)

func getHome() string {                                                                 usr, _ := user.Current()
    return usr.HomeDir
}

/// this function is called on a job node type and is used to log information from an executed job
// parameters: ch (channel used to return data), data (a byte array containing the data from execution), error (any exit status from the execution), store (a store of node information)
// returns: nil
func (job JobNode) logger(ch chan<- Data, data []byte, err error, store fsm.Store) {
    go func() {
        var logFile, _ = os.OpenFile("/etc/odin/logs/" + job.ID, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
        logrus.SetOutput(io.MultiWriter(logFile, os.Stdout))
        logrus.SetFormatter(&logrus.TextFormatter{})
        if job.ID != "" {
            logrus.WithFields(logrus.Fields{
                "node": store.ServerID,
                "id": job.ID,
                "uid": job.UID,
                "gid": job.GID,
                "language": job.Lang,
                "file": job.File,
            }).Info("executed")
        }
        if err != nil {
            logrus.WithFields(logrus.Fields{
                "node": store.ServerID,
                "id": job.ID,
                "uid": job.UID,
                "gid": job.GID,
                "language": job.Lang,
                "file": job.File,
                "error": err,
            }).Warn("failed")
        }
        ch <- Data{
            error:  err,
            output: data,
        }
    }()
}

// this function is called on a job node type and is used to run a job like a shell would run a command
// parameters: ch (channel used to return data), store (a store of node information)
// returns: nil
func (job JobNode) runCommand(ch chan<- Data, store fsm.Store) {
    URI := resources.UnmarsharlYaml(resources.ReadFileBytes(getHome() + "/odin-config.yml")).Mongo.Address
    os.Setenv("ODIN_EXEC_ENV", "True")
    os.Setenv("ODIN_MONGODB", URI)
    var cmd *exec.Cmd
    if job.Lang == "go" {
        cmd = exec.Command(job.File)
    } else {
        cmd = exec.Command(job.Lang, job.File)
    }
    cmd.SysProcAttr = &syscall.SysProcAttr{}
    cmd.SysProcAttr.Credential = &syscall.Credential{Uid: job.UID, Gid: job.GID}
    data, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println(cmd.Stderr)
    }
    go job.logger(ch, data, err, store)
}
