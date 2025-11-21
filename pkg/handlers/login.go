package handlers

import (
	"calendorario/pkg/auth"
	"calendorario/pkg/database"
	"calendorario/pkg/requestcontext"
	"calendorario/views"
	"errors"

	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}

	rc := requestcontext.FromRequest(r)
	user, err := rc.User()

	if err != nil {
		http.Redirect(w, r, views.DestLogin, http.StatusSeeOther)
	} else {
		switch user.Role {
		case database.RoleAdministrator:
			http.Redirect(w, r, views.DestAdmin, http.StatusSeeOther)
		case database.RoleSecretary:
			http.Redirect(w, r, views.DestSecretary, http.StatusSeeOther)
		case database.RoleTeacher:
			http.Redirect(w, r, views.DestTeacher, http.StatusSeeOther)
		}
	}
}

const errorQueryParam = "error"

func LoginGet(w http.ResponseWriter, r *http.Request) {
	rc := requestcontext.FromRequest(r)
	_, err := rc.User()

	if errors.Is(err, auth.ErrCookieExpired) {
		views.Login(false, true).Render(r.Context(), w)
	} else if err != nil {
		views.Login(r.URL.Query().Has(errorQueryParam), false).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(string(views.KeyUsername))
	password := r.FormValue(string(views.KeyPassword))

	var destLoginError = views.DestLogin + "?" + errorQueryParam

	rc := requestcontext.FromRequest(r)
	user, err := rc.Database.GetUserWithUsername(r.Context(), username)
	if err != nil {
		http.Redirect(w, r, destLoginError, http.StatusSeeOther)
		return
	}

	err = auth.SetAuthenticatedUser(w, &user, []byte(password))
	if err != nil {
		http.Redirect(w, r, destLoginError, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutGet(w http.ResponseWriter, r *http.Request) {
	auth.UnsetAuthenticatedUser(w)
	http.Redirect(w, r, views.DestLogin, http.StatusSeeOther)
}
