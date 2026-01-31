package mempool

import (
	"sync"
	"github.com/irlan/quantumpay-go/internal/core"
)

type Mempool struct {
	mu           sync.RWMutex
	transactions []*core.Transaction
}

func New() *Mempool {
	return &Mempool{
		transactions: make([]*core.Transaction, 0),
	}
}

// Add memasukkan transaksi baru ke antrean
func (m *Mempool) Add(tx *core.Transaction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.transactions = append(m.transactions, tx)
}

// GetTransactions mengambil semua transaksi untuk diproses Engine
// Audit: Ini menyelesaikan error "GetTransactions undefined"
func (m *Mempool) GetTransactions() []*core.Transaction {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	// Salin slice agar thread-safe
	txs := make([]*core.Transaction, len(m.transactions))
	copy(txs, m.transactions)
	return txs
}

// Clear mengosongkan mempool setelah transaksi masuk ke blok
// Audit: Ini menyelesaikan error "Clear undefined"
func (m *Mempool) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.transactions = make([]*core.Transaction, 0)
}

func (m *Mempool) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.transactions)
}
