package handlers

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/routes"
	"net/http"
	"strconv"
	"time"
)

func AdminTermGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		routes.TermEditPage(database.Term{}, true).Render(r.Context(), w)
		return
	}

	s := session.FromContext(r.Context())
	term, err := s.Database.GetTerm(r.Context(), int64(id))
	if err != nil {
		routes.TermEditPage(database.Term{}, true).Render(r.Context(), w)
		return
	}

	routes.TermEditPage(term, false).Render(r.Context(), w)
}

func AdminTermPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue(routes.KeyTermID))
	startDate, _ := time.Parse(time.DateOnly, r.FormValue(routes.KeyTermStartDate))
	endDate, _ := time.Parse(time.DateOnly, r.FormValue(routes.KeyTermEndDate))

	s := session.FromContext(r.Context())

	if r.Form.Has(routes.FlagCreate) {
		s.Database.CreateTerm(r.Context(), database.CreateTermParams{
			Name:      r.FormValue(routes.KeyTermName),
			StartDate: startDate,
			EndDate:   endDate,
		})
	} else if r.Form.Has(routes.FlagUpdate) {
		s.Database.UpdateTerm(r.Context(), database.UpdateTermParams{
			ID:        int64(id),
			Name:      r.FormValue(routes.KeyTermName),
			StartDate: startDate,
			EndDate:   endDate,
		})
	} else if r.Form.Has(routes.FlagDelete) {
		s.Database.DeleteTerm(r.Context(), int64(id))
	}

	http.Redirect(w, r, routes.DestAdmin, http.StatusSeeOther)
}

func AdminLoadTermGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Redirect(w, r, routes.DestAdmin, http.StatusSeeOther)
	}

	session.SetTermCookie(w, id)

	http.Redirect(w, r, routes.DestAdminCalendar, http.StatusSeeOther)
}
