const crypto = require('crypto');
const { ec: EC } = require('elliptic');
const ec = new EC('secp256k1');

function hashBlockData(block) {
  return crypto
    .createHash('sha256')
    .update(
      block.index +
      block.prevHash +
      block.timestamp +
      JSON.stringify(block.transactions) +
      block.gasUsed +
      block.chainId
    )
    .digest('hex');
}

function signBlock(block, privateKeyHex) {
  const key = ec.keyFromPrivate(privateKeyHex, 'hex');
  const hash = hashBlockData(block);
  const signature = key.sign(hash);

  return signature.toDER('hex');
}

function verifyBlockSignature(block) {
  if (!block.signature || !block.validator) return false;

  const key = ec.keyFromPublic(block.validator, 'hex');
  const hash = hashBlockData(block);

  return key.verify(hash, block.signature);
}

module.exports = {
  signBlock,
  verifyBlockSignature
};

