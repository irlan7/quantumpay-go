package p2pv1

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"sort"
	"sync"
)

/*
State = world state (deterministic)
- simpan balances
- eksekusi block
- hitung state root hash
- TIDAK tahu header / node / p2p (ANTI-CYCLE)
*/

// ===============================
// State structure
// ===============================

type State struct {
	mu       sync.RWMutex
	balances map[string]uint64
	height   uint64
}

// ===============================
// Constructor
// ===============================

func NewState() *State {
	return &State{
		balances: make(map[string]uint64),
		height:   0,
	}
}

// ===============================
// Snapshot (rollback safety)
// ===============================

func (s *State) Snapshot() *State {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cp := make(map[string]uint64, len(s.balances))
	for k, v := range s.balances {
		cp[k] = v
	}

	return &State{
		balances: cp,
		height:   s.height,
	}
}

// ===============================
// Apply block
// ===============================

func (s *State) ApplyBlock(b BlockMsg) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if b.Height != s.height+1 {
		return errors.New("invalid block height")
	}

	for _, tx := range b.Txs {
		if tx.Amount == 0 {
			continue
		}
		if s.balances[tx.From] < tx.Amount {
			return errors.New("insufficient balance")
		}

		s.balances[tx.From] -= tx.Amount
		s.balances[tx.To] += tx.Amount
	}

	s.height = b.Height
	return nil
}

// ===============================
// State Root Hash
// ===============================

// RootHash menghitung hash state deterministik
func (s *State) RootHash() []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	addrs := make([]string, 0, len(s.balances))
	for addr := range s.balances {
		addrs = append(addrs, addr)
	}
	sort.Strings(addrs)

	h := sha256.New()
	var buf [8]byte

	for _, addr := range addrs {
		h.Write([]byte(addr))
		binary.BigEndian.PutUint64(buf[:], s.balances[addr])
		h.Write(buf[:])
	}

	return h.Sum(nil)
}

// ===============================
// Read-only helpers
// ===============================

func (s *State) Height() uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.height
}
