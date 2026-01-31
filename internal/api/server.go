package api

import (
	"encoding/json"
	"net/http"

	"github.com/irlan/quantumpay-go/internal/blockchain"
	"github.com/irlan/quantumpay-go/internal/mempool"
	"github.com/irlan/quantumpay-go/internal/state"
)

type ServerConfig struct {
	Chain   *blockchain.Blockchain
	Mempool *mempool.Mempool
	State   *state.WorldState
}

// Handler untuk cek saldo
func (cfg *ServerConfig) HandleBalance(w http.ResponseWriter, r *http.Request) {
	// Ambil address dari query param ?address=0x...
	addr := r.URL.Query().Get("address")
	if addr == "" {
		http.Error(w, "missing address parameter", http.StatusBadRequest)
		return
	}

	// FIX: Ganti GetAccount() -> GetBalance()
	// Karena BadgerDB tidak menyimpan struct Account, cuma saldo.
	balance := cfg.State.GetBalance(addr)

	// Buat struct respon on-the-fly
	resp := struct {
		Address string `json:"address"`
		Balance uint64 `json:"balance"`
	}{
		Address: addr,
		Balance: balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Note: Pastikan di main.go handler ini dipasang:
// http.HandleFunc("/balance", serverConfig.HandleBalance)
