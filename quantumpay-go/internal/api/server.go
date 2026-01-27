package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	// Pastikan path ini benar
	"github.com/irlan/quantumpay-go/internal/crypto"
)

// KONFIGURASI FILE DB
const DB_FILE = "ledger_data.json"

type ServerConfig struct {
	Port string
}

type LedgerData struct {
	Balances  map[string]float64 `json:"balances"`
	TxHistory []Transaction      `json:"history"`
}

type Transaction struct {
	TxHash    string  `json:"tx_hash"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
	Status    string  `json:"status"`
}

// GLOBAL STATE
var (
	balances  = map[string]float64{"0x00000000000000000000": 1000000}
	txHistory = []Transaction{}
	mutex     = &sync.Mutex{}
)

// --- DB FUNCTIONS ---
func SaveData() {
	mutex.Lock()
	defer mutex.Unlock()
	data := LedgerData{Balances: balances, TxHistory: txHistory}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(DB_FILE, file, 0644)
}

func LoadData() {
	file, err := os.ReadFile(DB_FILE)
	if err != nil { return }
	var data LedgerData
	if err := json.Unmarshal(file, &data); err == nil {
		balances = data.Balances
		txHistory = data.TxHistory
	}
	log.Printf("📂 [DB] LOADED: %d Acc | %d Tx", len(balances), len(txHistory))
}

// --- SERVER SETUP ---
func StartServer(config *ServerConfig) {
	LoadData()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	
	// WALLET ROUTES
	mux.HandleFunc("/wallet/create", handleCreateWallet)
	mux.HandleFunc("/wallet/import", handleImportWallet) // <--- ROUTE BARU

	// FINANCE ROUTES
	mux.HandleFunc("/balance", handleGetBalance)
	mux.HandleFunc("/faucet", handleFaucet)
	mux.HandleFunc("/tx/send", handleSendTx)
	
	// EXPLORER ROUTES
	mux.HandleFunc("/tx/history", handleGetHistory)
	mux.HandleFunc("/tx/detail", handleGetTxDetail)

	log.Printf("🚀 [API] Server listening on %s", config.Port)
	if err := http.ListenAndServe(config.Port, mux); err != nil {
		log.Printf("🔥 [API] Error: %v", err)
	}
}

// --- HANDLERS ---

func handleHealth(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, map[string]string{"status": "online", "ver": "2.2"})
}

func handleCreateWallet(w http.ResponseWriter, r *http.Request) {
	wallet, err := crypto.NewKeyPair()
	if err != nil { http.Error(w, "Error", 500); return }
	
	// Cek apakah alamat ini sudah pernah ada? Kalau belum, kasih bonus.
	if _, exists := balances[wallet.Address()]; !exists {
		balances[wallet.Address()] = 10.0
		SaveData()
	}

	response := map[string]string{
		"address":     wallet.Address(),
		"public_key":  wallet.GetPublicKeyHex(),
		"private_key": wallet.PrivateKeyHex(),
		"mnemonic":    wallet.Mnemonic,
		"created_at":  time.Now().Format(time.RFC3339),
	}
	jsonResponse(w, response)
}

// HANDLER BARU: IMPORT WALLET DARI 12 KATA
func handleImportWallet(w http.ResponseWriter, r *http.Request) {
	phrase := r.URL.Query().Get("phrase")
	if phrase == "" {
		http.Error(w, "Mnemonic phrase required", 400)
		return
	}

	// Panggil crypto untuk memulihkan kunci
	wallet, err := crypto.NewKeyPairFromMnemonic(phrase)
	if err != nil {
		jsonResponse(w, map[string]string{"error": "Invalid Mnemonic Phrase"})
		return
	}

	// Cek saldo yang tersimpan (jangan direset ke 0 atau 10, ambil apa adanya)
	// Jika user baru import tapi datanya belum ada di server ini, default 0
	if _, exists := balances[wallet.Address()]; !exists {
		balances[wallet.Address()] = 0.0 
	}

	response := map[string]string{
		"address":     wallet.Address(),
		"public_key":  wallet.GetPublicKeyHex(),
		"private_key": wallet.PrivateKeyHex(),
		"mnemonic":    phrase, // Balikin lagi frasanya
		"message":     "Wallet recovered successfully",
	}
	jsonResponse(w, response)
}

func handleGetBalance(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	jsonResponse(w, map[string]interface{}{"address": address, "balance": balances[address], "symbol": "QPY"})
}

func handleFaucet(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	balances[address] += 100.0
	
	txHash := fmt.Sprintf("0x%x", time.Now().UnixNano())
	newTx := Transaction{
		TxHash: txHash, From: "FAUCET", To: address, Amount: 100.0,
		Timestamp: time.Now().Format(time.RFC3339), Status: "Success",
	}
	txHistory = append(txHistory, newTx)
	SaveData()

	jsonResponse(w, map[string]interface{}{"message": "Success", "new_balance": balances[address]})
}

func handleSendTx(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amountStr := r.URL.Query().Get("amount")
	var amount float64
	fmt.Sscanf(amountStr, "%f", &amount)

	if balances[from] < amount {
		jsonResponse(w, map[string]interface{}{"status": "failed", "error": "Saldo Kurang"})
		return
	}

	balances[from] -= amount
	balances[to] += amount

	txHash := "0x" + fmt.Sprintf("%x", time.Now().UnixNano())
	newTx := Transaction{
		TxHash: txHash, From: from, To: to, Amount: amount,
		Timestamp: time.Now().Format(time.RFC3339), Status: "Success",
	}
	txHistory = append([]Transaction{newTx}, txHistory...)
	SaveData()

	log.Printf("💸 [TX] %s -> %s | %.2f", from, to, amount)
	jsonResponse(w, map[string]interface{}{"status": "success", "tx_hash": txHash, "new_balance": balances[from]})
}

func handleGetHistory(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	myHistory := []Transaction{}
	for _, tx := range txHistory {
		if tx.From == address || tx.To == address {
			myHistory = append(myHistory, tx)
		}
	}
	jsonResponse(w, myHistory)
}

func handleGetTxDetail(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	for _, tx := range txHistory {
		if tx.TxHash == hash {
			jsonResponse(w, tx)
			return
		}
	}
	jsonResponse(w, map[string]string{"error": "Not Found"})
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}
