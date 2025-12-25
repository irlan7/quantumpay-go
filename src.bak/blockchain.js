const Block = require('./block');
const QHC = require('./consensus/qhc'); // pastikan file ini ada dan mengekspor class QHC
const Transaction = require('./transaction');

class Blockchain {
  constructor() {
    this.chain = [this.createGenesisBlock()];
    this.difficulty = 3;
    this.pendingTransactions = [];
    this.miningReward = 10;
    this.consensus = new QHC();
  }

  createGenesisBlock() {
    return new Block(0, Date.now(), "Genesis Block", "0");
  }

  getLatestBlock() {
    return this.chain[this.chain.length - 1];
  }

  // nodes = array of node identifiers (mis. ['NodeA','NodeB'])
  addBlock(newBlock, nodes = []) {
    let leader = "Unknown";
    if (nodes.length > 0) {
      leader = this.consensus.selectLeader(nodes);
      // gunakan backticks supaya ${leader} dievaluasi
      console.log(`Consensus selected leader: ${leader}`);
    }

    newBlock.previousHash = this.getLatestBlock().hash;
    newBlock.mineBlock(this.difficulty);

    console.log(`Block mined: ${newBlock.hash}`);
    this.chain.push(newBlock);
  }

addTransaction(tx) {
     if(!tx.fromAddress || !tx.toAddress){
         throw new Error("Transaksi harus punya fromAddress & toAddress");
     }

     tx.hash = tx.calculateHash();
     this.pendingTransactions.push(tx);
}

minePendingTransactions(minerAddress, nodes = []{
    let leader = "Unknown";
    if (nodes && nodes.length > 0) {
    leader = this.consensus.selectLeader(nodes);
    console.log(`Consensus selected leader: ${leader}`);
  }
    //add reward Transaction
    const rewardTx = new Transaction(null, minerAddress, this.miningReward || 10);
    this.pendingTransactions.push(rewardTx);

    const newBlock = new Block(this.chain.length, Date.now(), this.pendingTransactions);
    newBlock.previousHash = this.getLatestBlock().hash;
    newBlock.mineBlock(this.diffculty);

   
   console.log(`Block mined: ${newBlock.hash}`);

   this.chain.push(newBlock);

    //Reset pending transactions + Reward tx for next round
    this.pendingTransactions = [
        new Transaction(null, leader, this.miningReward || 10)
    ];
}


  isChainValid() {
    for (let i = 1; i < this.chain.length; i++) {
      const current = this.chain[i];
      const previous = this.chain[i - 1];

      if (current.hash !== current.calculateHash()) return false;
      if (current.previousHash !== previous.hash) return false;
    }
    return true;
  }
}

module.exports = Blockchain;
