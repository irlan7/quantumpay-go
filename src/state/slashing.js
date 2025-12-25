class SlashingEngine {
  constructor(state) {
    this.state = state;
  }

  slash(validatorId, reason, evidence = {}) {
    const validator = this.state.validators[validatorId];
    if (!validator || validator.slashed) return false;

    const rates = this.state.params.slashing;
    let penaltyRate = 0;

    if (reason === "doubleSign") penaltyRate = rates.doubleSign;
    if (reason === "downtime") penaltyRate = rates.downtime;
    if (reason === "censorship") penaltyRate = rates.censorship;

    const penalty = Math.floor(validator.stake * penaltyRate);

    validator.stake -= penalty;
    validator.slashed = true;
    validator.status = "jailed";

    this.state.supply.staked -= penalty;
    this.state.supply.burned =
      (this.state.supply.burned || 0) + penalty;

    validator.slashHistory = validator.slashHistory || [];
    validator.slashHistory.push({
      reason,
      penalty,
      height: this.state.height,
      evidence,
      time: Date.now(),
    });

    return {
      validatorId,
      penalty,
      reason,
    };
  }
}

module.exports = SlashingEngine;

