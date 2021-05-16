// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var src_proto_infrastructure_rpc_v1_proxy_pb = require('../../../../../src/proto/infrastructure/rpc/v1/proxy_pb.js');
var src_proto_domain_proxy_v1_proxy_pb = require('../../../../../src/proto/domain/proxy/v1/proxy_pb.js');

function serialize_infrastructure_rpc_v1_StatsRequest(arg) {
  if (!(arg instanceof src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest)) {
    throw new Error('Expected argument of type infrastructure.rpc.v1.StatsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_v1_StatsRequest(buffer_arg) {
  return src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_v1_StatsResponse(arg) {
  if (!(arg instanceof src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse)) {
    throw new Error('Expected argument of type infrastructure.rpc.v1.StatsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_v1_StatsResponse(buffer_arg) {
  return src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var StatsService = exports.StatsService = {
  stats: {
    path: '/infrastructure.rpc.v1.Stats/Stats',
    requestStream: false,
    responseStream: false,
    requestType: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest,
    responseType: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse,
    requestSerialize: serialize_infrastructure_rpc_v1_StatsRequest,
    requestDeserialize: deserialize_infrastructure_rpc_v1_StatsRequest,
    responseSerialize: serialize_infrastructure_rpc_v1_StatsResponse,
    responseDeserialize: deserialize_infrastructure_rpc_v1_StatsResponse,
  },
};

exports.StatsClient = grpc.makeGenericClientConstructor(StatsService);
