// src/consensus/doublesign.js

class DoubleSignDetector {
  constructor() {
    // map: height -> validator -> blockHash
    this.records = {};
  }

  record({ height, validator, blockHash }) {
    if (!this.records[height]) {
      this.records[height] = {};
    }

    const prev = this.records[height][validator];

    if (prev && prev !== blockHash) {
      return {
        doubleSign: true,
        height,
        validator,
        blocks: [prev, blockHash]
      };
    }

    this.records[height][validator] = blockHash;
    return { doubleSign: false };
  }

  // cleanup old heights
  prune(finalizedHeight, keep = 50) {
    Object.keys(this.records)
      .map(Number)
      .filter(h => h < finalizedHeight - keep)
      .forEach(h => delete this.records[h]);
  }
}

module.exports = DoubleSignDetector;

