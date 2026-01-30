package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// --- IMPORT PATH LENGKAP (Pastikan go.mod Anda bernama github.com/irlan/quantumpay-go) ---
	"github.com/irlan/quantumpay-go/internal/api"
	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/engine"
	"github.com/irlan/quantumpay-go/internal/grpc/server"
	"github.com/irlan/quantumpay-go/internal/grpc/service"
	"github.com/irlan/quantumpay-go/internal/mempool"
)

// --- ADAPTER PATTERN (ANTI IMPORT CYCLE) ---
// Adapter ini menjembatani Blockchain Core dengan gRPC Service
// tanpa membuat service mengimpor package main.
type ChainAdapter struct {
	Chain *blockchain.Blockchain
}

// Height mengembalikan tinggi blok saat ini
func (a *ChainAdapter) Height() uint64 {
	// Pastikan method Height() ada di package blockchain Anda
	return a.Chain.Height()
}

// GetBlockByHeight mengambil data blok berdasarkan tinggi
func (a *ChainAdapter) GetBlockByHeight(height uint64) (any, bool) {
	blk, err := a.Chain.GetBlockByHeight(height)
	if err != nil {
		return nil, false
	}
	return blk, true
}

// --- FUNGSI UTAMA ---
func main() {
	// 1. Konfigurasi Argumen CLI
	apiAddr := flag.String("api-addr", ":8080", "Port untuk REST API (Wallet & Health)")
	grpcAddr := flag.String("grpc-addr", ":9090", "Port untuk gRPC Node Communication")
	enableGRPC := flag.Bool("grpc", true, "Aktifkan gRPC Server")
	flag.Parse()

	fmt.Println("\n==================================================")
	fmt.Println("   🚀 QUANTUMPAY SOVEREIGN NODE v4.5 STARTING    ")
	fmt.Println("==================================================")

	// 2. Inisialisasi Core System
	log.Println("⚙️  [CORE] Initializing Blockchain & Mempool...")
	chain := blockchain.NewBlockchain() // Load database blok lokal
	mp := mempool.New()                 // Siapkan kolam transaksi memori
	eng := engine.New(chain, mp)        // Siapkan mesin konsensus

	// 3. Jalankan gRPC Server (Background Goroutine)
	if *enableGRPC {
		go func() {
			// Hubungkan Adapter ke Service
			adapter := &ChainAdapter{Chain: chain}
			svc := service.NewNodeService(adapter)
			
			// Start Server
			srv := server.NewServer(*grpcAddr)
			log.Printf("📡 [gRPC] Server Listening on %s", *grpcAddr)
			if err := srv.Start(svc); err != nil {
				log.Printf("🔥 [gRPC] Failed to start: %v", err)
			}
		}()
	}

	// 4. Jalankan REST API Server (Background Goroutine)
	// INI YANG MENGAKTIFKAN "Create Wallet" (PQC) DAN "Health Check"
	go func() {
		// Kita menggunakan konfigurasi dari package internal/api
		cfg := &api.ServerConfig{Port: *apiAddr}
		
		log.Printf("🌐 [API] Starting REST Server on %s...", *apiAddr)
		// Fungsi ini memblokir, jadi harus di dalam goroutine
		api.StartServer(cfg)
	}()

	// 5. Jalankan Mining Loop / Block Production (Simulasi)
	go func() {
		log.Println("⛏️  [MINER] Consensus Engine Started")
		for {
			// Simulasi block time 5 detik
			time.Sleep(5 * time.Second)

			// Coba produksi blok baru
			blk, err := eng.ProduceBlock()
			if err == nil {
				// Log sukses mining
				log.Printf("📦 [MINED] New Block #%d | Hash: %x | Tx: %d", 
					blk.Height, blk.Hash[:4], len(blk.Transactions))
			} else {
				// Log jika idle (misal tidak ada transaksi)
				// log.Println("💤 [MINER] Idle...") 
			}
		}
	}()

	// 6. Graceful Shutdown (Mencegah aplikasi langsung mati)
	// Menunggu sinyal CTRL+C dari terminal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	// Blokir main thread di sini sampai ada sinyal berhenti
	<-quit

	fmt.Println("\n🛑 [SHUTDOWN] Saving data and stopping node...")
	
	// Tambahkan logika penyimpanan state terakhir di sini jika perlu
	// chain.Close()
	
	fmt.Println("👋 QuantumPay Node Stopped. Goodbye!")
}
