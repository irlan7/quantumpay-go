package p2pv1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"errors"
	"math/big"
)

/*
Vote signature verification
- ECDSA for now
- upgradeable to PQC later
*/

func VerifyVoteSignature(v Vote, pub []byte) error {
	if len(pub) == 0 || len(v.Signature) == 0 {
		return errors.New("missing key or signature")
	}

	x, y := elliptic.Unmarshal(elliptic.P256(), pub)
	if x == nil {
		return errors.New("invalid public key")
	}

	pubKey := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	h := sha256.Sum256([]byte(v.Hash))
	r := new(big.Int).SetBytes(v.Signature[:32])
	s := new(big.Int).SetBytes(v.Signature[32:])

	if !ecdsa.Verify(pubKey, h[:], r, s) {
		return errors.New("invalid vote signature")
	}
	return nil
}
