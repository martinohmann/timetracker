package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "show")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "All intervals\n")
}

func TestShowYearCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "show", "year", "--year", "2018")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "All intervals between 2018/01/01 00:00:00 and 2019/01/01 00:00:00\n")
}

func TestShowMonthCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "show", "month", "--year", "2018", "--month", "12")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "All intervals between 2018/12/01 00:00:00 and 2019/01/01 00:00:00\n")
}

func TestShowWeekCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "show", "week", "--date", "2018-12-01")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "All intervals between 2018/11/26 00:00:00 and 2018/12/03 00:00:00\n")
}

func TestShowDateCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "show", "date", "--date", "2018-12-01")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "All intervals between 2018/12/01 00:00:00 and 2018/12/02 00:00:00\n")
}

func TestShowStartCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "show", "--start", "2018-12-01", "--end", "")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "All intervals since 2018/12/01 00:00:00\n")
}

func TestShowEndCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "show", "--end", "2018-12-01", "--start", "")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "All intervals until 2018/12/01 00:00:00\n")
}
