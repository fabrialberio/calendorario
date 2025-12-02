package admin

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"net/http"

	"github.com/a-h/templ"
)

type Handler struct {
	Database *database.Queries
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	s := session.FromContext(r.Context())
	authUser, _ := s.User()

	user, _ := h.Database.GetUser(r.Context(), authUser.ID)
	terms, _ := h.Database.ListTerms(r.Context())

	templ.Handler(View(user, terms)).ServeHTTP(w, r)
}
