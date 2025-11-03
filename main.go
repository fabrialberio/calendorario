package main

import (
	"calendorario/pkg/database"
	"calendorario/pkg/handlers"
	"calendorario/pkg/middleware"
	"embed"

	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed public
var publicFS embed.FS

func main() {
	db := createDatabase()
	defer db.Close()

	mux := http.NewServeMux()
	setupRoutes(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.WithLogging(mux),
	}

	log.Println("Server started on port 8080.")
	log.Fatal(server.ListenAndServe())
}

func createDatabase() *database.DB {
	dsn := fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_CONTAINER_NAME"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := database.New(dsn)
	for n_retries := 0; err != nil; n_retries++ {
		if n_retries == 5 {
			log.Fatalf("Error creating database after %v retries: %v", n_retries, err)
		}

		time.Sleep(5 * time.Second)
		log.Printf("Error creating database, retrying: %v", err)
		db, err = database.New(dsn)
	}
	log.Println("Database created successfully.")

	return db
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", handlers.IndexGet)
	mux.HandleFunc("GET /login", handlers.LoginGet)
	mux.HandleFunc("POST /login", handlers.LoginPost)

	mux.Handle("GET /public/", http.FileServerFS(publicFS))
}
