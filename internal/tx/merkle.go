package tx

import (
	"crypto/sha256"
	"encoding/hex"
)

func ComputeTxRoot(txs []Transaction) string {
	h := sha256.New()
	for _, tx := range txs {
		h.Write([]byte(tx.From))
		h.Write([]byte(tx.To))
	}
	return hex.EncodeToString(h.Sum(nil))
}
