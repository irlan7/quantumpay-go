package types

import "testing"

func TestDefaultGasSpec(t *testing.T) {
	spec := DefaultGasSpec()

	if spec.MinGas == 0 {
		t.Fatal("MinGas must be > 0")
	}
	if spec.MaxGas == 0 {
		t.Fatal("MaxGas must be > 0")
	}
	if spec.MinGas > spec.MaxGas {
		t.Fatalf("MinGas (%d) must not exceed MaxGas (%d)", spec.MinGas, spec.MaxGas)
	}
}
