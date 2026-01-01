package tx

import (
	"errors"

	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/state"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidValue        = errors.New("invalid transaction value")
)

// ValidateBasic melakukan validasi dasar transaksi
// ❗ TANPA nonce & fee (biar stabil dulu)
func ValidateBasic(tx *core.Transaction, ws *state.WorldState) error {
	if tx == nil {
		return errors.New("nil transaction")
	}

	if tx.Value == 0 {
		return ErrInvalidValue
	}

	from := ws.GetAccount(tx.From)
	if from.Balance < tx.Value {
		return ErrInsufficientBalance
	}

	return nil
}
