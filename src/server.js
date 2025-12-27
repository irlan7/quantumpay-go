/**
 * QuantumPay Mainnet Server
 * Mode: STRICT FINALITY + STRICT SLASHING
 * Launch Block Height: 1
 */

const fs = require("fs");
const path = require("path");

const Blockchain = require("./blockchain/blockchain");
const config = require("../config/config.json");

// ================================
// NODE ID & VALIDATORS
// ================================
const NODE_ID = process.env.NODE_ID || "node1";

// Validator set (HARUS sama di semua node)
const VALIDATORS = ["node1", "node2", "node3"];

// ================================
// DATA DIRECTORY SETUP
// ================================
const dataDir = path.join(
  config.paths.dataDir,
  NODE_ID,
  "chain"
);

if (!fs.existsSync(dataDir)) {
  fs.mkdirSync(dataDir, { recursive: true });
}

// ================================
// INIT BLOCKCHAIN
// ================================
const blockchain = new Blockchain({
  nodeId: NODE_ID,
  validators: VALIDATORS,
  config
});

// ================================
// BOOT LOG
// ================================
console.log("=======================================");
console.log("🚀 QuantumPay Mainnet Node Starting");
console.log("=======================================");
console.log("Node ID           :", NODE_ID);
console.log("Chain ID          :", config.network.chainId);
console.log("Consensus         :", config.consensus.type);
console.log("Finality Mode     : STRICT");
console.log("Validators        :", VALIDATORS.join(", "));
console.log("Launch Block      : #1");
console.log("=======================================\n");

// ================================
// BLOCK PRODUCTION LOOP
// ================================
const BLOCK_TIME_MS = config.network.blockTime * 1000;

console.log(`[BOOT] Block time = ${config.network.blockTime}s`);
console.log("[BOOT] Strict Finality & Slashing ENABLED\n");

setInterval(() => {
  try {
    blockchain.tick();
  } catch (err) {
    console.error("[FATAL NODE ERROR]", err.message);
    process.exit(1); // STRICT MODE = HALT
  }
}, BLOCK_TIME_MS);

// ================================
// GRACEFUL SHUTDOWN
// ================================
process.on("SIGINT", () => {
  console.log("\n[SHUTDOWN] SIGINT received");
  console.log("[SHUTDOWN] Saving state & exiting...");
  process.exit(0);
});

process.on("SIGTERM", () => {
  console.log("\n[SHUTDOWN] SIGTERM received");
  console.log("[SHUTDOWN] Saving state & exiting...");
  process.exit(0);
});

