package mempool

import (
	"errors"

	"github.com/irlan/quantumpay-go/internal/types"
)

type Policy struct {
	Gas types.GasSpec
}

func DefaultPolicy() *Policy {
	return &Policy{
		Gas: types.DefaultGasSpec(),
	}
}

func (p *Policy) ValidateGas(gas uint64) error {
	if gas < p.Gas.MinGas {
		return errors.New("gas below minimum")
	}
	if gas > p.Gas.MaxGas {
		return errors.New("gas above maximum")
	}
	return nil
}
