// src/blockchain/slashing.js
const fs = require('fs');
const path = require('path');

const STAKING_FILE = path.join(__dirname, '../../data/staking.json');

const SLASHING = {
  DOUBLE_SIGN: 0.05,   // 5%
  INVALID_BLOCK: 0.1,  // 10%
  DOWNTIME: 0.01       // 1%
};

function load() {
  return JSON.parse(fs.readFileSync(STAKING_FILE));
}

function save(data) {
  fs.writeFileSync(STAKING_FILE, JSON.stringify(data, null, 2));
}

function slashValidator(address, reason) {
  const data = load();
  const v = data.validators[address];
  if (!v || !v.active) throw new Error('Validator not active');

  const rate = SLASHING[reason];
  if (!rate) throw new Error('Unknown slashing reason');

  const penalty = Math.floor(v.stake * rate);
  v.stake -= penalty;
  v.slashed += penalty;

  if (v.stake <= 0) {
    v.active = false;
  }

  save(data);
  return penalty;
}

module.exports = {
  slashValidator,
  SLASHING
};

