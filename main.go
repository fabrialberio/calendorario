package main

import (
	"calendorario/pkg/database"
	"calendorario/pkg/middleware"
	"calendorario/pkg/session"
	"calendorario/routes"
	"calendorario/routes/admin"
	adminCalendar "calendorario/routes/admin/calendar"
	adminTimetableClass "calendorario/routes/admin/timetableclass"
	"calendorario/routes/class"
	"calendorario/routes/index"
	"calendorario/routes/loaddate"
	"calendorario/routes/loadterm"
	"calendorario/routes/login"
	"calendorario/routes/logout"
	"calendorario/routes/month"
	"calendorario/routes/term"
	"calendorario/routes/vacation"
	"calendorario/routes/week"
	"context"

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
	setupRoutes(mux, db)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.WithLogging(mux),
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

func setupRoutes(mux *http.ServeMux, db *database.Queries) {
	mux.Handle("GET /public/", http.FileServerFS(publicFS))
	mux.Handle(routes.RouteLogin, &login.Handler{Database: db})
	mux.Handle(routes.RouteLogout, &logout.Handler{})

	adminMux := http.NewServeMux()
	adminMux.Handle(routes.RouteAdmin, &admin.Handler{Database: db})
	adminMux.Handle(routes.RouteAdminCalendar, &adminCalendar.Handler{Database: db})
	adminMux.Handle(routes.RouteAdminTimetableClass, &adminTimetableClass.Handler{Database: db})

	mux.Handle(routes.RouteAdmin, middleware.WithAuthenticatedUserCheck(
		func(u *database.User) bool { return u.Role == database.RoleAdministrator },
		adminMux,
	))

	loggedInMux := http.NewServeMux()
	loggedInMux.Handle("/", &index.Handler{})
	loggedInMux.Handle(routes.RouteClass, &class.Handler{Database: db})
	loggedInMux.Handle(routes.RouteLoadDate, &loaddate.Handler{})
	loggedInMux.Handle(routes.RouteLoadTerm, &loadterm.Handler{})
	loggedInMux.Handle(routes.RouteMonth, &month.Handler{Database: db})
	loggedInMux.Handle(routes.RouteTerm, &term.Handler{Database: db})
	loggedInMux.Handle(routes.RouteVacation, &vacation.Handler{Database: db})
	loggedInMux.Handle(routes.RouteWeek, &week.Handler{Database: db})

	mux.Handle("/", middleware.WithAuthenticatedUserCheck(
		func(u *database.User) bool { return true },
		loggedInMux,
	))
}
