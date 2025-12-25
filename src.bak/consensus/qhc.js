// src/consensus/qhc.js
// QHC = Quantum Hybrid Consensus (contoh sederhana untuk pengembangan)
// Implementasi minimal: hitung hybrid score (integrity + efficiency) & pilih leader

class QHC {
  constructor(opts = {}) {
    // bobot default, bisa diubah saat inisialisasi
    this.weightIntegrity = opts.weightIntegrity ?? 0.6;
    this.weightEfficiency = opts.weightEfficiency ?? 0.4;
  }

  // Proof-of-Integrity: hitung skor kejujuran node
  // validations: { valid: <number>, total: <number> }
  calculateIntegrity(nodeId, validations = { valid: 0, total: 0 }) {
    // hindari pembagian nol
    return validations.total > 0 ? (validations.valid / validations.total) : 0;
  }

  // Proof-of-Efficiency: hitung skor efisiensi berdasarkan latency & penggunaan energi
  // Sebagai contoh sederhana: latencyMs terbalik (lebih kecil lebih baik), energyUsage terbalik
  calculateEfficiency(latencyMs = 1000, energyUsage = 1000) {
    // normalisasi sederhana ke range 0..1
    const latScore = 1 / (1 + latencyMs);        // semakin kecil latency -> semakin besar score
    const eneScore = 1 / (1 + energyUsage);      // semakin kecil energyUsage -> semakin besar score
    return (latScore + eneScore) / 2;            // rata-rata sederhana
  }

  // Hitung hybrid score gabungan untuk satu node
  // node = { id, validations, latencyMs, energyUsage }
  calculateHybridScore(node = {}) {
    const integrityScore = this.calculateIntegrity(node.id, node.validations || { valid: 0, total: 0 });
    const efficiencyScore = this.calculateEfficiency(node.latencyMs || 1000, node.energyUsage || 1000);
    return (integrityScore * this.weightIntegrity) + (efficiencyScore * this.weightEfficiency);
  }

  // Pilih node leader dari array nodes
  // nodes = [{ id, validations, latencyMs, energyUsage }, ...]
  selectLeader(nodes = []) {
    if (!Array.isArray(nodes) || nodes.length === 0) return 'Unknown';
    let best = null;
    let bestScore = -Infinity;
    for (const node of nodes) {
      const score = this.calculateHybridScore(node);
      if (score > bestScore) {
        bestScore = score;
        best = node.id ?? node;
      }
    }
    return best;
  }
}

module.exports = QHC;
