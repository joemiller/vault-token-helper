package cmd

import (
	"os"
	"testing"

	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

func TestStoreCmd(t *testing.T) {
	// TODO: get this working in CI for all platforms
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	app = "test"
	err := os.Setenv("VAULT_ADDR", "https://foo")
	assert.Nil(t, err)

	stdOut, stdErr, err := execCommand("store")
	defer func() { _, _, _ = execCommand("erase") }()
	assert.Nil(t, err)

	goldie.Assert(t, t.Name()+".stdout", stdOut)
	goldie.Assert(t, t.Name()+".stderr", stdErr)
}
