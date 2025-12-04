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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.FormValue(keyID))
	if err != nil {
		http.Redirect(w, r, routes.RouteAdmin, http.StatusSeeOther)
	}

	session.SetSelectedTermID(w, id)

	http.Redirect(w, r, routes.RouteAdminCalendar, http.StatusSeeOther)
}
