package engine

import (
	"sync"

	"github.com/irlan/quantumpay-go/internal/block"
	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

type Engine struct {
	mu    sync.Mutex
	chain *blockchain.Blockchain
	mp    *mempool.Mempool
}

func New(chain *blockchain.Blockchain, mp *mempool.Mempool) *Engine {
	return &Engine{
		chain: chain,
		mp:    mp,
	}
}

// ProduceBlock builds a new block from mempool and appends it to chain
func (e *Engine) ProduceBlock() (*core.Block, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// ✅ Adapter: Blockchain → ChainView
	view := blockchain.NewView(e.chain)

	// ✅ Builder only sees ChainView + Mempool
	builder := block.NewBuilder(view, e.mp)

	blk := builder.Build()
	if blk == nil {
		return nil, nil
	}

	// Append block (blockchain owns mutation)
	e.chain.AddBlock(blk)

	return blk, nil
}
