package staking

import "errors"

// -----------------------------
// Types
// -----------------------------

// RewardSource represents rewards available in an epoch
type RewardSource struct {
	GasRewards uint64 // gas fees after burn
	Incentives uint64 // protocol incentives (optional)
}

// ValidatorParams defines validator economic parameters
type ValidatorParams struct {
	CommissionNumerator   uint64
	CommissionDenominator uint64
}

// StakeSnapshot represents stake distribution at epoch end
type StakeSnapshot struct {
	ValidatorStake uint64
	DelegatorStake uint64
}

// RewardDistribution is the final reward outcome
type RewardDistribution struct {
	ValidatorReward uint64
	DelegatorReward uint64
}

// -----------------------------
// Core Logic
// -----------------------------

// CalculateTotalRewards sums all reward sources
func CalculateTotalRewards(src RewardSource) uint64 {
	return src.GasRewards + src.Incentives
}

// SplitValidatorCommission applies validator commission
func SplitValidatorCommission(
	totalRewards uint64,
	params ValidatorParams,
) (validatorCut uint64, remaining uint64, err error) {

	if params.CommissionDenominator == 0 {
		return 0, 0, errors.New("invalid commission denominator")
	}

	validatorCut = (totalRewards * params.CommissionNumerator) / params.CommissionDenominator
	remaining = totalRewards - validatorCut
	return validatorCut, remaining, nil
}

// DistributeRewards splits rewards between validator and delegators
func DistributeRewards(
	src RewardSource,
	params ValidatorParams,
	stake StakeSnapshot,
) (RewardDistribution, error) {

	totalRewards := CalculateTotalRewards(src)

	if totalRewards == 0 {
		return RewardDistribution{}, nil
	}

	validatorCommission, delegatorPool, err :=
		SplitValidatorCommission(totalRewards, params)
	if err != nil {
		return RewardDistribution{}, err
	}

	// No delegators → validator takes all
	if stake.DelegatorStake == 0 {
		return RewardDistribution{
			ValidatorReward: totalRewards,
			DelegatorReward: 0,
		}, nil
	}

	// Delegators receive remaining rewards
	return RewardDistribution{
		ValidatorReward: validatorCommission,
		DelegatorReward: delegatorPool,
	}, nil
}
