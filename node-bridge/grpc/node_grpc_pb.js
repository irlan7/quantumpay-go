// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var node_pb = require('./node_pb.js');

function serialize_quantumpay_node_v1_GetBalanceRequest(arg) {
  if (!(arg instanceof node_pb.GetBalanceRequest)) {
    throw new Error('Expected argument of type quantumpay.node.v1.GetBalanceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_GetBalanceRequest(buffer_arg) {
  return node_pb.GetBalanceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_GetBalanceResponse(arg) {
  if (!(arg instanceof node_pb.GetBalanceResponse)) {
    throw new Error('Expected argument of type quantumpay.node.v1.GetBalanceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_GetBalanceResponse(buffer_arg) {
  return node_pb.GetBalanceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_GetBlockRequest(arg) {
  if (!(arg instanceof node_pb.GetBlockRequest)) {
    throw new Error('Expected argument of type quantumpay.node.v1.GetBlockRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_GetBlockRequest(buffer_arg) {
  return node_pb.GetBlockRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_GetBlockResponse(arg) {
  if (!(arg instanceof node_pb.GetBlockResponse)) {
    throw new Error('Expected argument of type quantumpay.node.v1.GetBlockResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_GetBlockResponse(buffer_arg) {
  return node_pb.GetBlockResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_GetHeightRequest(arg) {
  if (!(arg instanceof node_pb.GetHeightRequest)) {
    throw new Error('Expected argument of type quantumpay.node.v1.GetHeightRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_GetHeightRequest(buffer_arg) {
  return node_pb.GetHeightRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_GetHeightResponse(arg) {
  if (!(arg instanceof node_pb.GetHeightResponse)) {
    throw new Error('Expected argument of type quantumpay.node.v1.GetHeightResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_GetHeightResponse(buffer_arg) {
  return node_pb.GetHeightResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_GetStatusRequest(arg) {
  if (!(arg instanceof node_pb.GetStatusRequest)) {
    throw new Error('Expected argument of type quantumpay.node.v1.GetStatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_GetStatusRequest(buffer_arg) {
  return node_pb.GetStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_GetStatusResponse(arg) {
  if (!(arg instanceof node_pb.GetStatusResponse)) {
    throw new Error('Expected argument of type quantumpay.node.v1.GetStatusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_GetStatusResponse(buffer_arg) {
  return node_pb.GetStatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_HealthRequest(arg) {
  if (!(arg instanceof node_pb.HealthRequest)) {
    throw new Error('Expected argument of type quantumpay.node.v1.HealthRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_HealthRequest(buffer_arg) {
  return node_pb.HealthRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_quantumpay_node_v1_HealthResponse(arg) {
  if (!(arg instanceof node_pb.HealthResponse)) {
    throw new Error('Expected argument of type quantumpay.node.v1.HealthResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_quantumpay_node_v1_HealthResponse(buffer_arg) {
  return node_pb.HealthResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// =====================
// Core Node Service
// =====================
var NodeServiceService = exports.NodeServiceService = {
  // Chain status
getStatus: {
    path: '/quantumpay.node.v1.NodeService/GetStatus',
    requestStream: false,
    responseStream: false,
    requestType: node_pb.GetStatusRequest,
    responseType: node_pb.GetStatusResponse,
    requestSerialize: serialize_quantumpay_node_v1_GetStatusRequest,
    requestDeserialize: deserialize_quantumpay_node_v1_GetStatusRequest,
    responseSerialize: serialize_quantumpay_node_v1_GetStatusResponse,
    responseDeserialize: deserialize_quantumpay_node_v1_GetStatusResponse,
  },
  // Blockchain data
getHeight: {
    path: '/quantumpay.node.v1.NodeService/GetHeight',
    requestStream: false,
    responseStream: false,
    requestType: node_pb.GetHeightRequest,
    responseType: node_pb.GetHeightResponse,
    requestSerialize: serialize_quantumpay_node_v1_GetHeightRequest,
    requestDeserialize: deserialize_quantumpay_node_v1_GetHeightRequest,
    responseSerialize: serialize_quantumpay_node_v1_GetHeightResponse,
    responseDeserialize: deserialize_quantumpay_node_v1_GetHeightResponse,
  },
  getBlock: {
    path: '/quantumpay.node.v1.NodeService/GetBlock',
    requestStream: false,
    responseStream: false,
    requestType: node_pb.GetBlockRequest,
    responseType: node_pb.GetBlockResponse,
    requestSerialize: serialize_quantumpay_node_v1_GetBlockRequest,
    requestDeserialize: deserialize_quantumpay_node_v1_GetBlockRequest,
    responseSerialize: serialize_quantumpay_node_v1_GetBlockResponse,
    responseDeserialize: deserialize_quantumpay_node_v1_GetBlockResponse,
  },
  // Account state
getBalance: {
    path: '/quantumpay.node.v1.NodeService/GetBalance',
    requestStream: false,
    responseStream: false,
    requestType: node_pb.GetBalanceRequest,
    responseType: node_pb.GetBalanceResponse,
    requestSerialize: serialize_quantumpay_node_v1_GetBalanceRequest,
    requestDeserialize: deserialize_quantumpay_node_v1_GetBalanceRequest,
    responseSerialize: serialize_quantumpay_node_v1_GetBalanceResponse,
    responseDeserialize: deserialize_quantumpay_node_v1_GetBalanceResponse,
  },
  // Health / liveness
health: {
    path: '/quantumpay.node.v1.NodeService/Health',
    requestStream: false,
    responseStream: false,
    requestType: node_pb.HealthRequest,
    responseType: node_pb.HealthResponse,
    requestSerialize: serialize_quantumpay_node_v1_HealthRequest,
    requestDeserialize: deserialize_quantumpay_node_v1_HealthRequest,
    responseSerialize: serialize_quantumpay_node_v1_HealthResponse,
    responseDeserialize: deserialize_quantumpay_node_v1_HealthResponse,
  },
};

exports.NodeServiceClient = grpc.makeGenericClientConstructor(NodeServiceService, 'NodeService');
