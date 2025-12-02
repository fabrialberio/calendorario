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

	today := time.Now()

	s := session.FromContext(r.Context())
	term, _ := h.Database.GetTerm(r.Context(), int64(s.SelectedTermID))
	vacations, _ := h.Database.ListVacationsWithTermID(r.Context(), int64(s.SelectedTermID))

	templ.Handler(View(today.Year(), today.Month(), today, term, vacations)).ServeHTTP(w, r)
}
