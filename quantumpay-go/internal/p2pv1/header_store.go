package p2pv1

import "sync"

type HeaderStore struct {
	mu      sync.RWMutex
	headers map[uint64]HeaderMsg
}

func NewHeaderStore() *HeaderStore {
	return &HeaderStore{
		headers: make(map[uint64]HeaderMsg),
	}
}

func (s *HeaderStore) Has(height uint64) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.headers[height]
	return ok
}

func (s *HeaderStore) Add(h HeaderMsg) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.headers[h.Height] = h
}
