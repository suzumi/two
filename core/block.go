package core

type Block struct {
	Index        uint64
	Timestamp    string
	MerkleRoot   uint64
	PrevHash     string
	Nonce        uint64
	Transactions []*Transaction
}
