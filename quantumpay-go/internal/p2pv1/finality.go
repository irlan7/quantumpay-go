package p2pv1

import "sync"

/*
FinalityGadget:
- mengelola checkpoint
- menentukan justified & finalized
- TIDAK tahu peer / p2p / tx (ANTI-CYCLE)
*/

const CheckpointInterval = 10

type Checkpoint struct {
	Height uint64
	Hash   string

	Justified bool
	Finalized bool
}

type FinalityGadget struct {
	mu sync.RWMutex

	checkpoints map[uint64]*Checkpoint
	lastFinal   uint64
}

// ===============================
// Constructor
// ===============================

func NewFinalityGadget() *FinalityGadget {
	return &FinalityGadget{
		checkpoints: make(map[uint64]*Checkpoint),
	}
}

// ===============================
// On new header
// ===============================

func (f *FinalityGadget) OnHeader(h HeaderMsg) {
	// only checkpoint heights
	if h.Height%CheckpointInterval != 0 {
		return
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	if _, exists := f.checkpoints[h.Height]; exists {
		return
	}

	f.checkpoints[h.Height] = &Checkpoint{
		Height:    h.Height,
		Hash:      h.Hash,
		Justified: true, // locally justified
	}
}

// ===============================
// Try finalize
// ===============================

func (f *FinalityGadget) TryFinalize() {
	f.mu.Lock()
	defer f.mu.Unlock()

	// finalize h if h and h-interval justified
	for height, cp := range f.checkpoints {
		if !cp.Justified || cp.Finalized {
			continue
		}

		parentHeight := height - CheckpointInterval
		parent, ok := f.checkpoints[parentHeight]
		if !ok || !parent.Justified {
			continue
		}

		cp.Finalized = true
		f.lastFinal = height
	}
}

// ===============================
// Query helpers
// ===============================

func (f *FinalityGadget) LastFinalized() uint64 {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.lastFinal
}
