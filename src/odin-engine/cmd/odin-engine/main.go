package main

import (
    "bytes"
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/user"
    "os/signal"
    "time"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/resources"
)


// define constsant default values to be used by the engine
const (
    DefaultRaftAddr = ":12000"
    DefaultHttpAddr = ":3939"
    retainSnapshotCount = 2
    raftTimeout = 10 * time.Second
)

// define variables to be used in setting up the engine
var (
    httpAddr string
    raftAddr string
    joinAddr string
    nodeID   string
)

// this function is used to define options to be used when running the engine
// parameters: nil
// returns: nil
func init() {
    flag.StringVar(&httpAddr, "http", DefaultHttpAddr, "Set Http bind address")
    flag.StringVar(&raftAddr, "raft", DefaultRaftAddr, "Set Raft bind address")
    flag.StringVar(&joinAddr, "join", "", "Set join address, if any")
    flag.StringVar(&nodeID, "id", "", "Node ID")
}


// this function is used to setup the odin-engine as a single node cluster, as well as allowing for further nodes to joining it. This is all done with the use of flags when initially running the engine.
// parameters: nil
// returns: nil
func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
        os.Exit(1)
    }
    raftDir := flag.Arg(0)
    if raftDir == "" {
        fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
        os.Exit(1)
    }
    os.MkdirAll(raftDir, 0700)

    s := fsm.NewStore()
    s.RaftDir = raftDir
    s.RaftBind = raftAddr
    if err := s.Open(joinAddr == "", nodeID); err != nil {
        log.Fatalf("%v", err)
    }

    if httpAddr == "" {
        usr, _ := user.Current()
        config := resources.UnmarsharlYaml(resources.ReadFileBytes(usr.HomeDir + "/odin-config.yml"))
        setOdinEnv(config.Mongo.Address)
        httpAddr = config.OdinVars.Master + ":" + config.OdinVars.Port
    }
    service := newService(httpAddr, *s)
    go service.Start()
    if joinAddr != "" {
	if err := join(joinAddr, raftAddr, nodeID); err != nil {
	    log.Fatalf("failed to join node at %s: %s", joinAddr, err.Error())
	}
    }
    log.Println("started successfully ...")
    terminate := make(chan os.Signal, 1)
    signal.Notify(terminate, os.Interrupt)
    <-terminate
    leave(nodeID)
    log.Println("exiting ...")
}


// this function is used to make a POST request to join a new node to the cluster
// parameters: joinAddr (an address string used to join the cluster), raftAddr (an address string used to be identified by Raft), nodeID (a string used to signify the node by ID)
// returns: error (if an error occurs during adding the node to the cluster), otherwise nil
func join(joinAddr, raftAddr, nodeID string) error {
	b, err := json.Marshal(map[string]string{"addr": raftAddr, "id": nodeID})
	if err != nil {
	    return err
	}
	resp, err := http.Post(fmt.Sprintf("http://localhost%s/join", joinAddr), "application-type/json", bytes.NewReader(b))
	if err != nil {
	    return err
	}
	defer resp.Body.Close()

	return nil
}

// this function is used to make a POST request to remove a node from a cluster
// parameters: nodeID (a string used to signify the node by ID)
// returns: error (if an error occurs during adding the node to the cluster), otherwise nil
func leave(nodeID string) error {
	b, err := json.Marshal(map[string]string{"id": nodeID})
	if err != nil {
	    return err
	}
	resp, err := http.Post(fmt.Sprintf("http://localhost%s/leave",":3939"), "application-type/json", bytes.NewReader(b))
        _, _ = ioutil.ReadAll(resp.Body)
	if err != nil {
	    return err
	}
	defer resp.Body.Close()

        return nil
}
