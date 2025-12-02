package session

import (
	"net/http"
	"strconv"
	"time"
)

const termCookieName = "selected_term"

func GetTermCookie(r *http.Request) (int, error) {
	cookie, err := r.Cookie(termCookieName)
	if err != nil {
		return 0, ErrCookieNotFound
	}

	return strconv.Atoi(cookie.Value)
}

func SetTermCookie(w http.ResponseWriter, termID int) {
	setCookie(w, termCookieName, strconv.Itoa(termID))
}

func UnsetTermCookie(w http.ResponseWriter) {
	unsetCookie(w, termCookieName)
}

const dateCookieName = "selected_date"

func GetDateCookie(r *http.Request) (time.Time, error) {
	cookie, err := r.Cookie(dateCookieName)
	if err != nil {
		return time.Time{}, ErrCookieNotFound
	}

	return time.Parse(time.DateOnly, cookie.Value)
}

func SetDateCookie(w http.ResponseWriter, date time.Time) {
	setCookie(w, dateCookieName, date.Format(time.DateOnly))
}

func UnsetDateCookie(w http.ResponseWriter) {
	unsetCookie(w, dateCookieName)
}

func setCookie(w http.ResponseWriter, name string, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	})
}

func unsetCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
}
