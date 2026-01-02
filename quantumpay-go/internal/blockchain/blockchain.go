package blockchain

import (
	"errors"
	"sync"

	"github.com/irlan/quantumpay-go/internal/core"
)

var (
	ErrEmptyChain = errors.New("blockchain is empty")
)

/*
Blockchain
- hanya menyimpan urutan block
- tidak tahu state
- tidak tahu executor
- pure chain storage
*/
type Blockchain struct {
	mu     sync.RWMutex
	blocks []*core.Block
}

// NewBlockchain creates empty chain
func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: make([]*core.Block, 0),
	}
}

// AddBlock appends block to chain
func (bc *Blockchain) AddBlock(b *core.Block) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.blocks = append(bc.blocks, b)
}

// Height returns current chain height
func (bc *Blockchain) Height() uint64 {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return uint64(len(bc.blocks))
}

// LastHash returns hash of last block
func (bc *Blockchain) LastHash() []byte {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if len(bc.blocks) == 0 {
		return nil
	}
	return bc.blocks[len(bc.blocks)-1].Hash
}

// LastBlock returns last block
func (bc *Blockchain) LastBlock() (*core.Block, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if len(bc.blocks) == 0 {
		return nil, ErrEmptyChain
	}
	return bc.blocks[len(bc.blocks)-1], nil
}

// GetBlockByHeight returns block at height (0-based)
func (bc *Blockchain) GetBlockByHeight(h uint64) (*core.Block, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if h >= uint64(len(bc.blocks)) {
		return nil, ErrEmptyChain
	}
	return bc.blocks[h], nil
}

// GetBlockByHash scans chain for hash
func (bc *Blockchain) GetBlockByHash(hash []byte) (*core.Block, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	for _, b := range bc.blocks {
		if string(b.Hash) == string(hash) {
			return b, nil
		}
	}
	return nil, ErrEmptyChain
}
