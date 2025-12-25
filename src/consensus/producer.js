const crypto = require('crypto');
const { getLeader } = require('./leader');

class BlockProducer {
  constructor({ blockchain, p2p, nodeId, interval = 5000 }) {
    this.blockchain = blockchain;
    this.p2p = p2p;
    this.nodeId = nodeId;
    this.interval = interval;
    this.timer = null;
  }

  start() {
    if (this.timer) return;

    console.log('⛏️ Block producer started');
    this.timer = setInterval(() => {
      try {
        this.tryProduceBlock();
      } catch (err) {
        console.error('Block production error:', err.message);
      }
    }, this.interval);
  }

  stop() {
    if (this.timer) clearInterval(this.timer);
    this.timer = null;
  }

  tryProduceBlock() {
    const state = this.blockchain.getState();
    const leader = getLeader(state.validators, state.height);

    if (leader !== this.nodeId) return; // bukan giliran

    const block = this.createBlock(state);
    this.blockchain.addBlock(block);

    if (this.p2p) {
      this.p2p.broadcast('block:new', block);
    }

    console.log(`✅ Block #${block.height} produced by ${this.nodeId}`);
  }

  createBlock(state) {
    const prevHash = state.lastBlockHash;
    const height = state.height + 1;
    const timestamp = Date.now();

    const body = {
      height,
      prevHash,
      timestamp,
      txs: [], // tx pool nanti
      producer: this.nodeId
    };

    const hash = crypto
      .createHash('sha256')
      .update(JSON.stringify(body))
      .digest('hex');

    return {
      ...body,
      hash,
      signature: `SIG_${this.nodeId}_${hash.slice(0, 16)}`
    };
  }
}

module.exports = BlockProducer;

