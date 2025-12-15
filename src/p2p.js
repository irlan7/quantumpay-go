// ===============================
//  p2p.js — QuantumPay Chain P2P (patched for server.js compatibility)
// ===============================

const WebSocket = require('ws');

const peers = []; // active websocket connections
let blockchain = undefined;

// ----------------------------------
//  Enum Tipe Pesan
// ----------------------------------
const MESSAGE_TYPES = {
  REQUEST_CHAIN: 'REQUEST_CHAIN',
  REQUEST_BLOCK: 'REQUEST_BLOCK',
  SEND_CHAIN: 'SEND_CHAIN',
  SEND_BLOCK: 'SEND_BLOCK',
  TRANSACTION: 'TRANSACTION',
};

// ----------------------------------
//  Internal helpers
// ----------------------------------
function safeParse(msg) {
  try {
    return JSON.parse(msg);
  } catch (e) {
    return null;
  }
}

function sendMessage(ws, message) {
  try {
    ws.send(JSON.stringify(message));
  } catch (err) {
    console.warn('sendMessage failed:', err && err.message ? err.message : err);
  }
}

function broadcast(message) {
  peers.forEach((peer) => {
    if (peer && peer.readyState === WebSocket.OPEN) {
      sendMessage(peer, message);
    }
  });
}

function responseChain() {
  return {
    type: MESSAGE_TYPES.SEND_CHAIN,
    data: blockchain && typeof blockchain.getChain === 'function' ? safeCallGetChain() : [],
  };
}

function responseLatestBlock() {
  return {
    type: MESSAGE_TYPES.SEND_BLOCK,
    data: blockchain && typeof blockchain.getLatestBlock === 'function' ? safeCallGetLatest() : null,
  };
}

function safeCallGetChain() {
  try {
    const maybe = blockchain.getChain();
    return typeof maybe.then === 'function' ? undefined : maybe; // if async, ignore here
  } catch (e) {
    return undefined;
  }
}

function safeCallGetLatest() {
  try {
    const maybe = blockchain.getLatestBlock();
    return typeof maybe.then === 'function' ? undefined : maybe;
  } catch (e) {
    return undefined;
  }
}

// ----------------------------------
//  Init / set blockchain
// ----------------------------------
function init(options = {}) {
  const p2pPort = options.p2pPort || options.port || process.env.P2P_PORT || 6001;
  const chainRef = options.chainRef || options.blockchain;
  if (chainRef) blockchain = chainRef;

  const server = new WebSocket.Server({ port: Number(p2pPort) });
  console.log(`🚀 P2P server running on ws://localhost:${p2pPort}`);

  server.on('connection', (ws) => {
    console.log('🔗 New peer connected');
    initConnection(ws);
  });

  server.on('error', (err) => {
    console.error('P2P server error:', err && err.message ? err.message : err);
  });

  return server;
}

function setBlockchain(chainRef) {
  blockchain = chainRef;
  console.log('🔧 blockchain reference set on p2p module');
}

// ----------------------------------
//  Koneksi ke peer lain (single)
// ----------------------------------
function connectToPeer(peerUrl) {
  if (!peerUrl) return;
  try {
    console.log(`🔌 Connecting to peer: ${peerUrl}`);
    const ws = new WebSocket(peerUrl);

    ws.on('open', () => initConnection(ws));
    ws.on('error', (err) => {
      console.log('❌ Connection failed:', peerUrl, err && err.message ? err.message : '');
      // ensure ws closed/cleanup
      try { ws.close(); } catch (_) {}
    });
  } catch (err) {
    console.warn('connectToPeer failed:', err && err.message ? err.message : err);
  }
}

