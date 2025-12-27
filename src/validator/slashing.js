class Slashing {
  constructor(state, config) {
    this.state = state;
    this.config = config.slashing;
  }

  slash(validatorId, percent, reason) {
    const v = this.state.validators[validatorId];
    if (!v) return;

    const amount = Math.floor(v.stake * (percent / 100));
    v.stake -= amount;

    v.slashed = true;
    v.slashReason = reason;

    console.log(
      `[SLASH] Validator ${validatorId} slashed ${percent}% for ${reason}`
    );
  }

  jail(validatorId, permanent = false) {
    const v = this.state.validators[validatorId];
    if (!v) return;

    v.jailed = true;
    v.permanentJail = permanent;

    console.log(
      `[JAIL] Validator ${validatorId} jailed${permanent ? " permanently" : ""}`
    );
  }

  onDoubleSign(validatorId) {
    const rule = this.config.doubleSign;
    this.slash(validatorId, rule.slashPercent, "DOUBLE_SIGN");
    if (rule.jail) this.jail(validatorId, rule.permanent);
  }

  onFinalizedViolation(validatorId) {
    const rule = this.config.finalizedViolation;
    this.slash(validatorId, rule.slashPercent, "FINALIZED_VIOLATION");
    if (rule.jail) this.jail(validatorId, rule.permanent);
  }

  onDowntime(validatorId, missedBlocks) {
    const rule = this.config.downtime;
    if (missedBlocks >= rule.maxMissedBlocks) {
      this.slash(validatorId, rule.slashPercent, "DOWNTIME");
      this.jail(validatorId, false);
    }
  }
}

module.exports = Slashing;

