package core

import (
	"crypto/sha256"
	"encoding/binary"
)

// Block merepresentasikan blok immutable di QuantumPay
type Block struct {
	Height       uint64
	PrevHash     []byte
	Transactions []*Transaction
	Timestamp    uint64
	Hash         []byte
}

// NewBlock membuat block baru dan langsung menghitung hash
func NewBlock(
	height uint64,
	prevHash []byte,
	txs []*Transaction,
	timestamp uint64,
) *Block {

	b := &Block{
		Height:       height,
		PrevHash:     prevHash,
		Transactions: txs,
		Timestamp:    timestamp,
	}

	b.Hash = b.calculateHash()
	return b
}

// calculateHash menghitung hash block secara deterministik
func (b *Block) calculateHash() []byte {
	h := sha256.New()

	// Height
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, b.Height)
	h.Write(buf)

	// PrevHash
	h.Write(b.PrevHash)

	// Timestamp
	binary.BigEndian.PutUint64(buf, b.Timestamp)
	h.Write(buf)

	// Transactions
	for _, tx := range b.Transactions {
		h.Write(tx.Hash())
	}

	return h.Sum(nil)
}
