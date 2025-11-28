package interval

import "time"

// Interval between start and end, both included.
type DateInterval struct {
	Start time.Time
	end   time.Time
}

func Dates(start time.Time, end time.Time) DateInterval {
	return DateInterval{start.Truncate(time.Hour), end.Truncate(time.Hour)}
}

// Wether date is contained in the interval.
func (i DateInterval) Contains(date time.Time) bool {
	return i.Start.Compare(date) <= 0 && i.end.Compare(date) >= 0
}

// Wether two date intervals overlap.
func (i DateInterval) Overlaps(interval DateInterval) bool {
	return i.end.Compare(interval.Start) >= 0 && i.Start.Compare(interval.end) <= 0
}

// Lenght of interval in days.
func (i DateInterval) Days() int {
	return int(i.end.AddDate(0, 0, 1).Sub(i.Start).Hours() / 24)
}
