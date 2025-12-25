// src/consensus/voting.js

class Voting {
  constructor({ nodeId, validators, logger }) {
    this.nodeId = nodeId;
    this.validators = validators;
    this.logger = logger;

    // votes[height] = Set(nodeId)
    this.votes = {};
  }

  maybeVote(block) {
    const height = block.height;

    if (!this.votes[height]) {
      this.votes[height] = new Set();
    }

    if (this.votes[height].has(this.nodeId)) {
      return null;
    }

    this.votes[height].add(this.nodeId);

    this.logger(`[VOTE] ${this.nodeId} votes for block #${height}`);

    return {
      height,
      voter: this.nodeId,
    };
  }

  addVote(vote) {
    const { height, voter } = vote;

    if (!this.votes[height]) {
      this.votes[height] = new Set();
    }

    this.votes[height].add(voter);
  }

  hasQuorum(height) {
    const count = this.votes[height]?.size || 0;
    const total = this.validators.length;

    return count >= Math.ceil((2 / 3) * total);
  }

  getVoteCount(height) {
    return this.votes[height]?.size || 0;
  }
}

module.exports = Voting;

