package mempool

import (
	"testing"

	"github.com/irlan/quantumpay-go/internal/types"
)

func TestPolicyValidateGas(t *testing.T) {
	p := &Policy{
		Gas: types.GasSpec{
			MinGas: 100,
			MaxGas: 1000,
		},
	}

	if err := p.ValidateGas(50); err == nil {
		t.Fatal("expected error for gas below minimum")
	}

	if err := p.ValidateGas(1500); err == nil {
		t.Fatal("expected error for gas above maximum")
	}

	if err := p.ValidateGas(500); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
