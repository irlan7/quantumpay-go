package block

import (
	"time"

	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

type ChainView interface {
	Height() uint64
	TipHash() []byte
}

type Builder struct {
	view    ChainView
	mempool *mempool.Mempool
}

func NewBuilder(view ChainView, mp *mempool.Mempool) *Builder {
	return &Builder{
		view:    view,
		mempool: mp,
	}
}

func (b *Builder) Build() (*core.Block, error) {
	parentHash := b.view.TipHash()
	height := b.view.Height() + 1
	timestamp := uint64(time.Now().Unix())

	txs := b.mempool.PopAll()

	block := core.NewBlock(
		height,
		parentHash,
		txs,
		timestamp,
	)

	return block, nil
}
