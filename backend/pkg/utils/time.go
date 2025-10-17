package utils

import "time"

// NowUTC returns current time in UTC
func NowUTC() time.Time {
	return time.Now().UTC()
}

// StartOfDay returns the start of the day (00:00:00) for the given time
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day (23:59:59) for the given time
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

// AddDays adds the specified number of days to the given time
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// IsSameDay checks if two times are on the same day
func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// DaysSince returns the number of days between two times
func DaysSince(from, to time.Time) int {
	from = StartOfDay(from)
	to = StartOfDay(to)
	return int(to.Sub(from).Hours() / 24)
}
