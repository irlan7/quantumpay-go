package p2pv1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"errors"
)

// SignatureECDSA is ASN.1 compatible
type SignatureECDSA struct {
	R, S []byte
}

// GenerateKeyPair generates ECDSA P-256 keypair
func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

// SignMessage signs arbitrary message bytes
func SignMessage(priv *ecdsa.PrivateKey, msg []byte) ([]byte, error) {
	if priv == nil {
		return nil, errors.New("nil private key")
	}

	hash := sha256.Sum256(msg)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return nil, err
	}

	return asn1.Marshal(SignatureECDSA{
		R: r.Bytes(),
		S: s.Bytes(),
	})
}

// VerifyMessage verifies ECDSA signature
func VerifyMessage(pub *ecdsa.PublicKey, msg []byte, sig []byte) bool {
	if pub == nil {
		return false
	}

	var parsed SignatureECDSA
	if _, err := asn1.Unmarshal(sig, &parsed); err != nil {
		return false
	}

	hash := sha256.Sum256(msg)
	r := new(bigInt).SetBytes(parsed.R)
	s := new(bigInt).SetBytes(parsed.S)

	return ecdsa.Verify(pub, hash[:], r, s)
}

/*
We wrap big.Int to avoid importing math/big in other files
and reduce accidental import cycles
*/
type bigInt struct {
	bytes []byte
}

func (b *bigInt) SetBytes(v []byte) *bigInt {
	b.bytes = v
	return b
}

func (b *bigInt) Bytes() []byte {
	return b.bytes
}
