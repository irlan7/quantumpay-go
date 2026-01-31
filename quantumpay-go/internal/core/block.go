package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// BlockHeader menyimpan metadata blok
type BlockHeader struct {
	Timestamp int64  `json:"timestamp"`
	PrevHash  []byte `json:"prev_hash"`
	Height    uint64 `json:"height"`
	Nonce     uint64 `json:"nonce"`
}

// Block menyimpan data transaksi dan header
type Block struct {
	Header       BlockHeader    `json:"header"`
	Transactions []*Transaction `json:"transactions"` // Mengacu ke struct di transaction.go
	Hash         []byte         `json:"hash"`
}

// NewBlock membuat instance blok baru
func NewBlock(prevHash []byte, height uint64, txs []*Transaction) *Block {
	block := &Block{
		Header: BlockHeader{
			Timestamp: time.Now().Unix(),
			PrevHash:  prevHash,
			Height:    height,
		},
		Transactions: txs,
	}
	block.Hash = block.CalculateHash()
	return block
}

// CalculateHash menghasilkan hash SHA-256 dari header dan transaksi
func (b *Block) CalculateHash() []byte {
	// Gabungkan data Header
	record := fmt.Sprintf("%d%x%d%d", 
		b.Header.Timestamp, 
		b.Header.PrevHash, 
		b.Header.Height, 
		b.Header.Nonce,
	)

	// Gabungkan ID Transaksi (Tx Hash)
	for _, tx := range b.Transactions {
		// Asumsi Transaction punya method Hash() atau kita pakai fieldnya
		record += fmt.Sprintf("%s%s%d%d", tx.From, tx.To, tx.Value, tx.Nonce)
	}

	h := sha256.New()
	h.Write([]byte(record))
	return h.Sum(nil)
}

// HashString helper untuk melihat hash dalam bentuk hex
func (b *Block) HashString() string {
	return hex.EncodeToString(b.Hash)
}
