const fs = require('fs');
const path = require('path');

const DATA_DIR = path.join(__dirname, '../../data');
const REP_FILE = path.join(DATA_DIR, 'peer_reputation.json');

class PeerReputation {
  constructor() {
    this.peers = new Map();
    this._load();
  }

  _load() {
    if (!fs.existsSync(DATA_DIR)) fs.mkdirSync(DATA_DIR);
    if (!fs.existsSync(REP_FILE)) return;

    const raw = JSON.parse(fs.readFileSync(REP_FILE));
    for (const peerId in raw) {
      this.peers.set(peerId, raw[peerId]);
    }
  }

  _save() {
    const obj = {};
    for (const [id, data] of this.peers) {
      obj[id] = data;
    }
    fs.writeFileSync(REP_FILE, JSON.stringify(obj, null, 2));
  }

  _initPeer(peerId) {
    if (!this.peers.has(peerId)) {
      this.peers.set(peerId, {
        peerId,
        score: 100,
        strikes: 0,
        banned: false,
        lastSeen: Date.now()
      });
    }
  }

  record(peerId, delta, reason = '') {
    this._initPeer(peerId);
    const peer = this.peers.get(peerId);

    peer.score += delta;
    peer.lastSeen = Date.now();

    if (delta < 0) peer.strikes += 1;

    if (peer.score <= 0) peer.banned = true;
    if (peer.score < 20) peer.banned = true;

    this.peers.set(peerId, peer);
    this._save();

    console.log(`[REPUTATION] ${peerId}: ${delta} (${reason}) → ${peer.score}`);
  }

  isBanned(peerId) {
    return this.peers.get(peerId)?.banned === true;
  }

  getScore(peerId) {
    return this.peers.get(peerId)?.score ?? 100;
  }

  getTopPeers(limit = 5) {
    return [...this.peers.values()]
      .filter(p => !p.banned)
      .sort((a, b) => b.score - a.score)
      .slice(0, limit);
  }
}

module.exports = PeerReputation;

