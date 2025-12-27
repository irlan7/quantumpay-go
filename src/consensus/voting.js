class Voting {
  constructor(finality) {
    this.finality = finality;
  }

  maybeVote(block) {
    if (this.finality.isFinalized(block.height)) {
      console.warn(
        `[VOTE BLOCKED] Block #${block.height} already finalized`
      );
      return false;
    }

    console.log(`[VOTE] Vote for block #${block.height}`);
    return true;
  }
}

module.exports = Voting;

