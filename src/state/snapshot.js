const crypto = require('../crypto/hash');

class StateSnapshot {
  constructor(options = {}) {
    const {
      height = 0,
      stateRoot = null,
      timestamp = Date.now()
    } = options;

    this.height = height;
    this.stateRoot = stateRoot;
    this.timestamp = timestamp;
  }

  hash() {
    return crypto(
      this.height,
      this.stateRoot,
      this.timestamp
    );
  }
}

module.exports = StateSnapshot;

