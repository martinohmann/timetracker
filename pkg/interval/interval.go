package interval

import (
	"time"
)

type Interval struct {
	ID    int
	Tag   string
	Start time.Time
	End   time.Time
}

func New(tag string, start, end time.Time) Interval {
	if !end.IsZero() && end.Before(start) {
		start, end = end, start
	}

	return Interval{Tag: tag, Start: start, End: end}
}

func (i Interval) IsOpen() bool {
	return i.End.IsZero()
}

func (i Interval) IsClosed() bool {
	return !i.IsOpen()
}

func (i Interval) Duration() time.Duration {
	if i.End.IsZero() {
		return time.Now().Sub(i.Start)
	}

	return i.End.Sub(i.Start)
}

func (i Interval) Equal(other Interval) bool {
	return i.Start.Equal(other.Start) && i.End.Equal(other.End)
}

func (i Interval) Before(other Interval) bool {
	if i.Start.Before(other.Start) || (i.Start.Equal(other.Start) && i.End.Before(other.End)) {
		return true
	}

	return false
}

func (i Interval) After(other Interval) bool {
	return !i.Equal(other) && !i.Before(other)
}
