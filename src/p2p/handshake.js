function handleHandshake(socket, node) {
  socket.write(JSON.stringify({
    type: 'HELLO',
    chainId: node.chainId,
    height: node.blockchain.getHeight()
  }));
}

module.exports = { handleHandshake };

