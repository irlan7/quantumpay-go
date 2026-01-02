package main

import (
	"log"

	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/persistence"
	"github.com/irlan/quantumpay-go/internal/state"
)

func main() {
	log.Println("🚀 QuantumPay Node starting...")

	storage := persistence.NewFileStorage("./data")

	blocks, err := storage.LoadBlocks()
	if err != nil {
		log.Fatal(err)
	}

	chain := blockchain.NewBlockchain()
	for _, b := range blocks {
		chain.AddBlock(b)
	}

	ws := state.NewWorldState()

	engine := blockchain.NewEngine(chain, ws)

	engine.ProduceBlock()

	_ = storage.SaveBlocks(chain.Blocks())

	log.Println("🟢 Node finished successfully")
}
