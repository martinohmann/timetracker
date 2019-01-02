package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmendCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "amend", "1")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), `interval with ID 1 updated`)
}
