package blockchain

import "github.com/irlan/quantumpay-go/internal/core"

// view adalah adapter read-only dari Blockchain
// untuk memenuhi interface block.ChainView
type view struct {
	chain *Blockchain
}

// NewView membuat ChainView tanpa import cycle
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

// GetBlockByHeight mengembalikan block atau nil jika tidak ada
func (v *view) GetBlockByHeight(h uint64) *core.Block {
	blk, err := v.chain.GetBlockByHeight(h)
	if err != nil {
		return nil
	}
	return blk
}
