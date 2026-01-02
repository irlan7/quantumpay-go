package block

import "github.com/irlan/quantumpay-go/internal/core"

type ChainView interface {
	Height() uint64
	LastHash() []byte
	GetBlockByHeight(height uint64) *core.Block
}
