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
// Menggunakan BadgerDB State untuk pengecekan saldo
func ValidateBasic(tx *core.Transaction, ws *state.WorldState) error {
	// 1. Sanity Check
	if tx == nil {
		return errors.New("nil transaction")
	}

	// 2. Value Check (Tidak boleh kirim 0 atau negatif)
	if tx.Value == 0 {
		return ErrInvalidValue
	}

	// 3. Balance Check (Integrasi BadgerDB)
	// Kita ambil saldo langsung dari database disk
	currentBalance := ws.GetBalance(tx.From)

	// Cek apakah saldo cukup
	if currentBalance < tx.Value {
		return ErrInsufficientBalance
	}

	// Catatan: Verifikasi Tanda Tangan (PQC) dilakukan terpisah
	// biasanya di level Block Validation atau Mempool admission.
	
	return nil
}
