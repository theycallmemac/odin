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

const (
    DefaultRaftAddr = ":12000"
    DefaultHttpAddr = ":3939"
    retainSnapshotCount = 2
    raftTimeout = 10 * time.Second
)

var (
    httpAddr string
    raftAddr string
    joinAddr string
    nodeID   string
)

func init() {
    flag.StringVar(&httpAddr, "http", DefaultHttpAddr, "Set Http bind address")
    flag.StringVar(&raftAddr, "raft", DefaultRaftAddr, "Set Raft bind address")
    flag.StringVar(&joinAddr, "join", "", "Set join address, if any")
    flag.StringVar(&nodeID, "id", "", "Node ID")
}

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
    fmt.Println(leave(nodeID))
    log.Println("exiting ...")
}

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

func leave(nodeID string) error {
	b, err := json.Marshal(map[string]string{"id": nodeID})
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://localhost%s/leave",":3939"), "application-type/json", bytes.NewReader(b))
        data, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

        return nil
}
