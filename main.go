package main

import (
	"calendorario/pkg/database"
	"calendorario/pkg/handlers"
	"calendorario/pkg/middleware"
	"calendorario/pkg/session"
	"calendorario/routes"
	"calendorario/routes/index"
	"calendorario/routes/login"
	"calendorario/routes/logout"
	"calendorario/routes/term"
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
		Handler: middleware.WithSession(db, middleware.WithLogging(mux)),
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
	for retries := 0; err != nil; retries++ {
		if retries == 5 {
			log.Fatalf("Error creating database after %v retries: %v", retries, err)
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
		hash, err := session.HashPassword(os.Getenv("ADMIN_PASSWORD"))
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
	mux.Handle("/", &index.Handler{})
	mux.Handle(routes.DestLogin, &login.Handler{})
	mux.Handle(routes.DestLogout, &logout.Handler{})
	mux.Handle("GET /public/", http.FileServerFS(publicFS))

	mux.HandleFunc("GET "+routes.DestMonth, handlers.CalendarGet)

	adminMux := http.NewServeMux()
	adminMux.Handle("GET "+routes.DestAdmin, templ.Handler(routes.TermsPage()))

	adminMux.Handle(routes.DestAdminTerm+"/{id}", &term.Handler{})
	adminMux.HandleFunc("GET "+routes.DestAdminLoadTerm+"/{id}", handlers.AdminLoadTermGet)

	adminMux.HandleFunc("GET "+routes.DestAdminVacation+"/{id}", handlers.AdminVacationGet)
	adminMux.HandleFunc("POST "+routes.DestAdminVacation, handlers.AdminVacationPost)

	adminMux.Handle("GET "+routes.DestAdminCalendar, templ.Handler(routes.CalendarPage(time.Now().Year(), time.Now().Month())))
	adminMux.Handle("GET "+routes.DestAdminTimetableClass, templ.Handler(routes.TimetableClassPage(time.Now())))

	mux.Handle(routes.DestAdmin, middleware.WithUserCheck(
		func(u *database.User) bool { return u.Role == database.RoleAdministrator },
		adminMux,
	))
}
