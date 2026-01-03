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
	// ===== FLAGS =====
	enableP2P := flag.Bool("p2p", false, "enable p2p networking (v1 placeholder)")
	p2pPort := flag.Int("p2p-port", 7001, "p2p listen port (reserved)")
	p2pPeer := flag.String("p2p-peer", "", "p2p peer address (reserved)")
	flag.Parse()

	log.Println("[NODE] Starting QuantumPay node")

	// ===== CORE (STABLE) =====
	chain := blockchain.NewBlockchain()
	mp := mempool.New()
	eng := engine.New(chain, mp)

	// ===== ENGINE LOOP (STABLE) =====
	go func() {
		for {
			blk, err := eng.ProduceBlock()
			if err != nil {
				log.Println("[ENGINE] block production error:", err)
				continue
			}
			log.Printf(
				"[ENGINE] New block produced height=%d hash=%x\n",
				blk.Height,
				blk.Hash,
			)
			time.Sleep(5 * time.Second)
		}
	}()

	// ===== P2P PLACEHOLDER (NO IMPORT, NO CYCLE) =====
	if *enableP2P {
		log.Printf("[P2P] enabled (v1 placeholder)")
		log.Printf("[P2P] listen port : %d", *p2pPort)

		if *p2pPeer != "" {
			log.Printf("[P2P] peer target  : %s", *p2pPeer)
		}

		log.Println("[P2P] runtime NOT wired yet (expected in P2P v1 phase)")
	}

	// ===== GRACEFUL SHUTDOWN =====
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("[NODE] Shutdown signal received")
	log.Println("[NODE] Exiting cleanly")
}
