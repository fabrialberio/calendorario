package handlers

import (
	"calendorario/views"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	views.Login().Render(r.Context(), w)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

}
