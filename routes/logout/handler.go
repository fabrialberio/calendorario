package logout

import (
	"calendorario/pkg/session"
	"calendorario/routes"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	session.UnsetAuthenticatedUser(w)
	session.UnsetTermCookie(w)
	http.Redirect(w, r, routes.DestLogin, http.StatusSeeOther)
}
