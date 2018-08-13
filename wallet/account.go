package wallet

type (
	Account struct {
		privateKey *PrivateKey
		publicKey  []byte
		address    string
	}
)

func NewAccountFromPrivateKey(passPhrase string) *Account {
	p := NewPrivateKey(passPhrase)
	pubKey := p.PublicKey()
	addr := p.Address()

	return &Account{
		privateKey: p,
		publicKey:  pubKey,
		address:    addr,
	}

}
