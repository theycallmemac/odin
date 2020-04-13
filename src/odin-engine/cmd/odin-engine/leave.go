package main

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/fsm"
)

type leaveResource struct {
    addr string
    store fsm.Store
}

func (rs leaveResource) Routes(s *Service) chi.Router {
    // establish new chi router
    r := chi.NewRouter()
    rs.addr = s.addr
    rs.store = s.store
    // define routes under the leave endpoint
    r.Post("/", rs.Leave)
    return r
}

// this function is used to create a new job
func (rs leaveResource) Leave(w http.ResponseWriter, r *http.Request) {
    m := map[string]string{}
    if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
	w.WriteHeader(http.StatusBadRequest)
	return
    }
    nodeID, _ := m["id"]
    rs.store.Leave(nodeID)
}

