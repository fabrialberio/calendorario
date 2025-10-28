package main

import (
	mw "calendorario/pkg/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	setupRoutes(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mw.WithLogging(mux),
	}

	log.Println("Server started on port 8080.")
	log.Fatal(server.ListenAndServe())
}

func setupRoutes(mux *http.ServeMux) {

}
