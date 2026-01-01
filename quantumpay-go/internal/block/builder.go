package block

import "github.com/irlan/quantumpay-go/internal/core"

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(
	height uint64,
	parentHash []byte,
	txs []*core.Transaction,
) *core.Block {

	return core.NewBlock(height, parentHash, txs)
}
