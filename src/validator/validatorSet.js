class ValidatorSet {
  constructor() {
    this.validators = new Map();
  }

  register({ pubKey, stake }) {
    this.validators.set(pubKey, {
      stake,
      jailed: false,
      slashed: 0
    });
  }

  slash(pubKey, amount) {
    const v = this.validators.get(pubKey);
    if (!v) return;

    v.stake -= amount;
    v.slashed += amount;

    if (v.stake < 0) v.stake = 0;
  }

  jail(pubKey) {
    const v = this.validators.get(pubKey);
    if (v) v.jailed = true;
  }

  get(pubKey) {
    return this.validators.get(pubKey);
  }
}

module.exports = ValidatorSet;

