package subject

import (
	"calendorario/pkg/database"
	"calendorario/pkg/templates"
	"calendorario/routes"
	"net/http"
	"strconv"
)

const (
	keyID            = routes.KeyID
	keyName          = "name"
	keyColorHexValue = "color_hex_value"
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
		View(database.Subject{}, true).Render(r.Context(), w)
		return
	}

	subject, err := h.Database.GetSubject(r.Context(), int64(id))
	if err != nil {
		View(database.Subject{}, true).Render(r.Context(), w)
		return
	}

	View(subject, false).Render(r.Context(), w)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue(keyID))

	if r.Form.Has(templates.FlagCreate) {
		h.Database.CreateSubject(r.Context(), database.CreateSubjectParams{
			Name:          r.FormValue(keyName),
			ColorHexValue: []byte(r.FormValue(keyColorHexValue)),
		})
	} else if r.Form.Has(templates.FlagUpdate) {
		h.Database.UpdateSubject(r.Context(), database.UpdateSubjectParams{
			ID:            int64(id),
			Name:          r.FormValue(keyName),
			ColorHexValue: []byte(r.FormValue(keyColorHexValue)),
		})
	} else if r.Form.Has(templates.FlagDelete) {
		h.Database.DeleteSubject(r.Context(), int64(id))
	}

	http.Redirect(w, r, routes.RouteAdminSubjects, http.StatusSeeOther)
}
