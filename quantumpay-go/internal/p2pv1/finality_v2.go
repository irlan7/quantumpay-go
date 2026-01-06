package p2pv1

import (
	"errors"
	"sync"
)

/*
Finality v2:
- checkpoint voting
- validator signature based
- no fork-choice interference
*/

type Vote struct {
	ValidatorID string
	Checkpoint  uint64
	Hash        string
	Signature   []byte
}

type CheckpointV2 struct {
	Height    uint64
	Hash      string
	Votes     map[string]Vote
	PowerSum  uint64
	Justified bool
	Finalized bool
}

type FinalityV2 struct {
	mu sync.RWMutex

	validators map[string]Validator
	totalPower uint64

	checkpoints map[uint64]*CheckpointV2
	lastFinal   uint64
}

// ===============================
// Constructor
// ===============================

func NewFinalityV2(validators []Validator) *FinalityV2 {
	vmap := make(map[string]Validator)
	var total uint64

	for _, v := range validators {
		vmap[v.ID] = v
		total += v.Power
	}

	return &FinalityV2{
		validators:  vmap,
		totalPower:  total,
		checkpoints: make(map[uint64]*CheckpointV2),
	}
}

// ===============================
// On Vote
// ===============================

func (f *FinalityV2) OnVote(v Vote) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	val, ok := f.validators[v.ValidatorID]
	if !ok {
		return errors.New("unknown validator")
	}

	cp, ok := f.checkpoints[v.Checkpoint]
	if !ok {
		cp = &CheckpointV2{
			Height: v.Checkpoint,
			Hash:   v.Hash,
			Votes:  make(map[string]Vote),
		}
		f.checkpoints[v.Checkpoint] = cp
	}

	// idempotent: one vote per validator
	if _, exists := cp.Votes[v.ValidatorID]; exists {
		return nil
	}

	// verify signature (external helper)
	if err := VerifyVoteSignature(v, val.PublicKey); err != nil {
		return err
	}

	cp.Votes[v.ValidatorID] = v
	cp.PowerSum += val.Power

	// justify if >2/3 power
	if cp.PowerSum*3 > f.totalPower*2 {
		cp.Justified = true
	}

	// finalize if parent justified
	parent, ok := f.checkpoints[v.Checkpoint-CheckpointInterval]
	if ok && parent.Justified && cp.Justified {
		cp.Finalized = true
		f.lastFinal = cp.Height
	}

	return nil
}

// ===============================
// Queries
// ===============================

func (f *FinalityV2) LastFinalized() uint64 {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.lastFinal
}
