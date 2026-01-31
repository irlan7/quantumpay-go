package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	// "crypto/rand" <--- Baris ini SUDAH SAYA HAPUS agar tidak error
	"crypto/sha256"
	"encoding/hex"
	"errors"

	// Library 12 Kata (Pastikan sudah 'go get' di terminal)
	"github.com/tyler-smith/go-bip39" 
)

// KeyPair menyimpan Kunci, Alamat, dan Mnemonic (Kata Rahasia)
type KeyPair struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Mnemonic   string 
}

// OPSI 1: Buat Dompet Baru (Generate New)
func NewKeyPair() (*KeyPair, error) {
	// 1. Buat Entropy 128 bit (Sumber Acak untuk 12 Kata)
	entropy, err := bip39.NewEntropy(128) 
	if err != nil {
		return nil, err
	}

	// 2. Ubah jadi Mnemonic (Kalimat Rahasia)
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	// 3. Buat Kunci dari Mnemonic tersebut
	return NewKeyPairFromMnemonic(mnemonic)
}

// OPSI 2: Pulihkan Dompet Lama (Import from Mnemonic)
func NewKeyPairFromMnemonic(mnemonic string) (*KeyPair, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("mnemonic tidak valid")
	}

	// Generate Seed dari Mnemonic
	seed := bip39.NewSeed(mnemonic, "") // Password kosong ("") standard BIP39

	// Gunakan Seed sebagai sumber acak untuk membuat Private Key
	reader := NewSeedReader(seed) 
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, reader)
	if err != nil {
		return nil, err
	}

	// Gabungkan koordinat X & Y untuk Public Key
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return &KeyPair{
		PrivateKey: private,
		PublicKey:  pubKey,
		Mnemonic:   mnemonic,
	}, nil
}

// --- HELPER KHUSUS: SEED READER ---
// Ini trik agar seed bisa dibaca oleh ecdsa.GenerateKey sebagai stream
type SeedReader struct {
	seed []byte
	pos  int
}

func NewSeedReader(seed []byte) *SeedReader {
	return &SeedReader{seed: seed}
}

func (r *SeedReader) Read(p []byte) (n int, err error) {
	// Salin seed ke buffer p secara berulang (Deterministik)
	count := copy(p, r.seed[r.pos:])
	
	if count < len(p) {
		// Jika seed habis dibaca, reset posisi (Looping)
		// Dalam produksi asli biasanya pakai HKDF/BIP32, tapi ini cukup untuk devnet
		r.pos = 0 
		return count, nil 
	}
	r.pos += count
	return count, nil
}

// --- FUNGSI STANDAR (Address & Hex) ---

func (w *KeyPair) Address() string {
	hash := sha256.Sum256(w.PublicKey)
	addressBytes := hash[len(hash)-20:]
	return "0x" + hex.EncodeToString(addressBytes)
}

func (w *KeyPair) PrivateKeyHex() string {
	return hex.EncodeToString(w.PrivateKey.D.Bytes())
}

func (w *KeyPair) GetPublicKeyHex() string {
	return hex.EncodeToString(w.PublicKey)
}
