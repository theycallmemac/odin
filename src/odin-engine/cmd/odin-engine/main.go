package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "os/user"
    "os/signal"
    "time"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/resources"
)

const (
    DefaultRaftAddr = ":12000"
    retainSnapshotCount = 2
    raftTimeout = 10 * time.Second
)

var (
    raftAddr string
    joinAddr string
    nodeID   string
)

func init() {
    flag.StringVar(&raftAddr, "raddr", DefaultRaftAddr, "Set Raft bind address")
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

    s := newStore()
    s.RaftDir = raftDir
    s.RaftBind = raftAddr
    if err := s.Open(joinAddr == "", nodeID); err != nil {
        log.Fatalf("%v", err)
    }
    usr, _ := user.Current()
    config := resources.UnmarsharlYaml(resources.ReadFileBytes(usr.HomeDir + "/odin-config.yml"))

    Start(config.OdinVars.Master + ":" + config.OdinVars.Port)
    setOdinEnv(config.Mongo.Address)

    log.Println("started successfully ...")
    terminate := make(chan os.Signal, 1)
    signal.Notify(terminate, os.Interrupt)
    <-terminate
    log.Println("exiting ...")
}

