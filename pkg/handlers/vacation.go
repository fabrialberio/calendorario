package handlers

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/routes"
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
		routes.VacationEditPage(defaultVacation, true).Render(r.Context(), w)
		return
	}

	vacation, err := s.Database.GetVacation(r.Context(), int64(id))
	if err != nil {
		routes.VacationEditPage(defaultVacation, true).Render(r.Context(), w)
		return
	}

	routes.VacationEditPage(vacation, false).Render(r.Context(), w)
}

func AdminVacationPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue(routes.KeyVacationID))
	startDate, _ := time.Parse(time.DateOnly, r.FormValue(routes.KeyVacationStartDate))
	endDate, _ := time.Parse(time.DateOnly, r.FormValue(routes.KeyVacationEndDate))
	termID, _ := strconv.Atoi(r.FormValue(routes.KeyVacationTermID))

	s := session.FromContext(r.Context())

	if r.Form.Has(routes.FlagCreate) {
		s.Database.CreateVacation(r.Context(), database.CreateVacationParams{
			Name:      r.FormValue(routes.KeyVacationName),
			StartDate: startDate,
			EndDate:   endDate,
			TermID:    int64(termID),
		})
	} else if r.Form.Has(routes.FlagUpdate) {
		s.Database.UpdateVacation(r.Context(), database.UpdateVacationParams{
			ID:        int64(id),
			Name:      r.FormValue(routes.KeyVacationName),
			StartDate: startDate,
			EndDate:   endDate,
			TermID:    int64(termID),
		})
	} else if r.Form.Has(routes.FlagDelete) {
		s.Database.DeleteVacation(r.Context(), int64(id))
	}

	http.Redirect(w, r, routes.RouteAdminCalendar, http.StatusSeeOther)
}
