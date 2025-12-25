/**
 * Treasury
 * - Holds slashed funds
 * - Used for ecosystem funding
 */
class Treasury {
  constructor() {
    this.balance = 0;
  }

  credit(amount) {
    this.balance += amount;
  }

  debit(amount) {
    if (amount > this.balance) {
      throw new Error('Treasury insufficient balance');
    }
    this.balance -= amount;
  }
}

module.exports = Treasury;

