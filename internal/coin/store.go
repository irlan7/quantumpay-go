package coin

type BalanceStore interface {
	Get(addr string) Amount
	Set(addr string, amt Amount)
}

type memStore struct {
	data map[string]Amount
}

func NewMemStore() BalanceStore {
	return &memStore{data: make(map[string]Amount)}
}

func (m *memStore) Get(addr string) Amount {
	if v, ok := m.data[addr]; ok {
		return v
	}
	return Zero()
}

func (m *memStore) Set(addr string, amt Amount) {
	m.data[addr] = amt
}
