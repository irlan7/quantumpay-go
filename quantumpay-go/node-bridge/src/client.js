const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

const PROTO_PATH = path.join(__dirname, '../grpc/node.proto');

const packageDef = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const proto = grpc.loadPackageDefinition(packageDef);

// ⬇️ SESUAI package quantumpay.node.v1
const NodeService =
  proto.quantumpay.node.v1.NodeService;

const client = new NodeService(
  '127.0.0.1:9090', // gRPC Go node
  grpc.credentials.createInsecure()
);

module.exports = client;
