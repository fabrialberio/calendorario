package admin

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
	authUser, _ := session.AuthenticatedUser(r)

	user, _ := h.Database.GetUser(r.Context(), authUser.ID)
	terms, _ := h.Database.ListTerms(r.Context())

	View(user, termID, terms).Render(r.Context(), w)
}
