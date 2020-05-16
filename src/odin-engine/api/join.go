package api

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
)


// create resouce type to be used by the router
// consists of a base http address and a store in the finite state machine
type joinResource struct {
    addr string
    store fsm.Store
}

func (rs joinResource) Routes(s *Service) chi.Router {
    // establish new chi router
    r := chi.NewRouter()
    rs.addr = s.addr
    rs.store = s.store
    // define routes under the join endpoint
    r.Post("/", rs.Join)
    return r
}

// this function is used to join a node to an exisiting cluster
func (rs joinResource) Join(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{}
        if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
	if len(m) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	remoteAddr, ok := m["addr"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nodeID, ok := m["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := rs.store.Join(nodeID, remoteAddr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

