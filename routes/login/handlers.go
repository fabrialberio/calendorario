package login

import (
	"calendorario/pkg/session"
	"calendorario/routes"
	"errors"

	"net/http"
)

const errorQueryParam = "error"

func Get(w http.ResponseWriter, r *http.Request) {
	s := session.FromContext(r.Context())
	_, err := s.User()

	if errors.Is(err, session.ErrCookieExpired) {
		Login(false, true).Render(r.Context(), w)
	} else if err != nil {
		Login(r.URL.Query().Has(errorQueryParam), false).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(string(routes.KeyUsername))
	password := r.FormValue(string(routes.KeyPassword))

	var destLoginError = routes.DestLogin + "?" + errorQueryParam

	s := session.FromContext(r.Context())
	user, err := s.Database.GetUserWithUsername(r.Context(), username)
	if err != nil {
		http.Redirect(w, r, destLoginError, http.StatusSeeOther)
		return
	}

	err = session.SetAuthenticatedUser(w, &user, []byte(password))
	if err != nil {
		http.Redirect(w, r, destLoginError, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
