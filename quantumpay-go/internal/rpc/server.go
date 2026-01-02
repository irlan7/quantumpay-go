package rpc

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/irlan/quantumpay-go/internal/blockchain"
)

type Server struct {
	chain *blockchain.Blockchain
	addr  string
}

func NewServer(chain *blockchain.Blockchain, addr string) *Server {
	return &Server{
		chain: chain,
		addr:  addr,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/status", s.handleStatus)
	mux.HandleFunc("/block/latest", s.handleLatestBlock)
	mux.HandleFunc("/block", s.handleBlockByHeight)

	log.Printf("🌐 RPC server listening on %s\n", s.addr)
	if err := http.ListenAndServe(s.addr, mux); err != nil {
		log.Fatal(err)
	}
}

/* ---------------- Handlers ---------------- */

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]string{
		"status": "ok",
	})
}

func (s *Server) handleStatus(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]interface{}{
		"height": s.chain.Height(),
	})
}

func (s *Server) handleLatestBlock(w http.ResponseWriter, _ *http.Request) {
	block := s.chain.LastBlock()
	if block == nil {
		http.Error(w, "no blocks", http.StatusNotFound)
		return
	}
	writeJSON(w, block)
}

func (s *Server) handleBlockByHeight(w http.ResponseWriter, r *http.Request) {
	heightStr := r.URL.Query().Get("height")
	if heightStr == "" {
		http.Error(w, "missing height", http.StatusBadRequest)
		return
	}

	height, err := strconv.ParseUint(heightStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid height", http.StatusBadRequest)
		return
	}

	block, ok := s.chain.GetBlockByHeight(height)
	if !ok {
		http.Error(w, "block not found", http.StatusNotFound)
		return
	}

	writeJSON(w, block)
}

/* ---------------- Helpers ---------------- */

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
