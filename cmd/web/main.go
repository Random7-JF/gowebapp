package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hit: /")
		w.Write([]byte("Hello"))
	})
	router.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit: /{name}")
		id := r.PathValue("name")
		w.Write([]byte(fmt.Sprintf("Hello, %s", id)))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
