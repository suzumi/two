package wallet

type (
	Account struct {
		privateKey *PrivateKey
		PublicKey  []byte `json:"public_key"`
		Address    string `json:"address"`
	}
)

func NewAccountFromPrivateKey(passPhrase string) *Account {
	p := NewPrivateKey(passPhrase)
	pubKey := p.PublicKey()
	addr := p.Address()

	return &Account{
		privateKey: p,
		PublicKey:  pubKey,
		Address:    addr,
	}

}
