package coin

import "errors"

var ErrInsufficientBalance = errors.New("insufficient balance")

type BalanceKeeper struct {
	store BalanceStore
}

func NewBalanceKeeper(store BalanceStore) *BalanceKeeper {
	return &BalanceKeeper{store: store}
}

func (bk *BalanceKeeper) BalanceOf(addr string) Amount {
	return bk.store.Get(addr)
}

func (bk *BalanceKeeper) Credit(addr string, amt Amount) {
	if amt.IsNegative() {
		panic("credit negative amount")
	}
	cur := bk.store.Get(addr)
	bk.store.Set(addr, cur.Add(amt))
}

func (bk *BalanceKeeper) Debit(addr string, amt Amount) error {
	if amt.IsNegative() {
		panic("debit negative amount")
	}
	cur := bk.store.Get(addr)
	if cur.Cmp(amt) < 0 {
		return ErrInsufficientBalance
	}
	bk.store.Set(addr, cur.Sub(amt))
	return nil
}

func (bk *BalanceKeeper) Transfer(from, to string, amt Amount) error {
	if err := bk.Debit(from, amt); err != nil {
		return err
	}
	bk.Credit(to, amt)
	return nil
}
