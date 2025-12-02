package term

import (
	"calendorario/pkg/database"
	"calendorario/pkg/templates"
	"calendorario/routes"
	"net/http"
	"strconv"
	"time"
)

const (
	keyID        = routes.KeyID
	keyName      = "name"
	keyStartDate = "start_date"
	keyEndDate   = "end_date"
)

type Handler struct {
	Database *database.Queries
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPost:
		h.Post(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue(keyID))
	if err != nil {
		View(database.Term{}, true).Render(r.Context(), w)
		return
	}

	term, err := h.Database.GetTerm(r.Context(), int64(id))
	if err != nil {
		View(database.Term{}, true).Render(r.Context(), w)
		return
	}

	View(term, false).Render(r.Context(), w)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue(keyID))
	startDate, _ := time.Parse(time.DateOnly, r.FormValue(keyStartDate))
	endDate, _ := time.Parse(time.DateOnly, r.FormValue(keyEndDate))

	if r.Form.Has(templates.FlagCreate) {
		h.Database.CreateTerm(r.Context(), database.CreateTermParams{
			Name:      r.FormValue(keyName),
			StartDate: startDate,
			EndDate:   endDate,
		})
	} else if r.Form.Has(templates.FlagUpdate) {
		h.Database.UpdateTerm(r.Context(), database.UpdateTermParams{
			ID:        int64(id),
			Name:      r.FormValue(keyName),
			StartDate: startDate,
			EndDate:   endDate,
		})
	} else if r.Form.Has(templates.FlagDelete) {
		h.Database.DeleteTerm(r.Context(), int64(id))
	}

	http.Redirect(w, r, routes.RouteAdmin, http.StatusSeeOther)
}
