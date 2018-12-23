package interval

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Interval struct {
	gorm.Model
	From time.Time
	To   time.Time
}

func New(tag string, from, to time.Time) Interval {
	if !to.IsZero() && to.Before(from) {
		from, to = from, to
	}

	return Interval{From: from, To: to}
}

func (i Interval) IsOpen() bool {
	return i.To.IsZero()
}

func (i Interval) Duration() time.Duration {
	if i.To.IsZero() {
		return time.Now().Sub(i.From)
	}

	return i.To.Sub(i.From)
}

func (i Interval) Equal(other Interval) bool {
	return i.From.Equal(other.From) && i.To.Equal(other.To)
}

func (i Interval) Before(other Interval) bool {
	if i.From.Before(other.From) || (i.From.Equal(other.From) && i.To.Before(other.To)) {
		return true
	}

	return false
}

func (i Interval) After(other Interval) bool {
	return !i.Equal(other) && !i.Before(other)
}
