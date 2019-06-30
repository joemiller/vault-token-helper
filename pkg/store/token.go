package store

import "time"

// Token represents a Vault token stored in a backend credential store
type Token struct {
	VaultAddr    string
	Token        string
	createDate   time.Time
	lastModified time.Time
}

func (t Token) Created() time.Time {
	return t.createDate
}

func (t Token) LastModified() time.Time {
	return t.lastModified
}
