package blockchain

import (
	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/state"
)

type Engine struct {
	Chain *Blockchain
	State *state.WorldState
}

func NewEngine(chain *Blockchain, state *state.WorldState) *Engine {
	return &Engine{
		Chain: chain,
		State: state,
	}
}

func (e *Engine) ExecuteBlock(block *core.Block) error {
	for _, tx := range block.Transactions {
		if err := e.State.ApplyTransaction(tx); err != nil {
			return err
		}
	}
	e.Chain.AddBlock(block)
	return nil
}
