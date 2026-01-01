package core

import "errors"

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidNonce        = errors.New("invalid nonce")
)
