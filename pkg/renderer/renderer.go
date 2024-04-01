package renderer

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateInfo struct {
	Data map[string]interface{}
}

func (ti *TemplateInfo) Render(w http.ResponseWriter) {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		log.Println("Error parsing", err)
	}
	w.Header().Set("Content-Type", "text/html")
	err = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
}
