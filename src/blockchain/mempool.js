class Mempool {
  constructor(maxSize = 5000) {
    this.maxSize = maxSize;
    this.pool = [];
  }

  add(tx) {
    if (this.pool.length >= this.maxSize) {
      // drop lowest gas tx
      this.pool.sort((a, b) => a.gas - b.gas);
      this.pool.shift();
    }

    this.pool.push({
      tx,
      gas: tx.gas,
      timestamp: Date.now()
    });
  }

  getPrioritized(limitGas) {
    // sort by gas DESC, then FIFO
    this.pool.sort((a, b) => {
      if (b.gas !== a.gas) return b.gas - a.gas;
      return a.timestamp - b.timestamp;
    });

    const selected = [];
    let usedGas = 0;

    for (const item of this.pool) {
      if (usedGas + item.gas > limitGas) continue;
      selected.push(item.tx);
      usedGas += item.gas;
    }

    // remove selected from mempool
    this.pool = this.pool.filter(
      p => !selected.includes(p.tx)
    );

    return selected;
  }

  size() {
    return this.pool.length;
  }
}

module.exports = Mempool;

