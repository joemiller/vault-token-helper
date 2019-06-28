package store_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/99designs/keyring"
	"github.com/joemiller/vault-token-helper/pkg/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO only run if pass binary is available

func setup(t *testing.T) (string, func(t *testing.T)) {
	pwd, err := os.Getwd()
	require.Nil(t, err)

	tmp := os.TempDir()
	if runtime.GOOS == "darwin" {
		// XXX: on macos we place the tempdir under /tmp to avoid "gpg: can't connect to the agent: File name too long" - https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=847206
		tmp = "/tmp"
	}
	tmpdir, err := ioutil.TempDir(tmp, "vault-token-helper-pass-test")
	require.Nil(t, err)

	// Create a temporary GPG dir
	gnupghome := filepath.Join(tmpdir, ".gnupg")
	os.Mkdir(gnupghome, os.FileMode(int(0700)))
	os.Setenv("GNUPGHOME", gnupghome)
	os.Unsetenv("GPG_AGENT_INFO")
	os.Unsetenv("GPG_TTY")

	// import and trust the test key
	cmd := exec.Command("gpg", "--import", filepath.Join(pwd, "fixtures", "test-gpg.key"))
	out, err := cmd.CombinedOutput()
	require.Nil(t, err, string(out))

	cmd = exec.Command("gpg", "--import-ownertrust", filepath.Join(pwd, "fixtures", "test-ownertrust-gpg.txt"))
	out, err = cmd.CombinedOutput()
	require.Nil(t, err, string(out))

	// initialize a 'pass' directory under the tmpdir using the gpg key-id created in the prior step
	passdir := filepath.Join(tmpdir, ".password-store")
	os.Setenv("PASSWORD_STORE_DIR", passdir)
	cmd = exec.Command("pass", "init", "test@example.com")
	err = cmd.Run()
	require.Nil(t, err)

	return passdir, func(t *testing.T) {
		os.RemoveAll(tmpdir)
	}
}

func TestPassStore(t *testing.T) {
	passdir, teardown := setup(t)
	defer teardown(t)

	st, err := store.New(keyring.Config{
		ServiceName:     "test",
		PassPrefix:      "vault-test",
		PassDir:         passdir,
		AllowedBackends: []keyring.BackendType{keyring.PassBackend},
	})
	assert.Nil(t, err)
	require.NotNil(t, st) // stop the tests if the store is nil, else everything following will panic

	// Store a token
	token1 := store.Token{
		VaultAddr: "https://localhost:8200",
		Token:     "token-foo",
	}

	err = st.Store(token1)
	assert.Nil(t, err)

	// GetAll tokens
	tokens, err := st.List()
	assert.NotEmpty(t, tokens)

	// Get a token by addr. Mixed case addr should be normalized for a successful lookup
	v1, err := st.Get("httpS://LOCALhost:8200/")
	assert.Nil(t, err)
	assert.Equal(t, token1, v1)

	// Erase
	err = st.Erase("https://localhost:8200")
	assert.Nil(t, err)

	// empty token store
	tokens, err = st.List()
	assert.Nil(t, err)
	assert.Empty(t, tokens)
}
