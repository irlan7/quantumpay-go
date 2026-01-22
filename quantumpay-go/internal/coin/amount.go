package coin

import "math/big"

type Amount struct {
	i *big.Int
}

func Zero() Amount {
	return Amount{big.NewInt(0)}
}

func NewAmountFromInt64(v int64) Amount {
	if v < 0 {
		panic("amount cannot be negative")
	}
	return Amount{big.NewInt(v)}
}

func (a Amount) Int() *big.Int {
	return new(big.Int).Set(a.i)
}

func (a Amount) IsNegative() bool {
	return a.i.Sign() < 0
}

func (a Amount) IsZero() bool {
	return a.i.Sign() == 0
}

func (a Amount) Add(b Amount) Amount {
	return Amount{new(big.Int).Add(a.i, b.i)}
}

func (a Amount) Sub(b Amount) Amount {
	if a.i.Cmp(b.i) < 0 {
		panic("amount underflow")
	}
	return Amount{new(big.Int).Sub(a.i, b.i)}
}

func (a Amount) Cmp(b Amount) int {
	return a.i.Cmp(b.i)
}

func (a Amount) String() string {
	return a.i.String()
}
