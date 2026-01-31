package core

// Transaction didefinisikan HANYA DI SINI (Single Source of Truth)
type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Value  uint64 `json:"value"`
	Nonce  uint64 `json:"nonce"`
	Sig    []byte `json:"sig,omitempty"`    // Tanda tangan kriptografi
	Pubkey []byte `json:"pubkey,omitempty"` // Kunci publik pengirim
}

// NewTransaction constructor helper (opsional, tapi bagus untuk kerapihan)
func NewTransaction(from, to string, value, nonce uint64) *Transaction {
	return &Transaction{
		From:  from,
		To:    to,
		Value: value,
		Nonce: nonce,
	}
}
