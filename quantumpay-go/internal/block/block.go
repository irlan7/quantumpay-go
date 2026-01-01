package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/irlan/quantumpay-go/internal/tx"
)

type BlockHeader struct {
	Height     uint64
	PrevHash   string
	StateRoot  string // placeholder (akan diisi nanti)
	TxRoot     string // placeholder
	Proposer   string
}

type Block struct {
	Header       BlockHeader
	Transactions []*tx.Transaction
	Hash         string
}

func (b *Block) CalculateHash() string {
	data :=
		strconv.FormatUint(b.Header.Height, 10) +
			b.Header.PrevHash +
			b.Header.StateRoot +
			b.Header.TxRoot +
			b.Header.Proposer

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
