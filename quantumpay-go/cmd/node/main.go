package main

import (
	"flag"
	"fmt" // Import fmt ditambahkan untuk spanduk visual
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
	// ===== FLAGS (SAFE) =====
	enableP2P := flag.Bool("p2p", false, "enable p2p networking (placeholder)")
	p2pPort := flag.Int("p2p-port", 7001, "p2p listen port")
	p2pPeer := flag.String("p2p-peer", "", "p2p peer address (ip:port)")

	enableGRPC := flag.Bool("grpc", false, "enable grpc server (SAFE placeholder)")
	grpcAddr := flag.String("grpc-addr", "127.0.0.1:9090", "grpc listen address")

	flag.Parse()

	// ==========================================================
	// 🚀 SPANDUK IDENTITAS JARINGAN (VISUAL AUDIT)
	// Ditambahkan untuk transparansi Genesis Hash bagi semua user
	// ==========================================================
	fmt.Println("\n==========================================================")
	fmt.Println("🚀 QUANTUMPAY NETWORK NODE V1.1")
	fmt.Println("🔗 Chain ID     : 77077 [FROZEN]")
	fmt.Println("💎 Genesis Hash : 0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a")
	fmt.Println("==========================================================\n")

	log.Println("[NODE] Starting QuantumPay node")

	// ===== CORE BLOCKCHAIN (UNTOUCHED) =====
	chain := blockchain.NewBlockchain()
	mp := mempool.New()
	eng := engine.New(chain, mp)

	// ===== P2P PLACEHOLDER (SAFE) =====
	if *enableP2P {
		log.Println("[P2P] enabled (v1 placeholder)")
		log.Printf("[P2P] listen port: %d\n", *p2pPort)

		if *p2pPeer != "" {
			log.Printf("[P2P] peer target: %s\n", *p2pPeer)
		}
	}

	// ===== gRPC PLACEHOLDER (SAFE & ISOLATED) =====
	if *enableGRPC {
		// ⚠️ TIDAK ADA import grpc / protobuf
		// ⚠️ TIDAK ADA pemanggilan rpc
		// ⚠️ HANYA lifecycle placeholder
		log.Println("[gRPC] enabled (SAFE placeholder)")
		log.Printf("[gRPC] planned listen addr: %s\n", *grpcAddr)
		log.Println("[gRPC] server not started yet (no-op)")
		log.Println("[gRPC] integration will be wired in P1.3+ without touching core")
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
