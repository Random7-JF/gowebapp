package handlers

import (
	"fmt"
	"net/http"
)

type State struct {
	Info string
}

func (s *State) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func (s *State) Name(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Write([]byte(fmt.Sprintf("Name: %s", name)))
}
