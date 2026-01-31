package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/irlan/quantumpay-go/internal/api"
	"github.com/irlan/quantumpay-go/internal/blockchain"
	// "github.com/irlan/quantumpay-go/internal/core" <-- INI DIHAPUS AGAR TIDAK ERROR
	"github.com/irlan/quantumpay-go/internal/engine"
	"github.com/irlan/quantumpay-go/internal/grpc/server"
	"github.com/irlan/quantumpay-go/internal/grpc/service"
	"github.com/irlan/quantumpay-go/internal/mempool"
	"github.com/irlan/quantumpay-go/internal/state"
)

// -------------------------------------------------------------
// ADAPTER LOKAL (Penghubung Blockchain -> gRPC Service)
// -------------------------------------------------------------

// ChainAdapter membungkus Blockchain agar sesuai dengan interface di service
type ChainAdapter struct {
	Chain *blockchain.Blockchain
}

// Height mengembalikan tinggi chain saat ini
func (a *ChainAdapter) Height() uint64 {
	return a.Chain.Height()
}

// GetBlockByHeight implementasi interface: want (any, bool)
func (a *ChainAdapter) GetBlockByHeight(height uint64) (any, bool) {
	// a.Chain.GetBlockByHeight mengembalikan (*core.Block, error)
	blk, err := a.Chain.GetBlockByHeight(height)
	if err != nil {
		return nil, false
	}
	// Kita return 'blk' langsung. Karena return type fungsi ini 'any',
	// Go otomatis membungkus *core.Block menjadi any.
	// Tidak perlu import 'core' di file ini.
	return blk, true
}

// -------------------------------------------------------------
// MAIN ENTRY POINT
// -------------------------------------------------------------
func main() {
	// 1. Konfigurasi
	apiAddr := flag.String("api-addr", ":8081", "API Port (Locked for Wallet)")
	grpcAddr := flag.String("grpc-addr", ":9090", "gRPC Port")
	dbPath := flag.String("db-path", "./data/quantum_db", "Path Database BadgerDB")
	
	// Alamat Sovereign (Pemilik Genesis)
	sovereignAddr := "0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a"
	flag.Parse()

	fmt.Println("🌟 QUANTUMPAY MAINNET: DAWN OF SOVEREIGNTY")
	fmt.Printf("👑 Sovereign Address: %s\n", sovereignAddr)

	// 2. Init WorldState (BadgerDB Persistence)
	ws, err := state.NewWorldState(*dbPath)
	if err != nil {
		log.Fatalf("❌ Gagal inisialisasi BadgerDB: %v", err)
	}
	defer ws.Close() // Pastikan DB tertutup rapi saat exit

	// 3. Init Blockchain
	// Memasukkan WorldState dan Alamat Founder untuk Genesis
	chain := blockchain.NewBlockchain(ws, sovereignAddr)
	
	// --- AMBIL GENESIS HASH ---
	// Variabel 'genesisBlock' otomatis terdeteksi sebagai *core.Block
	// karena return value dari chain.GetBlockByHeight.
	genesisBlock, err := chain.GetBlockByHeight(0)
	if err != nil {
		log.Fatalf("❌ CRITICAL: Genesis block not found in BadgerDB!")
	}
	genesisHash := genesisBlock.Hash
	
	fmt.Printf("🏛️  GENESIS HASH SECURED: %x\n", genesisHash)
	fmt.Println("--------------------------------------------------")

	// 4. Init Engine & Mempool
	mp := mempool.New()
	eng := engine.New(chain, mp)

	// 5. Jalankan gRPC Server (Port 9090)
	go func() {
		// Gunakan Adapter Lokal
		adapter := &ChainAdapter{Chain: chain}
		
		// Injeksi adapter ke Service
		svc := service.NewNodeService(adapter)
		
		// Start Server
		srv := server.NewServer(*grpcAddr)
		log.Printf("📡 [gRPC] Gateway Open: %s", *grpcAddr)
		srv.Start(svc)
	}()

	// 6. Jalankan API Server (Port 8081)
	go func() {
		cfg := &api.ServerConfig{Port: *apiAddr}
		log.Printf("🌐 [API] Interface Live: %s", *apiAddr)
		api.StartServer(cfg) 
	}()

	// 7. Mining Loop (The Pulse of the Network)
	go func() {
		log.Println("⛏️  [MINER] Consensus Engine Activated")
		for {
			// ProduceBlock dari engine baru
			blk, err := eng.ProduceBlock()
			
			if err == nil && blk != nil {
				// Akses Height lewat Header
				log.Printf("📦 Block Minted #%d | Hash: %x | Tx: %d", 
					blk.Header.Height, 
					blk.Hash[:4], 
					len(blk.Transactions))
			}
			
			// Jeda antar blok
			time.Sleep(5 * time.Second)
		}
	}()

	// 8. Graceful Shutdown (Agar data BadgerDB tidak korup)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("\n🛑 Node Safely Hibernating... Flushing data to SSD.")
}
