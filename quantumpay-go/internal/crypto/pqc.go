package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/cloudflare/circl/sign/mldsa/mldsa65"
	"golang.org/x/crypto/hkdf"
)

// SovereignSignature menyimpan bukti tanda tangan hibrida v4.5
type SovereignSignature struct {
	ClassicSig []byte `json:"classic_sig"`
	QuantumSig []byte `json:"quantum_sig"`
}

// SignHybrid: Audit Perbaikan Error Sign (Argument io.Reader)
func SignHybrid(msg []byte, privClassic ed25519.PrivateKey, privPQ *mldsa65.PrivateKey) (SovereignSignature, error) {
	// 1. Classic: Ed25519
	sigClassic := ed25519.Sign(privClassic, msg)

	// 2. Quantum: ML-DSA-65
	// FIX: Sign membutuhkan (rand, msg, opts). Kita pakai rand.Reader.
	sigPQ, err := privPQ.Sign(rand.Reader, msg, nil)
	if err != nil {
		return SovereignSignature{}, fmt.Errorf("mldsa sign failed: %w", err)
	}

	return SovereignSignature{
		ClassicSig: sigClassic,
		QuantumSig: sigPQ,
	}, nil
}

// VerifyHybrid: Audit Perbaikan Error Verify (Standalone Function)
func VerifyHybrid(msg []byte, sig SovereignSignature, pubClassic ed25519.PublicKey, pubPQ *mldsa65.PublicKey) bool {
	// Verifikasi Ed25519
	if !ed25519.Verify(pubClassic, msg, sig.ClassicSig) {
		return false
	}

	// Verifikasi ML-DSA-65
	// FIX: Gunakan mldsa65.Verify, bukan pubPQ.Verify
	return mldsa65.Verify(pubPQ, msg, nil, sig.QuantumSig)
}

// DeriveHybridKey: Mengolah kunci dengan HKDF (Standar 10/10)
func DeriveHybridKey(classicSS, pqSS []byte) ([]byte, error) {
	info := []byte("QuantumPay-v4.5|Shield-KDF-AES256")
	combined := append(classicSS, pqSS...)
	
	h := hkdf.New(sha256.New, combined, nil, info)
	key := make([]byte, 32)
	if _, err := io.ReadFull(h, key); err != nil {
		return nil, fmt.Errorf("hkdf failed: %w", err)
	}
	return key, nil
}

// EncryptPayload: Enkripsi AES-GCM (Production Ready)
func EncryptPayload(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
