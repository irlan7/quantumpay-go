package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/engine"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

func main() {
	// ---- Flags (future-proof, P2P disabled by default)
	enableP2P := flag.Bool("p2p", false, "enable p2p networking (disabled by default)")
	p2pPort := flag.Int("p2p-port", 7001, "p2p listen port (unused for now)")
	flag.Parse()

	log.Println("[NODE] Starting QuantumPay node")

	// ---- Core components (STABLE)
	chain := blockchain.NewBlockchain()
	mp := mempool.New()

	eng := engine.New(chain, mp)

	// ---- Produce blocks in background
	go func() {
		for {
			blk, err := eng.ProduceBlock()
			if err != nil {
				log.Println("[ENGINE] block production error:", err)
				continue
			}
			log.Printf("[ENGINE] New block produced height=%d hash=%x\n",
				blk.Height, blk.Hash)
			time.Sleep(5 * time.Second)
		}
	}()

	// ---- P2P placeholder (INTENTIONALLY NO IMPORT)
	if *enableP2P {
		log.Printf("[P2P] flag enabled but P2P runtime is not wired yet (port=%d)\n", *p2pPort)
		log.Println("[P2P] This is expected in P2P v1 phase")
	}

	// ---- Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("[NODE] Shutdown signal received")
	log.Println("[NODE] Exiting cleanly")
}
