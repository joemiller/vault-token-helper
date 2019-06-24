package cmd

import (
	"testing"

	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

// ./vault-token-helper
// Should return without error and print help usage to stderr
func TestRoot_noArgs(t *testing.T) {
	stdOut, stdErr, err := execCommand("")
	assert.Nil(t, err)
	goldie.Assert(t, t.Name()+".stdout", stdOut)
	goldie.Assert(t, t.Name()+".stderr", stdErr)
}
