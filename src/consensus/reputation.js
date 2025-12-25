const MAX_REPUTATION = 100;
const MIN_REPUTATION = 0;

class ReputationManager {
  constructor() {
    this.reputations = new Map(); // validatorId => score
  }

  initValidator(validatorId) {
    if (!this.reputations.has(validatorId)) {
      this.reputations.set(validatorId, 100);
    }
  }

  get(validatorId) {
    return this.reputations.get(validatorId) ?? 0;
  }

  increase(validatorId, amount = 1) {
    const score = this.get(validatorId);
    this.reputations.set(
      validatorId,
      Math.min(MAX_REPUTATION, score + amount)
    );
  }

  decrease(validatorId, amount = 5) {
    const score = this.get(validatorId);
    this.reputations.set(
      validatorId,
      Math.max(MIN_REPUTATION, score - amount)
    );
  }

  isHealthy(validatorId) {
    return this.get(validatorId) >= 50;
  }
}

module.exports = new ReputationManager();

