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

func WithSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := session.GetAuthenticatedUser(r)
		selectedTermID, _ := session.GetTermCookie(r)
		selectedDate, _ := session.GetDateCookie(r)

		ctx := r.Context()
		s := session.NewContext(ctx, user, err, selectedTermID, &selectedDate)
		r = r.WithContext(s)

		next.ServeHTTP(w, r)
	})
}

type UserCheckerFunc func(user *database.User) bool

func WithAuthenticatedUserCheck(checker UserCheckerFunc, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := session.FromContext(r.Context())
		user, err := s.User()
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
