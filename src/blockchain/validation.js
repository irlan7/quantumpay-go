const { verifyBlockSignature } = require('../crypto/signature');
const validators = require('../validators/validators.json');

function isAuthorizedValidator(pubKey, chainId) {
  return (
    validators.chainId === chainId &&
    validators.validators.includes(pubKey)
  );
}

function validateChain(chain, chainId) {
  for (let i = 1; i < chain.length; i++) {
    const block = chain[i];
    const prev = chain[i - 1];

    if (block.prevHash !== prev.hash) return false;
    if (block.chainId !== chainId) return false;
    if (!isAuthorizedValidator(block.validator, chainId)) return false;
    if (!verifyBlockSignature(block)) return false;
  }
  return true;
}

module.exports = {
  validateChain
};