// ----------------------------------
//  Koneksi ke banyak peers (array) - expected by server.js
// ----------------------------------
function connectToPeers(peersArray) {
  if (!Array.isArray(peersArray)) {
    console.warn('connectToPeers expects an array');
    return;
  }
  for (const p of peersArray) {
    // normalize: ensure ws:// or wss://
    let url = String(p || '').trim();
    if (!url) continue;
    if (!/^wss?:\/\//i.test(url)) {
      // try to convert http(s) -> ws(s)
      if (/^https:\/\//i.test(url)) url = url.replace(/^https:\/\//i, 'wss://');
      else if (/^http:\/\//i.test(url)) url = url.replace(/^http:\/\//i, 'ws://');
      else url = 'ws://' + url;
    }
    connectToPeer(url);
  }
}

// ----------------------------------
//  Setup koneksi peer
// ----------------------------------
function initConnection(ws) {
  if (!ws) return;
  peers.push(ws);
  console.log(`📡 Total peers: ${peers.length}`);

  ws.on('message', (msg) => {
    // msg can be Buffer or string
    const text = typeof msg === 'string' ? msg : msg.toString();
    const message = safeParse(text);
    if (!message) {
      console.warn('Received non-JSON message (ignored)');
      return;
    }
    handleMessage(ws, message);
  });

  ws.on('close', () => closeConnection(ws));
  ws.on('error', () => closeConnection(ws));

  // Send latest block (best-effort)
  try {
    const latest = responseLatestBlock();
    sendMessage(ws, latest);
  } catch (err) {
    // ignore
  }
}

// ----------------------------------
//  Bersihkan koneksi mati
// ----------------------------------
function closeConnection(ws) {
  console.log('⚠️ Peer disconnected');
  const index = peers.indexOf(ws);
  if (index !== -1) peers.splice(index, 1);
}

// ----------------------------------
//  Implementasi broadcastChain()
// ----------------------------------
function broadcastChain() {
  if (!blockchain || typeof blockchain.getChain !== 'function') {
    console.warn('broadcastChain called but blockchain.getChain not available');
    return;
  }
  console.log('📢 Broadcasting chain to all peers...');
  broadcast({
    type: MESSAGE_TYPES.SEND_CHAIN,
    data: blockchain.getChain(),
  });
}

// ----------------------------------
//  Broadcast transaction
// ----------------------------------
function broadcastTransaction(tx) {
  if (!tx) return;
  try {
    broadcast({
      type: MESSAGE_TYPES.TRANSACTION,
      data: tx,
    });
  } catch (err) {
    console.warn('broadcastTransaction failed:', err && err.message ? err.message : err);
  }
}

// ----------------------------------
//  Handle message dari peer
// ----------------------------------
function handleMessage(ws, message) {
  if (!message || !message.type) {
    console.warn('Invalid message format from peer');
    return;
  }

  switch (message.type) {
    case MESSAGE_TYPES.REQUEST_CHAIN:
      console.log('📥 Peer meminta chain');
      sendMessage(ws, responseChain());
      break;

    case MESSAGE_TYPES.REQUEST_BLOCK:
      console.log('📥 Peer meminta block terakhir');
      sendMessage(ws, responseLatestBlock());
      break;

    case MESSAGE_TYPES.SEND_CHAIN:
      console.log('🔄 Menerima chain dari peer');
      if (message.data && Array.isArray(message.data)) {
        try {
          if (blockchain && typeof blockchain.replaceChain === 'function') {
            blockchain.replaceChain(message.data);
          } else if (blockchain && Array.isArray(blockchain.chain) && message.data.length > blockchain.chain.length) {
            blockchain.chain = message.data;
          } else {
            console.warn('No replaceChain available or chain not longer -> ignoring');
          }
        } catch (err) {
          console.warn('Error while applying received chain:', err && err.message ? err.message : err);
        }
      } else {
        console.warn('Received SEND_CHAIN with invalid data');
      }
      break;

    case MESSAGE_TYPES.SEND_BLOCK:
      console.log('🔄 Menerima block baru dari peer');
      try {
        if (!blockchain) {
          console.warn('No blockchain instance to apply block');
          break;
        }
        const block = message.data;
        const localLatest = typeof blockchain.getLatestBlock === 'function' ? blockchain.getLatestBlock() : (blockchain.chain && blockchain.chain[blockchain.chain.length - 1]);
        if (!localLatest) {
          console.warn('No local latest block available');
          break;
        }

        if (block && typeof block.index === 'number') {
          if (block.index === localLatest.index + 1) {
            console.log('🧱 Menambahkan block dari peer');
            if (typeof blockchain.addBlockFromPeer === 'function') {
              blockchain.addBlockFromPeer(block);
            } else if (typeof blockchain.addBlock === 'function') {
              blockchain.addBlock(block);
            } else {
              // fallback: push
              blockchain.chain = blockchain.chain || [];
              blockchain.chain.push(block);
            }
          } else if (block.index > localLatest.index) {
            console.log('📣 Block peer lebih maju → request full chain');
            broadcast({ type: MESSAGE_TYPES.REQUEST_CHAIN });
          } else {
            console.log('ℹ️ Block lebih lama, diabaikan');
          }
        }
      } catch (err) {
        console.warn('Error handling SEND_BLOCK:', err && err.message ? err.message : err);
      }
      break;

    case MESSAGE_TYPES.TRANSACTION:
      console.log('📤 Menerima transaction broadcast from peer');
      try {
        if (blockchain && typeof blockchain.addTransaction === 'function') {
          blockchain.addTransaction(message.data);
        } else {
          console.warn('No addTransaction method on blockchain; ignoring tx');
        }
      } catch (err) {
        console.warn('Error applying received transaction:', err && err.message ? err.message : err);
      }
      break;

    default:
      console.log('❓ Unknown message type:', message.type);
  }
}

// ----------------------------------
//  Exports (compatible names for server.js)
// ----------------------------------
module.exports = {
  init,               // init({ p2pPort, chainRef })
  setBlockchain,      // set blockchain reference later
  connectToPeer,      // connectToPeer(url) - kept for compatibility
  connectToPeers,     // connectToPeers([urls])
  broadcastChain,     // broadcastChain()
  broadcastTransaction, // broadcastTransaction(tx)
  peers,              // array of ws
};

