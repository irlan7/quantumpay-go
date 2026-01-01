package crypto

import (
	"crypto/sha256"
)

func hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}

// BuildMerkleRoot membangun merkle root dari list hash
func BuildMerkleRoot(leaves [][]byte) []byte {
	if len(leaves) == 0 {
		return nil
	}

	// copy agar immutable
	level := make([][]byte, len(leaves))
	copy(level, leaves)

	for len(level) > 1 {
		var next [][]byte

		for i := 0; i < len(level); i += 2 {
			if i+1 < len(level) {
				combined := append(level[i], level[i+1]...)
				next = append(next, hash(combined))
			} else {
				// duplicate last if odd
				next = append(next, hash(level[i]))
			}
		}

		level = next
	}

	return level[0]
}
