package engine

import (
	"log"

	"github.com/irlan/quantumpay-go/internal/block"
	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

/*
   ===============================
   Engine (Consensus Core)
   ===============================
   Bertanggung jawab menghubungkan Blockchain, Mempool, dan Block Builder.
*/
type Engine struct {
	chain   *blockchain.Blockchain
	mempool *mempool.Mempool
}

// New Engine Constructor
func New(chain *blockchain.Blockchain, mp *mempool.Mempool) *Engine {
	return &Engine{
		chain:   chain,
		mempool: mp,
	}
}

/*
   ===============================
   Produce Block (Mining)
   ===============================
   Proses:
   1. Ambil data terbaru dari Chain (Height, LastHash).
   2. Ambil transaksi dari Mempool.
   3. Panggil Builder untuk menyusun blok.
   4. Lakukan Mining (Finalisasi).
   5. Simpan ke Chain.
*/
func (e *Engine) ProduceBlock() (*core.Block, error) {
	// 1. Siapkan Data Dasar
	// Height diambil dari panjang chain saat ini
	currentHeight := e.chain.Height() 
	lastHash := e.chain.LastHash()

	// 2. Ambil Transaksi dari Mempool
	// (Sementara kita kosongkan dulu atau ambil semua jika Mempool punya method GetAll)
	// TODO: Nanti sambungkan dengan logika e.mempool.SelectTransactions()
	txs := []*core.Transaction{} 

	// 3. Inisialisasi Builder (FIX ERROR: NewBuilder tidak butuh argumen lagi)
	builder := block.NewBuilder()

	// 4. Buat Proposal Blok (FIX ERROR: Gunakan CreateBlock dengan 3 argumen)
	// Urutan: prevHash, height, transactions
	blk := builder.CreateBlock(currentHeight, lastHash, txs)

	if blk == nil {
		return nil, nil // Gagal membuat blok
	}

	// 5. Lakukan Mining / Finalisasi (FIX ERROR: builder.Build undefined)
	// Kita ganti jadi MineBlock sesuai update file builder.go
	builder.MineBlock(blk)

	// 6. Simpan ke Blockchain (Persist ke BadgerDB via Chain)
	e.chain.AddBlock(blk)

	log.Printf("⚙️  [ENGINE] Block produced successfully. Height: %d | Hash: %x", blk.Header.Height, blk.Hash[:4])

	return blk, nil
}
