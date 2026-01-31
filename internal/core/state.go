package core

type Account struct {
	Balance uint64
	Nonce   uint64
}

type WorldState struct {
	Accounts map[string]*Account
}

func NewWorldState() *WorldState {
	return &WorldState{
		Accounts: make(map[string]*Account),
	}
}
