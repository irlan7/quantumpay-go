const Evidence = require('./evidence');

class Slashing {
  constructor({
    validatorSet,
    treasury,
    burnAddress = '*BURN*'
  }) {
    this.validatorSet = validatorSet;
    this.treasury = treasury;
    this.burnAddress = burnAddress;
  }

  /**
   * Handle double-sign slashing
   */
  handleDoubleSign({ snapshotA, snapshotB, reporter }) {
    if (!Evidence.doubleSign({ snapshotA, snapshotB })) {
      return false;
    }

    const pubKey = snapshotA.signer;
    const validator = this.validatorSet.get(pubKey);
    if (!validator || validator.jailed) return false;

    // Total slash amount
    const totalSlash = validator.stake * 0.3;

    // Distribution
    const reporterReward = totalSlash * 0.4;
    const treasuryShare = totalSlash * 0.4;
    const burnAmount = totalSlash * 0.2;

    // Apply slashing
    this.validatorSet.slash(pubKey, totalSlash);
    this.validatorSet.jail(pubKey);

    // Redistribute
    if (reporter) {
      reporter.balance =
        (reporter.balance || 0) + reporterReward;
    }

    this.treasury.credit(treasuryShare);

    console.log(`🔥 Validator ${pubKey} slashed`);
    console.log(`🎯 Reporter reward: ${reporterReward}`);
    console.log(`🏦 Treasury funded: ${treasuryShare}`);
    console.log(`🔥 Burned: ${burnAmount}`);

    return {
      slashed: totalSlash,
      reporterReward,
      treasuryShare,
      burnAmount
    };
  }
}

module.exports = Slashing;

