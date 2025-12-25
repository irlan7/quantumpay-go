const cryptoHash = require('../crypto/hash');

/**
 * Deterministic Account Merkle Tree
 * Leaf = hash(address, balance, nonce)
 */
class MerkleTree {
  constructor(accounts) {
    this.accounts = accounts;
    this.leaves = accounts.map(acc =>
      cryptoHash(acc.address, acc.balance, acc.nonce)
    );

    this.layers = [this.leaves];
    this._build();
  }

  _build() {
    let level = this.leaves;

    while (level.length > 1) {
      const next = [];

      for (let i = 0; i < level.length; i += 2) {
        if (i + 1 < level.length) {
          next.push(cryptoHash(level[i], level[i + 1]));
        } else {
          next.push(level[i]); // carry
        }
      }

      this.layers.push(next);
      level = next;
    }
  }

  root() {
    return this.layers[this.layers.length - 1][0];
  }

  proof(index) {
    const proof = [];

    for (let level = 0; level < this.layers.length - 1; level++) {
      const layer = this.layers[level];
      const isRight = index % 2 === 1;
      const pairIndex = isRight ? index - 1 : index + 1;

      if (pairIndex < layer.length) {
        proof.push({
          hash: layer[pairIndex],
          position: isRight ? 'left' : 'right'
        });
      }

      index = Math.floor(index / 2);
    }

    return proof;
  }

  static verify({ address, balance, nonce, proof, root }) {
    let hash = cryptoHash(address, balance, nonce);

    for (const step of proof) {
      hash =
        step.position === 'left'
          ? cryptoHash(step.hash, hash)
          : cryptoHash(hash, step.hash);
    }

    return hash === root;
  }
}

module.exports = MerkleTree;

