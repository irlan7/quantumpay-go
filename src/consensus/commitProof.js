const crypto = require('crypto');

class CommitProof {
  static create(blockHash, voters) {
    return crypto
      .createHash('sha256')
      .update(blockHash + voters.join(','))
      .digest('hex');
  }
}

module.exports = CommitProof;

