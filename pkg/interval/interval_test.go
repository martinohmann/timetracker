package interval

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClose(t *testing.T) {
	i := &Interval{}

	assert.False(t, i.IsClosed())

	i.Close()

	assert.True(t, i.IsClosed())
	assert.NotEqual(t, time.Time{}, i.End)
}

func TestAfter(t *testing.T) {
	i1 := &Interval{
		Start: time.Date(2018, time.December, 25, 0, 0, 0, 0, time.Local),
		End:   time.Date(2018, time.December, 26, 0, 0, 0, 0, time.Local),
	}
	i2 := &Interval{
		Start: time.Date(2018, time.December, 25, 0, 0, 0, 0, time.Local),
		End:   time.Date(2018, time.December, 26, 0, 0, 0, 0, time.Local),
	}
	i3 := &Interval{
		Start: time.Date(2018, time.December, 25, 0, 0, 0, 0, time.Local),
		End:   time.Date(2018, time.December, 27, 0, 0, 0, 0, time.Local),
	}
	i4 := &Interval{
		Start: time.Date(2018, time.December, 24, 0, 0, 0, 0, time.Local),
		End:   time.Date(2018, time.December, 26, 0, 0, 0, 0, time.Local),
	}

	assert.False(t, i1.After(i2))
	assert.True(t, i3.After(i1))
	assert.False(t, i4.After(i1))
}

func TestDuration(t *testing.T) {
	i := &Interval{
		Start: time.Date(2018, time.December, 25, 0, 0, 0, 0, time.Local),
		End:   time.Date(2018, time.December, 25, 0, 0, 10, 0, time.Local),
	}

	assert.Equal(t, 10*time.Second, i.Duration())

	i = &Interval{
		Start: time.Now().AddDate(0, 0, -1),
	}

	assert.True(t, i.Duration() >= time.Hour*24)
}
