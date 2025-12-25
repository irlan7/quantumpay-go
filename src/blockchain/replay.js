const { validateChain } = require('./validation');

function replayChain(blocks, chainId) {
  if (!validateChain(blocks, chainId)) {
    throw new Error('Replay failed: invalid chain');
  }
  return blocks;
}

module.exports = { replayChain };

