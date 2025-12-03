package middleware

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/routes"

	"log"
	"net/http"
)

func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	})
}

type UserCheckerFunc func(user *database.User) bool

func WithAuthenticatedUserCheck(checker UserCheckerFunc, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := session.AuthenticatedUser(r)
		if err != nil {
			http.Redirect(w, r, routes.RouteLogout, http.StatusSeeOther)
			return
		}

		if !checker(user) {
			http.Error(w, "Invalid role.", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
