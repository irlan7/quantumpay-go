package p2pv1

import (
	"errors"
	"sync"
)

/*
BlockStore bertugas:
- menyimpan block body
- memastikan idempotency
- menandai block sudah dieksekusi atau belum
- TIDAK melakukan eksekusi (anti-cycle)
*/

// ===============================
// Data Structures
// ===============================

// StoredBlock menyimpan block + status eksekusi
type StoredBlock struct {
	Block    BlockMsg
	Executed bool
}

// BlockStore adalah penyimpanan block thread-safe
type BlockStore struct {
	mu     sync.RWMutex
	blocks map[string]*StoredBlock // hash -> stored block
}

// ===============================
// Constructor
// ===============================

func NewBlockStore() *BlockStore {
	return &BlockStore{
		blocks: make(map[string]*StoredBlock),
	}
}

// ===============================
// Query helpers
// ===============================

// Has mengecek apakah block ada (executed / not)
func (bs *BlockStore) Has(hash string) bool {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	_, ok := bs.blocks[hash]
	return ok
}

// IsExecuted mengecek apakah block sudah dieksekusi
func (bs *BlockStore) IsExecuted(hash string) bool {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	sb, ok := bs.blocks[hash]
	return ok && sb.Executed
}

// Get mengambil block (read-only)
func (bs *BlockStore) Get(hash string) (BlockMsg, bool) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	sb, ok := bs.blocks[hash]
	if !ok {
		return BlockMsg{}, false
	}
	return sb.Block, true
}

// ===============================
// Mutation helpers
// ===============================

// Add menyimpan block secara idempotent
func (bs *BlockStore) Add(b BlockMsg) error {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	// idempotent: block sudah ada → skip
	if _, exists := bs.blocks[b.Hash]; exists {
		return nil
	}

	// anti-cycle: block tidak boleh menunjuk dirinya sendiri
	if b.Hash == b.ParentHash {
		return errors.New("block parent cycle")
	}

	bs.blocks[b.Hash] = &StoredBlock{
		Block:    b,
		Executed: false,
	}
	return nil
}

// MarkExecuted menandai block sudah dieksekusi
func (bs *BlockStore) MarkExecuted(hash string) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if sb, ok := bs.blocks[hash]; ok {
		sb.Executed = true
	}
}

// ===============================
// Execution helpers
// ===============================

// NextUnexecuted mencari block berikutnya yang siap dieksekusi
// parentHash = hash block terakhir yang SUDAH dieksekusi
func (bs *BlockStore) NextUnexecuted(parentHash string) (*BlockMsg, bool) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	for _, sb := range bs.blocks {
		if sb.Executed {
			continue
		}

		// parent-first rule
		if sb.Block.ParentHash == parentHash {
			return &sb.Block, true
		}
	}
	return nil, false
}
