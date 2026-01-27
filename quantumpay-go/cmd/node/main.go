package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// IMPORT PATH LENGKAP (Anti Import Cycle)
	"github.com/irlan/quantumpay-go/internal/api"
	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/engine"
	"github.com/irlan/quantumpay-go/internal/grpc/server"
	"github.com/irlan/quantumpay-go/internal/grpc/service"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

// Adapter untuk gRPC (Wajib ada agar tidak error interface)
type ChainAdapter struct {
	Chain *blockchain.Blockchain
}

func (a *ChainAdapter) Height() uint64 {
	return a.Chain.Height()
}

func (a *ChainAdapter) GetBlockByHeight(height uint64) (any, bool) {
	blk, err := a.Chain.GetBlockByHeight(height)
	if err != nil { return nil, false }
	return blk, true
}

func main() {
	// 1. Konfigurasi
	apiAddr := flag.String("api-addr", ":8080", "API Port")
	grpcAddr := flag.String("grpc-addr", ":9090", "gRPC Port")
	enableGRPC := flag.Bool("grpc", true, "Enable gRPC")
	flag.Parse()

	fmt.Println("🚀 QUANTUMPAY NODE STARTING...")

	// 2. Init Core
	chain := blockchain.NewBlockchain()
	mp := mempool.New()
	eng := engine.New(chain, mp)

	// 3. Jalankan gRPC Server (Background)
	if *enableGRPC {
		go func() {
			adapter := &ChainAdapter{Chain: chain}
			svc := service.NewNodeService(adapter)
			srv := server.NewServer(*grpcAddr)
			log.Printf("📡 [gRPC] Running on %s", *grpcAddr)
			srv.Start(svc)
		}()
	}

	// 4. Jalankan API Server (Background) - INI YANG TADI ERROR
	go func() {
		cfg := &api.ServerConfig{Port: *apiAddr}
		// Sekarang aman karena import path sudah benar
		api.StartServer(cfg) 
	}()

	// 5. Mining Loop (Background)
	go func() {
		log.Println("⛏️  [MINER] Engine Started")
		for {
			blk, err := eng.ProduceBlock()
			if err == nil {
				log.Printf("📦 New Block #%d [%x]", blk.Height, blk.Hash[:4])
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// 6. Tahan Aplikasi Agar Tidak Mati
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("\n🛑 Shutting down node...")
}
