// src/test/test_consensus.js
const QHC = require('../consensus/qhc');

const q = new QHC();
const nodes = [
  { id: 'NodeA', validations: { valid: 95, total: 100 }, latencyMs: 20, energyUsage: 50 },
  { id: 'NodeB', validations: { valid: 80, total: 100 }, latencyMs: 5, energyUsage: 200 },
  { id: 'NodeC', validations: { valid: 99, total: 100 }, latencyMs: 100, energyUsage: 10 }
];

console.log('Leader:', q.selectLeader(nodes));
console.log('Scores:');
for (const n of nodes) {
  console.log(n.id, q.calculateHybridScore(n));
}
