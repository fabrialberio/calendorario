package handlers

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/routes"
	"errors"

	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}

	s := session.FromContext(r.Context())
	user, err := s.User()

	if err != nil {
		http.Redirect(w, r, routes.DestLogin, http.StatusSeeOther)
	} else {
		switch user.Role {
		case database.RoleAdministrator:
			http.Redirect(w, r, routes.DestAdmin, http.StatusSeeOther)
		case database.RoleSecretary:
			http.Redirect(w, r, routes.DestSecretary, http.StatusSeeOther)
		case database.RoleTeacher:
			http.Redirect(w, r, routes.DestTeacher, http.StatusSeeOther)
		}
	}
}

const errorQueryParam = "error"

func LoginGet(w http.ResponseWriter, r *http.Request) {
	s := session.FromContext(r.Context())
	_, err := s.User()

	if errors.Is(err, session.ErrCookieExpired) {
		routes.Login(false, true).Render(r.Context(), w)
	} else if err != nil {
		routes.Login(r.URL.Query().Has(errorQueryParam), false).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
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

func LogoutGet(w http.ResponseWriter, r *http.Request) {
	session.UnsetAuthenticatedUser(w)
	session.UnsetTermCookie(w)
	http.Redirect(w, r, routes.DestLogin, http.StatusSeeOther)
}
