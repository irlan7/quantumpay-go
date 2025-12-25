class VotePool {
  constructor() {
    this.votes = new Map(); // blockHash -> Set(validatorPubKey)
  }

  addVote(blockHash, validator) {
    if (!this.votes.has(blockHash)) {
      this.votes.set(blockHash, new Set());
    }
    this.votes.get(blockHash).add(validator);
  }

  count(blockHash) {
    return this.votes.get(blockHash)?.size || 0;
  }

  clear(blockHash) {
    this.votes.delete(blockHash);
  }
}

module.exports = new VotePool();

