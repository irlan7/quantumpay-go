const { startDiscovery } = require('./discovery');
const { handleHandshake } = require('./handshake');

class P2PNode {
  constructor({ blockchain, chainId, port }) {
    this.blockchain = blockchain;
    this.chainId = chainId;
    this.port = port;
    this.peers = new Set();
  }

  start() {
    startDiscovery(this);
    console.log(`🌐 P2P listening on port ${this.port}`);
  }

  onPeer(socket) {
    handleHandshake(socket, this);
    this.peers.add(socket);
  }
}

module.exports = P2PNode;

