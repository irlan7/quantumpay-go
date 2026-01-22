package coin

import "errors"

// CoinLedger manages balances and total supply
type CoinLedger struct {
	balances    map[string]uint64
	totalSupply uint64
	maxSupply   uint64
}

// NewCoinLedger initializes ledger
func NewCoinLedger(maxSupply uint64) *CoinLedger {
	return &CoinLedger{
		balances:  make(map[string]uint64),
		maxSupply: maxSupply,
	}
}

// --------------------
// Read-only API
// --------------------

func (c *CoinLedger) BalanceOf(addr string) uint64 {
	return c.balances[addr]
}

func (c *CoinLedger) TotalSupply() uint64 {
	return c.totalSupply
}

// --------------------
// Mint
// --------------------

func (c *CoinLedger) Mint(to string, amount uint64) error {
	if amount == 0 {
		return errors.New("mint amount must be > 0")
	}
	if c.totalSupply+amount > c.maxSupply {
		return errors.New("mint exceeds max supply")
	}

	c.balances[to] += amount
	c.totalSupply += amount
	return nil
}

// --------------------
// Burn
// --------------------

func (c *CoinLedger) Burn(from string, amount uint64) error {
	if amount == 0 {
		return errors.New("burn amount must be > 0")
	}
	if c.balances[from] < amount {
		return errors.New("insufficient balance to burn")
	}

	c.balances[from] -= amount
	c.totalSupply -= amount
	return nil
}

// --------------------
// Transfer
// --------------------

func (c *CoinLedger) Transfer(from, to string, amount uint64) error {
	if from == to {
		return errors.New("cannot transfer to self")
	}
	if amount == 0 {
		return errors.New("transfer amount must be > 0")
	}
	if c.balances[from] < amount {
		return errors.New("insufficient balance")
	}

	c.balances[from] -= amount
	c.balances[to] += amount
	return nil
}
