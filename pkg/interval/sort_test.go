package interval

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSortByDate(t *testing.T) {
	intervals := []Interval{
		{
			Tag:   "foo",
			Start: time.Date(2018, time.December, 30, 0, 0, 0, 0, time.Local),
			End:   time.Date(2018, time.December, 31, 0, 0, 0, 0, time.Local),
		},
		{
			Start: time.Date(2018, time.December, 24, 0, 0, 0, 0, time.Local),
			End:   time.Date(2018, time.December, 26, 0, 0, 0, 0, time.Local),
		},
		{
			Start: time.Date(2018, time.December, 25, 0, 0, 0, 0, time.Local),
			End:   time.Date(2018, time.December, 27, 0, 0, 0, 0, time.Local),
		},
		{
			Tag:   "bar",
			Start: time.Date(2018, time.December, 30, 0, 0, 0, 0, time.Local),
			End:   time.Date(2018, time.December, 31, 0, 0, 0, 0, time.Local),
		},
	}

	given := make([]Interval, len(intervals))
	copy(given, intervals)

	sort.Sort(SortByDate(intervals))

	assert.Equal(t, given[0], intervals[3])
	assert.Equal(t, given[1], intervals[0])
	assert.Equal(t, given[2], intervals[1])
	assert.Equal(t, given[3], intervals[2])
}
