package cmd

import (
	"os"
	"runtime"
	"testing"

	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

func TestBackends(t *testing.T) {
	// TODO: get this working in CI for all platforms
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	stdOut, stdErr, err := execCommand("backends")
	assert.Nil(t, err)

	var goldfileBaseName string
	switch runtime.GOOS {
	case "darwin":
		goldfileBaseName = t.Name() + ".darwin"
	case "linux":
		goldfileBaseName = t.Name() + ".linux"
	}

	goldie.Assert(t, goldfileBaseName+".stdout", stdOut)
	goldie.Assert(t, goldfileBaseName+".stderr", stdErr)
}
