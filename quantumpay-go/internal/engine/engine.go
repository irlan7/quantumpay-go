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
   ChainView Adapter (ANTI CYCLE)
   ===============================
   Mengadaptasi *blockchain.Blockchain
   -> block.ChainView
*/
type chainViewAdapter struct {
	chain *blockchain.Blockchain
}

func (v *chainViewAdapter) Height() uint64 {
	return v.chain.Height()
}

func (v *chainViewAdapter) LastHash() []byte {
	return v.chain.LastHash()
}

func (v *chainViewAdapter) GetBlockByHeight(h uint64) *core.Block {
	blk, err := v.chain.GetBlockByHeight(h)
	if err != nil {
		return nil
	}
	return blk
}

/*
   ===============================
   Engine
   ===============================
*/
type Engine struct {
	chain   *blockchain.Blockchain
	mempool *mempool.Mempool
}

func New(
	chain *blockchain.Blockchain,
	mempool *mempool.Mempool,
) *Engine {
	return &Engine{
		chain:   chain,
		mempool: mempool,
	}
}

/*
   ===============================
   Produce Block
   ===============================
*/
func (e *Engine) ProduceBlock() (*core.Block, error) {
	view := &chainViewAdapter{chain: e.chain}

	builder := block.NewBuilder(
		view,
		e.mempool,
	)

	blk := builder.Build()

	if blk == nil {
		return nil, nil
	}

	// Add block ke chain (void function)
	e.chain.AddBlock(blk)

	log.Printf("[ENGINE] Block produced height=%d", e.chain.Height())

	return blk, nil
}
