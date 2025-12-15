const express = require('express');
const bodyParser = require('body-parser');

const Blockchain = require('./blockchain/blockchain');
const Transaction = require('./blockchain/transaction');

const P2PServer = require('./p2p/p2p-server');

const HTTP_PORT = process.env.HTTP_PORT || 3001;
const P2P_PORT = process.env.P2P_PORT || 6001;

const app = express();
app.use(bodyParser.json());

const blockchain = new Blockchain();
const p2pServer = new P2PServer(blockchain);

/**
 * Create new transaction (with signature verification)
 */
app.post('/transactions/new', (req, res) => {
  try {
    const { from, to, amount, signature } = req.body;

    const tx = new Transaction(from, to, amount, signature);

    if (!tx.isValid()) {
      return res.status(400).json({ error: 'Invalid signature' });
    }

    blockchain.addTransaction(tx);
    p2pServer.broadcastTransaction(tx);

    res.json({ message: 'Transaction accepted', tx });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

/**
 * Get blockchain
 */
app.get('/blocks', (req, res) => {
  res.json(blockchain.chain);
});

app.listen(HTTP_PORT, () => {
  console.log(`HTTP server listening on port ${HTTP_PORT}`);
});

p2pServer.listen(P2P_PORT);

