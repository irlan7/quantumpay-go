const Blockchain = require('../blockchain');
const Block = require('../block');

// Buat instance blockchain
const qpay = new Blockchain();

// Simulasi node jaringan
const nodes = [
    { id: "NodeA", validations: { valid: 95, total: 100 }, latencyMs: 30, energyUsage: 50 },
    { id: "NodeB", validations: { valid: 80, total: 100 }, latencyMs: 10, energyUsage: 200 },
    { id: "NodeC", validations: { valid: 99, total: 100 }, latencyMs: 100, energyUsage: 10 }
];

// Tambah block pertama → consensus memilih leader
console.log("Mining block 1...");
let block1 = new Block(1, Date.now(), { amount: 10 });
qpay.addBlock(block1, nodes);

// Tambah block kedua
console.log("Mining block 2...");
let block2 = new Block(2, Date.now(), { amount: 20 });
qpay.addBlock(block2, nodes);

// Cek apakah chain valid
console.log("\nBlockchain valid:", qpay.isChainValid());
