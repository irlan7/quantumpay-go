const fs = require('fs');
const path = require('path');

class PeerManager {
  constructor({ chainId, selfAddress, maxPeers = 20 }) {
    this.chainId = chainId;
    this.selfAddress = selfAddress;
    this.maxPeers = maxPeers;
    this.peers = new Set();
  }

  loadBootstrapPeers() {
    const file = path.join(__dirname, 'bootstrap.json');
    if (!fs.existsSync(file)) return [];

    const data = JSON.parse(fs.readFileSync(file));
    if (data.chainId !== this.chainId) return [];

    return data.bootstraps || [];
  }

  addPeer(address) {
    if (!address) return;
    if (address === this.selfAddress) return;
    if (this.peers.size >= this.maxPeers) return;

    this.peers.add(address);
  }

  removePeer(address) {
    this.peers.delete(address);
  }

  getPeers() {
    return Array.from(this.peers);
  }

  mergePeers(peerList = []) {
    peerList.forEach(p => this.addPeer(p));
  }
}

module.exports = PeerManager;

