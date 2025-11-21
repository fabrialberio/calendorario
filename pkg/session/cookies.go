package session

import (
	"calendorario/pkg/database"
	"net/http"
	"strconv"
)

const termCookieName = "current_term"

func SetTermCookie(w http.ResponseWriter, term database.Term) {
	http.SetCookie(w, &http.Cookie{
		Name:     termCookieName,
		Value:    strconv.Itoa(int(term.ID)),
		HttpOnly: true,
	})
}
