package coin

import "testing"

func TestBalanceTransfer(t *testing.T) {
	store := NewMemStore()
	bk := NewBalanceKeeper(store)

	bk.Credit("alice", NewAmountFromInt64(100))

	if err := bk.Transfer("alice", "bob", NewAmountFromInt64(40)); err != nil {
		t.Fatal(err)
	}

	if bk.BalanceOf("alice").String() != "60" {
		t.Fatal("alice balance wrong")
	}

	if bk.BalanceOf("bob").String() != "40" {
		t.Fatal("bob balance wrong")
	}
}
