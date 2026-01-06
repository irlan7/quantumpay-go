package state

import (
	"crypto/sha256"
	"encoding/binary"
	"sort"
)

// ComputeStateRoot menghitung hash deterministik dari seluruh account state
func (ws *WorldState) ComputeStateRoot() []byte {
	// Ambil semua address lalu sort agar deterministik
	keys := make([]string, 0, len(ws.accounts))
	for addr := range ws.accounts {
		keys = append(keys, addr)
	}
	sort.Strings(keys)

	h := sha256.New()

	for _, addr := range keys {
		acc := ws.accounts[addr]

		h.Write([]byte(addr))
		h.Write(uint64ToBytes(acc.Balance))
		h.Write(uint64ToBytes(acc.Nonce))
	}

	return h.Sum(nil)
}

// helper lokal — sengaja di sini agar TIDAK import core
func uint64ToBytes(v uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, v)
	return buf
}
