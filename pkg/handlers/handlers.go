package handlers

import (
	"fmt"
	"net/http"

	"github.com/Random7-JF/gowebapp/pkg/renderer"
)

type State struct {
	Info string
}

func (s *State) Index(w http.ResponseWriter, r *http.Request) {
	info := renderer.TemplateInfo{}
	info.Render(w)

}

func (s *State) Name(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Write([]byte(fmt.Sprintf("Name: %s", name)))
}
