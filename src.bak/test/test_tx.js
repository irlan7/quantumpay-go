const Blockchain = require("../blockchain");
const Transaction = require("../transaction");

let qpay = new Blockchain();

// buat transaksi
const tx1 = new Transaction("Alice", "Bob", 50);
qpay.addTransaction(tx1);

// buat nodes dummy untuk consensus
const nodes = [
    { id: "NodeA", integrityScore: 0.8, efficiencyScore: 0.4 },
    { id: "NodeB", integrityScore: 0.7, efficiencyScore: 0.6 },
    { id: "NodeC", integrityScore: 0.9, efficiencyScore: 0.7 }
];

qpay.minePendingTransactions("MinerX", nodes);

console.log("\nBlockchain valid:", qpay.isChainValid());
console.log(JSON.stringify(qpay, null, 4));
