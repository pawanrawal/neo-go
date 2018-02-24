package wallet

type Account struct {
	wallet  Wallet
	nep2key string
	//	key keypair
}

func (a *Account) Decrypted() bool {
	// TODO - Implement
	return true
}

func (a *Account) HasKey() bool {
	return len(a.nep2key) > 0
}

func (a *Account) VerifyPassword(password string) bool {
	// TODO - See what has to be done here.
	return true
}

func (a *Account) GetKey(password string) bool {
	// TODO - Make this return the actual key after you figure out DS for it.
	if len(a.nep2key) == 0 {
		return false
	}

	// TODO - GetPrivateKeyFromNEP2
}
