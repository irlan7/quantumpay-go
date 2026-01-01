package rpc

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/irlan/quantumpay-go/internal/blockchain"
)

type Server struct {
	chain      *blockchain.Blockchain
	addr       string
	clients    map[string]int
	mu         sync.Mutex
	limit      int
	window     time.Duration
}

func NewServer(chain *blockchain.Blockchain, addr string) *Server {
	return &Server{
		chain:   chain,
		addr:    addr,
		clients: make(map[string]int),
		limit:   60,               // 60 request
		window:  time.Minute,      // per menit
	}
}

func (s *Server) Start() {
	go s.resetLimiter()

	http.HandleFunc("/status", s.withLimit(s.handleStatus))
	http.HandleFunc("/block/latest", s.withLimit(s.handleLatestBlock))
	http.HandleFunc("/block", s.withLimit(s.handleGetBlockByHeight))

	log.Printf("🌐 RPC server listening on %s\n", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, nil))
}

// --------------------
// Rate Limiter
// --------------------

func (s *Server) withLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		s.mu.Lock()
		s.clients[ip]++
		count := s.clients[ip]
		s.mu.Unlock()

		if count > s.limit {
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next(w, r)
	}
}

func (s *Server) resetLimiter() {
	ticker := time.NewTicker(s.window)
	for range ticker.C {
		s.mu.Lock()
		s.clients = make(map[string]int)
		s.mu.Unlock()
	}
}

// --------------------
// Handlers
// --------------------

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"height": s.chain.Height(),
	}
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) handleLatestBlock(w http.ResponseWriter, r *http.Request) {
	block := s.chain.LastBlock()
	json.NewEncoder(w).Encode(block)
}

func (s *Server) handleGetBlockByHeight(w http.ResponseWriter, r *http.Request) {
	heightStr := r.URL.Query().Get("height")
	if heightStr == "" {
		http.Error(w, "missing height parameter", http.StatusBadRequest)
		return
	}

	// ✅ FIX UTAMA ADA DI SINI
	height, err := strconv.ParseUint(heightStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid height parameter", http.StatusBadRequest)
		return
	}

	block, err := s.chain.GetBlockByHeight(height)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(block)
}
