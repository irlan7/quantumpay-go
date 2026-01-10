package state

import (
	"sync"

	"github.com/irlan/quantumpay-go/internal/core"
)

type Account struct {
	Balance uint64
	Nonce   uint64
}

type WorldState struct {
	mu       sync.RWMutex
	accounts map[string]*Account
}

func NewWorldState() *WorldState {
	return &WorldState{
		accounts: make(map[string]*Account),
	}
}

// internal helper (TIDAK export)
func (ws *WorldState) getOrCreate(addr string) *Account {
	acc, ok := ws.accounts[addr]
	if !ok {
		acc = &Account{}
		ws.accounts[addr] = acc
	}
	return acc
}

// Read-only (aman untuk RPC / validator)
func (ws *WorldState) GetAccount(addr string) Account {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	acc, ok := ws.accounts[addr]
	if !ok {
		return Account{}
	}
	return *acc // copy (AMAN)
}

// PURE STATE TRANSITION
func (ws *WorldState) ApplyTransaction(tx *core.Transaction) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	from := ws.getOrCreate(tx.From)
	to := ws.getOrCreate(tx.To)

	// rule 1: balance cukup
	if from.Balance < tx.Value {
		return ErrInsufficientBalance
	}

	// rule 2: nonce harus tepat
	if from.Nonce != tx.Nonce {
		return ErrInvalidNonce
	}

	// apply transition
	from.Balance -= tx.Value
	from.Nonce++

	to.Balance += tx.Value

	return nil
}
