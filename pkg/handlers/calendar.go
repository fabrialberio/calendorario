package handlers

import (
	"calendorario/views"
	"net/http"
	"time"
)

func CalendarGet(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse(time.DateOnly, r.FormValue(views.KeyCalendarDate))
	if err != nil {
		date = time.Now()
	}

	views.Calendar(date.Year(), date.Month()).Render(r.Context(), w)
}
