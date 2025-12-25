const Blockchain = require("../blockchain");
const Block = require("../block");
const QHC = require("../consensus/qhc");

// Initialize blockchain
const qpay = new Blockchain();

// Sample validator nodes
const nodes = [
    { id: "NodeA", integrityScore: 0.98, efficiencyScore: 0.80 },
    { id: "NodeB", integrityScore: 0.70, efficiencyScore: 0.90 },
    { id: "NodeC", integrityScore: 0.90, efficiencyScore: 0.70 }
];

console.log("Mining block 1...");
qpay.addBlock(new Block(1, Date.now(), { amount: 10 }), nodes);

console.log("Mining block 2...");
qpay.addBlock(new Block(2, Date.now(), { amount: 20 }), nodes);

console.log("\nBlockchain valid:", qpay.isChainValid());
