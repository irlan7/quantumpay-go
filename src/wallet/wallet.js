const signHash = require('../crypto/sign');
const cryptoHash = require('../crypto/hash');

class Wallet {
  constructor() {
    // keypair init (existing)
  }

  signTransaction(dataHash) {
    return signHash({
      privateKey: this.privateKey,
      dataHash: cryptoHash.domain('TX', dataHash)
    });
  }

  signSnapshot(snapshotHash) {
    return signHash({
      privateKey: this.privateKey,
      dataHash: cryptoHash.domain('SNAPSHOT', snapshotHash)
    });
  }
}

module.exports = Wallet;

