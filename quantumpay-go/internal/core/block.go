package core

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
)

// Block adalah unit immutable di blockchain
type Block struct {
	Height       uint64
	PrevHash     []byte
	Transactions []*Transaction
	Timestamp    uint64
	Hash         []byte
}

// NewBlock membuat block baru dan menghitung hash
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
	b.Hash = b.computeHash()
	return b
}

// computeHash → deterministic block hash
func (b *Block) computeHash() []byte {
	h := sha256.New()

	// Height
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, b.Height)
	h.Write(buf)

	// Prev hash
	h.Write(b.PrevHash)

	// Transaction hashes
	for _, tx := range b.Transactions {
		h.Write(tx.Hash()) // ✅ FIX UTAMA
	}

	// Timestamp
	binary.BigEndian.PutUint64(buf, b.Timestamp)
	h.Write(buf)

	return h.Sum(nil)
}

// HashHex helper (RPC / debug)
func (b *Block) HashHex() string {
	return hex.EncodeToString(b.Hash)
}
