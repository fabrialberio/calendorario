package dates

import (
	"time"
)

// Interval between start and end, both included.
type DateInterval struct {
	Start time.Time
	End   time.Time
}

func Interval(start time.Time, end time.Time) DateInterval {
	return DateInterval{truncateDate(start), truncateDate(end)}
}

// Wether date is contained in the interval.
func (i DateInterval) Contains(date time.Time) bool {
	date = truncateDate(date)
	return i.Start.Compare(date) <= 0 && i.End.Compare(date) >= 0
}

// Wether two date intervals overlap.
func (i DateInterval) Overlaps(interval DateInterval) bool {
	return i.End.Compare(interval.Start) >= 0 && i.Start.Compare(interval.End) <= 0
}

// Lenght of interval in days.
func (i DateInterval) Days() int {
	return int(i.End.AddDate(0, 0, 1).Sub(i.Start).Hours() / 24)
}

func (i DateInterval) String() string {
	showYear := i.Start.Year() != i.End.Year()

	return DayString(i.Start, false, true, showYear) + " - " +
		DayString(i.End, false, true, showYear)
}

func truncateDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}
