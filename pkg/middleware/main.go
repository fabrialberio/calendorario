package middleware

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/views"

	"log"
	"net/http"
)

func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	})
}

func WithContext(database *database.Queries, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := session.GetAuthenticatedUser(r)

		ctx := r.Context()
		s := session.NewContext(ctx, database, user, err)
		r = r.WithContext(s)

		next.ServeHTTP(w, r)
	})
}

type UserCheckerFunc func(user *database.User) bool

func WithUserCheck(checker UserCheckerFunc, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := session.FromContext(r.Context())
		user, err := s.User()
		if err != nil {
			http.Redirect(w, r, views.DestLogout, http.StatusSeeOther)
			return
		}

		if !checker(user) {
			http.Error(w, "Status 403 forbidden.", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
