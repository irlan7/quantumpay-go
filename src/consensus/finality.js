class Finality {
  constructor(state, slashing) {
    this.state = state;
    this.slashing = slashing;
  }

  checkVote(vote) {
    const finalized = this.state.finalizedHeight;

    if (
      finalized !== null &&
      vote.height <= finalized
    ) {
      this.slashing.onFinalizedViolation(vote.validatorId);
      throw new Error(
        `FATAL: Vote below finalized height (${finalized})`
      );
    }
  }

  updateFinalized(height) {
    if (
      this.state.finalizedHeight === null ||
      height > this.state.finalizedHeight
    ) {
      this.state.finalizedHeight = height;
      console.log(`[FINALITY] Finalized height = ${height}`);
    }
  }
}

module.exports = Finality;

