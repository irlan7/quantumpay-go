const WebSocket = require('ws');

const MESSAGE_TYPES = {
  CHAIN: 'CHAIN',
  TRANSACTION: 'TRANSACTION'
};

class P2PServer {
  constructor(blockchain, transactionPool) {
    this.blockchain = blockchain;
    this.transactionPool = transactionPool;
    this.sockets = [];
  }

  listen() {
    const p2pPort = process.env.P2P_PORT || 6001;
    const server = new WebSocket.Server({ port: p2pPort });

    server.on('connection', socket => {
      this.connectSocket(socket);
    });

    console.log(`P2P server listening on port ${p2pPort}`);
  }

  connectSocket(socket) {
    this.sockets.push(socket);
    console.log('New peer connected');

    this.messageHandler(socket);
    this.sendChain(socket);
  }

  messageHandler(socket) {
    socket.on('message', data => {
      try {
        const message = JSON.parse(data);

        switch (message.type) {
          case MESSAGE_TYPES.CHAIN:
            this.blockchain.replaceChain(message.chain);
            break;

          case MESSAGE_TYPES.TRANSACTION:
            this.transactionPool.addTransaction(message.transaction);
            break;

          default:
            console.log('Unknown message type');
        }
      } catch (err) {
        console.error('P2P message error:', err.message);
      }
    });
  }

  sendChain(socket) {
    socket.send(JSON.stringify({
      type: MESSAGE_TYPES.CHAIN,
      chain: this.blockchain.chain
    }));
  }

  sendTransaction(socket, transaction) {
    socket.send(JSON.stringify({
      type: MESSAGE_TYPES.TRANSACTION,
      transaction
    }));
  }

  broadcastChain() {
    this.sockets.forEach(socket => this.sendChain(socket));
  }

  broadcastTransaction(transaction) {
    this.sockets.forEach(socket =>
      this.sendTransaction(socket, transaction)
    );
  }
}

module.exports = P2PServer;

