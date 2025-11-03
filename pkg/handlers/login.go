package handlers

import (
	"calendorario/views"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("error") {
		views.Login(true, false).Render(r.Context(), w)
	} else {
		views.Login(false, false).Render(r.Context(), w)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

}
