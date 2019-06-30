package store

import (
	"testing"

	"github.com/99designs/keyring"
	"github.com/stretchr/testify/assert"
)

// Test the reference implementation of the keyring lib using its
// in-memory ArrayKeyRing backend

func TestStore(t *testing.T) {
	st := &Store{
		cfg: keyring.Config{},
		kr:  &keyring.ArrayKeyring{},
	}

	// should be empty
	tokens, err := st.List()
	assert.Nil(t, err)
	assert.Empty(t, tokens)

	// Get of a missing item should not return an error
	_, err = st.Get("https://localhost:8200")
	assert.Nil(t, err)

	// Store a token
	token1 := Token{
		VaultAddr: "https://localhost:8200",
		Token:     "token-foo",
	}

	err = st.Store(token1)
	assert.Nil(t, err)

	// GetAll tokens
	tokens, err = st.List()
	assert.NotEmpty(t, tokens)

	// Get a token by addr. Mixed case addr should be normalized for a successful lookup
	token2, err := st.Get("httpS://LOCALhost:8200/")
	assert.Nil(t, err)
	assert.Equal(t, token1, token2)

	// Erase
	err = st.Erase("https://localhost:8200")
	assert.Nil(t, err)

	// empty token store
	tokens, err = st.List()
	assert.Nil(t, err)
	assert.Empty(t, tokens)
}
