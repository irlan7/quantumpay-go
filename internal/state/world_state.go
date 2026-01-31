package state

import (
	"encoding/binary"
	"errors"
	"log"
	"sync"

	"github.com/dgraph-io/badger/v4"
	"github.com/irlan/quantumpay-go/internal/core"
)

type WorldState struct {
	db *badger.DB
	mu sync.RWMutex
}

func NewWorldState(dbPath string) (*WorldState, error) {
	opts := badger.DefaultOptions(dbPath)
	opts.Logger = nil 
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &WorldState{db: db}, nil
}

func (ws *WorldState) Close() error {
	return ws.db.Close()
}

// ==========================================
// 1. FITUR METADATA (FIX Error Blockchain)
// ==========================================

// SaveMetadata menyimpan data sistem (misal: LastHash, Height)
func (ws *WorldState) SaveMetadata(key string, value []byte) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	
	// Tambahkan prefix "meta:" agar tidak bentrok dengan alamat wallet
	dbKey := []byte("meta:" + key)
	
	return ws.db.Update(func(txn *badger.Txn) error {
		return txn.Set(dbKey, value)
	})
}

// GetMetadata mengambil data sistem
func (ws *WorldState) GetMetadata(key string) []byte {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	dbKey := []byte("meta:" + key)
	var result []byte

	err := ws.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(dbKey)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			result = append([]byte{}, val...) // Copy data
			return nil
		})
	})

	if err != nil {
		return nil
	}
	return result
}

// ==========================================
// 2. LOGIKA SALDO & GENESIS
// ==========================================

// InitializeGenesis: Menggantikan InitialDistribute (Nama Baru)
func (ws *WorldState) InitializeGenesis(sovereignAddr, founderAddr string) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	return ws.db.Update(func(txn *badger.Txn) error {
		// Cek apakah genesis sudah pernah jalan
		if _, err := txn.Get([]byte(sovereignAddr)); err == nil {
			return nil
		}

		totalSupply := uint64(210_000_000)
		founderAlloc := totalSupply * 10 / 100
		sovereignAlloc := totalSupply - founderAlloc

		// Set Founder
		if err := txn.Set([]byte(founderAddr), uint64ToBytes(founderAlloc)); err != nil {
			return err
		}
		log.Printf("💰 [GENESIS] Founder 10%%: %d QP -> %s", founderAlloc, founderAddr)

		// Set Sovereign
		if err := txn.Set([]byte(sovereignAddr), uint64ToBytes(sovereignAlloc)); err != nil {
			return err
		}
		log.Printf("👑 [GENESIS] Sovereign 90%%: %d QP -> %s", sovereignAlloc, sovereignAddr)

		return nil
	})
}

func (ws *WorldState) GetBalance(address string) uint64 {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	var balance uint64
	err := ws.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(address))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			balance = bytesToUint64(val)
			return nil
		})
	})
	if err != nil { return 0 }
	return balance
}

// ExecuteTransaction untuk update saldo saat blok baru masuk
func (ws *WorldState) ExecuteTransaction(tx *core.Transaction) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	return ws.db.Update(func(txn *badger.Txn) error {
		senderBal := uint64(0)
		if item, err := txn.Get([]byte(tx.From)); err == nil {
			item.Value(func(v []byte) error { senderBal = bytesToUint64(v); return nil })
		}

		if senderBal < tx.Value {
			return errors.New("insufficient balance")
		}

		newSenderBal := senderBal - tx.Value
		if err := txn.Set([]byte(tx.From), uint64ToBytes(newSenderBal)); err != nil { return err }

		recipientBal := uint64(0)
		if item, err := txn.Get([]byte(tx.To)); err == nil {
			item.Value(func(v []byte) error { recipientBal = bytesToUint64(v); return nil })
		}

		newRecipientBal := recipientBal + tx.Value
		if err := txn.Set([]byte(tx.To), uint64ToBytes(newRecipientBal)); err != nil { return err }

		return nil
	})
}

func uint64ToBytes(val uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	return b
}

func bytesToUint64(val []byte) uint64 {
	if len(val) < 8 { return 0 }
	return binary.BigEndian.Uint64(val)
}
