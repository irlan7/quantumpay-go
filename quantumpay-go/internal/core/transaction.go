package core

import (
	"crypto/sha256"
	"encoding/hex"
)

type Transaction struct {
	From  string
	To    string
	Value uint64
	Fee   uint64
	Nonce uint64
}

func (tx *Transaction) Hash() []byte {
	h := sha256.New()
	h.Write([]byte(tx.From))
	h.Write([]byte(tx.To))
	h.Write(uint64ToBytes(tx.Value))
	h.Write(uint64ToBytes(tx.Fee))
	h.Write(uint64ToBytes(tx.Nonce))
	return h.Sum(nil)
}

func (tx *Transaction) HashHex() string {
	return hex.EncodeToString(tx.Hash())
}

func uint64ToBytes(v uint64) []byte {
	b := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		b[7-i] = byte(v >> (i * 8))
	}
	return b
}
