package p2pv1

import (
	"errors"
	"sync"
)

/*
TxPool = mempool sederhana
- anti duplicate
- anti cycle
- tidak tahu peer / node
*/

type Transaction struct {
	Hash   string
	From   string
	To     string
	Amount uint64
	Nonce  uint64
}

// ===============================
// TxPool structure
// ===============================

type TxPool struct {
	mu   sync.RWMutex
	txs  map[string]Transaction // hash -> tx
	size int
	max  int
}

// ===============================
// Constructor
// ===============================

func NewTxPool(max int) *TxPool {
	return &TxPool{
		txs: make(map[string]Transaction),
		max: max,
	}
}

// ===============================
// Add transaction
// ===============================

func (tp *TxPool) Add(tx Transaction, st *State) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	// capacity guard (anti spam)
	if len(tp.txs) >= tp.max {
		return errors.New("tx pool full")
	}

	// duplicate guard
	if _, exists := tp.txs[tx.Hash]; exists {
		return errors.New("duplicate tx")
	}

	// basic sanity
	if tx.Amount == 0 {
		return errors.New("zero amount")
	}

	// optional state-based check (READ ONLY)
	if st != nil {
		if st.balances[tx.From] < tx.Amount {
			return errors.New("insufficient balance")
		}
	}

	tp.txs[tx.Hash] = tx
	tp.size++
	return nil
}

// ===============================
// Remove transaction (after block commit)
// ===============================

func (tp *TxPool) Remove(hash string) {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	delete(tp.txs, hash)
}

// ===============================
// Get batch for block producer
// ===============================

func (tp *TxPool) Pending(limit int) []Transaction {
	tp.mu.RLock()
	defer tp.mu.RUnlock()

	if limit <= 0 {
		limit = 1
	}

	out := make([]Transaction, 0, limit)
	for _, tx := range tp.txs {
		out = append(out, tx)
		if len(out) >= limit {
			break
		}
	}
	return out
}

// ===============================
// Size
// ===============================

func (tp *TxPool) Size() int {
	tp.mu.RLock()
	defer tp.mu.RUnlock()
	return tp.size
}
