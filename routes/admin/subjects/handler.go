package subjects

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"net/http"
)

type Handler struct {
	Database *database.Queries
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	termID, _ := session.SelectedTermID(r)
	term, _ := h.Database.GetTerm(r.Context(), int64(termID))

	View(term).Render(r.Context(), w)
}
