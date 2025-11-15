package main

import (
	"calendorario/pkg/auth"
	"calendorario/pkg/database"
	"calendorario/pkg/handlers"
	"calendorario/pkg/middleware"
	"calendorario/views"
	"context"

	"github.com/a-h/templ"
	_ "github.com/lib/pq"

	"database/sql"
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
	addAdminUserIfNotExists(db)

	mux := http.NewServeMux()
	setupRoutes(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.WithContext(db, middleware.WithLogging(mux)),
	}

	log.Println("Server started on port 8080.")
	log.Fatal(server.ListenAndServe())
}

func createDatabase() *database.Queries {
	dsn := fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_CONTAINER_NAME"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	conn, err := sql.Open("postgres", dsn)
	for n_retries := 0; err != nil; n_retries++ {
		if n_retries == 5 {
			log.Fatalf("Error creating database after %v retries: %v", n_retries, err)
		}

		time.Sleep(5 * time.Second)
		log.Printf("Error creating database, retrying: %v", err)
		conn, err = sql.Open("postgres", dsn)
	}
	log.Println("Database created successfully.")

	return database.New(conn)
}

func addAdminUserIfNotExists(db *database.Queries) {
	_, err := db.GetUserWithUsername(context.Background(), "admin")
	if err == sql.ErrNoRows {
		hash, err := auth.HashPassword(os.Getenv("ADMIN_PASSWORD"))
		if err != nil {
			log.Fatalf("Error hashing admin password: %v", err)
		}

		_, err = db.CreateUser(context.Background(), database.CreateUserParams{
			Username:     "admin",
			Name:         "Administrator",
			Surname:      "",
			Role:         database.RoleAdministrator,
			PasswordHash: hash,
		})

		if err != nil {
			log.Fatalf("Error creating admin user: %v", err)
		}
		log.Println("Admin user created successfully.")
	} else if err != nil {
		log.Fatalf("Error getting admin user: %v", err)
	}
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET "+views.DestLogin, handlers.LoginGet)
	mux.HandleFunc("POST "+views.DestLogin, handlers.LoginPost)
	mux.HandleFunc("GET "+views.DestLogout, handlers.LogoutGet)

	mux.Handle("GET "+views.DestCalendar, templ.Handler(views.Calendar(2025, time.January)))

	mux.Handle("GET "+views.DestTimetableClass, templ.Handler(views.TimetableClass()))

	mux.Handle("GET /public/", http.FileServerFS(publicFS))
}
