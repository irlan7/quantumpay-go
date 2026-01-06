package p2pv1

import "time"

func ValidateHeaderBasic(h HeaderMsg, pool *HeaderPool) error {
	if h.Height == 0 {
		return ErrInvalidHeight
	}

	if h.Height > 1 && !pool.Has(h.ParentHash) {
		return ErrMissingParent
	}

	if h.Timestamp > time.Now().Unix()+30 {
		return ErrFutureTimestamp
	}

	return nil
}
