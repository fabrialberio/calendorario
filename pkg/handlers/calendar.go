package handlers

import (
	"calendorario/pkg/session"
	"calendorario/views"
	"net/http"
	"time"
)

func CalendarGet(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse(time.DateOnly, r.FormValue(views.KeyCalendarDate))
	if err != nil {
		date = time.Now()
	}

	s := session.FromContext(r.Context())
	term, _ := s.Database.GetTerm(r.Context(), int64(s.TermID))
	vacations, _ := s.Database.ListVacationsWithTermID(r.Context(), term.ID)

	views.Calendar(date.Year(), date.Month(), time.Now(), term, vacations).Render(r.Context(), w)
}
