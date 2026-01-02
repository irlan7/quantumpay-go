package blockchain

import (
	"log"

	"github.com/irlan/quantumpay-go/internal/block"
	"github.com/irlan/quantumpay-go/internal/state"
)

type Engine struct {
	Chain *Blockchain
	State *state.WorldState
}

// Constructor
func NewEngine(chain *Blockchain, state *state.WorldState) *Engine {
	return &Engine{
		Chain: chain,
		State: state,
	}
}

// Produce + execute + commit ONE block
func (e *Engine) ProduceBlock() error {
	view := NewChainView(e.Chain)

	builder := block.NewBuilder(view, e.State)

	newBlock, err := builder.Build()
	if err != nil {
		return err
	}

	// Execute block
	err = block.Execute(newBlock, e.State)
	if err != nil {
		return err
	}

	// Commit block
	e.Chain.AddBlock(newBlock)

	log.Printf("✅ Block committed. Height: %d\n", e.Chain.Height())

	return nil
}
