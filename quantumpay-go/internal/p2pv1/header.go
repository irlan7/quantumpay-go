package p2pv1

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
HeaderMsg = metadata block
TIDAK mengandung logic state / execution / p2p (ANTI-CYCLE)

Dipakai oleh:
- header_pool.go (fork-choice)
- handler.go (validation)
- node.go (commit hasil execution)
*/

// ===============================
// Header structure
// ===============================

type HeaderMsg struct {
	Hash       string `json:"hash"`
	ParentHash string `json:"parent_hash"`
	Height     uint64 `json:"height"`
	Timestamp  int64  `json:"timestamp"`

	// PATCH: state root hash (deterministic)
	StateRoot []byte `json:"state_root"`
}

// ===============================
// Header hashing
// ===============================

// ComputeHash menghitung hash header (tanpa signature)
// Dipakai saat membuat block / header baru
func (h HeaderMsg) ComputeHash() string {
	hasher := sha256.New()

	hasher.Write([]byte(h.ParentHash))
	hasher.Write(uint64ToBytes(h.Height))
	hasher.Write(int64ToBytes(h.Timestamp))

	// include state root (jika ada)
	if len(h.StateRoot) > 0 {
		hasher.Write(h.StateRoot)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

// ===============================
// Helper (local, no import cycle)
// ===============================

func uint64ToBytes(v uint64) []byte {
	var b [8]byte
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b[:]
}

func int64ToBytes(v int64) []byte {
	var b [8]byte
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b[:]
}
