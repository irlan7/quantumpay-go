package crypto

import (
	"crypto/ed25519"
	"testing"
	"github.com/cloudflare/circl/sign/mldsa/mldsa65"
)

func TestQuantumSovereignShield(t *testing.T) {
	msg := []byte("Auth: Mint 500 qBNB for Sovereign Owner")

	// 1. Generate Keypairs
	// FIX: mldsa65.GenerateKey membutuhkan parameter (rand). nil = default rand.
	pubC, privC, _ := ed25519.GenerateKey(nil)
	pubQ, privQ, _ := mldsa65.GenerateKey(nil)

	// 2. Sign
	sig, err := SignHybrid(msg, privC, privQ)
	if err != nil {
		t.Fatalf("Signing Error: %v", err)
	}

	// 3. Verify
	if !VerifyHybrid(msg, sig, pubC, pubQ) {
		t.Fatal("❌ Security Failure: Valid signature was rejected!")
	}

	// 4. Test Integrity (Pesan Palsu)
	fakeMsg := []byte("Auth: Mint 999999 qBNB for Hacker")
	if VerifyHybrid(fakeMsg, sig, pubC, pubQ) {
		t.Fatal("❌ Critical Security Leak: Fake message accepted!")
	}

	t.Log("✅ PASS: QuantumPay v4.5 Hybrid Shield is Active & Verified.")
}
