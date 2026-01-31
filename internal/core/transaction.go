package core

import (
	"crypto/sha256"
	"errors"
	"fmt"

	// Core HANYA mengimpor layer abstraksi crypto kita
	"github.com/irlan/quantumpay-go/internal/crypto"
)

var (
	ErrInvalidSignature = errors.New("pqc signature verification failed")
)

// Transaction: Struktur data murni
type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Value  uint64 `json:"value"`
	Nonce  uint64 `json:"nonce"`
	Sig    []byte `json:"sig,omitempty"`
	Pubkey []byte `json:"pubkey,omitempty"`
}

// Hash: ID unik transaksi
func (tx *Transaction) Hash() []byte {
	record := fmt.Sprintf("%s%s%d%d", tx.From, tx.To, tx.Value, tx.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return h.Sum(nil)
}

// Sign: Sekarang menerima []byte, bukan objek library yang rumit
func (tx *Transaction) Sign(privateKeyBytes []byte) error {
	msg := tx.Hash()
	
	// Panggil layer crypto yang sudah kita perbaiki di Langkah 1
	sig, err := crypto.SignMessage(msg, privateKeyBytes)
	if err != nil {
		return err
	}
	
	tx.Sig = sig
	// Catatan: Pubkey idealnya diset saat inisialisasi Wallet, 
	// tapi di sini kita asumsikan sudah ada atau diset manual.
	return nil
}

// VerifyPQC: Validasi aman tanpa menyentuh library Dilithium langsung
func (tx *Transaction) VerifyPQC() error {
	if len(tx.Sig) == 0 || len(tx.Pubkey) == 0 {
		return errors.New("missing signature or public key")
	}

	// Delegasikan verifikasi ke pakar kriptografi (package crypto)
	valid := crypto.VerifySignature(tx.Pubkey, tx.Hash(), tx.Sig)
	if !valid {
		return ErrInvalidSignature
	}

	return nil
}
