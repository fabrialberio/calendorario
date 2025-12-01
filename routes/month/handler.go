package month

import (
	"calendorario/pkg/session"
	"calendorario/routes"
	"net/http"
	"time"
)

const EventColorVacation = "#d8dbd1"

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
	date, err := time.Parse(time.DateOnly, r.FormValue(routes.KeyCalendarDate))
	if err != nil {
		date = time.Now()
	}

	s := session.FromContext(r.Context())
	term, _ := s.Database.GetTerm(r.Context(), int64(s.TermID))
	vacations, _ := s.Database.ListVacationsWithTermID(r.Context(), term.ID)

	innerView(date.Year(), date.Month(), time.Now(), term, vacations).Render(r.Context(), w)
}
