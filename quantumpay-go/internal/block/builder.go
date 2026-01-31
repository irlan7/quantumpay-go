package block

import (
	"fmt" // Wajib ada untuk Printf di MineBlock

	"github.com/irlan/quantumpay-go/internal/core"
)

// Builder bertugas menyusun blok (Mining candidate)
type Builder struct {
	// Nanti bisa tambah konfigurasi miner address disini
}

// NewBuilder membuat instance builder baru
func NewBuilder() *Builder {
	return &Builder{}
}

// CreateBlock menyusun proposal blok baru
// NOTE: Argumen input disesuaikan dengan kebutuhan (Height, PrevHash, Txs)
func (b *Builder) CreateBlock(height uint64, prevHash []byte, txs []*core.Transaction) *core.Block {
	
	// FIX ERROR: "Too many arguments" & "Type mismatch"
	// Urutan yang benar sesuai core/block.go adalah:
	// 1. prevHash ([]byte)
	// 2. height (uint64)
	// 3. txs ([]*core.Transaction)
	
	blk := core.NewBlock(prevHash, height, txs)

	return blk
}

// MineBlock simulasi Proof of Work / Finalisasi blok
func (b *Builder) MineBlock(blk *core.Block) {
	// Gunakan fmt agar log mining terlihat di terminal VPS
	fmt.Printf("⛏️  Mining Block #%d [Tx: %d]...\n", blk.Header.Height, len(blk.Transactions))
	
	// Logika Mining Sederhana (Sovereign Authority)
	// Di Mainnet nanti, kita bisa ganti ini dengan PoW atau PoS.
	// Untuk sekarang, kita set Nonce dan hitung Hash final.
	
	blk.Header.Nonce = 0 // Reset nonce
	
	// Hitung Hash Final agar valid di Blockchain
	blk.Hash = blk.CalculateHash()
	
	fmt.Printf("✅ Block Sealed: %x\n", blk.Hash[:8])
}
