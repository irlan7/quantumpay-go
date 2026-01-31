package blockchain

import (
	"encoding/json" // Sekarang akan TERPAKAI di AddBlock
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/state"
)

var (
	ErrBlockNotFound   = errors.New("block not found in storage")
	ErrGenesisMismatch = errors.New("genesis hash mismatch")
)

const (
	MainnetGenesisHash = "1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a"
	MainnetTimestamp   = 1736985600
	TotalInitialSupply = 210000000
)

type Blockchain struct {
	mu    sync.RWMutex
	state *state.WorldState
	tip   []byte
}

func NewBlockchain(ws *state.WorldState, founderAddr string) *Blockchain {
	bc := &Blockchain{state: ws}

	genesisHash, _ := ws.GetMetadata("genesis_hash")
	if genesisHash == nil {
		fmt.Println("🏗️  MINTING GENESIS BLOCK...")
		bc.createGenesis(founderAddr)
	} else {
		// Validasi Hash
		if fmt.Sprintf("%x", genesisHash) != MainnetGenesisHash {
			log.Printf("⚠️ WARNING: Local genesis differs from declaration. Expected %s", MainnetGenesisHash)
		}
		bc.tip = genesisHash
		fmt.Printf("📜 CHAIN LOADED. GENESIS: %x\n", genesisHash)
	}
	return bc
}

func (bc *Blockchain) createGenesis(founderAddr string) {
	// 1. Distribusi Awal
	bc.state.InitialDistribute(founderAddr, TotalInitialSupply)

	// 2. Buat Blok Genesis
	genesisBlock := &core.Block{
		Header: core.BlockHeader{
			Height:    0,
			Timestamp: MainnetTimestamp,
			PrevHash:  make([]byte, 32),
		},
		// FIX: Gunakan []*core.Transaction (Slice of Pointers)
		Transactions: []*core.Transaction{}, 
	}

	// 3. Hitung Hash (Sekarang CalculateHash sudah Public/Exported)
	genesisBlock.Hash = genesisBlock.CalculateHash()
	hashStr := fmt.Sprintf("%x", genesisBlock.Hash)

	// Validasi Hardcoded
	if hashStr != MainnetGenesisHash {
		log.Printf("⚠️ GENERATED HASH: %s (Check core logic)", hashStr)
	}

	// Simpan
	bc.tip = genesisBlock.Hash
	bc.AddBlock(genesisBlock) // Panggil fungsi AddBlock di bawah
	bc.state.SaveMetadata("genesis_hash", genesisBlock.Hash)

	log.Printf("🏛️  GENESIS SECURED. SUPPLY: %d", TotalInitialSupply)
}

// AddBlock: Fungsi yang tadi hilang/undefined
func (bc *Blockchain) AddBlock(b *core.Block) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	key := fmt.Sprintf("block_%d", b.Header.Height)
	// FIX: json sekarang terpakai di sini
	data, _ := json.Marshal(b)

	bc.state.SaveMetadata(key, data)
	bc.state.SaveMetadata("last_height", []byte(fmt.Sprintf("%d", b.Header.Height)))
	bc.tip = b.Hash
}

// Height: Fungsi yang dipanggil oleh view.go
func (bc *Blockchain) Height() uint64 {
	data, _ := bc.state.GetMetadata("last_height")
	if data == nil {
		return 0
	}
	var h uint64
	fmt.Sscanf(string(data), "%d", &h)
	return h + 1
}

// LastHash: Fungsi yang dipanggil oleh view.go
func (bc *Blockchain) LastHash() []byte {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return bc.tip
}

// GetBlockByHeight: Fungsi yang dipanggil oleh view.go
func (bc *Blockchain) GetBlockByHeight(h uint64) (*core.Block, error) {
	key := fmt.Sprintf("block_%d", h)
	data, err := bc.state.GetMetadata(key)
	if err != nil || data == nil {
		return nil, ErrBlockNotFound
	}
	var b core.Block
	json.Unmarshal(data, &b)
	return &b, nil
}
