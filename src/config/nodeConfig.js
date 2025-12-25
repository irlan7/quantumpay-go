/**
 * Node Configuration
 * Determines behavior based on NODE_MODE
 */

const NODE_MODE = process.env.NODE_MODE || 'full';

const MODES = {
  ARCHIVE: 'archive',
  FULL: 'full',
  LIGHT: 'light'
};

const nodeConfig = {
  mode: NODE_MODE,

  isArchive() {
    return this.mode === MODES.ARCHIVE;
  },

  isFull() {
    return this.mode === MODES.FULL;
  },

  isLight() {
    return this.mode === MODES.LIGHT;
  },

  allowPruning() {
    return this.isFull();
  },

  allowMining() {
    return this.isFull() || this.isArchive();
  },

  allowConsensus() {
    return this.isFull() || this.isArchive();
  },

  allowFullChainAPI() {
    return this.isArchive();
  }
};

module.exports = nodeConfig;

