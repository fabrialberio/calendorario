package login

import (
	"calendorario/pkg/session"
	"calendorario/routes"
	"errors"

	"net/http"
)

type Handler struct{}

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

const errorQueryParam = "error"

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	s := session.FromContext(r.Context())
	_, err := s.User()

	if errors.Is(err, session.ErrCookieExpired) {
		View(false, true).Render(r.Context(), w)
	} else if err != nil {
		View(r.URL.Query().Has(errorQueryParam), false).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
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
