package loaddate

import (
	"calendorario/pkg/session"
	"calendorario/routes"
	"net/http"
	"time"
)

const keyDate = routes.KeyDate

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	date, err := time.Parse(time.DateOnly, r.FormValue(keyDate))
	if err != nil {
		return
	}

	session.SetSelectedDate(w, date)

	http.Redirect(w, r, routes.RouteAdminTimetableClass, http.StatusSeeOther)
}
