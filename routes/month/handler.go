package month

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/routes"
	"net/http"
	"time"
)

const EventColorVacation = "#d8dbd1"

const keyDate = routes.KeyDate

type Handler struct {
	Database *database.Queries
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	date, _ := time.Parse(time.DateOnly, r.FormValue(keyDate))

	session.SetSelectedDate(w, date)

	termID, _ := session.SelectedTermID(r)
	term, _ := h.Database.GetTerm(r.Context(), int64(termID))
	vacations, _ := h.Database.ListVacationsWithTermID(r.Context(), term.ID)

	innerView(date.Year(), date.Month(), time.Now(), term, vacations).Render(r.Context(), w)
}
