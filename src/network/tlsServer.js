'use strict';

const https = require('https');
const fs = require('fs');
const express = require('express');
const { getPeerIdFromCert } = require('./peerIdentity');

function startTLSServer({
  port,
  certPath,
  keyPath,
  caPath,
  onVote
}) {
  const app = express();
  app.use(express.json());

  app.post('/vote', (req, res) => {
    const peerCert = req.socket.getPeerCertificate();
    if (!peerCert || !peerCert.raw) {
      return res.status(401).json({ error: 'No client cert' });
    }

    const peerId = getPeerIdFromCert(peerCert);
    const vote = req.body;

    onVote(vote, peerId);
    res.json({ ok: true });
  });

  const server = https.createServer(
    {
      key: fs.readFileSync(keyPath),
      cert: fs.readFileSync(certPath),
      ca: fs.readFileSync(caPath),
      requestCert: true,
      rejectUnauthorized: true
    },
    app
  );

  server.listen(port, () => {
    console.log(`[P2P-TLS] Listening on ${port}`);
  });

  return server;
}

module.exports = { startTLSServer };

