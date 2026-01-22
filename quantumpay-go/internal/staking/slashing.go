package staking

import "errors"

// -------------------------------
// Slashing Parameters (v1 - Locked)
// -------------------------------

const (
	// Slashing ratios are expressed as fractions
	// Example: 5% = 5 / 100
	DoubleSignSlashNumerator   uint64 = 5
	DoubleSignSlashDenominator uint64 = 100

	DowntimeSlashNumerator   uint64 = 1
	DowntimeSlashDenominator uint64 = 100

	// Burn ratio of slashed amount (e.g. 50%)
	SlashBurnNumerator   uint64 = 50
	SlashBurnDenominator uint64 = 100
)

// -------------------------------
// Slashing Types
// -------------------------------

type SlashReason int

const (
	SlashDoubleSign SlashReason = iota + 1
	SlashDowntime
)

// -------------------------------
// Slashing Result
// -------------------------------

type SlashResult struct {
	SlashedAmount uint64
	BurnedAmount  uint64
	RewardAmount  uint64
	RemainingStake uint64
}

// -------------------------------
// Core Slashing Logic
// -------------------------------

// CalculateSlashAmount computes slashed amount based on reason
func CalculateSlashAmount(stake uint64, reason SlashReason) (uint64, error) {
	if stake == 0 {
		return 0, errors.New("cannot slash zero stake")
	}

	switch reason {
	case SlashDoubleSign:
		return (stake * DoubleSignSlashNumerator) / DoubleSignSlashDenominator, nil
	case SlashDowntime:
		return (stake * DowntimeSlashNumerator) / DowntimeSlashDenominator, nil
	default:
		return 0, errors.New("unknown slashing reason")
	}
}

// ApplySlashing applies slashing and returns economic breakdown
func ApplySlashing(stake uint64, reason SlashReason) (SlashResult, error) {
	slashAmount, err := CalculateSlashAmount(stake, reason)
	if err != nil {
		return SlashResult{}, err
	}

	if slashAmount > stake {
		return SlashResult{}, errors.New("slashing exceeds stake")
	}

	burned := (slashAmount * SlashBurnNumerator) / SlashBurnDenominator
	reward := slashAmount - burned
	remaining := stake - slashAmount

	return SlashResult{
		SlashedAmount:  slashAmount,
		BurnedAmount:  burned,
		RewardAmount:  reward,
		RemainingStake: remaining,
	}, nil
}
