package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/irlan/quantumpay-go/internal/tx"
)

type Block struct {
	Index        uint64
	Transactions []tx.Transaction
	PrevHash     string
	Hash         string
}

func (b *Block) CalculateHash() string {
	data := strconv.FormatUint(b.Index, 10) + b.PrevHash

	for _, t := range b.Transactions {
		data += t.From +
			t.To +
			strconv.FormatUint(t.Amount, 10) +
			strconv.FormatUint(t.Nonce, 10)
	}

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
