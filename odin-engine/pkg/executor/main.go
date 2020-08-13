package executor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"strconv"
	"strings"

	"github.com/theycallmemac/odin/odin-engine/pkg/fsm"
	"github.com/theycallmemac/odin/odin-engine/pkg/repository"
	"github.com/theycallmemac/odin/odin-engine/pkg/resources"
)

// Queue is a type used to contain an array of JobNode's
type Queue []JobNode

// JobNode is a type used to define a job as a node in a Queue
type JobNode struct {
	ID       string
	UID      uint32
	GID      uint32
	Lang     string
	File     string
	Schedule []int
	Runs     int
	Links    string
}

// Data is a type used to define channels for execution
type Data struct {
	output []byte
	error  error
}

// makePutRequest is used to make a put request to a given url
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
	response.Body.Close()
	return string(bodyBytes)
}

// executeYaml is used to run a job like straight from the command line tool
// parameters: filename (a string containing the path to the local file to execute), done (a boolean channel), store (a store of node information)
// returns: nil
func executeYaml(repo repository.Repository, filename string, done chan bool, httpAddr string, store fsm.Store) {
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
		go job.runCommand(repo, singleChannel, httpAddr, store)
		res := <-singleChannel
		ReviewError(res.error, "bool")
		done <- true
		return
	}
	done <- false
	return
}

// executeLang is used to execute a file in /etc/odin/$id
// parameters: contentsJSON (byte array containing uid, gid, language and file information), store (a store of node information)
// returns: boolean (returns true if the file exists and is executed)
func executeLang(repo repository.Repository, contentsJSON []byte, done chan bool, httpAddr string, store fsm.Store) {
	var queue Queue
	json.Unmarshal(contentsJSON, &queue)
	go queue.BatchRun(repo, httpAddr, store)
	go queue.UpdateRuns(httpAddr)
	done <- true
	return
}

// Execute is used to decide which of the executeLang and exectureYaml functions to use
// parameters: contents (byte array containing uid, gid, language and file information), process (int used to decide the function to use in the code), httpAddr (an address string for the engine), store (a store of node information)
// returns: boolean (returns true if one of the functions executes successfully, false otherwise)
func Execute(repo repository.Repository, contents []byte, process int, httpAddr string, store fsm.Store) bool {
	done := make(chan bool)
	switch process {
	case 0:
		go executeLang(repo, contents, done, httpAddr, store)
	case 1:
		go executeYaml(repo, string(contents), done, httpAddr, store)
	}
	return <-done
}
