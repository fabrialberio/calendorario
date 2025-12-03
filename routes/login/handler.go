package login

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/routes"
	"errors"

	"net/http"
)

const (
	KeyUsername     = "username"
	KeyPassword     = "password"
	errorQueryParam = "error"
)

type Handler struct {
	Database *database.Queries
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPost:
		h.Post(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	_, err := session.AuthenticatedUser(r)

	if errors.Is(err, session.ErrCookieExpired) {
		View(false, true).Render(r.Context(), w)
	} else if err != nil {
		View(r.URL.Query().Has(errorQueryParam), false).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(string(KeyUsername))
	password := r.FormValue(string(KeyPassword))

	var destLoginError = routes.RouteLogin + "?" + errorQueryParam

	user, err := h.Database.GetUserWithUsername(r.Context(), username)
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
