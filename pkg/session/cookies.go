package session

import (
	"net/http"
	"strconv"
	"time"
)

const selectedTermCookie = "selected_term"

func SelectedTermID(r *http.Request) (int, error) {
	cookie, err := r.Cookie(selectedTermCookie)
	if err != nil {
		return 0, ErrCookieNotFound
	}

	return strconv.Atoi(cookie.Value)
}

func SetSelectedTermID(w http.ResponseWriter, termID int) {
	setCookie(w, selectedTermCookie, strconv.Itoa(termID))
}

func UnsetSelectedTermID(w http.ResponseWriter) {
	unsetCookie(w, selectedTermCookie)
}

const selectedDateCookie = "selected_date"

func SelectedDate(r *http.Request) (time.Time, error) {
	cookie, err := r.Cookie(selectedDateCookie)
	if err != nil {
		return time.Time{}, ErrCookieNotFound
	}

	return time.Parse(time.DateOnly, cookie.Value)
}

func SetSelectedDate(w http.ResponseWriter, date time.Time) {
	setCookie(w, selectedDateCookie, date.Format(time.DateOnly))
}

func UnsetSelectedDate(w http.ResponseWriter) {
	unsetCookie(w, selectedDateCookie)
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
