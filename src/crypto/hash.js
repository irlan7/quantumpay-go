const crypto = require('crypto');

const DOMAIN = 'QuantumPay-Block';

const cryptoHash = (...inputs) => {
  const hash = crypto.createHash('sha256');
  hash.update(
    DOMAIN +
    inputs.map(i => JSON.stringify(i)).join('|')
  );
  return hash.digest('hex');
};

module.exports = cryptoHash;

