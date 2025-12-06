package timetableclass

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"net/http"
	"time"
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
	classes, _ := h.Database.ListClasses(r.Context())

	View(date, time.Now(), term, classes).Render(r.Context(), w)
}
