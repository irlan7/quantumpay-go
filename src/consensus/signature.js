const cryptoHash = require('../crypto/hash');
const DOMAIN = require('../crypto/domain');

function consensusHash(data) {
  return cryptoHash(DOMAIN.CONSENSUS, data);
}

module.exports = consensusHash;

