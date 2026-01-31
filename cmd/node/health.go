package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse adalah struktur data untuk respon API
type HealthResponse struct {
	Status    string `json:"status"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
}

// HealthCheckHandler menangani permintaan cek status server
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Header CORS wajib ada agar React bisa akses
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	// Respon jika method OPTIONS (Preflight browser)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	response := HealthResponse{
		Status:    "online",
		Version:   "4.5-sovereign",
		Timestamp: time.Now().Format(time.RFC3339),
		Service:   "QuantumPay-Node",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
