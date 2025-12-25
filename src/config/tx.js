// src/config/tx.js
module.exports = {
  CHAIN_ID: 'quantumpay-mainnet-1',

  // Gas rules
  GAS: {
    MIN_GAS_PRICE: 1,
    MAX_GAS_PER_TX: 100,
    MAX_GAS_PER_BLOCK: 1000
  },

  // Nonce expiry (anti replay + spam)
  NONCE_TTL_BLOCKS: 100
};

