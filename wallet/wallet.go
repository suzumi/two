package wallet

import (
	"os"
	"encoding/json"
	"io"
)

const Version = "1.0"

type Wallet struct {
	Version string  `json:"version"`
	Account Account `json:"account"`
	path    string
	rw      io.ReadWriter
}

func NewWallet(location string) (*Wallet, error) {
	file, err := os.Create(location)
	if err != nil {
		return nil, err
	}
	return createWallet(file), nil
}

func createWallet(rw io.ReadWriter) *Wallet {
	var path string
	if f, ok := rw.(*os.File); ok {
		path = f.Name()
	}
	return &Wallet{
		Version: Version,
		Account: Account{},
		path:    path,
		rw:      rw,
	}
}

func (wlt *Wallet) CreateAccount(passPhrase string) error {
	account := NewAccountFromPrivateKey(passPhrase)
	wlt.AddAccount(account)
	return wlt.Save()
}

func (wlt *Wallet) Save() error {
	return json.NewEncoder(wlt.rw).Encode(wlt)
}

func (wlt *Wallet) AddAccount(account *Account) {
	wlt.Account = *account
}

func (wlt *Wallet) JSON() ([]byte, error) {
	return json.MarshalIndent(wlt, " ", " ")
}
