package dateutil

import (
	"time"

	"github.com/araddon/dateparse"
)

// DateFormat and DurationPrecision are variables so that they can be
// customized at runtime if needed.
var (
	// DateFormat defines the display date format
	DateFormat = "2006/01/02 15:04:05"

	// DurationPrecision defines the precision for formatted durations
	DurationPrecision = time.Second
)

// ParseDate parses a date from given string. It will return the optional
// default value if s is empty and an error if s does not match any known date
// format
func ParseDate(s string, defaultValue ...time.Time) (time.Time, error) {
	if s == "" && len(defaultValue) > 0 {
		return defaultValue[0], nil
	}

	return dateparse.ParseLocal(s)
}

// BeginOfDay returns a date which represents the beginning of the day
func BeginOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// Format formats date using the defined DateFormat
func Format(date time.Time) string {
	return date.Format(DateFormat)
}

// FormatDuration formats a duration
func FormatDuration(duration time.Duration) string {
	return duration.Truncate(DurationPrecision).String()
}
