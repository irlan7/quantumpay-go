const EC = require('elliptic').ec;
const ec = new EC('secp256k1');

function verifySignature(publicKeyHex, data, signatureHex) {
  try {
    const key = ec.keyFromPublic(publicKeyHex, 'hex');
    return key.verify(data, signatureHex);
  } catch (err) {
    console.error('Signature verify error:', err.message);
    return false;
  }
}

module.exports = { verifySignature };

