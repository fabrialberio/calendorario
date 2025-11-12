package handlers

import (
	"calendorario/pkg/auth"
	"calendorario/pkg/requestcontext"
	"calendorario/views"
	"errors"

	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	rc := requestcontext.FromRequest(r)
	_, err := rc.AuthenticatedUser()

	if errors.Is(err, auth.ErrCookieExpired) {
		views.Login(false, true).Render(r.Context(), w)
	} else if err != nil {
		views.Login(r.URL.Query().Has("error"), false).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, views.DestCalendar, http.StatusSeeOther)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(string(views.KeyUsername))
	password := r.FormValue(string(views.KeyPassword))

	var destLoginError = views.DestLogin + "?error"

	rc := requestcontext.FromRequest(r)
	user, err := rc.Database().GetUserWithUsername(r.Context(), username)
	if err != nil {
		http.Redirect(w, r, destLoginError, http.StatusSeeOther)
		return
	}

	err = auth.SetAuthenticatedUser(w, &user, []byte(password))
	if err != nil {
		http.Redirect(w, r, destLoginError, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, views.DestCalendar, http.StatusSeeOther)
}

func LogoutGet(w http.ResponseWriter, r *http.Request) {
	auth.UnsetAuthenticatedUser(w)
	http.Redirect(w, r, views.DestLogin, http.StatusSeeOther)
}
