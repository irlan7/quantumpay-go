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
	p2pPort := flag.Int("p2p-port", 7001, "p2p listen port")
	p2pPeer := flag.String("p2p-peer", "", "p2p peer address (ip:port)")
	flag.Parse()

	log.Println("[NODE] Starting QuantumPay node")

	// ===== CORE =====
	chain := blockchain.NewBlockchain()
	mp := mempool.New()
	eng := engine.New(chain, mp)

	// ===== P2P PLACEHOLDER (SAFE) =====
	if *enableP2P {
		log.Println("[P2P] enabled (v1 placeholder)")
		log.Printf("[P2P] listen port : %d\n", *p2pPort)

		if *p2pPeer != "" {
			log.Printf("[P2P] peer target : %s\n", *p2pPeer)
		}

		log.Println("[P2P] runtime NOT wired yet (expected in P2P v1 phase)")
	}

	// ===== BLOCK PRODUCTION LOOP (STABLE) =====
	go func() {
		for {
			blk, err := eng.ProduceBlock()
			if err != nil {
				log.Println("[ENGINE] block production error:", err)
				time.Sleep(2 * time.Second)
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

	// ===== GRACEFUL SHUTDOWN =====
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("[NODE] Shutdown signal received")
	log.Println("[NODE] Exiting cleanly")
}
