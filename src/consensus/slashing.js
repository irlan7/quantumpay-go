'use strict';

class Slashing {
  constructor({ state }) {
    if (!state || !state.validators) {
      throw new Error('Slashing requires state with validators');
    }

    this.state = state;

    // default params (fallback)
    this.state.params = this.state.params || {};
    this.state.params.slashing =
      this.state.params.slashing || {
        doubleSign: 0.05,   // 5%
        invalidVote: 0.02  // 2%
      };

    this.state.supply = this.state.supply || {
      burned: 0
    };
  }

  // =========================
  // INTERNAL HELPER
  // =========================
  _burn(amount) {
    if (amount <= 0) return;
    this.state.supply.burned =
      (this.state.supply.burned || 0) + amount;
  }

  _recordSlash(validator, record) {
    validator.slashHistory =
      validator.slashHistory || [];
    validator.slashHistory.push(record);
  }

  // =========================
  // DOUBLE-SIGN SLASHING
  // =========================
  slashDoubleSign(validatorId, evidence) {
    const validator = this.state.validators[validatorId];
    if (!validator) return null;
    if (validator.jailed) return null;

    const rate =
      this.state.params.slashing.doubleSign;

    const penalty = Math.max(
      1,
      Math.floor(validator.stake * rate)
    );

    validator.stake -= penalty;
    validator.jailed = true;

    this._burn(penalty);

    this._recordSlash(validator, {
      type: 'DOUBLE_SIGN',
      penalty,
      height: evidence.height,
      blocks: evidence.blocks,
      time: Date.now()
    });

    return {
      validatorId,
      penalty,
      reason: 'DOUBLE_SIGN'
    };
  }

  // =========================
  // INVALID VOTE SLASHING
  // =========================
  slashInvalidVote(validatorId, evidence) {
    const validator = this.state.validators[validatorId];
    if (!validator) return null;
    if (validator.jailed) return null;

    const rate =
      this.state.params.slashing.invalidVote;

    const penalty = Math.max(
      1,
      Math.floor(validator.stake * rate)
    );

    validator.stake -= penalty;
    validator.jailed = true;

    this._burn(penalty);

    this._recordSlash(validator, {
      type: 'INVALID_VOTE',
      penalty,
      reason: evidence.reason,
      height: evidence.blockHeight,
      time: Date.now()
    });

    return {
      validatorId,
      penalty,
      reason: 'INVALID_VOTE'
    };
  }

  // =========================
  // QUERY HELPERS
  // =========================
  isJailed(validatorId) {
    const v = this.state.validators[validatorId];
    return !!(v && v.jailed);
  }

  getSlashHistory(validatorId) {
    const v = this.state.validators[validatorId];
    return v?.slashHistory || [];
  }
}

module.exports = Slashing;

