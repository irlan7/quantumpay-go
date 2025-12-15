const crypto = require('crypto');

const MINE_RATE = 3000; // ms — target waktu mine
const DEFAULT_DIFFICULTY = 3;

class Block {
  constructor({ timestamp, lastHash, hash, data, nonce, difficulty }) {
    this.timestamp = timestamp;
    this.lastHash = lastHash;
    this.hash = hash;
    this.data = data;
    this.nonce = nonce;
    this.difficulty = difficulty;
  }

  // -------------------------------
  // GENESIS BLOCK
  // -------------------------------
  static genesis() {
    return new Block({
      timestamp: 1,
      lastHash: '----',
      hash: 'genesis-hash',
      data: [],
      nonce: 0,
      difficulty: DEFAULT_DIFFICULTY
    });
  }

  // -------------------------------
  // HASHING UTILITY
  // -------------------------------
  static hashBlock({ timestamp, lastHash, data, nonce, difficulty }) {
    const blockString =
      `${timestamp}${lastHash}${JSON.stringify(data)}${nonce}${difficulty}`;
    return crypto.createHash('sha256').update(blockString).digest('hex');
  }

  static blockHash(block) {
    return Block.hashBlock({
      timestamp: block.timestamp,
      lastHash: block.lastHash,
      data: block.data,
      nonce: block.nonce,
      difficulty: block.difficulty
    });
  }

  // -------------------------------
  // DIFFICULTY ADJUSTMENT
  // -------------------------------
  static adjustDifficulty(lastBlock, currentTime) {
    if (!lastBlock || typeof lastBlock.difficulty !== 'number') {
      return DEFAULT_DIFFICULTY;
    }

    let newDifficulty =
      (currentTime - lastBlock.timestamp) > MINE_RATE
        ? lastBlock.difficulty - 1
        : lastBlock.difficulty + 1;

    return newDifficulty < 1 ? 1 : newDifficulty;
  }

  // -------------------------------
  // PROOF OF WORK
  // -------------------------------
  static mineBlock({ lastBlock, data }) {
    let hash, timestamp;
    let nonce = 0;
    let difficulty = lastBlock.difficulty;

    do {
      nonce++;
      timestamp = Date.now();
      difficulty = Block.adjustDifficulty(lastBlock, timestamp);
      hash = Block.hashBlock({
        timestamp,
        lastHash: lastBlock.hash,
        data,
        nonce,
        difficulty
      });

    } while (hash.substring(0, difficulty) !== '0'.repeat(difficulty));

    return new Block({
      timestamp,
      lastHash: lastBlock.hash,
      hash,
      data,
      nonce,
      difficulty
    });
  }
}

Block.DEFAULT_DIFFICULTY = DEFAULT_DIFFICULTY;

module.exports = Block;
