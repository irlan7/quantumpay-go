package coin

type GenesisAllocation struct {
	Address string
	Amount  Amount
	Type    string
}

type GenesisState struct {
	Allocations []GenesisAllocation
	TotalSupply Amount
}
