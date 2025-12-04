package calendar

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"net/http"
	"time"

	"github.com/a-h/templ"
)

type Handler struct {
	Database *database.Queries
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	date, err := session.SelectedDate(r)
	if err != nil {
		date = time.Now()
	}

	termID, _ := session.SelectedTermID(r)
	term, _ := h.Database.GetTerm(r.Context(), int64(termID))
	vacations, _ := h.Database.ListVacationsWithTermID(r.Context(), int64(termID))

	templ.Handler(View(date.Year(), date.Month(), time.Now(), term, vacations)).ServeHTTP(w, r)
}
