package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/engine"
	"github.com/irlan/quantumpay-go/internal/grpc/server"
	"github.com/irlan/quantumpay-go/internal/mempool"
	"github.com/irlan/quantumpay-go/internal/state"
)

func main() {
	log.Println("🚀 QuantumPay Node starting...")

	// --- Core ---
	chain := blockchain.NewBlockchain()
	ws := state.NewWorldState()
	mp := mempool.New()

	eng := engine.New(chain, ws, mp)

	// --- gRPC ---
	grpcServer := server.New(":9090")

	go func() {
		if err := grpcServer.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// --- Block production (sementara 1x) ---
	eng.ProduceBlock()
	log.Println("📦 Blockchain height:", chain.Height())

	// --- Graceful shutdown ---
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	grpcServer.Stop()
	log.Println("🟢 Node shutdown cleanly")
}
