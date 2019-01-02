package dateutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		given        string
		defaultValue []time.Time
		expected     time.Time
		expectError  bool
	}{
		{
			given:        "",
			defaultValue: []time.Time{time.Time{}},
			expected:     time.Time{},
		},
		{
			given:    "2018-12-25 12:34:59",
			expected: time.Date(2018, time.December, 25, 12, 34, 59, 0, time.Local),
		},
		{
			given:       "foo",
			expectError: true,
		},
	}

	for _, tt := range tests {
		res, err := ParseDate(tt.given, tt.defaultValue...)

		if tt.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res)
		}
	}
}

func TestBeginOfDay(t *testing.T) {
	expected := time.Date(2018, time.December, 29, 0, 0, 0, 0, time.Local)
	res := BeginOfDay(2018, time.December, 29)

	assert.Equal(t, expected, res)
}

func TestFormat(t *testing.T) {
	date := time.Date(2018, time.December, 29, 12, 34, 56, 99, time.Local)

	assert.Equal(t, "2018/12/29 12:34:56", Format(date))
}

func TestFormatDuration(t *testing.T) {
	duration := time.Hour*3 + time.Minute + time.Second*2 + time.Millisecond*100

	assert.Equal(t, "3h1m2s", FormatDuration(duration))
}
