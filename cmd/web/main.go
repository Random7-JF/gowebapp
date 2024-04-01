package main

import (
	"net/http"

	"github.com/Random7-JF/gowebapp/pkg/handlers"
	"github.com/Random7-JF/gowebapp/pkg/middleware"
)

func main() {
	router := http.NewServeMux()
	state := handlers.State{
		Info: "Test",
	}
	router.HandleFunc("/", state.Index)
	router.HandleFunc("/hello/{name}", state.Name)

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
