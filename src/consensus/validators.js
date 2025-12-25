class ValidatorSet {
  constructor() {
    this.validators = new Map();
  }

  register({ id, pubKey, stake }) {
    this.validators.set(id, {
      id,
      pubKey,
      stake,
      reputation: 100,
      active: true,
      slashed: false
    });
  }

  getActive(minStake, minRep) {
    return [...this.validators.values()].filter(v =>
      v.active &&
      !v.slashed &&
      v.stake >= minStake &&
      v.reputation >= minRep
    );
  }

  update(v) {
    this.validators.set(v.id, v);
  }
}

module.exports = ValidatorSet;

