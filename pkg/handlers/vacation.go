package handlers

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/views"
	"net/http"
	"strconv"
	"time"
)

func AdminVacationGet(w http.ResponseWriter, r *http.Request) {
	s := session.FromContext(r.Context())
	defaultVacation := database.Vacation{
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 1),
		TermID:    int64(s.TermID),
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		views.VacationEditPage(defaultVacation, true).Render(r.Context(), w)
		return
	}

	vacation, err := s.Database.GetVacation(r.Context(), int64(id))
	if err != nil {
		views.VacationEditPage(defaultVacation, true).Render(r.Context(), w)
		return
	}

	views.VacationEditPage(vacation, false).Render(r.Context(), w)
}

func AdminVacationPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue(views.KeyTermID))
	startDate, _ := time.Parse(time.DateOnly, r.FormValue(views.KeyTermStartDate))
	endDate, _ := time.Parse(time.DateOnly, r.FormValue(views.KeyTermEndDate))
	termID, _ := strconv.Atoi(r.FormValue(views.KeyVacationTermID))

	s := session.FromContext(r.Context())

	if r.Form.Has(views.FlagCreate) {
		s.Database.CreateVacation(r.Context(), database.CreateVacationParams{
			Name:      r.FormValue(views.KeyTermName),
			StartDate: startDate,
			EndDate:   endDate,
			TermID:    int64(termID),
		})
	} else if r.Form.Has(views.FlagUpdate) {
		s.Database.UpdateVacation(r.Context(), database.UpdateVacationParams{
			ID:        int64(id),
			Name:      r.FormValue(views.KeyTermName),
			StartDate: startDate,
			EndDate:   endDate,
			TermID:    int64(termID),
		})
	} else if r.Form.Has(views.FlagDelete) {
		s.Database.DeleteVacation(r.Context(), int64(id))
	}

	http.Redirect(w, r, views.DestAdminCalendar, http.StatusSeeOther)
}
