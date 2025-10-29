package main

import (
	"calendorario/pkg/middleware"
	"calendorario/views"

	"github.com/a-h/templ"

	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	setupRoutes(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.WithLogging(mux),
	}

	log.Println("Server started on port 8080.")
	log.Fatal(server.ListenAndServe())
}

func setupRoutes(mux *http.ServeMux) {
	mux.Handle("/", templ.Handler(views.Hello("person")))
}
