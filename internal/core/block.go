package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type BlockHeader struct {
	Timestamp int64  `json:"timestamp"`
	PrevHash  []byte `json:"prev_hash"`
	Height    uint64 `json:"height"`
	Nonce     uint64 `json:"nonce"`
}

type Block struct {
	Header       BlockHeader    `json:"header"`
	Transactions []*Transaction `json:"transactions"`
	Hash         []byte         `json:"hash"`
}

func NewBlock(prevHash []byte, height uint64, txs []*Transaction) *Block {
	blk := &Block{
		Header: BlockHeader{
			Timestamp: time.Now().Unix(),
			PrevHash:  prevHash,
			Height:    height,
		},
		Transactions: txs,
	}
	blk.Hash = blk.CalculateHash()
	return blk
}

func (b *Block) CalculateHash() []byte {
	record := fmt.Sprintf("%d%x%d%d", 
		b.Header.Timestamp, 
		b.Header.PrevHash, 
		b.Header.Height, 
		b.Header.Nonce,
	)
	for _, tx := range b.Transactions {
		record += fmt.Sprintf("%x", tx.Hash())
	}
	h := sha256.New()
	h.Write([]byte(record))
	return h.Sum(nil)
}

func (b *Block) ValidateTransactions() error {
	for i, tx := range b.Transactions {
		// Memanggil VerifyPQC yang sudah aman di transaction.go
		if err := tx.VerifyPQC(); err != nil {
			return fmt.Errorf("invalid tx at index %d: %v", i, err)
		}
	}
	return nil
}

func (b *Block) HashString() string {
	return hex.EncodeToString(b.Hash)
}
