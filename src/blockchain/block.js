const crypto = require('crypto');

class Block {
  constructor({ index, prevHash, timestamp, transactions, gasUsed, chainId, validator, signature }) {
    this.index = index;
    this.prevHash = prevHash;
    this.timestamp = timestamp;
    this.transactions = transactions;
    this.gasUsed = gasUsed;
    this.chainId = chainId;

    this.validator = validator;   // public key
    this.signature = signature;   // hex DER

    this.hash = this.computeHash();
  }

  computeHash() {
    return crypto
      .createHash('sha256')
      .update(
        this.index +
        this.prevHash +
        this.timestamp +
        JSON.stringify(this.transactions) +
        this.gasUsed +
        this.chainId
      )
      .digest('hex');
  }

  static genesis(chainId) {
    return new Block({
      index: 0,
      prevHash: '0',
      timestamp: Date.now(),
      transactions: [],
      gasUsed: 0,
      chainId,
      validator: 'GENESIS',
      signature: 'GENESIS'
    });
  }
}

module.exports = Block;

