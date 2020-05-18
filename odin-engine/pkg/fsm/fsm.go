package fsm

import (
    "fmt"
    "log"
    "net"
    "os"
    "regexp"
    "strings"
    "time"

    "github.com/hashicorp/raft"
)

// define constant values to be used by the finite state machine
const (
    retainSnapshotCount = 2
    raftTimeout = 10 * time.Second
)

// create a Store type which contains node and raft information
type Store struct {
    RaftDir     string
    RaftBind    string
    Raft        *raft.Raft
    ServerID    string
    NumericalID int
    PeersLength int
}

// this function is called on a store type and is used to boostrap the cluster
// parameters: enableSingle (a boolean used to allow single node cluster setup), localID (a string used to define the ID of a server)
// returns: error if an error occurs during this bootstrapping process, otherwise nil
func (s *Store) Open(enableSingle bool, localID string) error {
    config := raft.DefaultConfig()
    config.LocalID = raft.ServerID(localID)
    s.ServerID = localID
    log.Printf("Open, local ID [%v]", config.LocalID)
    addr, err := net.ResolveTCPAddr("tcp", s.RaftBind)
    if err != nil {
        return err
    }
    transport, err := raft.NewTCPTransport(s.RaftBind, addr, 3, 10*time.Second, os.Stderr)
    if err != nil {
        return err
    }
    snapshots, err := raft.NewFileSnapshotStore(s.RaftDir, retainSnapshotCount, os.Stderr)
    if err != nil {
        return fmt.Errorf("file snapshot store: %s", err)
    }
    var logStore raft.LogStore
    var stableStore raft.StableStore
    logStore = raft.NewInmemStore()
    stableStore = raft.NewInmemStore()
    ra, err := raft.NewRaft(config, (*fsm)(s), logStore, stableStore, snapshots, transport)
    if err != nil {
        return fmt.Errorf("new raft: %s", err)
    }
    s.Raft = ra
    if enableSingle {
        configuration := raft.Configuration{
            Servers: []raft.Server{
                {
                    ID:      config.LocalID,
                    Address: transport.LocalAddr(),
                },
            },
        }
        ra.BootstrapCluster(configuration)
    }
    return nil
}

// this function is called on a store type and is used to join a node to the cluster
// parameters: nodeID (a string used to represent a node by name), addr (a string used to denote the address to be used by the server)
// returns: error if an error occurs during this bootstrapping process, otherwise nil
func (s *Store) Join(nodeID, addr string) error {
    configFuture := s.Raft.GetConfiguration()
    if err := configFuture.Error(); err != nil {
        log.Printf("failed to get raft configuration: %v", err)
        return err
    }
    for _, srv := range configFuture.Configuration().Servers {
        if srv.ID == raft.ServerID(nodeID) || srv.Address == raft.ServerAddress(addr) {
            if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(nodeID) {
                log.Printf("node %s at %s already member of cluster, ignoring join request", nodeID, addr)
                return nil
            }

            future := s.Raft.RemoveServer(srv.ID, 0, 0)
            if err := future.Error(); err != nil {
                return fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
            }
        }
    }
    f := s.Raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
    if f.Error() != nil {
        return f.Error()
    }
    s.Raft.Apply([]byte("new voter"), raftTimeout)
    log.Printf("node %s at %s joined successfully", nodeID, addr)
    return nil
}

// this function is called on a store type and is used to remove a node from the cluster
// parameters: nodeID (a string used to represent a node by name)
// returns: error if an error occurs during this bootstrapping process, otherwise nil
func (s *Store) Leave(nodeID string) error {
	log.Printf("received leave request for remote node %s", nodeID)
        cf := s.Raft.GetConfiguration()
	if err := cf.Error(); err != nil {
		log.Printf("failed to get raft configuration")
		return err
	}
	for _, srv := range cf.Configuration().Servers {
	    if srv.ID == raft.ServerID(nodeID) {
	        f := s.Raft.RemoveServer(srv.ID, 0, 0)
		if err := f.Error(); err != nil {
		    log.Printf("failed to remove server %s", nodeID)
		    return err
		}
		log.Printf("node %s leaved successfully", nodeID)
		return nil
	    }
	}
	log.Printf("node %s not exists in raft group", nodeID)
	return nil
}

// this function is used to initialize a Store tstruct
// parameters: nil
// returns: *Store (a pointer to a newly initialized store)
func NewStore() *Store {
    return &Store{NumericalID: -1, PeersLength: -1}
}

// this function is used to get the numerical ID of a node from the list of peers
// parameters: ID (a string identifier of the node), peers (an array of current nodes in the cluster)
// returns: int (the numeric value of a node's indentifier), otherwise -1
func GetNumericalID(ID string, peers []string) int {
    for i, value := range peers {
        if value == ID {
            return i
        }
    }
    return -1
}


// this function is used to get the list of peers from a raft config
// parameters: rawConfig (latest_configuration section of the raft stats)
//returns: []string (a string array of all nodes in the cluster)
func PeersList(rawConfig string) []string {
    peers := []string{}
    re := regexp.MustCompile(`ID:[0-9A-z]*`)
    for _, peer := range re.FindAllString(rawConfig, -1) {
        peers = append(peers, strings.Replace(peer, "ID:", "", -1))
    }
    return peers
}

