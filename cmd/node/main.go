package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/irlan/quantumpay-go/internal/api"
	"github.com/irlan/quantumpay-go/internal/blockchain"
	// "github.com/irlan/quantumpay-go/internal/core" // Keep disabled to avoid unused error
	"github.com/irlan/quantumpay-go/internal/engine"
	"github.com/irlan/quantumpay-go/internal/grpc/server"
	"github.com/irlan/quantumpay-go/internal/grpc/service"
	"github.com/irlan/quantumpay-go/internal/mempool"
	"github.com/irlan/quantumpay-go/internal/state"
)

// -------------------------------------------------------------
// ADAPTER LOKAL (Penghubung Blockchain -> gRPC Service)
// -------------------------------------------------------------

type ChainAdapter struct {
	Chain *blockchain.Blockchain
}

func (a *ChainAdapter) Height() uint64 {
	return a.Chain.Height()
}

// GetBlockByHeight: FIX Error "undefined field or method"
// Kita gunakan View karena Blockchain struct tidak mengekspos method ini secara langsung
func (a *ChainAdapter) GetBlockByHeight(height uint64) (any, bool) {
	// Buat View read-only yang aman (Thread-Safe)
	view := blockchain.NewView(a.Chain)
	
	blk := view.GetBlockByHeight(height)
	if blk == nil {
		return nil, false
	}
	// Go otomatis membungkus *core.Block menjadi any
	return blk, true
}

// -------------------------------------------------------------
// MAIN ENTRY POINT
// -------------------------------------------------------------
func main() {
	// 1. Konfigurasi Flag
	apiAddr := flag.String("api-addr", ":8081", "API Port")
	grpcAddr := flag.String("grpc-addr", ":9090", "gRPC Port")
	dbPath := flag.String("db-path", "./data/quantum_db", "Path BadgerDB")
	
	// --- IDENTITAS SOVEREIGN & FOUNDER (PQC COMPLIANT) ---
	// Sovereign (90% Supply): Pemilik Node / Mining Power
	sovereignAddr := "0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a"
	
	// Founder (10% Supply): Alokasi Awal untuk Tim Pengembang
	// (Ganti dengan address PQC Founder yang valid)
	founderAddr := "0xFounderAddressAllocationQuantumShieldProtected" 

	flag.Parse()

	fmt.Println("🌟 QUANTUMPAY MAINNET: DAWN OF SOVEREIGNTY")
	fmt.Printf("👑 Sovereign Address: %s\n", sovereignAddr)
	fmt.Printf("💰 Founder Address : %s\n", founderAddr)

	// 2. Init WorldState (BadgerDB Persistence)
	ws, err := state.NewWorldState(*dbPath)
	if err != nil {
		log.Fatalf("❌ Gagal inisialisasi BadgerDB: %v", err)
	}
	defer ws.Close()

	// 3. Init Blockchain
	// FIX: Menggunakan 3 Argumen (State, Sovereign, Founder) dan menangkap Error
	chain, err := blockchain.NewBlockchain(ws, sovereignAddr, founderAddr)
	if err != nil {
		log.Fatalf("❌ Gagal memuat Blockchain: %v", err)
	}
	
	// --- VERIFIKASI GENESIS ---
	// Gunakan View untuk mengambil blok 0
	view := blockchain.NewView(chain)
	genesisBlock := view.GetBlockByHeight(0)
	
	if genesisBlock == nil {
		// Seharusnya tidak mungkin terjadi jika NewBlockchain sukses
		log.Fatalf("❌ CRITICAL: Genesis block hilang!")
	}
	
	fmt.Printf("🏛️  GENESIS HASH: %x\n", genesisBlock.Hash)
	fmt.Println("--------------------------------------------------")

	// 4. Init Engine & Mempool
	mp := mempool.New()
	eng := engine.New(chain, mp)

	// 5. Jalankan gRPC Server (Gateway Node-to-Node)
	go func() {
		adapter := &ChainAdapter{Chain: chain}
		svc := service.NewNodeService(adapter)
		srv := server.NewServer(*grpcAddr)
		
		log.Printf("📡 [gRPC] Gateway Open: %s", *grpcAddr)
		// Error handling untuk gRPC listener bisa ditambahkan di sini
		if err := srv.Start(svc); err != nil {
			log.Printf("⚠️ gRPC Server Error: %v", err)
		}
	}()

	// 6. Jalankan API Server (Wallet Interface)
	// FIX: Implementasi HTTP Server manual agar anti-error 'undefined api.StartServer'
	go func() {
		// Konfigurasi Handler
		cfg := &api.ServerConfig{
			Chain:   chain, // Tambahkan Chain agar API bisa baca blok
			Mempool: mp,    // Agar API bisa kirim TX
			State:   ws,    // Agar API bisa cek saldo
		}

		mux := http.NewServeMux()
		
		// Map Handler (Pastikan method ini ada di api/server.go)
		mux.HandleFunc("/balance", cfg.HandleBalance)
		// mux.HandleFunc("/send", cfg.HandleTransaction) // Uncomment jika sudah ada

		server := &http.Server{
			Addr:    *apiAddr,
			Handler: mux,
		}

		log.Printf("🌐 [API] Interface Live: %s", *apiAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ API Server Error: %v", err)
		}
	}()

	// 7. Mining Loop (Consensus Engine)
	go func() {
		log.Println("⛏️  [MINER] Consensus Engine Activated")
		for {
			// Engine otomatis mengambil TX dari Mempool
			blk, err := eng.ProduceBlock()
			
			if err == nil && blk != nil {
				log.Printf("📦 Block Minted #%d | Hash: %x | Tx: %d", 
					blk.Header.Height, 
					blk.Hash[:4], // Print 4 byte pertama hash biar rapi
					len(blk.Transactions))
			}
			
			// Block Time: 5 Detik
			time.Sleep(5 * time.Second)
		}
	}()

	// 8. Graceful Shutdown (PENTING untuk BadgerDB)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n🛑 Node Safely Hibernating... Flushing data to SSD.")
	
	// Berikan waktu 2 detik untuk proses cleanup
	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// ws.Close() dipanggil otomatis oleh defer di atas
}
