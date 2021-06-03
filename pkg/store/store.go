package store

import (
	"net/url"
	"os"

	"github.com/99designs/keyring"
	"github.com/PuerkitoBio/purell"
)

// purellFlags are used to normalize VAULT_ADDR using the purell lib
const purellFlags = purell.FlagsSafe | purell.FlagsUsuallySafeGreedy | purell.FlagRemoveDuplicateSlashes

// Store is a backend token store
type Store struct {
	cfg keyring.Config
	kr  keyring.Keyring
}

// SupportedBackends is the list of backends we have support and have tested.
var SupportedBackends = []keyring.BackendType{
	// Windows
	keyring.WinCredBackend,
	// MacOS
	keyring.KeychainBackend,
	// Linux
	keyring.SecretServiceBackend,
	// KWalletBackend,
	keyring.PassBackend,
}

// New creates a new Store from a keyring.Config
func New(cfg keyring.Config) (*Store, error) {
	if os.Getenv("KEYRING_DEBUG") != "" {
		keyring.Debug = true
	}

	kr, err := keyring.Open(cfg)
	if err != nil {
		return nil, err
	}
	st := &Store{
		cfg: cfg,
		kr:  kr,
	}
	return st, nil
}

// Store saves the token in the token store, returning any errors that occur while trying to
// persist the token.
func (s *Store) Store(token Token) error {
	vaultAddr := encodeVaultAddr(token.VaultAddr)

	i := keyring.Item{
		Key:         vaultAddr,
		Data:        []byte(token.Token),
		Label:       "Vault-token: " + decodeVaultAddr(vaultAddr),
		Description: "Vault-token: " + decodeVaultAddr(vaultAddr),
		//KeychainNotSynchronizable: true,
	}
	return s.kr.Set(i)
}

// Get retrieves a token for the vaultAddr if one is available in the token store. A missing
// token is not an error. Errors are returned if there are errors communicating with the token store.
func (s *Store) Get(vaultAddr string) (Token, error) {
	vaultAddr = encodeVaultAddr(vaultAddr)

	i, err := s.kr.Get(vaultAddr)
	if err != nil {
		if err == keyring.ErrKeyNotFound {
			return Token{}, nil
		}
		return Token{}, err
	}

	t := Token{
		VaultAddr: decodeVaultAddr(vaultAddr),
		Token:     string(i.Data),
	}
	return t, nil
}

// List retrieves all tokens available in the token store.
// An empty store is not an error. Errors are returned if there are errors communicating
// with the token store.
// func (s *Store) List() ([]Token, error) {
func (s *Store) List() ([]string, error) {
	var keys []string

	list, err := s.kr.Keys()
	if err != nil {
		return []string{}, err
	}

	for _, i := range list {
		keys = append(keys, decodeVaultAddr(i))
	}
	return keys, nil
}

// TODO:
// func (s *Store) ListNew() ([]Token, error) {
// 	items, err := s.kr.List()
// 	if err != nil {
// 		return nil, err // TODO: maybe an empty slice instead
// 	}
// 	tokens := make([]Token, len(items))
// 	for idx, i := range items {
// 		tokens[idx] = Token{
// 			VaultAddr:    i.Key,
// 			createDate:   i.Created(),
// 			lastModified: i.LastModified(),
// 		}
// 	}
// 	return tokens, nil
// }

// Error erases the token for the vaultAddr from the token store. A missing token is not an error.
// Errors are returned if there are errors communicating with the token store.
func (s *Store) Erase(vaultAddr string) error {
	vaultAddr = encodeVaultAddr(vaultAddr)
	return s.kr.Remove(vaultAddr)
}

// AvailableBackends returns the available backends on this platform
func AvailableBackends() []string {
	backends := []string{}
	for _, i := range keyring.AvailableBackends() {
		for _, x := range SupportedBackends {
			if i == x {
				backends = append(backends, string(i))
			}
		}
	}
	return backends
}

func encodeVaultAddr(addr string) string {
	// step 1 - normalize url's to prevent duplicate entries, eg: downcase, remove redundant slashes, etc
	encoded, err := purell.NormalizeURLString(addr, purellFlags)
	if err != nil {
		return addr
	}

	// step 2 - some backends (ie: pass) store keyring items as files so the key (vault-addr)
	// must be a valid filename. We url-encode the vault-addr to support valid filenames yet
	// still being easy to read if someone were viewing the encoded names in the keychain.
	encoded = url.PathEscape(encoded)

	return encoded
}

func decodeVaultAddr(encoded string) string {
	decoded, _ := url.PathUnescape(encoded)
	return decoded
}
