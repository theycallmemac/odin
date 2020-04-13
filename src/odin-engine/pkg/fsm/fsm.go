package fsm

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "regexp"
    "strings"
    "time"

    "github.com/hashicorp/raft"
)


const (
    retainSnapshotCount = 2
    raftTimeout = 10 * time.Second
)

type Store struct {
    RaftDir     string
    RaftBind    string
    Raft        *raft.Raft
    ServerID    string
    NumericalID int
    PeersLength int
}

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

func NewStore() *Store {
    return &Store{NumericalID: -1, PeersLength: -1}
}

type fsm Store

func (f *fsm) Apply(l *raft.Log) interface{}  {
    stats := f.Raft.Stats()
    config := stats["latest_configuration"]
    peers := PeersList(config)
    f.PeersLength = len(peers)
    ID := f.ServerID
    f.NumericalID = GetNumericalID(ID, peers)
    log.Printf("apply ID [%s] [%d]", ID, f.NumericalID)
    return nil
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
    log.Printf("snapshot")
    return &fsmSnapshot{}, nil
}

func (f *fsm) Restore(rc io.ReadCloser) error {
    log.Printf("restore [%v]", rc)
    return nil
}

type fsmSnapshot struct {}

func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
    err := func() error {
        b := []byte("hello from persist")
        if _, err := sink.Write(b); err != nil {
            return err
        }
        return sink.Close()
    }()
    if err != nil {
        sink.Cancel()
    }
    return err
}

func (f *fsmSnapshot) Release() {}

func GetNumericalID(ID string, peers []string) int {
    for i, value := range peers {
        if value == ID {
                return i
        }
    }
    return -1
}

func PeersList(rawConfig string) []string {
    peers := []string{}
    re := regexp.MustCompile(`ID:[0-9A-z]*`)
    for _, peer := range re.FindAllString(rawConfig, -1) {
        peers = append(peers, strings.Replace(peer, "ID:", "", -1))
    }
    return peers
}

