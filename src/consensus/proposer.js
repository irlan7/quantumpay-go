// src/consensus/proposer.js

const crypto = require('crypto');

class ProposerSelector {
  constructor({ validators }) {
    this.validators = validators;
  }

  getActiveValidators() {
    return Object.entries(this.validators)
      .filter(([, v]) => !v.jailed && v.stake > 0)
      .map(([id, v]) => ({ id, stake: v.stake }));
  }

  // ===== MODE A: ROUND-ROBIN =====
  roundRobin(height) {
    const list = this.getActiveValidators();
    if (list.length === 0) return null;
    const index = height % list.length;
    return list[index].id;
  }

  // ===== MODE B: STAKE-WEIGHTED (deterministic) =====
  stakeWeighted(height) {
    const list = this.getActiveValidators();
    if (list.length === 0) return null;

    const totalStake = list.reduce((s, v) => s + v.stake, 0);
    const seed = crypto
      .createHash('sha256')
      .update(String(height))
      .digest('hex');

    // deterministic random in [0, totalStake)
    const r =
      parseInt(seed.slice(0, 12), 16) % totalStake;

    let acc = 0;
    for (const v of list) {
      acc += v.stake;
      if (r < acc) return v.id;
    }
    return list[list.length - 1].id;
  }
}

module.exports = ProposerSelector;

