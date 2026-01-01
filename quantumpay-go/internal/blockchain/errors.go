package blockchain

import "errors"

// ErrStateTransition muncul jika transaksi gagal diterapkan ke state
var ErrStateTransition = errors.New("state transition failed")
