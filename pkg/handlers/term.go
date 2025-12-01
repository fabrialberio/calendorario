package handlers

import (
	"calendorario/pkg/session"
	"calendorario/routes"
	"net/http"
	"strconv"
)

func AdminLoadTermGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Redirect(w, r, routes.DestAdmin, http.StatusSeeOther)
	}

	session.SetTermCookie(w, id)

	http.Redirect(w, r, routes.DestAdminCalendar, http.StatusSeeOther)
}
