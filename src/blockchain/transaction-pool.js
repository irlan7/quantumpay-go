// transaction-pool.js
// QuantumPay Chain — Blockchain v4.1

class TransactionPool {
  constructor() {
    // Menyimpan transaksi yang belum masuk block
    this.transactions = [];
  }

  /**
   * Tambah transaksi baru atau update transaksi existing
   * (berdasarkan hash transaksi)
   */
  addOrUpdateTransaction(transaction) {
    const txHash = transaction.calculateHash();
    const existingIndex = this.transactions.findIndex(
      t => t.calculateHash() === txHash
    );

    if (existingIndex >= 0) {
      console.log("[TP] Updating existing transaction in pool...");
      this.transactions[existingIndex] = transaction;
    } else {
      console.log("[TP] Adding new transaction to pool...");
      this.transactions.push(transaction);
    }
  }

  /**
   * Cek apakah wallet sudah memiliki transaksi pending sebelumnya
   */
  existingTransaction(fromAddress) {
    return this.transactions.find(tx => tx.fromAddress === fromAddress);
  }

  /**
   * Mengambil hanya transaksi valid
   * (Signature valid & format valid)
   */
  validTransactions() {
    return this.transactions.filter(tx => {
      try {
        return tx.isValid();
      } catch (err) {
        console.log("[TP] Invalid transaction removed:", err.message);
        return false;
      }
    });
  }

  /**
   * Hapus transaksi berdasarkan hash — digunakan setelah mining
   */
  removeTransactionByHash(hash) {
    this.transactions = this.transactions.filter(
      tx => tx.calculateHash() !== hash
    );
  }

  /**
   * Hapus semua transaksi — digunakan setelah block baru berhasil ditambang
   */
  clear() {
    console.log("[TP] Clearing transaction pool...");
    this.transactions = [];
  }

  /**
   * Sinkronisasi pool dari peer lain
   * (digunakan dalam P2P broadcast)
   */
  replacePool(newPool) {
    console.log("[TP] Replacing transaction pool with received pool...");
    this.transactions = newPool;
  }

  /**
   * Menghapus transaksi yang sudah masuk ke block
   */
  clearMinedTransactions(minedTransactions) {
    const minedHashes = minedTransactions.map(tx => tx.calculateHash());
    this.transactions = this.transactions.filter(
      tx => !minedHashes.includes(tx.calculateHash())
    );
  }
}

module.exports = TransactionPool;
