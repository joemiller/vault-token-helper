package cmd_test

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/joemiller/vault-token-helper/cmd"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if v := os.Getenv("GO_TEST_MODE"); v == "1" {
		// we are the vault-token-helper binary under test, execute rootCmd
		cmd.Execute()
	} else {
		// we are the test runner, run the Tests*
		os.Exit(m.Run())
	}
}

// execCmd executes the vault-token-helper with the provides args and returns
// stdout, stderr, and error.
// execCommand("list", "--debug") would be similar to executing the compiled program "vault-token-helper list --debug"
func execCmd(env []string, stdin string, args ...string) (string, string, error) {
	cmd := exec.Command(os.Args[0], args...)

	cmd.Env = env
	cmd.Env = append(cmd.Env, "GO_TEST_MODE=1")

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Stdin = strings.NewReader(stdin)

	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func TestGetCmd_MissingVAULT_ADDR(t *testing.T) {
	// TODO: get this working in CI for all platforms
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	stdin := ""
	env := []string{}
	stdout, stderr, err := execCmd(env, stdin, "get")

	assert.NotNil(t, err) // vault-token-helper should exit non-zero when VAULT_ADDR is not set
	assert.Equal(t, "", stdout)
	assert.Equal(t, "Error: Missing VAULT_ADDR environment variable\n", stderr)
}

func TestGetCmd_NoMatch(t *testing.T) {
	// TODO: get this working in CI for all platforms
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	stdin := ""
	env := []string{"VAULT_ADDR=https://foo.bar:8200"}
	stdout, stderr, err := execCmd(env, stdin, "get")

	assert.Nil(t, err) // vault-token-helper should exit 0 if no token is stored for the $VAULT_ADDR
	assert.Equal(t, "", stdout)
	assert.Equal(t, "", stderr)
}

func TestGetCmd_Match(t *testing.T) {
	// TODO: get this working in CI for all platforms
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	stdin := ""
	env := []string{"VAULT_ADDR=https://foo.bar:8200"}
	stdout, stderr, err := execCmd(env, stdin, "get")
	assert.Nil(t, err) // vault-token-helper should exit 0 if no token is stored for the $VAULT_ADDR
	assert.Equal(t, "", stdout)
	assert.Equal(t, "", stderr)

	stdin = "token-foo"
	_, _, err = execCmd(env, stdin, "store")
	assert.Nil(t, err)
	defer func() { _, _, _ = execCmd(env, "", "erase") }()

	stdout, stderr, err = execCmd(env, "", "get")
	assert.Nil(t, err)
	// spew.Dump(stdout)
	// spew.Dump(stderr)
	assert.Nil(t, err) // vault-token-helper should exit 0 if no token is stored for the $VAULT_ADDR
	assert.Equal(t, "token-foo", stdout)
	assert.Equal(t, "", stderr)
}
