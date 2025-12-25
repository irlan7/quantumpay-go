'use strict';

const dilithium = require('./dilithium_wasm');

/**
 * API STABLE
 * Consensus tidak perlu tahu WASM
 */

async function generateKeypair() {
  return dilithium.generateKeypair();
}

async function sign(message, privateKey) {
  // privateKey dipakai internal wasm
  return dilithium.sign(message);
}

async function verify(message, signature, publicKey) {
  return dilithium.verify(message, signature);
}

module.exports = {
  generateKeypair,
  sign,
  verify
};

