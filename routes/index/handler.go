package index

import (
	"calendorario/pkg/database"
	"calendorario/pkg/session"
	"calendorario/routes"

	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}

	s := session.FromContext(r.Context())
	user, err := s.User()

	if err != nil {
		http.Redirect(w, r, routes.DestLogin, http.StatusSeeOther)
	} else {
		switch user.Role {
		case database.RoleAdministrator:
			http.Redirect(w, r, routes.DestAdmin, http.StatusSeeOther)
		case database.RoleSecretary:
			http.Redirect(w, r, routes.DestSecretary, http.StatusSeeOther)
		case database.RoleTeacher:
			http.Redirect(w, r, routes.DestTeacher, http.StatusSeeOther)
		}
	}
}
