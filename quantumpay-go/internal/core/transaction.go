package core

import (
	"crypto/sha256"
)

type Transaction struct {
	From  string
	To    string
	Value uint64
	Nonce uint64
}

func (tx *Transaction) Hash() []byte {
	h := sha256.New()

	h.Write([]byte(tx.From))
	h.Write([]byte(tx.To))
	h.Write(Uint64ToBytes(tx.Value))
	h.Write(Uint64ToBytes(tx.Nonce))

	return h.Sum(nil)
}
