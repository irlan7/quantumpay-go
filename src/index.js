// src/index.js
// Entry point QuantumPay Chain – versi stabil & rapih

const express = require('express');
const bodyParser = require('body-parser');
const Blockchain = require('./blockchain');

const HTTP_PORT = process.env.HTTP_PORT || 3001;
const P2P_PORT = process.env.P2P_PORT || 6001;

const app = express();
app.use(bodyParser.json());

// ----------------------
// 1. Init Blockchain
// ----------------------
const blockchain = new Blockchain();

// ----------------------
// 2. Import P2P module
// ----------------------
let p2pRaw;
try {
  p2pRaw = require('./p2p');
} catch (e) {
  console.error('❌ ERROR: p2p.js gagal di-import:', e.message);
  process.exit(1);
}

// ----------------------
// 3. Normalisasi export P2P
// ----------------------
function initP2P(moduleExport) {
  // Export berupa function factory: module.exports = (chain, port) => instance
  if (typeof moduleExport === "function") {
    try {
      const instance = moduleExport(blockchain, P2P_PORT);
      if (instance && typeof instance === "object") return instance;
    } catch (_) {}
  }

  // Export berupa object: { P2pServer }
  if (typeof moduleExport === "object") {
    if (typeof moduleExport.P2pServer === "function") {
      return new moduleExport.P2pServer(blockchain, P2P_PORT);
    }
    if (typeof moduleExport.initP2PServer === "function") {
      return moduleExport.initP2PServer(blockchain, P2P_PORT);
    }
    if (moduleExport.default) {
      return initP2P(moduleExport.default);
    }
  }

  // Export berupa class langsung
  if (typeof moduleExport === 'function') {
    try {
      return new moduleExport(blockchain, P2P_PORT);
    } catch (_) {}
  }

  throw new Error("p2p.js tidak memiliki export yang valid.");
}

let p2pServer;
try {
  p2pServer = initP2P(p2pRaw);
} catch (e) {
  console.error('❌ Gagal membuat P2P server:', e.message);
  process.exit(1);
}

// ----------------------
// 4. Routes HTTP
// ----------------------

// GET chain
app.get('/chain', (req, res) => {
  res.json(blockchain.chain);
});

// POST mine block
app.post('/mine', (req, res) => {
  const data = req.body.data;

  if (!data)
    return res.status(400).json({ error: 'Missing body: { "data": ... }' });

  let newBlock;

  try {
    if (typeof blockchain.addBlock === 'function') {
      newBlock = blockchain.addBlock(data);
    } else if (typeof blockchain.mineBlock === 'function') {
      newBlock = blockchain.mineBlock(data);
    } else {
      return res.status(500).json({
        error: 'Blockchain tidak memiliki addBlock()/mineBlock()',
      });
    }
  } catch (e) {
    console.error('❌ Error add block:', e);
    return res.status(500).json({ error: e.message });
  }

  // ---- Broadcast ----
  if (p2pServer) {
    if (typeof p2pServer.syncChains === 'function') {
      p2pServer.syncChains();
    } else if (typeof p2pServer.broadcastChain === 'function') {
      p2pServer.broadcastChain();
    } else if (typeof p2pServer.broadcast === 'function') {
      p2pServer.broadcast(blockchain.chain);
    } else {
      console.warn('⚠️ Tidak ada fungsi broadcast di p2p.js');
    }
  }

  return res.json({
    message: 'Block added successfully',
    block: newBlock
  });
});

// GET peers
app.get('/peers', (req, res) => {
  if (!p2pServer) return res.json({ peers: [] });

  if (Array.isArray(p2pServer.sockets))
    return res.json({ peers: p2pServer.sockets.length });

  if (typeof p2pServer.getPeers === 'function')
    return res.json({ peers: p2pServer.getPeers() });

  return res.json({ debug: Object.keys(p2pServer) });
});

// health check
app.get('/health', (_, res) => res.json({ status: 'ok' }));

// ----------------------
// 5. Start HTTP + P2P
// ----------------------
app.listen(HTTP_PORT, () => {
  console.log(`🌐 HTTP server berjalan di port ${HTTP_PORT}`);

  if (!p2pServer) {
    console.warn('⚠️ P2P server tidak terbentuk.');
    return;
  }

  if (typeof p2pServer.listen === 'function') {
    p2pServer.listen();
    console.log(`🔗 P2P server listen() pada port ${P2P_PORT}`);
  } else if (typeof p2pServer.start === 'function') {
    p2pServer.start();
    console.log(`🔗 P2P server start() pada port ${P2P_PORT}`);
  } else {
    console.warn('⚠️ P2P instance tidak memiliki listen()/start().');
  }
});
