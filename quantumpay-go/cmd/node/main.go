package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/engine"
	"github.com/irlan/quantumpay-go/internal/grpc/server"
	"github.com/irlan/quantumpay-go/internal/grpc/service"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

// --- ADAPTER (JEMBATAN) ---
// Adapter ini mengubah *core.Block menjadi 'any' agar interface senang
type ChainAdapter struct {
	Chain *blockchain.Blockchain
}

func (a *ChainAdapter) Height() uint64 {
	return a.Chain.Height()
}

// Perhatikan return (any, bool) -> INI KUNCINYA
func (a *ChainAdapter) GetBlockByHeight(height uint64) (any, bool) {
	blk, err := a.Chain.GetBlockByHeight(height)
	if err != nil {
		return nil, false
	}
	// blk otomatis dianggap sebagai 'any'
	return blk, true
}

func main() {
	// Flags
	_ = flag.Bool("p2p", false, "enable p2p networking")
	enableGRPC := flag.Bool("grpc", true, "enable gRPC server")
	grpcAddr := flag.String("grpc-addr", "0.0.0.0:9090", "gRPC listen address")
	flag.Parse()

	// 1. Identitas
	fmt.Println("\n==========================================================")
	fmt.Println("🚀 QUANTUMPAY NETWORK NODE V1.1")
	fmt.Println("🔗 Chain ID     : 77077 [FROZEN]")
	fmt.Println("💎 Genesis Hash : 0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a")
	fmt.Println("==========================================================\n")

	// 2. Init Core
	chain := blockchain.NewBlockchain()
	mp := mempool.New()
	eng := engine.New(chain, mp)

	// 3. Aktivasi gRPC
	if *enableGRPC {
		// Pasang Adapter
		adapter := &ChainAdapter{Chain: chain}
		
		nodeSvc := service.NewNodeService(adapter)
		grpcSrv := server.NewServer(*grpcAddr)

		go func() {
			if err := grpcSrv.Start(nodeSvc); err != nil {
				log.Fatalf("[gRPC] Failed to start: %v", err)
			}
		}()
	}

	// 4. Engine Loop
	go func() {
		for {
			blk, err := eng.ProduceBlock()
			if err != nil {
				log.Println("[ENGINE] Wait:", err)
				time.Sleep(2 * time.Second)
				continue
			}
			log.Printf("[ENGINE] 🔨 New Block #%d Hash=%x\n", blk.Height, blk.Hash)
			time.Sleep(5 * time.Second)
		}
	}()

	// 5. Shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("[NODE] Shutting down...")
}
