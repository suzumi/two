package wallet

import (
	"os"
	"fmt"
	"encoding/json"
	"io"
)

const Version = "1.0"

type Wallet struct {
	Version string
	Account Account
	Path    string
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
		Path:    path,
		rw:      rw,
	}
}

func (wlt *Wallet) CreateAccount(passPhrase string) error {
	account := NewAccountFromPrivateKey(passPhrase)
	fmt.Printf("Save Before wallet Account: %s\n", wlt.Account)
	wlt.AddAccount(account)
	fmt.Printf("Save After wallet Account: %s\n", wlt.Account)
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
