package vacation

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/pkg/templates"
	"calendorario/routes"
	"net/http"
	"strconv"
	"time"
)

const (
	KeyID        = "id"
	KeyName      = "name"
	KeyStartDate = "start_date"
	KeyEndDate   = "end_date"
	KeyTermID    = "term_id"
)

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
	s := session.FromContext(r.Context())
	defaultVacation := database.Vacation{
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 1),
		TermID:    int64(s.TermID),
	}

	id, err := strconv.Atoi(r.FormValue(KeyID))
	if err != nil {
		View(defaultVacation, true).Render(r.Context(), w)
		return
	}

	vacation, err := s.Database.GetVacation(r.Context(), int64(id))
	if err != nil {
		View(defaultVacation, true).Render(r.Context(), w)
		return
	}

	View(vacation, false).Render(r.Context(), w)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue(KeyID))
	startDate, _ := time.Parse(time.DateOnly, r.FormValue(KeyStartDate))
	endDate, _ := time.Parse(time.DateOnly, r.FormValue(KeyEndDate))
	termID, _ := strconv.Atoi(r.FormValue(KeyTermID))

	s := session.FromContext(r.Context())

	if r.Form.Has(templates.FlagCreate) {
		s.Database.CreateVacation(r.Context(), database.CreateVacationParams{
			Name:      r.FormValue(KeyName),
			StartDate: startDate,
			EndDate:   endDate,
			TermID:    int64(termID),
		})
	} else if r.Form.Has(templates.FlagUpdate) {
		s.Database.UpdateVacation(r.Context(), database.UpdateVacationParams{
			ID:        int64(id),
			Name:      r.FormValue(KeyName),
			StartDate: startDate,
			EndDate:   endDate,
			TermID:    int64(termID),
		})
	} else if r.Form.Has(templates.FlagDelete) {
		s.Database.DeleteVacation(r.Context(), int64(id))
	}

	http.Redirect(w, r, routes.RouteAdminCalendar, http.StatusSeeOther)
}
