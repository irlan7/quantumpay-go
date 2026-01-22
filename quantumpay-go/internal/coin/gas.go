package coin

import (
	"errors"
)

// ----------------------------
// Gas Parameters (v1)
// ----------------------------

const (
	// Base gas price in QP smallest unit
	BaseGasPrice uint64 = 1

	// Gas burn ratio (30%)
	GasBurnRatioNumerator   uint64 = 30
	GasBurnRatioDenominator uint64 = 100
)

// ----------------------------
// Gas Usage Model
// ----------------------------

// GasUsage represents gas consumed by a transaction
type GasUsage struct {
	GasUsed uint64
}

// GasFeeResult represents the outcome of gas accounting
type GasFeeResult struct {
	TotalFee        uint64
	BurnedAmount    uint64
	ValidatorReward uint64
}

// ----------------------------
// Core Functions
// ----------------------------

// CalculateGasFee calculates total gas fee in QP
func CalculateGasFee(gasUsed uint64) uint64 {
	return gasUsed * BaseGasPrice
}

// SplitGasFee splits gas fee into burned and validator reward
func SplitGasFee(totalFee uint64) GasFeeResult {
	burned := (totalFee * GasBurnRatioNumerator) / GasBurnRatioDenominator
	reward := totalFee - burned

	return GasFeeResult{
		TotalFee:        totalFee,
		BurnedAmount:    burned,
		ValidatorReward: reward,
	}
}

// ValidateGasPayment ensures sender can pay gas
func ValidateGasPayment(balance uint64, gasUsed uint64) error {
	required := CalculateGasFee(gasUsed)
	if balance < required {
		return errors.New("insufficient balance to pay gas fee")
	}
	return nil
}
