package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStopCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "stop")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), `interval with ID 1 closed`)
}
