package tx

import (
	"crypto/sha256"
	"encoding/binary"
)

// HashTransaction computes deterministic transaction hash
func HashTransaction(t *Transaction) []byte {
	h := sha256.New()

	h.Write([]byte(t.From))
	h.Write([]byte(t.To))

	amountBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBuf, t.Amount)
	h.Write(amountBuf)

	nonceBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBuf, t.Nonce)
	h.Write(nonceBuf)

	return h.Sum(nil)
}
