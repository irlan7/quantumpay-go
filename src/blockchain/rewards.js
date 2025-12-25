const delegation = require("./delegation");

/**
 * Distribute rewards with delegation support
 */
function distributeRewardsWithDelegation({
  proposer,
  voters,
  totalFees,
  emissionEnabled = true,
  commissionRate = 0.10 // 10%
}) {
  if (!proposer) return;

  const baseReward =
    totalFees + (emissionEnabled ? EMISSION_PER_BLOCK : 0);

  // === COMMISSION ===
  const commission = Math.floor(baseReward * commissionRate);
  credit(proposer, commission);

  // === DELEGATOR POOL ===
  const delegatorPool = baseReward - commission;

  const delegators = delegation.getDelegators(proposer);
  const totalDelegated = delegation.getDelegatedStake(proposer);

  if (delegators.length > 0 && totalDelegated > 0) {
    delegators.forEach(d => {
      const share = Math.floor(
        (d.amount / totalDelegated) * delegatorPool
      );
      d.pendingRewards += share;
    });
  } else {
    // no delegators → validator gets all
    credit(proposer, delegatorPool);
  }
}

module.exports.distributeRewardsWithDelegation =
  distributeRewardsWithDelegation;

