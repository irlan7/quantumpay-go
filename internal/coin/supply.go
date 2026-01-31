package coin

import "errors"

var ErrSupplyExceeded = errors.New("supply limit exceeded")

type SupplyKeeper struct {
	total Amount
	max   Amount
}

func NewSupplyKeeper(max Amount) *SupplyKeeper {
	return &SupplyKeeper{total: Zero(), max: max}
}

func (s *SupplyKeeper) Mint(a Amount) error {
	next := s.total.Add(a)
	if next.Cmp(s.max) > 0 {
		return ErrSupplyExceeded
	}
	s.total = next
	return nil
}

func (s *SupplyKeeper) Total() Amount {
	return s.total
}
