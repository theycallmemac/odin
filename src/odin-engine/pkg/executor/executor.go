package executor

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "os/exec"
    "os/user"
    "net/http"
    "strings"
    "strconv"
    "syscall"

    "github.com/sirupsen/logrus"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/resources"
)

type Queue []JobNode

type JobNode struct {
    ID string
    UID uint32
    GID uint32
    Lang string
    File string
    Schedule []int
    Runs int
}

type Data struct {
    output []byte
    error  error
}

// this function is used to make a put request to a given url
// parameters: link (a string of the link to make a request to), data (a buffer to pass to the post request)
// returns: string (the result of a PUT to the provided link with the given data)
func makePutRequest(link string, data *bytes.Buffer) string {
    client := &http.Client{}
    req, _ := http.NewRequest("PUT", link, data)
    response, clientErr := client.Do(req)
    if clientErr != nil {
        fmt.Println(clientErr)
    }
    bodyBytes, _ := ioutil.ReadAll(response.Body)
    return string(bodyBytes)
}

// this function is called on a queue type and is used to run the batch loop to run all executions
// parameters:  store (a store of node information)
// returns: nil
func (queue Queue) batchRun(store fsm.Store) {
    for _, job := range queue {
        go func(job JobNode) {
            channel := make(chan Data)
            go job.runCommand(channel, store)
        }(job)
    }
}

// this function is called on a queue type and is used to update the run number for each job
// parameters: httpAddr (an address string for the node)
// returns: nil
func (queue Queue) updateRuns(httpAddr string) {
    for _, job := range queue {
        go func(job JobNode) {
            inc := job.Runs + 1
            go makePutRequest("http://localhost" + httpAddr + "/jobs/info/runs", bytes.NewBuffer([]byte(job.ID + " " + strconv.Itoa(inc) + " " + fmt.Sprint(job.UID))))
        }(job)
    }
}

// this function is called on a job node type and is used to log information from an executed job
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
    cmd := exec.Command(job.Lang, job.File)
    cmd.SysProcAttr = &syscall.SysProcAttr{}
    cmd.SysProcAttr.Credential = &syscall.Credential{Uid: job.UID, Gid: job.GID}
    data, err := cmd.CombinedOutput()

    go job.logger(ch, data, err, store)
}

// this function is used to run a job like straight from the command line tool
// parameters: filename (a string containing the path to the local file to execute), done (a boolean channel), store (a store of node information)
// returns: nil
func executeYaml(filename string, done chan bool, store fsm.Store) {
    if exists(filename) {
        var job JobNode
        singleChannel := make(chan Data)
        path := strings.Split(filename, "/")
        basePath := strings.Join(path[:len(path)-1], "/")
        job.Lang, job.File = resources.ExecutorYaml(filename)
        job.File = basePath + "/" + job.File
        uid, _ := strconv.ParseUint("0", 10, 32)
        group, _ := user.LookupGroup("odin")
        gid, _ := strconv.Atoi(group.Gid)
        job.UID = uint32(uid)
        job.GID = uint32(gid)
        go job.runCommand(singleChannel, store)
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
// parameters: contentsJSON (byte array containing uid, gid, language and file information), store (a store of node information)
// returns: boolean (returns true if the file exists and is executed)
func executeLang(contentsJSON []byte, done chan bool, httpAddr string, store fsm.Store) {
    var queue Queue
    json.Unmarshal(contentsJSON, &queue)
    go queue.batchRun(store)
    go queue.updateRuns(httpAddr)
    done<- true
    return
}

// this function is used to decide which of the executeLang and exectureYaml functions to use
// parameters: contents (byte array containing uid, gid, language and file information), process (int used to decide the function to use in the code), httpAddr (an address string for the engine), store (a store of node information)
// returns: boolean (returns true if one of the functions executes sucessfully, false otherwise)
func Execute(contents []byte, process int, httpAddr string, store fsm.Store) bool {
    done := make(chan bool)
    switch process {
        case 0:
            go executeLang(contents, done, httpAddr, store)
        case 1:
            go executeYaml(string(contents), done, store)
    }
    return <-done
}

