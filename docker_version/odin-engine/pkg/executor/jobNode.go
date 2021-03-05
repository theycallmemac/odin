package executor

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"
	"time"

	"github.com/theycallmemac/odin/odin-engine/pkg/fsm"
	"github.com/theycallmemac/odin/odin-engine/pkg/jobs"
	"github.com/theycallmemac/odin/odin-engine/pkg/resources"
)

func getHome() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

/// logger is called on a JobNode type and is used to log information from an executed job
// parameters: ch (channel used to return data), data (a byte array containing the data from execution), error (any exit status from the execution), store (a store of node information)
// returns: nil
func (job JobNode) logger(ch chan<- Data, data []byte, err error, store fsm.Store) {
	go func() {
		var logStatus string
		if job.ID != "" && err == nil {
			logStatus = "[EXECUTED]"
		}
		if err != nil {
			logStatus = "[FAILED]"
		}
		t := time.Now()
		f, _ := os.OpenFile("/etc/odin/logs/" + job.ID, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
                _, err := f.WriteString(t.Format("Jan 2 15:04:05") + " " + logStatus + ": " + job.ID + " on " + store.ServerID + "\n")
                if err != nil {
                    fmt.Println(err)
                    f.Close()
                }
		ch <- Data{
			error:  err,
			output: data,
		}
	}()
}

// runCommand is called on a JobNode type and is used to run a job like a shell would run a command
// parameters: ch (channel used to return data), store (a store of node information)
// returns: nil
func (job JobNode) runCommand(ch chan<- Data, httpAddr string, store fsm.Store) {
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
	if job.Links != "" {
		links := strings.Split(job.Links, ",")
		jobs.RunLinks(links, job.UID, httpAddr, store)
	}
	job.logger(ch, data, err, store)
}
