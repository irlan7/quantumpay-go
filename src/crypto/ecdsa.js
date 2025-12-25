'use strict';

const { ec: EC } = require('elliptic');
const ec = new EC('secp256k1');
const crypto = require('crypto');

function hashMessage(msg) {
  return crypto.createHash('sha256').update(msg).digest();
}

function verifyECDSA({ message, signature, pubKey }) {
  try {
    const key = ec.keyFromPublic(pubKey, 'hex');
    const msgHash = hashMessage(message);
    return key.verify(msgHash, signature);
  } catch (e) {
    return false;
  }
}

module.exports = {
  verifyECDSA
};

