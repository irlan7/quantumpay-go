package main

import (
	"flag"
	"log"

	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/mempool"
	"github.com/irlan/quantumpay-go/internal/rpc"
)

func main() {
	nodeID := flag.String("node-id", "node1", "Node ID")
	rpcPort := flag.String("rpc-port", "8080", "RPC port")
	dataDir := flag.String("data", "data/node1", "Data directory")
	flag.Parse()

	log.Println("🚀 QuantumPay Node starting...", *nodeID)

	// Storage
	storage := blockchain.NewFileStorage(*dataDir)

	// Blockchain
	chain := blockchain.NewBlockchain()
	blocks, err := storage.LoadBlockchain()
	if err != nil {
		log.Fatal(err)
	}
	chain.LoadBlocks(blocks)

	log.Printf("📦 Blockchain height: %d\n", chain.Height())

	// Mempool
	mp := mempool.New()

	// Engine
	engine := blockchain.NewEngine(chain, mp)

	// RPC
	rpcServer := rpc.NewServer(chain, ":"+*rpcPort)
	go rpcServer.Start()

	// Auto mine 1 block (testnet only)
	if chain.Height() == 0 {
		engine.BuildAndCommitBlock()
		storage.SaveBlockchain(chain.AllBlocks())
		log.Printf("✅ Block committed. Height: %d\n", chain.Height())
	}

	select {} // keep node alive
}
