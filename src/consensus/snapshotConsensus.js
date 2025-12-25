const StateSnapshot = require('../state/snapshot');
const verifySignature = require('../crypto/verify');
const Evidence = require('../validator/evidence');

/**
 * SnapshotConsensus
 * - Verifies multi-validator snapshot signatures
 * - Enforces threshold quorum
 * - Detects slashable behavior (double-sign)
 */
class SnapshotConsensus {
  constructor({
    validatorSet,
    slashingEngine,
    threshold = 0.67
  }) {
    this.validatorSet = validatorSet;
    this.slashingEngine = slashingEngine;
    this.threshold = threshold;

    // Track last signed snapshot per validator (for double-sign detection)
    this.lastSignedSnapshot = new Map();
  }

  /**
   * =========================
   * VERIFY SNAPSHOT CONSENSUS
   * =========================
   */
  verify(snapshot) {
    if (!snapshot || !snapshot.signatures) {
      console.warn('Snapshot has no signatures');
      return false;
    }

    const snapshotHash = StateSnapshot.hash(snapshot);

    let totalWeight = 0;
    let signedWeight = 0;

    // Hitung total validator weight
    for (const v of this.validatorSet.validators.values()) {
      if (!v.jailed) {
        totalWeight += v.stake;
      }
    }

    // Verifikasi signature satu per satu
    for (const [signer, signature] of Object.entries(snapshot.signatures)) {
      const validator = this.validatorSet.get(signer);
      if (!validator || validator.jailed) continue;

      const valid = verifySignature({
        publicKey: signer,
        data: snapshotHash,
        signature
      });

      if (!valid) {
        console.warn(`Invalid signature from ${signer}`);
        continue;
      }

      signedWeight += validator.stake;

      // =========================
      // DOUBLE-SIGN DETECTION
      // =========================
      const last = this.lastSignedSnapshot.get(signer);
      if (last) {
        const isDoubleSign = Evidence.doubleSign({
          snapshotA: last,
          snapshotB: snapshot
        });

        if (isDoubleSign) {
          console.error(`🔥 DOUBLE SIGN detected by ${signer}`);

          this.slashingEngine.handleDoubleSign({
            snapshotA: last,
            snapshotB: snapshot
          });

          // Setelah slash, abaikan signature ini
          continue;
        }
      }

      // Update last signed snapshot
      this.lastSignedSnapshot.set(signer, snapshot);
    }

    const quorum = signedWeight / totalWeight;

    if (quorum >= this.threshold) {
      console.log(`✅ Snapshot consensus reached (${(quorum * 100).toFixed(2)}%)`);
      return true;
    }

    console.warn(
      `❌ Snapshot consensus failed (${(quorum * 100).toFixed(2)}%)`
    );
    return false;
  }
}

module.exports = SnapshotConsensus;

