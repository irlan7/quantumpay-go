package state

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/dgraph-io/badger/v4"
	// Import core agar struct Transaction terbaca
	"github.com/irlan/quantumpay-go/internal/core"
)

// HAPUS BAGIAN VAR ERROR DISINI KARENA SUDAH ADA DI errors.go
// (ErrInsufficientBalance, ErrInvalidNonce, dll sudah didefinisikan di errors.go)

// Account struct
type Account struct {
	Balance uint64 `json:"balance"`
	Nonce   uint64 `json:"nonce"`
}

// WorldState struct
type WorldState struct {
	mu       sync.RWMutex
	accounts map[string]*Account
	db       *badger.DB
}

// NewWorldState inisialisasi
func NewWorldState(dbPath string) (*WorldState, error) {
	opts := badger.DefaultOptions(dbPath).WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to open BadgerDB: %v", err)
	}

	ws := &WorldState{
		accounts: make(map[string]*Account),
		db:       db,
	}

	if err := ws.loadStateFromDB(); err != nil {
		return nil, err
	}

	return ws, nil
}

// --- PERSISTENCE ---

func (ws *WorldState) saveAccountToDB(addr string, acc *Account) error {
	data, err := json.Marshal(acc)
	if err != nil {
		return err
	}
	return ws.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(addr), data)
	})
}

func (ws *WorldState) loadStateFromDB() error {
	return ws.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := string(item.Key())
			// Skip metadata internal
			if len(key) < 20 {
				continue
			}
			err := item.Value(func(v []byte) error {
				var acc Account
				if err := json.Unmarshal(v, &acc); err == nil {
					ws.accounts[key] = &acc
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// --- METADATA ---

func (ws *WorldState) SaveMetadata(key string, value []byte) error {
	return ws.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
}

func (ws *WorldState) GetMetadata(key string) ([]byte, error) {
	var val []byte
	err := ws.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, _ = item.ValueCopy(nil)
		return nil
	})
	if err == badger.ErrKeyNotFound {
		return nil, nil
	}
	return val, err
}

// --- STATE LOGIC ---

func (ws *WorldState) InitialDistribute(addr string, amount uint64) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	acc := &Account{Balance: amount, Nonce: 0}
	ws.accounts[addr] = acc
	ws.saveAccountToDB(addr, acc)
	fmt.Printf("💰 Initial Credit to %s: %d QP\n", addr, amount)
}

func (ws *WorldState) GetAccount(addr string) Account {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	acc, ok := ws.accounts[addr]
	if !ok {
		return Account{Balance: 0, Nonce: 0}
	}
	return *acc
}

func (ws *WorldState) getOrCreate(addr string) *Account {
	acc, ok := ws.accounts[addr]
	if !ok {
		acc = &Account{}
		ws.accounts[addr] = acc
	}
	return acc
}

func (ws *WorldState) ApplyTransaction(tx *core.Transaction) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	from := ws.getOrCreate(tx.From)
	to := ws.getOrCreate(tx.To)

	// Menggunakan ErrInsufficientBalance dari errors.go
	if from.Balance < tx.Value {
		return ErrInsufficientBalance
	}
	// Menggunakan ErrInvalidNonce dari errors.go
	if from.Nonce != tx.Nonce {
		return fmt.Errorf("%w: expected %d, got %d", ErrInvalidNonce, from.Nonce, tx.Nonce)
	}

	from.Balance -= tx.Value
	from.Nonce++
	to.Balance += tx.Value

	if err := ws.saveAccountToDB(tx.From, from); err != nil {
		return err
	}
	if err := ws.saveAccountToDB(tx.To, to); err != nil {
		return err
	}

	return nil
}

func (ws *WorldState) Close() {
	if ws.db != nil {
		ws.db.Close()
	}
}
