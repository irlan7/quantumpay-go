package service

import "github.com/irlan/quantumpay-go/internal/core"

type ChainAPI interface {
	Height() uint64
	GetBlockByHeight(uint64) (*core.Block, bool)
}
