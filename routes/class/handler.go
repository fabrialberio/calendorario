package class

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/pkg/templates"
	"calendorario/routes"
	"net/http"
	"strconv"
)

const (
	keyID        = routes.KeyID
	keyGrade     = "grade"
	keySection   = "section"
	keyTermID    = "term_id"
	keyProgramID = "program_id"
)

type Handler struct{}

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
	s := session.FromContext(r.Context())
	initialClass := database.Class{
		Grade:  1,
		TermID: int64(s.SelectedTermID),
	}

	id, err := strconv.Atoi(r.FormValue(keyID))
	if err != nil {
		View(initialClass, true).Render(r.Context(), w)
		return
	}

	class, err := s.Database.GetClass(r.Context(), int64(id))
	if err != nil {
		View(initialClass, true).Render(r.Context(), w)
		return
	}

	View(class, false).Render(r.Context(), w)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue(keyID))
	grade, _ := strconv.Atoi(r.FormValue(keyGrade))
	termID, _ := strconv.Atoi(r.FormValue(keyTermID))
	programID, _ := strconv.Atoi(r.FormValue(keyProgramID))

	s := session.FromContext(r.Context())

	if r.Form.Has(templates.FlagCreate) {
		s.Database.CreateClass(r.Context(), database.CreateClassParams{
			Grade:     int32(grade),
			Section:   r.FormValue(keySection),
			TermID:    int64(termID),
			ProgramID: int64(programID),
		})
	} else if r.Form.Has(templates.FlagUpdate) {
		s.Database.UpdateClass(r.Context(), database.UpdateClassParams{
			ID:        int64(id),
			Grade:     int32(grade),
			Section:   r.FormValue(keySection),
			TermID:    int64(termID),
			ProgramID: int64(programID),
		})
	} else if r.Form.Has(templates.FlagDelete) {
		s.Database.DeleteClass(r.Context(), int64(id))
	}

	http.Redirect(w, r, routes.RouteAdmin, http.StatusSeeOther)
}
