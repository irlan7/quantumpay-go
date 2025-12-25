// src/consensus/leaderElection.js
const crypto = require('crypto');

class LeaderElection {
  constructor(validatorSet) {
    this.validatorSet = validatorSet;
  }

  electLeader({ height, lastBlockHash }) {
    const validators = this.validatorSet.getActiveValidators();
    if (validators.length === 0) return null;

    const totalStake = this.validatorSet.totalStake();

    // Deterministic random seed
    const seed = crypto
      .createHash('sha256')
      .update(lastBlockHash + height)
      .digest('hex');

    // Convert seed to number
    const rand = parseInt(seed.slice(0, 16), 16) / 0xffffffffffffffff;

    let cumulative = 0;
    for (const v of validators) {
      cumulative += v.stake / totalStake;
      if (rand <= cumulative) {
        return v.address;
      }
    }

    // Fallback (should not happen)
    return validators[validators.length - 1].address;
  }
}

module.exports = LeaderElection;

