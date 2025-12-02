package loadterm

import (
	"calendorario/pkg/session"
	"calendorario/routes"
	"net/http"
	"strconv"
)

const keyID = routes.KeyID

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue(keyID))
	if err != nil {
		http.Redirect(w, r, routes.RouteAdmin, http.StatusSeeOther)
	}

	session.SetTermCookie(w, id)

	http.Redirect(w, r, routes.RouteAdminCalendar, http.StatusSeeOther)
}
