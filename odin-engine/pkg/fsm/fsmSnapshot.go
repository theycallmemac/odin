package fsm

import (
	"io"
	"log"

	"github.com/hashicorp/raft"
)

// define a type store by the name fsm (finite state machine)
type fsm Store

// Apply is called on a fsm and is used to apply the information from the latest Join to the store
// parameters: l (a log consists of a byte array and a raftTimeout)
//returns: interface{} (this will always be nil)
func (f *fsm) Apply(l *raft.Log) interface{} {
	stats := f.Raft.Stats()
	config := stats["latest_configuration"]
	peers := PeersList(config)
	f.PeersLength = len(peers)
	ID := f.ServerID
	f.NumericalID = GetNumericalID(ID, peers)
	log.Printf("apply ID [%s] [%d]", ID, f.NumericalID)
	return nil
}

// Snapshot is called on an fsm and will be used to perform snapshots
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	log.Printf("snapshot")
	return &fsmSnapshot{}, nil
}

// Restore is called on an fsm and will be used to perform restorations
func (f *fsm) Restore(rc io.ReadCloser) error {
	log.Printf("restore [%v]", rc)
	return nil
}

// define an empty struct by the name fsmSnapshot
type fsmSnapshot struct{}

// Release is called on an fsmSnapshot and is used to release and flush any resources
func (f *fsmSnapshot) Release() {}

// Persist is called on an fsmSnapshot and is used to write to file
// parameters: (a raft.SnapshotSink used to write snapshots)
// returns: error
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
