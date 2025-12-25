class TransactionPool {
  constructor() {
    this.transactions = [];
  }

  add(tx) {
    this.transactions.push(tx);
  }

  clear() {
    this.transactions = [];
  }

  validTransactions() {
    return this.transactions.filter(tx => tx && typeof tx === 'object');
  }
}

module.exports = TransactionPool;

