package state

import "errors"

// GasMeter adalah interface runtime untuk konsumsi gas.
// Diletakkan di execution agar state logic tidak tergantung ke types/mempool.
type GasMeter interface {
	Consume(amount uint64) error
	Remaining() uint64
}

// SimpleGasMeter adalah implementasi dasar GasMeter.
type SimpleGasMeter struct {
	remaining uint64
}

func NewSimpleGasMeter(limit uint64) *SimpleGasMeter {
	return &SimpleGasMeter{remaining: limit}
}

func (g *SimpleGasMeter) Consume(amount uint64) error {
	if amount > g.remaining {
		return errors.New("out of gas")
	}
	g.remaining -= amount
	return nil
}

func (g *SimpleGasMeter) Remaining() uint64 {
	return g.remaining
}
