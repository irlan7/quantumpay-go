package coin

import "testing"

func TestAmountBasic(t *testing.T) {
	a := NewAmountFromInt64(10)
	b := NewAmountFromInt64(5)

	if a.Add(b).String() != "15" {
		t.Fatal("add failed")
	}

	if a.Sub(b).String() != "5" {
		t.Fatal("sub failed")
	}

	if !Zero().IsZero() {
		t.Fatal("zero failed")
	}
}
