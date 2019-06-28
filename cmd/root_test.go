package cmd

import (
	"os"
	"testing"

	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

// ./vault-token-helper
// Should return without error and print help usage to stderr
func TestRoot_noArgs(t *testing.T) {
	// TODO: get this working in CI for all platforms. It currently fails under the windows2019 image on azure pipelines although
	//       passes when using the windows10 image in ./vagrant/windows
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	stdOut, stdErr, err := execCommand("")
	assert.Nil(t, err)
	goldie.Assert(t, t.Name()+".stdout", stdOut)
	goldie.Assert(t, t.Name()+".stderr", stdErr)
}
