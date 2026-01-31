package api

import (
	"encoding/json"
	"log" // Kita pakai log, bukan fmt
	"net/http"
	"time"

	"github.com/irlan/quantumpay-go/internal/core"
	"github.com/irlan/quantumpay-go/internal/mempool"
	"github.com/irlan/quantumpay-go/internal/state"
)

// ServerConfig menyimpan pointer ke komponen vital node
type ServerConfig struct {
	Port    string
	Mempool *mempool.Mempool  // Akses ke antrean transaksi
	State   *state.WorldState // Akses ke database saldo (BadgerDB)
}

// StartServer menyalakan HTTP REST API
func StartServer(cfg *ServerConfig) {
	mux := http.NewServeMux()

	// 1. Health Check (Untuk memastikan Node Hidup)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"status":  "Mainnet Active",
			"version": "1.0.0-Sovereign",
			"uptime":  "running",
		})
	})

	// 2. Cek Saldo (Langsung dari BadgerDB SSD)
	mux.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request) {
		addr := r.URL.Query().Get("address")
		if addr == "" {
			http.Error(w, "Address required", http.StatusBadRequest)
			return
		}

		// Ambil data akun dari State
		acc := cfg.State.GetAccount(addr)

		response := map[string]interface{}{
			"address": addr,
			"balance": acc.Balance, // Satuan terkecil (misal: Satoshi/Quantum)
			"nonce":   acc.Nonce,
		}
		jsonResponse(w, response)
	})

	// 3. Kirim Transaksi (Pintu Gerbang PQC Wallet)
	mux.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		// Validasi Method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed, use POST", http.StatusMethodNotAllowed)
			return
		}

		// Decode JSON Body ke Struct Transaction Core
		var tx core.Transaction
		if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
			http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Validasi Data Minimal
		if tx.From == "" || tx.To == "" || tx.Value == 0 {
			http.Error(w, "Invalid transaction data (From, To, Value required)", http.StatusBadRequest)
			return
		}

		// Log aktivitas transaksi masuk
		log.Printf("🛡️  [API] Receiving Tx: %s -> %s | Value: %d", tx.From, tx.To, tx.Value)

		// Dorong ke Mempool (Antrean Mining)
		// Engine akan mengambil ini otomatis saat ProduceBlock
		cfg.Mempool.Add(&tx)

		response := map[string]interface{}{
			"status":    "success",
			"message":   "Transaction pushed to Mempool",
			"timestamp": time.Now().Unix(),
			"tx_hash":   "pending", // Hash akan digenerate saat mining
		}
		jsonResponse(w, response)
	})

	log.Printf("🚀 [API] Interface Live on %s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil {
		log.Printf("🔥 [API] Error: %v", err)
	}
}

// Helper untuk respons JSON standar
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS Enabled agar bisa diakses website
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
