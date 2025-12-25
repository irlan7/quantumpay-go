const fs = require("fs");
const path = require("path");

const DATA_DIR = path.join(__dirname, "../../data");
const STAKING_FILE = path.join(DATA_DIR, "staking.json");

const MIN_STAKE = 1000;
const UNBONDING_DELAY = 100; // block height

function ensureFile() {
  if (!fs.existsSync(DATA_DIR)) fs.mkdirSync(DATA_DIR, { recursive: true });
  if (!fs.existsSync(STAKING_FILE)) {
    fs.writeFileSync(
      STAKING_FILE,
      JSON.stringify({ validators: {}, totalStaked: 0 }, null, 2)
    );
  }
}

function load() {
  ensureFile();
  return JSON.parse(fs.readFileSync(STAKING_FILE));
}

function save(data) {
  fs.writeFileSync(STAKING_FILE, JSON.stringify(data, null, 2));
}

/* =========================
   CORE STAKING FUNCTIONS
========================= */

function bond(address, amount, height) {
  if (amount < MIN_STAKE) {
    throw new Error("Stake below minimum");
  }

  const data = load();

  if (!data.validators[address]) {
    data.validators[address] = {
      stake: 0,
      bonded: false,
      jailed: false,
      unbondHeight: null
    };
  }

  data.validators[address].stake += amount;
  data.validators[address].bonded = true;
  data.validators[address].jailed = false;
  data.validators[address].unbondHeight = null;

  data.totalStaked += amount;
  save(data);
}

function requestUnbond(address, height) {
  const data = load();
  const v = data.validators[address];

  if (!v || !v.bonded) {
    throw new Error("Validator not bonded");
  }

  v.unbondHeight = height + UNBONDING_DELAY;
  save(data);
}

function finalizeUnbond(address, height) {
  const data = load();
  const v = data.validators[address];

  if (!v || v.unbondHeight === null) {
    throw new Error("Unbond not requested");
  }

  if (height < v.unbondHeight) {
    throw new Error("Unbonding period not finished");
  }

  data.totalStaked -= v.stake;
  const released = v.stake;

  v.stake = 0;
  v.bonded = false;
  v.unbondHeight = null;

  save(data);
  return released;
}

function getActiveValidators() {
  const data = load();
  return Object.entries(data.validators)
    .filter(([_, v]) => v.bonded && !v.jailed && v.stake >= MIN_STAKE)
    .map(([address, v]) => ({
      address,
      stake: v.stake
    }));
}

function getTotalStake() {
  return load().totalStaked;
}

module.exports = {
  bond,
  requestUnbond,
  finalizeUnbond,
  getActiveValidators,
  getTotalStake,
  MIN_STAKE
};

