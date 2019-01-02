package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "add", "test")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), `interval with tag "test" added`)
}
