'use strict';

const { verifyECDSA } = require('./ecdsa');
const dilithium = require('../pq/dilithium');

/**
 * verifySignature
 *
 * @param {Object} params
 * @param {'ecdsa'|'dilithium'} params.sigType
 * @param {string} params.message        // domain-separated message
 * @param {string|Object} params.signature
 * @param {string|Object} params.pubKey
 *
 * @returns {Promise<boolean>}
 */
async function verifySignature({
  sigType,
  message,
  signature,
  pubKey
}) {
  try {
    if (!sigType || !message || !signature || !pubKey) {
      return false;
    }

    // =========================
    // ECDSA (secp256k1)
    // =========================
    if (sigType === 'ecdsa') {
      return verifyECDSA({
        message,
        signature,
        pubKey
      });
    }

    // =========================
    // DILITHIUM (Post-Quantum)
    // =========================
    if (sigType === 'dilithium') {
      // Real PQ verification (WASM / liboqs)
      return await dilithium.verify(
        message,
        signature,
        pubKey
      );
    }

    // =========================
    // UNKNOWN TYPE
    // =========================
    return false;
  } catch (err) {
    // HARD FAIL SAFE
    return false;
  }
}

module.exports = {
  verifySignature
};

