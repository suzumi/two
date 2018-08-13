package wallet

import (
	"os"
	"io"
	"encoding/json"
	"fmt"
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
	wlt.Add(account)
	fmt.Printf("Wallet: %s\n", wlt)
	return wlt.Save()
}

func (wlt *Wallet) Save() error {
	return json.NewEncoder(wlt.rw).Encode(wlt)
}

func (wlt *Wallet) Add(account *Account) {
	wlt.Account = *account
}
