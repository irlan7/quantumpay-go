// src/blockchain/blockchain.js

class Blockchain {
  constructor({ nodeId, voting, finality, logger }) {
    this.nodeId = nodeId;
    this.voting = voting;
    this.finality = finality;
    this.logger = logger;

    this.height = 0;
    this.state = {
      finalizedHeight: 0,
    };
  }

  produceBlock() {
    this.height++;

    const block = {
      height: this.height,
      proposer: this.nodeId,
      timestamp: Date.now(),
    };

    this.logger(
      `[BLOCK] Node ${this.nodeId} produced block #${block.height}`
    );

    // 🗳️ vote
    const vote = this.voting.maybeVote(block);
    if (vote) {
      this.voting.addVote(vote);
    }

    // ✅ check quorum
    if (this.voting.hasQuorum(block.height)) {
      this.finality.finalize(block.height);
    }

    return block;
  }
}

module.exports = Blockchain;

