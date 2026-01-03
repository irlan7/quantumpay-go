package main

import (
	"log"
	"time"

	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/engine"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

func main() {
	// Initialize blockchain
	chain := blockchain.NewBlockchain()

	// Initialize mempool
	mp := mempool.New()

	// Initialize engine
	eng := engine.New(chain, mp)

	log.Println("QuantumPay node started")

	// Simple block production loop (temporary, dev-mode)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		blk, err := eng.ProduceBlock()
		if err != nil {
			log.Println("ProduceBlock error:", err)
			continue
		}

		log.Printf("New block produced: height=%d hash=%x\n",
			chain.Height(),
			blk.Hash,
		)
	}
}
