package wallet

import (
	"crypto/rand"
	"errors"
)

const (
	version = "1.0"
)

type Wallet struct {
	path string

	password string
	name     string
	// Is this a string?
	version  string
	scrypt   ScryptParameters
	accounts map[[20]byte]Account
	extra    map[string]interface{}
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
// TODO - Move to a util package.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Create(path string, name string) *Wallet {
	return &Wallet{
		path:     path,
		name:     name,
		version:  version,
		accounts: make(map[[20]byte]Account),
		extra:    make(map[string]interface{}),
	}
}

func (w *Wallet) verifyPassword(password string) bool {
	var account *Account

	for _, a := range w.accounts {
		if !a.Decrypted() {
			account = a
			break
		}
	}

	if account == nil {
		for _, a := range w.accounts {
			if a.HasKey() {
				account = a
				break
			}

		}
	}

	if account == nil {
		return true
	}

	if account.Decrypted() {
		return account.VerifyPassword(w.password)
	}

	if account.GetKey() {
		return true
	}

	return false
}

func (w *Wallet) Unlock(password string) error {
	if !w.verifyPassword(password) {
		return errors.New("Password wrong")
	}
	w.password = password
	return nil
}

func (w *Wallet) GetPrivateKeyFromNEP2(nep2key string, password string) error {
	// TODO - Get this from neo-go-sdk
}
