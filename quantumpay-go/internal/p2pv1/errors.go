package p2pv1

import "errors"

var (
	ErrInvalidHeight   = errors.New("invalid height")
	ErrEmptyHash       = errors.New("empty hash")
	ErrMissingParent   = errors.New("missing parent hash")
	ErrFutureTimestamp = errors.New("timestamp in future")
)
