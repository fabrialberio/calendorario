package dates

import (
	"strconv"
	"time"
)

var (
	WeekdayNames = []string{
		"Lunedì",
		"Martedì",
		"Mercoledì",
		"Giovedì",
		"Venerdì",
		"Sabato",
		"Domenica",
	}
	WeekdayNicks = []string{"lun", "mar", "mer", "gio", "ven", "sab", "dom"}
	MonthNames   = map[time.Month]string{
		time.January:   "Gennaio",
		time.February:  "Febbraio",
		time.March:     "Marzo",
		time.April:     "Aprile",
		time.May:       "Maggio",
		time.June:      "Giugno",
		time.July:      "Luglio",
		time.August:    "Agosto",
		time.September: "Settembre",
		time.October:   "Ottobre",
		time.November:  "Novembre",
		time.December:  "Dicembre",
	}
	MonthNicks = map[time.Month]string{
		time.January:   "gen",
		time.February:  "feb",
		time.March:     "mar",
		time.April:     "apr",
		time.May:       "mag",
		time.June:      "giu",
		time.July:      "lug",
		time.August:    "ago",
		time.September: "set",
		time.October:   "ott",
		time.November:  "nov",
		time.December:  "dic",
	}
)

func DayString(date time.Time, showWeekDay bool, showDay bool, showYear bool) string {
	result := ""

	if showWeekDay {
		result += WeekdayNicks[date.Weekday()] + " "
	}
	if showDay {
		result += strconv.Itoa(date.Day()) + " "
	}
	result += MonthNicks[date.Month()]
	if showYear {
		result += " " + strconv.Itoa(date.Year())
	}

	return result
}

func MonthString(year int, month time.Month) string {
	return MonthNames[month] + " " + strconv.Itoa(year)
}
