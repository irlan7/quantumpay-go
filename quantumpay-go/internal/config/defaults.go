package config

// defaults.go
//
// Package config berisi *default parameters* yang bersifat:
// - stateless
// - tidak bergantung pada package lain (anti import-cycle)
// - opsional (node / engine boleh override)
//
// File ini TIDAK:
// - mengimpor mempool
// - mengimpor state
// - mengimpor execution
// - mengimpor p2p
//
// Aman untuk build dan future-proof.

import "time"

// GasDefaults mendefinisikan nilai default gas / fee minimum.
// Struct ini hanya data, TANPA logic.
type GasDefaults struct {
	// BaseGas adalah biaya minimum per transaksi
	BaseGas uint64

	// GasPerByte adalah biaya tambahan per byte payload tx
	GasPerByte uint64

	// MaxGasPerTx adalah batas atas gas tx tunggal
	MaxGasPerTx uint64

	// BlockGasLimit adalah total gas maksimum per block
	BlockGasLimit uint64
}

// MempoolDefaults berisi parameter dasar kebijakan mempool.
// Tidak ada dependency ke mempool package.
type MempoolDefaults struct {
	MaxTxPerAccount uint64
	MaxTotalTx      uint64
	MinGasPrice     uint64
}

// ConsensusDefaults (opsional) – disiapkan untuk future use.
// Tidak dipakai sekarang, tapi tidak merusak build.
type ConsensusDefaults struct {
	BlockTime       time.Duration
	MaxFutureBlocks uint64
}

// DefaultGas adalah default global gas spec.
// Digunakan jika node tidak override via config / flag.
var DefaultGas = GasDefaults{
	BaseGas:       21_000,     // mirip Ethereum baseline
	GasPerByte:    16,         // konservatif
	MaxGasPerTx:   5_000_000,  // anti spam
	BlockGasLimit: 30_000_000, // cukup longgar
}

// DefaultMempool adalah kebijakan default mempool.
var DefaultMempool = MempoolDefaults{
	MaxTxPerAccount: 64,
	MaxTotalTx:      50_000,
	MinGasPrice:     1, // unit terkecil
}

// DefaultConsensus hanya placeholder aman.
var DefaultConsensus = ConsensusDefaults{
	BlockTime:       5 * time.Second,
	MaxFutureBlocks: 2,
}
