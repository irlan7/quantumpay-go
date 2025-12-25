const { CHAIN_ID, NETWORK_ID } = require('../config/network');

function domainBase() {
  return `QPAY::${CHAIN_ID}::${NETWORK_ID}`;
}

module.exports = {
  TX: `${domainBase()}::TX::v1`,
  BLOCK: `${domainBase()}::BLOCK::v1`,
  SNAPSHOT: `${domainBase()}::SNAPSHOT::v1`,
  CONTRACT: (name, method) =>
    `${domainBase()}::CONTRACT::${name}::${method}::v1`
};

