const { ec } = require('../crypto/ec');
const { hash } = require('../crypto/hash');

class Transaction {
  constructor(from, to, amount, signature) {
    this.from = from;
    this.to = to;
    this.amount = amount;
    this.signature = signature;
  }

  calculateHash() {
    return hash(`${this.from}${this.to}${this.amount}`);
  }

  isValid() {
    // mining reward / system tx
    if (this.from === null) return true;

    if (!this.signature || !this.from) {
      return false;
    }

    try {
      const key = ec.keyFromPublic(this.from, 'hex');
      return key.verify(this.calculateHash(), this.signature);
    } catch (err) {
      console.error('Signature verify error:', err.message);
      return false;
    }
  }
}

module.exports = Transaction;

