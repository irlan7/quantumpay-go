package p2pv1

import (
	"sync"
	"time"
)

/*
TxGossip = kontrol penyebaran tx
- rate limit
- anti rebroadcast
- tidak tahu peer internals
*/

type TxGossip struct {
	mu sync.Mutex

	seen map[string]time.Time // txHash -> last seen
	ttl  time.Duration
}

// Constructor
func NewTxGossip(ttl time.Duration) *TxGossip {
	return &TxGossip{
		seen: make(map[string]time.Time),
		ttl:  ttl,
	}
}

// Allow menentukan apakah tx boleh digossip
func (tg *TxGossip) Allow(hash string) bool {
	tg.mu.Lock()
	defer tg.mu.Unlock()

	now := time.Now()

	// cleanup expired entries (bounded)
	for h, t := range tg.seen {
		if now.Sub(t) > tg.ttl {
			delete(tg.seen, h)
		}
	}

	if _, exists := tg.seen[hash]; exists {
		return false
	}

	tg.seen[hash] = now
	return true
}
