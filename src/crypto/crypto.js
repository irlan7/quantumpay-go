const crypto = require('crypto');
const EC = require('elliptic').ec;

// gunakan kurva bitcoin standard
const ec = new EC('secp256k1');

/**
 * Hash data menggunakan SHA256
 */
function sha256(data) {
  return crypto
    .createHash('sha256')
    .update(data)
    .digest('hex');
}

/**
 * Generate keypair (wallet)
 */
function generateKeyPair() {
  const keyPair = ec.genKeyPair();

  return {
    privateKey: keyPair.getPrivate('hex'),
    publicKey: keyPair.getPublic('hex')
  };
}

/**
 * Sign data hash
 */
function signHash(hash, privateKey) {
  const key = ec.keyFromPrivate(privateKey, 'hex');
  const signature = key.sign(hash, { canonical: true });

  return signature.toDER('hex');
}

/**
 * Verify signature
 */
function verifySignature(publicKey, signature, hash) {
  try {
    const key = ec.keyFromPublic(publicKey, 'hex');
    return key.verify(hash, signature);
  } catch (err) {
    return false;
  }
}

module.exports = {
  sha256,
  generateKeyPair,
  signHash,
  verifySignature
};

