// src/blockchain.js
// Minimal Blockchain implementation for development / testing.
// Provides: getChain, getLatestBlock, addTransaction, minePendingTransactions, replaceChain, isValidChain

const crypto = require('crypto');

class Block {
  constructor(index, timestamp, transactions = [], previousHash = '') {
    this.index = index;
    this.timestamp = timestamp;
    this.transactions = transactions;
    this.previousHash = previousHash;
    this.nonce = 0;
    this.hash = this.calculateHash();
  }

  calculateHash() {
    return crypto
      .createHash('sha256')
      .update(this.index + this.timestamp + JSON.stringify(this.transactions) + this.previousHash + this.nonce)
      .digest('hex');
  }

  // very small, demo PoW (blocking) — keep difficulty low for local testing
  mine(difficulty = 2) {
    const target = '0'.repeat(Math.max(0, Number(difficulty) || 2));
    while (!this.hash.startsWith(target)) {
      this.nonce += 1;
      this.hash = this.calculateHash();
    }
  }
}

class Blockchain {
  constructor(opts = {}) {
    this.difficulty = typeof opts.difficulty !== 'undefined' ? opts.difficulty : 2;
    this.miningReward = typeof opts.miningReward !== 'undefined' ? opts.miningReward : 1;
    this.chain = [this.createGenesisBlock()];
    this.pendingTransactions = [];
  }

  createGenesisBlock() {
    return new Block(0, Date.now(), ['genesis'], '0');
  }

  getLatestBlock() {
    return this.chain[this.chain.length - 1];
  }

  getChain() {
    return this.chain;
  }

  // Add a transaction to pending pool (no signature checks here)
  addTransaction(tx) {
    if (!tx || typeof tx !== 'object') throw new Error('Invalid transaction object');
    // basic shape validation
    if (!('from' in tx) || !('to' in tx) || !('amount' in tx)) {
      throw new Error('Transaction must have from, to, and amount');
    }
    this.pendingTransactions.push(tx);
    return true;
  }

  // Mine pending transactions into a new block.
  // Returns summary object or null if nothing to mine.
  minePendingTransactions(minerAddress = 'miner-demo') {
    if (!this.pendingTransactions.length) return null;

    const block = new Block(this.chain.length, Date.now(), this.pendingTransactions.slice(), this.getLatestBlock().hash);
    block.mine(this.difficulty);
    this.chain.push(block);

    // award miner (put reward as next pending transaction)
    this.pendingTransactions = [{ from: null, to: minerAddress, amount: this.miningReward }];

    return {
      index: block.index,
      hash: block.hash,
      transactionsCount: block.transactions.length,
      timestamp: block.timestamp,
    };
  }

  // Accept a block pushed from a peer — naive; for demo only
  addBlockFromPeer(blockData) {
    // Basic minimal checks: ensure property presence, then push
    if (!blockData || typeof blockData !== 'object') return false;
    if (typeof blockData.index !== 'number' || typeof blockData.hash !== 'string') return false;
    this.chain.push(blockData);
    return true;
  }

  // Replace chain if the new chain is longer and valid
  replaceChain(newChain) {
    if (!Array.isArray(newChain)) return false;
    if (newChain.length <= this.chain.length) return false;
    if (!this.isValidChain(newChain)) return false;
    this.chain = newChain;
    return true;
  }

  // Basic validation: check prev-hash continuity. Does NOT re-calc full POW.
  isValidChain(chainToValidate) {
    if (!Array.isArray(chainToValidate) || chainToValidate.length === 0) return false;
    // validate genesis match basic structure
    for (let i = 1; i < chainToValidate.length; i++) {
      const prev = chainToValidate[i - 1];
      const curr = chainToValidate[i];
      if (!prev || !curr) return false;
      if (curr.previousHash !== prev.hash) return false;
      // optionally you can re-check hashes:
      // const reconstructed = crypto.createHash('sha256').update(curr.index + curr.timestamp + JSON.stringify(curr.transactions) + curr.previousHash + (curr.nonce||0)).digest('hex');
      // if (reconstructed !== curr.hash) return false;
    }
    return true;
  }
}

module.exports = Blockchain;

