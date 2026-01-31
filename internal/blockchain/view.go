package blockchain

import (
	"github.com/irlan/quantumpay-go/internal/core"
)

// view adalah adapter read-only dari Blockchain
// Adapter ini aman digunakan oleh package lain tanpa import cycle
type view struct {
	chain *Blockchain
}

// NewView membuat ChainView baru
func NewView(chain *Blockchain) *view {
	return &view{chain: chain}
}

// Height mengembalikan tinggi chain saat ini
func (v *view) Height() uint64 {
	return v.chain.Height()
}

// LastHash mengembalikan hash block terakhir
func (v *view) LastHash() []byte {
	return v.chain.LastHash()
}

// GetBlockByHeight mengembalikan block berdasarkan nomor urut (height)
// FIX: Implementasi langsung akses slice internal dengan Thread Safety
func (v *view) GetBlockByHeight(h uint64) *core.Block {
	// 1. Kunci Mutex Read-Only (Karena kita ada di package yang sama, kita bisa akses mu)
	v.chain.mu.RLock()
	defer v.chain.mu.RUnlock()

	// 2. Cek apakah Height valid (tidak melebihi jumlah blok yang ada)
	// Ingat: Slice index dimulai dari 0. Jika len=10, index max=9.
	if h >= uint64(len(v.chain.blocks)) {
		return nil // Blok belum ada
	}

	// 3. Kembalikan pointer ke blok yang diminta
	return v.chain.blocks[h]
}
