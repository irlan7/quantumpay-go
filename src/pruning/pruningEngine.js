/**
 * PruningEngine
 * - Snapshot-aware
 * - Deterministic
 * - Safe for consensus
 */

class PruningEngine {
  constructor({
    blockchain,
    snapshot,
    safetyMargin = 50,
    mode = 'full' // 'full' | 'archive'
  }) {
    this.blockchain = blockchain;
    this.snapshot = snapshot;
    this.safetyMargin = safetyMargin;
    this.mode = mode;
  }

  /**
   * Check if pruning is allowed
   */
  canPrune() {
    if (this.mode === 'archive') return false;
    if (!this.snapshot) return false;

    if (
      typeof this.snapshot.height !== 'number' ||
      typeof this.snapshot.stateRoot !== 'string'
    ) {
      return false;
    }

    const currentHeight = this.blockchain.chain.length - 1;

    return this.snapshot.height < currentHeight;
  }

  /**
   * Execute pruning
   */
  prune() {
    if (!this.canPrune()) {
      console.log('⏭️ Pruning skipped (conditions not met)');
      return false;
    }

    const pruneBelow =
      this.snapshot.height - this.safetyMargin;

    if (pruneBelow <= 0) {
      console.log('⏭️ Nothing to prune yet');
      return false;
    }

    const originalLength = this.blockchain.chain.length;

    // Keep genesis block always
    this.blockchain.chain = [
      this.blockchain.chain[0],
      ...this.blockchain.chain.slice(pruneBelow)
    ];

    const pruned = originalLength - this.blockchain.chain.length;

    console.log(
      `🧹 Pruned ${pruned} blocks (below height ${pruneBelow})`
    );

    return true;
  }
}

module.exports = PruningEngine;

