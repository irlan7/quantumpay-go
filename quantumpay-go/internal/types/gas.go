package types

// GasSpec adalah konfigurasi batas gas (policy-level, BUKAN runtime).
type GasSpec struct {
	MinGas uint64
	MaxGas uint64
}

// DefaultGasSpec mengembalikan konfigurasi gas default jaringan.
func DefaultGasSpec() GasSpec {
	return GasSpec{
		MinGas: 21_000,
		MaxGas: 10_000_000,
	}
}
