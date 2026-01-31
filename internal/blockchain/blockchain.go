package blockchain

import (
	"encoding/binary"
	"fmt"
	"log"
	"sync"

	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/state"
)

type Blockchain struct {
	blocks []*core.Block
	state  *state.WorldState
	mu     sync.RWMutex
}

// NewBlockchain membuat instance blockchain baru dan me-load state dari BadgerDB
func NewBlockchain(ws *state.WorldState, sovereignAddr, founderAddr string) (*Blockchain, error) {
	bc := &Blockchain{
		blocks: []*core.Block{},
		state:  ws,
	}

	// 1. Cek Metadata di DB (Apakah ini node baru atau restart?)
	lastHash := ws.GetMetadata("lastHash")
	
	if lastHash == nil {
		// --- INI NODE BARU (Fresh Install) ---
		log.Println("⚡ Inisialisasi Genesis Block & Distribusi Awal...")

		// A. Distribusi Koin (Founder 10%, Sovereign 90%)
		// FIX: Ganti InitialDistribute -> InitializeGenesis
		if err := bc.state.InitializeGenesis(sovereignAddr, founderAddr); err != nil {
			return nil, fmt.Errorf("gagal distribusi genesis: %v", err)
		}

		// B. Buat Genesis Block
		genesisBlock := core.NewBlock([]byte{}, 0, nil) // PrevHash kosong
		bc.blocks = append(bc.blocks, genesisBlock)

		// C. Simpan Metadata Awal
		// FIX: Gunakan SaveMetadata yang baru kita buat
		ws.SaveMetadata("lastHash", genesisBlock.Hash)
		ws.SaveMetadata("height", uint64ToBytes(0))

	} else {
		// --- INI RESTART (Resume) ---
		currentHeightBytes := ws.GetMetadata("height")
		currentHeight := binary.BigEndian.Uint64(currentHeightBytes)
		log.Printf("🔄 Node Restarted. Last Hash: %x, Height: %d", lastHash, currentHeight)
		
		// Note: Idealnya kita load blok dari disk juga, tapi untuk tahap ini
		// kita fokus memulihkan state saldo dan metadata dulu.
	}

	return bc, nil
}

// AddBlock menambahkan blok baru ke chain
func (bc *Blockchain) AddBlock(blk *core.Block) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// 1. Validasi Transaksi PQC
	if err := blk.ValidateTransactions(); err != nil {
		return err
	}

	// 2. Update World State (Saldo)
	for _, tx := range blk.Transactions {
		if err := bc.state.ExecuteTransaction(tx); err != nil {
			return fmt.Errorf("transaksi gagal dieksekusi: %v", err)
		}
	}

	// 3. Tambahkan ke Memori
	bc.blocks = append(bc.blocks, blk)

	// 4. Update Metadata Persisten (BadgerDB)
	// FIX: SaveMetadata agar kalau mati lampu data aman
	if err := bc.state.SaveMetadata("lastHash", blk.Hash); err != nil {
		return err
	}
	if err := bc.state.SaveMetadata("height", uint64ToBytes(blk.Header.Height)); err != nil {
		return err
	}

	return nil
}

func (bc *Blockchain) Height() uint64 {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	if len(bc.blocks) == 0 {
		return 0
	}
	return bc.blocks[len(bc.blocks)-1].Header.Height
}

func (bc *Blockchain) LastHash() []byte {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	if len(bc.blocks) == 0 {
		return []byte{}
	}
	return bc.blocks[len(bc.blocks)-1].Hash
}

func uint64ToBytes(val uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	return b
}
