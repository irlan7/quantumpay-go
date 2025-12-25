const fs = require("fs");
const path = require("path");

const DATA_DIR = path.join(__dirname, "../../data");
const STAKING_FILE = path.join(DATA_DIR, "staking.json");

/* =========================
   SLASHING RULES (FINAL)
========================= */

const SLASH = {
  DOUBLE_SIGN: 0.05,     // 5%
  INVALID_BLOCK: 0.10,   // 10%
  FINALITY_FAULT: 0.20,  // 20%
  DOWNTIME: 0.01         // 1%
};

function load() {
  return JSON.parse(fs.readFileSync(STAKING_FILE));
}

function save(data) {
  fs.writeFileSync(STAKING_FILE, JSON.stringify(data, null, 2));
}

function slash(address, reason) {
  const data = load();
  const v = data.validators[address];

  if (!v || !v.bonded) {
    throw new Error("Validator not bonded");
  }

  const rate = SLASH[reason];
  if (!rate) {
    throw new Error("Unknown slashing reason");
  }

  const penalty = Math.floor(v.stake * rate);

  v.stake -= penalty;
  data.totalStaked -= penalty;

  // auto-jail for serious faults
  if (rate >= 0.10) {
    v.jailed = true;
    v.bonded = false;
  }

  save(data);

  console.warn(
    `⚠️ SLASHED ${address} | reason=${reason} | penalty=${penalty}`
  );

  return penalty;
}

module.exports = {
  slash,
  SLASH
};

