package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	// Pastikan path ini benar sesuai go.mod Bapak
	"github.com/irlan/quantumpay-go/internal/crypto"
)

const DB_FILE = "ledger_data.json"

type ServerConfig struct {
	Port string
}

type LedgerData struct {
	Balances  map[string]map[string]float64 `json:"balances"`
	TxHistory []Transaction                 `json:"history"`
}

type Transaction struct {
	TxHash    string                   `json:"tx_hash"`
	From      string                   `json:"from"`
	To        string                   `json:"to"`
	Amount    float64                  `json:"amount"`
	Symbol    string                   `json:"symbol"`
	Timestamp string                   `json:"timestamp"`
	Status    string                   `json:"status"`
	// PQC FIELD: Menyimpan bukti keamanan quantum
	Security  crypto.SovereignSignature `json:"security"`
}

var (
	balances  = make(map[string]map[string]float64)
	txHistory = []Transaction{}
	mutex     = &sync.Mutex{}
	adminKey  = "QUANTUM_SECRET_2026"
)

// --- DATABASE LOAD/SAVE ---

func LoadData() {
	file, err := os.ReadFile(DB_FILE)
	if err != nil {
		log.Println("⚠️ [DB] Membuat Ledger Baru.")
		return
	}
	var data LedgerData
	if err := json.Unmarshal(file, &data); err == nil {
		if data.Balances != nil { balances = data.Balances }
		txHistory = data.TxHistory
		log.Printf("📂 [DB] LOADED: %d Akun", len(balances))
	}
}

func SaveData() {
	mutex.Lock()
	defer mutex.Unlock()
	data := LedgerData{Balances: balances, TxHistory: txHistory}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(DB_FILE, file, 0644)
	log.Println("💾 [DB] Data Tersimpan Aman.")
}

// --- SERVER START ---

func StartServer(config *ServerConfig) {
	LoadData()
	mux := http.NewServeMux()
	
	// Daftarkan semua route
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/wallet/create", handleCreateWallet)
	mux.HandleFunc("/wallet/import", handleImportWallet)
	mux.HandleFunc("/balances", handleGetBalances)
	mux.HandleFunc("/faucet", handleFaucet)
	mux.HandleFunc("/tx/send", handleSendTx)
	mux.HandleFunc("/admin/mint-native", handleAdminMint)
	mux.HandleFunc("/tx/history", handleGetHistory)
	mux.HandleFunc("/tx/detail", handleGetTxDetail)

	log.Printf("🚀 [API] QuantumPay Node Active on %s", config.Port)
	if err := http.ListenAndServe(config.Port, mux); err != nil {
		log.Fatal(err)
	}
}

// --- HANDLERS (ANTI ERROR UNDEFINED) ---

func handleHealth(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w); jsonResponse(w, map[string]string{"status": "online", "security": "PQC_ENABLED"})
}

func handleAdminMint(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w); if r.Method == "OPTIONS" { return }

	var req struct {
		To string `json:"to"`; Amount float64 `json:"amount"`; Symbol string `json:"symbol"`; Key string `json:"key"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, "Error", 400); return }
	if req.Key != adminKey { http.Error(w, "Unauthorized", 401); return }

	mutex.Lock()
	if _, exists := balances[req.To]; !exists { balances[req.To] = make(map[string]float64) }
	balances[req.To][req.Symbol] += req.Amount
	mutex.Unlock()

	// --- INTEGRASI ANTI-QUANTUM ---
	txHash := fmt.Sprintf("mint_%d", time.Now().UnixNano())
	// Panggil fungsi SignHybrid dari modul crypto
	pqcProof := crypto.SignHybrid(txHash, "ADMIN_KEY")

	newTx := Transaction{
		TxHash: txHash, From: "VAULT", To: req.To, Amount: req.Amount, Symbol: req.Symbol,
		Timestamp: time.Now().Format(time.RFC3339), Status: "Success",
		Security: pqcProof, // Simpan bukti quantum ke ledger
	}
	txHistory = append([]Transaction{newTx}, txHistory...)
	SaveData()
	jsonResponse(w, map[string]string{"status": "success", "tx_hash": txHash})
}

func handleSendTx(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w); if r.Method == "OPTIONS" { return }
	
	from, to, sym := r.URL.Query().Get("from"), r.URL.Query().Get("to"), r.URL.Query().Get("symbol")
	if sym == "" { sym = "QPAY" }
	var amt float64
	fmt.Sscanf(r.URL.Query().Get("amount"), "%f", &amt)

	mutex.Lock()
	if balances[from] == nil || balances[from][sym] < amt { 
		mutex.Unlock(); jsonResponse(w, map[string]string{"error": "Saldo Kurang"}); return 
	}
	balances[from][sym] -= amt
	if _, exists := balances[to]; !exists { balances[to] = make(map[string]float64) }
	balances[to][sym] += amt
	mutex.Unlock()

	// --- INTEGRASI ANTI-QUANTUM ---
	txHash := fmt.Sprintf("tx_%d", time.Now().UnixNano())
	pqcProof := crypto.SignHybrid(txHash, "USER_KEY")

	newTx := Transaction{
		TxHash: txHash, From: from, To: to, Amount: amt, Symbol: sym,
		Timestamp: time.Now().Format(time.RFC3339), Status: "Success",
		Security: pqcProof,
	}
	txHistory = append([]Transaction{newTx}, txHistory...)
	SaveData()
	jsonResponse(w, map[string]string{"status": "success", "tx_hash": txHash})
}

func handleGetBalances(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w); if r.Method == "OPTIONS" { return }
	addr := r.URL.Query().Get("address")
	mutex.Lock()
	res, exists := balances[addr]
	if !exists { res = map[string]float64{"QPAY": 0} }
	mutex.Unlock()
	jsonResponse(w, res)
}

func handleCreateWallet(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w)
	// Panggil crypto lokal
	wallet, _ := crypto.NewKeyPair()
	addr := wallet.Address()
	mutex.Lock()
	if _, exists := balances[addr]; !exists { balances[addr] = map[string]float64{"QPAY": 10.0} }
	mutex.Unlock()
	SaveData()
	jsonResponse(w, map[string]string{"address": addr, "mnemonic": wallet.Mnemonic, "private_key": wallet.PrivateKeyHex()})
}

func handleImportWallet(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w); phrase := r.URL.Query().Get("phrase")
	wallet, _ := crypto.NewKeyPairFromMnemonic(phrase)
	jsonResponse(w, map[string]string{"address": wallet.Address(), "message": "Success"})
}

func handleFaucet(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w); addr := r.URL.Query().Get("address")
	mutex.Lock()
	if _, exists := balances[addr]; !exists { balances[addr] = make(map[string]float64) }
	balances[addr]["QPAY"] += 100.0
	mutex.Unlock()
	SaveData()
	jsonResponse(w, map[string]string{"message": "Success"})
}

func handleGetHistory(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w); addr := r.URL.Query().Get("address")
	userHistory := []Transaction{}
	for _, tx := range txHistory {
		if tx.From == addr || tx.To == addr { userHistory = append(userHistory, tx) }
	}
	jsonResponse(w, userHistory)
}

func handleGetTxDetail(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w); hash := r.URL.Query().Get("hash")
	for _, tx := range txHistory {
		if tx.TxHash == hash { jsonResponse(w, tx); return }
	}
	http.Error(w, "Not Found", 404)
}

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
