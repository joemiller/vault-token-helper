package cmd

import (
	"os"
	"testing"

	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

func TestGetCmd_Empty(t *testing.T) {
	// TODO: get this working in CI for all platforms
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	app = "test"

	stdOut, stdErr, err := execCommand("get")
	assert.NotNil(t, err)

	goldie.Assert(t, t.Name()+".stdout", stdOut)
	goldie.Assert(t, t.Name()+".stderr", stdErr)
}

func TestGetCmd_NotEmpty(t *testing.T) {
	// TODO: get this working in CI for all platforms
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	app = "test"
	err := os.Setenv("VAULT_ADDR", "https://foo")
	assert.Nil(t, err)

	_, _, err = execCommand("store")
	assert.Nil(t, err)
	defer func() { _, _, _ = execCommand("erase") }()

	stdOut, stdErr, err := execCommand("get")
	assert.Nil(t, err)

	goldie.Assert(t, t.Name()+".stdout", stdOut)
	goldie.Assert(t, t.Name()+".stderr", stdErr)
}
