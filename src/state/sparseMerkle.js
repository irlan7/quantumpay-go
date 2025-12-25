const cryptoHash = require('../crypto/hash');

const EMPTY_HASH = cryptoHash('0');

class SparseMerkleTree {
  constructor(depth = 256) {
    this.depth = depth;
    this.nodes = new Map(); // key: path, value: hash
  }

  _getBit(key, index) {
    const byte = parseInt(key.slice(index >> 3, (index >> 3) + 2), 16);
    return (byte >> (7 - (index % 8))) & 1;
  }

  _path(key, depth) {
    return key.slice(0, depth);
  }

  update(keyHash, valueHash) {
    let currentHash = valueHash;
    let path = '';

    for (let i = this.depth - 1; i >= 0; i--) {
      const bit = this._getBit(keyHash, i);
      const siblingPath = path + (bit ? '0' : '1');
      const siblingHash = this.nodes.get(siblingPath) || EMPTY_HASH;

      currentHash = bit
        ? cryptoHash(siblingHash, currentHash)
        : cryptoHash(currentHash, siblingHash);

      path += bit ? '1' : '0';
      this.nodes.set(path, currentHash);
    }

    this.root = currentHash;
    return this.root;
  }

  getRoot() {
    return this.root || EMPTY_HASH;
  }

  getProof(keyHash) {
    let proof = [];
    let path = '';

    for (let i = this.depth - 1; i >= 0; i--) {
      const bit = this._getBit(keyHash, i);
      const siblingPath = path + (bit ? '0' : '1');
      const siblingHash = this.nodes.get(siblingPath) || EMPTY_HASH;

      proof.push(siblingHash);
      path += bit ? '1' : '0';
    }

    return proof;
  }

  static verify({ keyHash, valueHash, proof, root }) {
    let hash = valueHash;

    for (let i = 0; i < proof.length; i++) {
      const bit =
        (parseInt(keyHash.slice((proof.length - 1 - i) >> 3, ((proof.length - 1 - i) >> 3) + 2), 16)
          >> (7 - ((proof.length - 1 - i) % 8))) & 1;

      hash = bit
        ? cryptoHash(proof[i], hash)
        : cryptoHash(hash, proof[i]);
    }

    return hash === root;
  }
}

module.exports = SparseMerkleTree;

