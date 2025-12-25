'use strict';

const https = require('https');
const fs = require('fs');

function sendVoteTLS({
  peerUrl,
  vote,
  certPath,
  keyPath,
  caPath
}) {
  const data = JSON.stringify(vote);

  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Content-Length': data.length
    },
    cert: fs.readFileSync(certPath),
    key: fs.readFileSync(keyPath),
    ca: fs.readFileSync(caPath),
    rejectUnauthorized: true
  };

  const req = https.request(
    `${peerUrl}/vote`,
    options,
    () => {}
  );

  req.on('error', () => {});
  req.write(data);
  req.end();
}

module.exports = { sendVoteTLS };

