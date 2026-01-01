package blockchain

import "github.com/irlan/quantumpay-go/internal/core"

type Blockchain struct {
	blocks []*core.Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: make([]*core.Block, 0),
	}
}

func (bc *Blockchain) AddBlock(block *core.Block) {
	bc.blocks = append(bc.blocks, block)
}

func (bc *Blockchain) Height() uint64 {
	return uint64(len(bc.blocks))
}

func (bc *Blockchain) LastBlock() *core.Block {
	if len(bc.blocks) == 0 {
		return nil
	}
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) LastHash() []byte {
	last := bc.LastBlock()
	if last == nil {
		return nil
	}
	return last.Hash
}

func (bc *Blockchain) GetBlockByHeight(h uint64) (*core.Block, bool) {
	if h == 0 || h > uint64(len(bc.blocks)) {
		return nil, false
	}
	return bc.blocks[h-1], true
}

func (bc *Blockchain) Blocks() []*core.Block {
	return bc.blocks
}
