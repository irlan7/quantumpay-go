package coin

type RewardCalculator interface {
	BlockReward(height uint64) Amount
}
