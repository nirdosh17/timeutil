package timeutil

import "time"

// BeginningOfMonth finds current month and moves date to first day of the month setting time components to zero.
//
//	e.g.  '2022-04-07 13:25:33.79' -> '2022-04-01 00:00:00'
func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// BeginningOfDay sets time components in a date to zero.
//
//	e.g '2022-04-07 13:25:33.79' -> '2022-04-07 00:00:00'
func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// BeginningOfHour resets time component smaller than hour (minute, second) to represent start of the hour (hh:00:00).
//
// e.g '2022-04-07 13:25:33.79' -> '2022-04-07 13:00:00'
func BeginningOfHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

// BeginningOfMinute resets time component smaller than minute (second, nano second) to represent start of the minute (hh:mm:00).
//
// e.g '2022-04-07 13:25:33.79' -> '2022-04-07 13:25:00'
func BeginningOfMinute(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMonth moves time to the end of month with nanosecond level precision.
//
//	e.g.  '2022-04-07 13:25:33.79' -> '2022-04-30 23:59:59'
func EndOfMonth(t time.Time) time.Time {
	// go to next month and reduce 1 NS to get 23:59:59.999999 i.e. end of day
	return MonthsAfter(t, 1).Add(-time.Nanosecond)
}

// EndOfDay moves time to the end of day (23:59:59).
//
//	e.g.  '2022-04-07 13:25:33.79' -> '2022-04-07 23:59:59'
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999, t.Location())
}

// EndOfHour moves time to the end of hour (hh:59:59).
//
//	e.g.  '2022-04-07 13:25:33.79' -> '2022-04-07 13:59:59'
func EndOfHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 59, 59, 999999, t.Location())
}

// EndOfMinute moves time to the end of hour (hh:mm:59).
//
//	e.g.  '2022-04-07 13:25:33.79' -> '2022-04-07 13:25:59'
func EndOfMinute(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 59, 999999, t.Location())
}

// MonthsAgo moves time back to N months without normalizing the result like time.AddDate method.
//
// If remaining months in current year is less than entered months, date will move to previous year.
//
// It sets the day to the first day of the month and time components to zero.
// Reason being month like Feb which only has 28 or 29 days, we can't travel back from March 30 to Feb 30.
func MonthsAgo(t time.Time, months int) time.Time {
	m := int(t.Month())
	y := t.Year()
	if m <= months {
		m = 12 - (months - m)
		y--
	} else {
		m -= months
	}
	return time.Date(y, time.Month(m), 1, 0, 0, 0, 0, t.Location())
}

// MonthsAfter moves time to N months in the future.
//
// If entered number of months is larger than remaining months in current year, it will spillover to next year.
//
// It sets the day to the first day of the month and time components to zero.
func MonthsAfter(t time.Time, months int) time.Time {
	m := int(t.Month())
	y := t.Year()
	if m+months > 12 {
		m = (m + months) - 12
		y++
	} else {
		m += months
	}
	return time.Date(y, time.Month(m), 1, 0, 0, 0, 0, t.Location())
}

// DaysAgo moves date to specified number of days in the past without altering time.
//
//	e.g.  DaysAgo("2022-04-14 13:25:33.79", 10) -> "2022-04-04 13:25:33.79"
func DaysAgo(t time.Time, numDays int) time.Time {
	return t.Add(-time.Duration(numDays*24) * time.Hour)
}

// DaysAfter moves date to specified number of days in the future without altering time.
//
//	e.g.  DaysAfter("2022-04-14 13:25:33.79", 1) -> "2022-04-15 13:25:33.79"
func DaysAfter(t time.Time, numDays int) time.Time {
	return t.Add(time.Duration(numDays*24) * time.Hour)
}
