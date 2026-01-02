package block

import (
	"time"

	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

type Builder struct {
	view ChainView
	mp   *mempool.Mempool
}

func NewBuilder(view ChainView, mp *mempool.Mempool) *Builder {
	return &Builder{
		view: view,
		mp:   mp,
	}
}

func (b *Builder) Build() *core.Block {
	txs := b.mp.PopAll()

	return core.NewBlock(
		b.view.Height()+1,
		b.view.LastHash(),
		txs,
		uint64(time.Now().Unix()),
	)
}
