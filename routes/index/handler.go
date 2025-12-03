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
		return
	}

	user, err := session.AuthenticatedUser(r)

	if err != nil {
		http.Redirect(w, r, routes.RouteLogin, http.StatusSeeOther)
	} else {
		switch user.Role {
		case database.RoleAdministrator:
			http.Redirect(w, r, routes.RouteAdmin, http.StatusSeeOther)
		case database.RoleSecretary:
			http.Redirect(w, r, routes.RouteSecretary, http.StatusSeeOther)
		case database.RoleTeacher:
			http.Redirect(w, r, routes.RouteTeacher, http.StatusSeeOther)
		}
	}
}
