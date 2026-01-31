package engine

import (
	"log"

	"github.com/irlan/quantumpay-go/internal/block"
	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

type Engine struct {
	chain   *blockchain.Blockchain
	mempool *mempool.Mempool
}

func New(chain *blockchain.Blockchain, mp *mempool.Mempool) *Engine {
	return &Engine{
		chain:   chain,
		mempool: mp,
	}
}

func (e *Engine) ProduceBlock() (*core.Block, error) {
	// 1. Ambil transaksi dari mempool
	pendingTxs := e.mempool.GetTransactions() 
	
	currentHeight := e.chain.Height()
	lastHash := e.chain.LastHash()

	builder := block.NewBuilder()
	
	// 2. Buat blok (Height, PrevHash, Txs)
	blk := builder.CreateBlock(currentHeight, lastHash, pendingTxs)

	if blk == nil {
		return nil, nil
	}

	// 3. Finalisasi Blok
	builder.MineBlock(blk)

	// 4. Simpan ke BadgerDB via Blockchain
	e.chain.AddBlock(blk)
	
	// 5. Audit Transaksi: Jika ada transaksi yang diproses, bersihkan mempool
	if len(pendingTxs) > 0 {
		e.mempool.Clear()
		log.Printf("✅ [ENGINE] Success! %d transactions secured in Block #%d", len(pendingTxs), blk.Header.Height)
	}

	return blk, nil
}
