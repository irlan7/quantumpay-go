package p2pv1

import (
	"errors"
	"sync"
)

type HeaderPool struct {
	mu       sync.RWMutex
	headers  map[string]HeaderMsg // hash -> header
	children map[string][]string  // parentHash -> []childHash
	tip      string               // current best tip hash
}

func NewHeaderPool() *HeaderPool {
	return &HeaderPool{
		headers:  make(map[string]HeaderMsg),
		children: make(map[string][]string),
	}
}

func (hp *HeaderPool) Add(h HeaderMsg) {
	hp.mu.Lock()
	defer hp.mu.Unlock()

	hash := h.Hash
	if _, exists := hp.headers[hash]; exists {
		return
	}

	hp.headers[hash] = h
	hp.children[h.ParentHash] = append(hp.children[h.ParentHash], hash)

	// fork-choice update
	if hp.tip == "" {
		hp.tip = hash
		return
	}

	if hp.isBetterChain(hash, hp.tip) {
		hp.tip = hash
	}
}

func (hp *HeaderPool) Tip() (HeaderMsg, error) {
	hp.mu.RLock()
	defer hp.mu.RUnlock()

	if hp.tip == "" {
		return HeaderMsg{}, errors.New("no tip")
	}
	return hp.headers[hp.tip], nil
}
func (hp *HeaderPool) isBetterChain(a, b string) bool {
	wa := hp.chainWeight(a)
	wb := hp.chainWeight(b)

	if wa != wb {
		return wa > wb
	}

	ha := hp.headers[a].Height
	hb := hp.headers[b].Height
	if ha != hb {
		return ha > hb
	}

	return hp.headers[a].Timestamp < hp.headers[b].Timestamp
}

func (hp *HeaderPool) chainWeight(tip string) int {
	visited := make(map[string]bool)
	weight := 0
	cur := tip

	for {
		if cur == "" || visited[cur] {
			break // anti-cycle
		}
		visited[cur] = true

		h, ok := hp.headers[cur]
		if !ok {
			break
		}

		weight++
		cur = h.ParentHash
	}

	return weight
}
