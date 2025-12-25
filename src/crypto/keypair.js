const { ec: EC } = require('elliptic');
const ec = new EC('secp256k1');

function generateKeypair() {
  const key = ec.genKeyPair();
  return {
    publicKey: key.getPublic('hex'),
    privateKey: key.getPrivate('hex')
  };
}

function loadKeypair(privateKeyHex) {
  return ec.keyFromPrivate(privateKeyHex, 'hex');
}

module.exports = {
  generateKeypair,
  loadKeypair
};

