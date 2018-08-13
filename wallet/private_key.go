package wallet

import (
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"io"
)

type PrivateKey struct {
	privKey []byte
}

func NewPrivateKey(passPhrase string) *PrivateKey {
	//p, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	s := sha256.New()
	io.WriteString(s, passPhrase)
	return &PrivateKey{s.Sum(nil)}
}

func (p *PrivateKey) PublicKey() []byte {
	s := sha256.New()
	s.Write(p.privKey)
	return s.Sum(nil)
}

func (p *PrivateKey) Address() string {
	addr := base58.Encode(p.PublicKey())
	return addr

	//hashed := []byte("")
	//r, s, err := ecdsa.Sign(rand.Reader, p.privKey, hashed)
	//if err != nil {
	//	fmt.Printf("Err: %s\n", err)
	//	return "", err
	//}
	//return r.Bytes()
}
