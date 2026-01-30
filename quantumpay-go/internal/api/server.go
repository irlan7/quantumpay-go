package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	// Import PQC Crypto package
	// Pastikan modul project Anda bernama: github.com/irlan/quantumpay-go
	"github.com/irlan/quantumpay-go/internal/crypto"
)

// --- KONFIGURASI DATABASE ---
const DB_FILE = "ledger_data.json"

// ServerConfig menyimpan konfigurasi server
type ServerConfig struct {
	Port string
}

// LedgerData struktur penyimpanan data json
type LedgerData struct {
	Balances  map[string]float64 `json:"balances"`
	TxHistory []Transaction      `json:"history"`
}

// Transaction struktur data transaksi
type Transaction struct {
	TxHash    string  `json:"tx_hash"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
	Status    string  `json:"status"`
}

// --- GLOBAL STATE (Thread-Safe) ---
var (
	// Genesis Account (Contoh)
	balances  = map[string]float64{"0x00000000000000000000": 1000000}
	txHistory = []Transaction{}
	mutex     = &sync.Mutex{}
)

// --- FUNGSI PERSISTENCE (DATABASE JSON) ---

func SaveData() {
	mutex.Lock()
	defer mutex.Unlock()
	
	data := LedgerData{Balances: balances, TxHistory: txHistory}
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("❌ [DB] Gagal menyimpan data:", err)
		return
	}
	
	// WriteFile dengan permission 0644 (Readable/Writable)
	_ = os.WriteFile(DB_FILE, file, 0644)
}

func LoadData() {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.ReadFile(DB_FILE)
	if err != nil {
		log.Println("⚠️ [DB] File ledger belum ada, membuat baru...")
		return
	}

	var data LedgerData
	if err := json.Unmarshal(file, &data); err == nil {
		if data.Balances != nil {
			balances = data.Balances
		}
		if data.TxHistory != nil {
			txHistory = data.TxHistory
		}
		log.Printf("📂 [DB] LOADED: %d Accounts | %d Transactions", len(balances), len(txHistory))
	} else {
		log.Println("❌ [DB] Gagal parsing JSON:", err)
	}
}

// --- SERVER SETUP & ROUTING ---

// StartServer dipanggil oleh main.go di cmd/node
func StartServer(config *ServerConfig) {
	// 1. Load Data Terakhir
	LoadData()

	// 2. Setup Router
	mux := http.NewServeMux()

	// HEALTH CHECK (Sesuai health.go)
	mux.HandleFunc("/api/health", handleHealth)
	
	// WALLET ROUTES (PQC)
	mux.HandleFunc("/api/wallet/create", handleCreateWallet)
	mux.HandleFunc("/api/wallet/import", handleImportWallet)

	// FINANCE ROUTES
	mux.HandleFunc("/api/balance", handleGetBalance)
	mux.HandleFunc("/api/faucet", handleFaucet)
	mux.HandleFunc("/api/tx/send", handleSendTx)
	
	// EXPLORER ROUTES
	mux.HandleFunc("/api/tx/history", handleGetHistory)
	mux.HandleFunc("/api/tx/detail", handleGetTxDetail)

	// 3. Jalankan Server
	log.Printf("🚀 [API] Sovereign Node Listening on Port %s", config.Port)
	log.Printf("🔗 [API] Health Check: http://localhost%s/api/health", config.Port)

	if err := http.ListenAndServe(config.Port, mux); err != nil {
		log.Fatalf("🔥 [API] Critical Error: %v", err)
	}
}

// --- HANDLER FUNCTIONS ---

// 1. Handle Health (Integrasi Status Node)
func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":    "online",
		"version":   "4.5-sovereign", // Sesuai versi QuantumPay
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "QuantumPay-Node",
	}
	jsonResponse(w, response)
}

// 2. Create Wallet (Generate PQC Keys)
func handleCreateWallet(w http.ResponseWriter, r *http.Request) {
	// Memanggil fungsi dari package internal/crypto
	wallet, err := crypto.NewKeyPair()
	if err != nil { 
		http.Error(w, "Gagal membuat wallet PQC", 500)
		return 
	}

	// --- [LOG KHUSUS PQC] ---
	// Ini akan muncul di Terminal saat tombol "Create Wallet" ditekan
	log.Printf("🛡️ [PQC] New Wallet Generated with ML-DSA Shield: %s", wallet.Address())
	
	mutex.Lock()
	// Bonus saldo awal untuk user baru (Logic Faucet otomatis untuk testing)
	if _, exists := balances[wallet.Address()]; !exists {
		balances[wallet.Address()] = 0.0 // Default 0
	}
	mutex.Unlock()
	SaveData()

	response := map[string]string{
		"address":     wallet.Address(),
		"public_key":  wallet.GetPublicKeyHex(),
		"private_key": wallet.PrivateKeyHex(),
		"mnemonic":    wallet.Mnemonic,
		"created_at":  time.Now().Format(time.RFC3339),
		"type":        "Quantum-Safe (ML-DSA)",
	}
	jsonResponse(w, response)
}

// 3. Import Wallet (Recover dari Mnemonic)
func handleImportWallet(w http.ResponseWriter, r *http.Request) {
	phrase := r.URL.Query().Get("phrase")
	if phrase == "" {
		http.Error(w, "Mnemonic phrase required", 400)
		return
	}

	// Memulihkan wallet menggunakan crypto package
	wallet, err := crypto.NewKeyPairFromMnemonic(phrase)
	if err != nil {
		jsonResponse(w, map[string]string{"status": "error", "message": "Invalid Mnemonic Phrase"})
		return
	}

	// Cek ledger, jangan reset saldo jika sudah ada
	mutex.Lock()
	if _, exists := balances[wallet.Address()]; !exists {
		balances[wallet.Address()] = 0.0
	}
	mutex.Unlock()
	SaveData()

	response := map[string]string{
		"address":     wallet.Address(),
		"public_key":  wallet.GetPublicKeyHex(),
		"private_key": wallet.PrivateKeyHex(), // Hanya dikirim sekali ke user
		"mnemonic":    phrase,
		"status":      "recovered",
		"message":     "Wallet recovered successfully",
	}
	jsonResponse(w, response)
}

// 4. Get Balance
func handleGetBalance(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	
	mutex.Lock()
	bal := balances[address]
	mutex.Unlock()

	jsonResponse(w, map[string]interface{}{
		"address": address, 
		"balance": bal, 
		"symbol": "QBNB", // Updated symbol
	})
}

// 5. Faucet (Untuk Testing)
func handleFaucet(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	
	mutex.Lock()
	balances[address] += 100.0
	
	// Catat Transaksi Faucet
	txHash := fmt.Sprintf("0x%x", time.Now().UnixNano())
	newTx := Transaction{
		TxHash:    txHash, 
		From:      "FAUCET_SYSTEM", 
		To:        address, 
		Amount:    100.0,
		Timestamp: time.Now().Format(time.RFC3339), 
		Status:    "Success",
	}
	txHistory = append(txHistory, newTx)
	mutex.Unlock()
	
	SaveData()

	jsonResponse(w, map[string]interface{}{
		"message": "Faucet Success", 
		"new_balance": balances[address],
		"tx_hash": txHash,
	})
}

// 6. Send Transaction
func handleSendTx(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amountStr := r.URL.Query().Get("amount")
	
	var amount float64
	fmt.Sscanf(amountStr, "%f", &amount)

	mutex.Lock()
	defer mutex.Unlock()

	if balances[from] < amount {
		jsonResponse(w, map[string]interface{}{"status": "failed", "error": "Insufficient Funds"})
		return
	}

	// Eksekusi Transfer
	balances[from] -= amount
	balances[to] += amount

	txHash := "0x" + fmt.Sprintf("%x", time.Now().UnixNano())
	newTx := Transaction{
		TxHash:    txHash, 
		From:      from, 
		To:        to, 
		Amount:    amount,
		Timestamp: time.Now().Format(time.RFC3339), 
		Status:    "Success",
	}
	
	// Tambah ke history (paling baru di atas)
	txHistory = append([]Transaction{newTx}, txHistory...)
	
	// Simpan ke File
	data := LedgerData{Balances: balances, TxHistory: txHistory}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(DB_FILE, file, 0644)

	log.Printf("💸 [TX] %s -> %s | %.2f QBNB", from, to, amount)
	jsonResponse(w, map[string]interface{}{
		"status": "success", 
		"tx_hash": txHash, 
		"new_balance": balances[from],
	})
}

// 7. Get History
func handleGetHistory(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	myHistory := []Transaction{}
	
	mutex.Lock()
	for _, tx := range txHistory {
		if tx.From == address || tx.To == address {
			myHistory = append(myHistory, tx)
		}
	}
	mutex.Unlock()

	jsonResponse(w, myHistory)
}

// 8. Get Tx Detail
func handleGetTxDetail(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	
	mutex.Lock()
	defer mutex.Unlock()

	for _, tx := range txHistory {
		if tx.TxHash == hash {
			jsonResponse(w, tx)
			return
		}
	}
	jsonResponse(w, map[string]string{"error": "Transaction Not Found"})
}

// --- HELPER FUNCTIONS ---

// jsonResponse mengirimkan data JSON dengan header CORS yang benar
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// PENTING: Mengizinkan akses dari Frontend React (Port berapapun)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("❌ [API] Failed to encode JSON response")
	}
}
