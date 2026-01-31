package tx

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Transaction struct {
	From   string
	To     string
	Amount uint64
	Nonce  uint64
	Fee    uint64
}

func (t *Transaction) Hash() string {
	data := t.From + t.To +
		strconv.FormatUint(t.Amount, 10) +
		strconv.FormatUint(t.Nonce, 10) +
		strconv.FormatUint(t.Fee, 10)

	h := sha256.Sum256([]byte(data))
	return hex.EncodeToString(h[:])
}
