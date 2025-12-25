const fs = require('fs');
const path = require('path');

const DATA_DIR = path.join(__dirname, '../../data');
const BLOCK_FILE = path.join(DATA_DIR, 'blocks.json');
const FORK_DIR = path.join(DATA_DIR, 'forks');

function loadChain() {
  if (!fs.existsSync(BLOCK_FILE)) return null;
  return JSON.parse(fs.readFileSync(BLOCK_FILE));
}

function saveChain(chain) {
  fs.mkdirSync(DATA_DIR, { recursive: true });
  fs.writeFileSync(BLOCK_FILE, JSON.stringify(chain, null, 2));
}

function saveFork(block, chain) {
  fs.mkdirSync(FORK_DIR, { recursive: true });
  const shortHash = block.hash.slice(0, 6);
  const filename = `fork_${block.height}_${shortHash}.json`;
  const filePath = path.join(FORK_DIR, filename);

  fs.writeFileSync(filePath, JSON.stringify(chain, null, 2));
  console.log(`🪓 Fork saved: ${filename}`);
}

module.exports = {
  loadChain,
  saveChain,
  saveFork
};

