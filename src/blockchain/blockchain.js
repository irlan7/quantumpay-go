const fs = require("fs");
const path = require("path");
const crypto = require("crypto");

class Blockchain {
  constructor({ nodeId, validators, config }) {
    this.nodeId = nodeId;
    this.validators = validators; // array of validator IDs
    this.config = config;

    this.chain = [];
    this.votes = {}; // height -> Set(validators)
    this.validatorActivity = {}; // validatorId -> lastSeenHeight

    this.stateFile = path.join(
      config.paths.dataDir,
      nodeId,
      "chain",
      "state.json"
    );

    this.state = {
      height: 0,
      finalizedHeight: 0,
      slashed: {}, // validatorId -> reason
    };

    this.loadState();
  }

  /* =========================
     STATE MANAGEMENT
  ========================== */

  loadState() {
    if (fs.existsSync(this.stateFile)) {
      this.state = JSON.parse(fs.readFileSync(this.stateFile));
    } else {
      this.persistState();
    }
  }

  persistState() {
    fs.mkdirSync(path.dirname(this.stateFile), { recursive: true });
    fs.writeFileSync(this.stateFile, JSON.stringify(this.state, null, 2));
  }

  /* =========================
     BLOCK CREATION
  ========================== */

  produceBlock() {
    const nextHeight = this.chain.length + 1;

    // 🚫 STRICT FINALITY: forbid fork below finalized height
    if (nextHeight <= this.state.finalizedHeight) {
      throw new Error(
        `[STRICT FINALITY] Cannot produce block below finalizedHeight=${this.state.finalizedHeight}`
      );
    }

    const prevHash =
      this.chain.length === 0
        ? "GENESIS"
        : this.chain[this.chain.length - 1].hash;

    const block = {
      height: nextHeight,
      prevHash,
      timestamp: Date.now(),
      producer: this.nodeId,
    };

    block.hash = this.hashBlock(block);
    this.chain.push(block);
    this.state.height = block.height;

    this.markValidatorActive(this.nodeId, block.height);

    console.log(`[BLOCK] Node ${this.nodeId} produced block #${block.height}`);

    this.persistState();
    return block;
  }

  hashBlock(block) {
    return crypto
      .createHash("sha256")
      .update(
        `${block.height}${block.prevHash}${block.timestamp}${block.producer}`
      )
      .digest("hex");
  }

  /* =========================
     VOTING & FINALITY
  ========================== */

  voteBlock(height, validatorId) {
    if (!this.votes[height]) {
      this.votes[height] = new Set();
    }

    // 🚨 DOUBLE SIGN DETECTION
    if (this.votes[height].has(validatorId)) {
      this.slashValidator(
        validatorId,
        `Double-sign at height ${height}`
      );
      return;
    }

    this.votes[height].add(validatorId);
    this.markValidatorActive(validatorId, height);

    console.log(
      `[VOTE] ${validatorId} votes for block #${height}`
    );

    this.checkFinality(height);
  }

  checkFinality(height) {
    const voteCount = this.votes[height]
      ? this.votes[height].size
      : 0;

    const threshold = Math.ceil(
      this.validators.length * this.config.consensus.finalityThreshold
    );

    if (voteCount >= threshold && height > this.state.finalizedHeight) {
      this.state.finalizedHeight = height;

      console.log(
        `[FINALITY] Block #${height} FINALIZED (${voteCount}/${threshold})`
      );

      this.persistState();
    }
  }

  /* =========================
     SLASHING (STRICT)
  ========================== */

  markValidatorActive(validatorId, height) {
    this.validatorActivity[validatorId] = height;
  }

  checkDowntime(currentHeight) {
    const maxMissed =
      this.config.consensus.maxMissedBlocks || 10;

    for (const v of this.validators) {
      const lastSeen = this.validatorActivity[v] || 0;
      if (currentHeight - lastSeen > maxMissed) {
        this.slashValidator(
          v,
          `Downtime: missed ${currentHeight - lastSeen} blocks`
        );
      }
    }
  }

  slashValidator(validatorId, reason) {
    if (this.state.slashed[validatorId]) return;

    this.state.slashed[validatorId] = {
      reason,
      time: Date.now(),
    };

    console.error(
      `[SLASHED] Validator ${validatorId} → ${reason}`
    );

    this.persistState();
  }

  /* =========================
     TICK (called by server)
  ========================== */

  tick() {
    try {
      this.produceBlock();
    } catch (e) {
      console.error(e.message);
    }

    const currentHeight = this.state.height;

    for (const v of this.validators) {
      this.voteBlock(currentHeight, v);
    }

    this.checkDowntime(currentHeight);
  }
}

module.exports = Blockchain;

