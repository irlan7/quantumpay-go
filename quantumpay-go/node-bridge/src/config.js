module.exports = {
  http: {
    port: process.env.HTTP_PORT || 8080
  },
  grpc: {
    target: process.env.GRPC_TARGET || "127.0.0.1:9090"
  }
};
