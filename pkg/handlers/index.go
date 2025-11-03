package handlers

import (
	"calendorario/views"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)

	/* user, err := appContext.AuthenticatedUser(r)
	if errors.Is(err, auth.ErrNoCookie) {
		if r.URL.Query().Has("errormsg") {
			data.ErrorMsg = "Utente o password errati"
		}
	} else if err != nil {
		data.ErrorMsg = "Sessione scaduta"
	} else {
		loginRedirect(w, r, user.RoleID)
		return
	} */

	views.Login().Render(r.Context(), w)
}
