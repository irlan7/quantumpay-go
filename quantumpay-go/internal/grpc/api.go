package grpcapi

import "github.com/irlan/quantumpay-go/internal/core"

type EngineAPI interface {
	ChainHeight() uint64
	GetBlock(height uint64) (*core.Block, bool)
	SubmitTransaction(tx *core.Transaction) error
}
