package service

// ChainAPI adalah kontrak agar Service bisa bicara dengan Blockchain
// Menggunakan 'any' adalah KUNCI ANTI IMPORT CYCLE
type ChainAPI interface {
	// Mendapatkan tinggi blok saat ini
	Height() uint64
	
	// Wajib 'any' agar cocok dengan Adapter di main.go
	GetBlockByHeight(height uint64) (any, bool)
}
