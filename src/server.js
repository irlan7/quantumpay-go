const Blockchain = require("./blockchain/blockchain");
const Voting = require("./consensus/voting");
const Finality = require("./consensus/finality");

const NODE_ID = process.env.NODE_ID || "node1";
const DATA_DIR = `data/${NODE_ID}`;

const validators = ["node1", "node2", "node3"];

const log = (msg) => console.log(msg);

const voting = new Voting({
  nodeId: NODE_ID,
  validators,
  logger: log,
});

const blockchain = new Blockchain({
  nodeId: NODE_ID,
  voting,
  finality: new Finality({
    state: { finalizedHeight: 0 },
    dataDir: DATA_DIR,
    logger: log,
  }),
  logger: log,
});

log(`[BOOT] Starting QuantumPay node: ${NODE_ID}`);
log(`[READY] Node ${NODE_ID} is running`);

setInterval(() => {
  blockchain.produceBlock();
}, 5000);

