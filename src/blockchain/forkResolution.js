const crypto = require('crypto');

/**
 * Calculate total gas used by a chain
 */
function totalGas(chain) {
  return chain.reduce((sum, block) => sum + (block.gasUsed || 0), 0);
}

/**
 * Deterministic fork choice rule
 */
function chooseBestChain(localChain, remoteChain) {
  if (!remoteChain || remoteChain.length === 0) return localChain;

  // 1. Higher total gas wins
  const localGas = totalGas(localChain);
  const remoteGas = totalGas(remoteChain);
  if (remoteGas !== localGas) {
    return remoteGas > localGas ? remoteChain : localChain;
  }

  // 2. Higher height wins
  if (remoteChain.length !== localChain.length) {
    return remoteChain.length > localChain.length ? remoteChain : localChain;
  }

  // 3. Earlier last block timestamp
  const localTs = localChain.at(-1).timestamp;
  const remoteTs = remoteChain.at(-1).timestamp;
  if (remoteTs !== localTs) {
    return remoteTs < localTs ? remoteChain : localChain;
  }

  // 4. Lexicographically smallest hash (final deterministic tie-break)
  const lHash = localChain.at(-1).hash;
  const rHash = remoteChain.at(-1).hash;
  return rHash < lHash ? remoteChain : localChain;
}

module.exports = {
  chooseBestChain
};

