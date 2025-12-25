const { verifySignature } = require('../crypto/hash');

class Transaction {
  constructor({ from, to, amount, nonce, gas, signature, chainId }) {
    this.from = from;
    this.to = to;
    this.amount = amount;
    this.nonce = nonce;
    this.gas = gas;
    this.chainId = chainId;
    this.signature = signature;
  }

  // 🔐 DOMAIN SEPARATION
  static signingPayload(tx) {
    return JSON.stringify({
      domain: 'QUANTUMPAY_TX',
      chainId: tx.chainId,
      from: tx.from,
      to: tx.to,
      amount: tx.amount,
      nonce: tx.nonce,
      gas: tx.gas
    });
  }

  static verify(tx, blockchain) {
    if (tx.chainId !== blockchain.chainId) return false;
    if (tx.gas < blockchain.MIN_GAS_FLOOR) return false;

    const expectedNonce = blockchain.getNonce(tx.from);
    if (tx.nonce !== expectedNonce) return false;

    return verifySignature(
      tx.from,
      Transaction.signingPayload(tx),
      tx.signature
    );
  }

  // ✅ INI YANG HILANG
  static fromJSON(obj) {
    return new Transaction(obj);
  }
}

module.exports = Transaction;

