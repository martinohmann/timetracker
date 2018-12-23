package dateutil

import (
	"fmt"
	"os"
	"time"

	"github.com/araddon/dateparse"
)

func ParseDate(s string, defaultValue ...time.Time) (time.Time, error) {
	if s == "" && len(defaultValue) > 0 {
		return defaultValue[0], nil
	}

	return dateparse.ParseLocal(s)
}

func MustParseDate(s string, defaultValue ...time.Time) time.Time {
	date, err := ParseDate(s, defaultValue...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return date
}

func BeginOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func EndOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 23, 59, 59, 99, time.Local)
}
