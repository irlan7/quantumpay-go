package mempool

import (
	"sync"

	"github.com/irlan/quantumpay-go/internal/core"
)

// Mempool menyimpan transaksi sementara
type Mempool struct {
	mu  sync.Mutex
	txs []*core.Transaction
}

// New membuat mempool baru
func New() *Mempool {
	return &Mempool{
		txs: make([]*core.Transaction, 0),
	}
}

// Add menambahkan transaksi ke mempool
func (m *Mempool) Add(tx *core.Transaction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.txs = append(m.txs, tx)
}

// PopAll mengambil dan mengosongkan mempool
func (m *Mempool) PopAll() []*core.Transaction {
	m.mu.Lock()
	defer m.mu.Unlock()

	txs := m.txs
	m.txs = make([]*core.Transaction, 0)
	return txs
}
