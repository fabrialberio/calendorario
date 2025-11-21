package session

import (
	"net/http"
	"strconv"
)

const termCookieName = "current_term"

func GetTermCookie(r *http.Request) (int, error) {
	cookie, err := r.Cookie(termCookieName)
	if err != nil {
		return 0, ErrCookieNotFound
	}

	return strconv.Atoi(cookie.Value)
}

func SetTermCookie(w http.ResponseWriter, termID int) {
	http.SetCookie(w, &http.Cookie{
		Name:     termCookieName,
		Value:    strconv.Itoa(termID),
		Path:     "/",
		HttpOnly: true,
	})
}

func UnsetTermCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     termCookieName,
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
}
