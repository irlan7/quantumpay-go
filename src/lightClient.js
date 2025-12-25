const axios = require('axios');
const cryptoHash = require('../crypto/hash');

class LightClient {
  constructor() {
    this.headers = [];
  }

  async syncHeaders({ peer, from = 0 }) {
    const res = await axios.get(`${peer}/headers?from=${from}`);
    const newHeaders = res.data;

    for (const header of newHeaders) {
      this._validateHeader(header);
      this.headers.push(header);
    }

    return this.headers.length;
  }

  _validateHeader(header) {
    const lastHeader = this.headers[this.headers.length - 1];

    if (lastHeader && header.lastHash !== lastHeader.hash) {
      throw new Error('Invalid header chain');
    }

    const validatedHash = cryptoHash(
      header.timestamp,
      header.lastHash,
      header.nonce,
      header.difficulty,
      header.stateRoot
    );

    if (validatedHash !== header.hash) {
      throw new Error('Invalid header hash');
    }
  }

  latestStateRoot() {
    if (!this.headers.length) return null;
    return this.headers[this.headers.length - 1].stateRoot;
  }
}

module.exports = LightClient;

