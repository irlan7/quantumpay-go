const WebSocket = require('ws');

class P2PServer {
  constructor(blockchain, transactionPool) {
    this.blockchain = blockchain;
    this.transactionPool = transactionPool;
    this.sockets = [];
  }

  listen(port) {
    const server = new WebSocket.Server({ port });
    server.on('connection', socket => this.connectSocket(socket));
    console.log(`P2P server listening on port ${port}`);
  }

  connectSocket(socket) {
    this.sockets.push(socket);
  }

  broadcastTransaction(transaction) {
    this.sockets.forEach(socket =>
      socket.send(JSON.stringify(transaction))
    );
  }
}

module.exports = P2PServer;

