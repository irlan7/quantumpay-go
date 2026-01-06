package state

import "errors"

// ==========================
// State Transition Errors
// ==========================

// ErrInsufficientBalance is returned when account balance
// is lower than required amount (value + gas)
var ErrInsufficientBalance = errors.New("insufficient balance")

// ErrInvalidNonce is returned when tx nonce
// does not match account nonce
var ErrInvalidNonce = errors.New("invalid nonce")

// ErrGasLimitExceeded is returned when tx gas limit
// is lower than required execution cost
var ErrGasLimitExceeded = errors.New("gas limit exceeded")

// ErrGasPriceTooLow is returned when tx gas price
// does not meet minimum network requirement
var ErrGasPriceTooLow = errors.New("gas price too low")
