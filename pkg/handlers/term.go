package handlers

import (
	"calendorario/pkg/database"
	"calendorario/pkg/requestcontext"
	"calendorario/views"
	"net/http"
	"strconv"
	"time"
)

func AdminTermGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		views.TermEditPage(database.Term{}, true).Render(r.Context(), w)
		return
	}

	rc := requestcontext.FromContext(r.Context())
	term, err := rc.Database.GetTerm(r.Context(), int64(id))
	if err != nil {
		views.TermEditPage(database.Term{}, true).Render(r.Context(), w)
		return
	}

	views.TermEditPage(term, false).Render(r.Context(), w)
}

func AdminTermPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue(views.KeyTermID))
	startDate, _ := time.Parse(time.DateOnly, r.FormValue(views.KeyTermStartDate))
	endDate, _ := time.Parse(time.DateOnly, r.FormValue(views.KeyTermEndDate))

	rc := requestcontext.FromContext(r.Context())

	if r.Form.Has(views.FlagCreate) {
		rc.Database.CreateTerm(r.Context(), database.CreateTermParams{
			Name:      r.FormValue(views.KeyTermName),
			StartDate: startDate,
			EndDate:   endDate,
		})
	} else if r.Form.Has(views.FlagUpdate) {
		rc.Database.UpdateTerm(r.Context(), database.UpdateTermParams{
			ID:        int64(id),
			Name:      r.FormValue(views.KeyTermName),
			StartDate: startDate,
			EndDate:   endDate,
		})
	} else if r.Form.Has(views.FlagDelete) {
		rc.Database.DeleteTerm(r.Context(), int64(id))
	}

	http.Redirect(w, r, views.DestAdmin, http.StatusSeeOther)
}
