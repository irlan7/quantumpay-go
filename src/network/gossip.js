// src/network/gossip.js

const axios = require('axios');

class Gossip {
  constructor({ peers }) {
    this.peers = peers || [];
  }

  broadcastVote(vote) {
    for (const peer of this.peers) {
      axios
        .post(`${peer}/vote`, vote, { timeout: 2000 })
        .catch(() => {});
    }
  }
}

module.exports = Gossip;

