package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartCommand(t *testing.T) {
	buf, err := executeCommand(rootCmd, "start", "foo")

	assert.NoError(t, err)

	assert.Contains(t, buf.String(), `interval with tag "foo" started`)
}
