// src/consensus/finality.js

const fs = require("fs");
const path = require("path");

class Finality {
  constructor({ state, dataDir, logger }) {
    this.state = state;
    this.dataDir = dataDir;
    this.logger = logger;
  }

  finalize(height) {
    if (this.state.finalizedHeight >= height) return;

    this.state.finalizedHeight = height;
    this.persist();

    this.logger(
      `[FINALITY] Block #${height} finalized (soft-enforce)`
    );
  }

  persist() {
    const file = path.join(this.dataDir, "chain", "state.json");

    fs.mkdirSync(path.dirname(file), { recursive: true });
    fs.writeFileSync(file, JSON.stringify(this.state, null, 2));
  }
}

module.exports = Finality;

