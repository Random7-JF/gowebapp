package main

import (
	"fmt"
	"net/http"

	"github.com/Random7-JF/gowebapp/pkg/middleware"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Index"))
	})
	router.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("name")
		w.Write([]byte(fmt.Sprintf("Hello, %s", id)))
	})

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Auth,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	server.ListenAndServe()
}
