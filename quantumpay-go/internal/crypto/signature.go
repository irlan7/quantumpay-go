package crypto

type SignatureVerifier interface {
	Verify(message []byte, sig []byte, pub []byte) bool
}
