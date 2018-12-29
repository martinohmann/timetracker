package interval

import (
	"time"
)

// Interval defines the interval type
type Interval struct {
	ID    int
	Tag   string
	Start time.Time
	End   time.Time
}

// IsClosed returns true in the interval has a non-zero end date
func (i *Interval) IsClosed() bool {
	return !i.End.IsZero()
}

// Closes closes the interval by setting its end date to time.Now()
func (i *Interval) Close() {
	i.End = time.Now()
}

// Duration returns the total duration of the interval. If the interval is not
// closed, it will return the currently elapsed time.
func (i *Interval) Duration() time.Duration {
	if i.IsClosed() {
		return i.End.Sub(i.Start)
	}

	return time.Now().Sub(i.Start)
}

// Equal returns true if the interval has the same start and end as other
func (i *Interval) Equal(other *Interval) bool {
	return i.Start.Equal(other.Start) && i.End.Equal(other.End)
}

// Before return true if the interval started before other, or if start times
// are the same and it end before other, false otherwise
func (i *Interval) Before(other *Interval) bool {
	if i.Start.Before(other.Start) || (i.Start.Equal(other.Start) && i.End.Before(other.End)) {
		return true
	}

	return false
}

// After returns true if interval start after other
func (i *Interval) After(other *Interval) bool {
	return !i.Equal(other) && !i.Before(other)
}
