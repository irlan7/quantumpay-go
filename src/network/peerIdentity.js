'use strict';

const crypto = require('crypto');

function getPeerIdFromCert(cert) {
  return crypto
    .createHash('sha256')
    .update(cert.raw)
    .digest('hex');
}

module.exports = {
  getPeerIdFromCert
};

