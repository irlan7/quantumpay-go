const { domainHash } = require('./hash');
const EC = require('elliptic').ec;
const ec = new EC('secp256k1');

function signWithDomain({ privateKey, domain, payload }) {
  const key = ec.keyFromPrivate(privateKey, 'hex');
  const hash = domainHash(domain, payload);
  return key.sign(hash).toDER('hex');
}

module.exports = { signWithDomain };

