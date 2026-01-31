package crypto

import (
	"crypto/rand"
	"errors"

	// Sub-package spesifik Mode 3
	"github.com/cloudflare/circl/sign/dilithium/mode3"
)

// SignMessage: Wrapper untuk menandatangani pesan (bytes) menggunakan Private Key (bytes)
func SignMessage(msg []byte, privateKeyBytes []byte) ([]byte, error) {
	// 1. Validasi Input Ukuran
	if len(privateKeyBytes) != mode3.PrivateKeySize {
		return nil, errors.New("ukuran private key tidak valid")
	}

	// 2. Load Private Key (Gunakan UnmarshalBinary)
	sk := new(mode3.PrivateKey)
	if err := sk.UnmarshalBinary(privateKeyBytes); err != nil {
		return nil, errors.New("gagal memuat private key: " + err.Error())
	}

	// 3. Lakukan Signing
	// Kita siapkan wadah signature sesuai ukuran standar Mode 3
	signature := make([]byte, mode3.SignatureSize)
	
	// Gunakan SignTo (API Resmi CIRCL)
	mode3.SignTo(sk, msg, signature)
	
	return signature, nil
}

// VerifySignature: Wrapper untuk verifikasi (PENTING untuk Node)
func VerifySignature(pubKeyBytes []byte, msg []byte, signature []byte) bool {
	// 1. Validasi Input Ukuran
	if len(pubKeyBytes) != mode3.PublicKeySize {
		return false
	}
	if len(signature) != mode3.SignatureSize {
		return false
	}

	// 2. Load Public Key (Gunakan UnmarshalBinary)
	pk := new(mode3.PublicKey)
	if err := pk.UnmarshalBinary(pubKeyBytes); err != nil {
		return false
	}

	// 3. Verifikasi menggunakan API mode3
	return mode3.Verify(pk, msg, signature)
}

// GenerateKeyPair: Membuat kunci baru (untuk Wallet)
func GenerateKeyPair() ([]byte, []byte, error) {
	// 1. Generate Key Objects
	pk, sk, err := mode3.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	
	// 2. Konversi ke Bytes (MarshalBinary)
	// Kita abaikan error di sini karena key yang baru digenerate pasti valid
	pkBytes, _ := pk.MarshalBinary()
	skBytes, _ := sk.MarshalBinary()

	return pkBytes, skBytes, nil
}
