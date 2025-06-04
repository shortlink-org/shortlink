// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var infrastructure_rpc_proxy_v1_proxy_pb = require('../../../../infrastructure/rpc/proxy/v1/proxy_pb.js');
var domain_proxy_v1_proxy_pb = require('../../../../domain/proxy/v1/proxy_pb.js');

function serialize_infrastructure_rpc_proxy_v1_StatsRequest(arg) {
  if (!(arg instanceof infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest)) {
    throw new Error('Expected argument of type infrastructure.rpc.proxy.v1.StatsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_proxy_v1_StatsRequest(buffer_arg) {
  return infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_proxy_v1_StatsResponse(arg) {
  if (!(arg instanceof infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse)) {
    throw new Error('Expected argument of type infrastructure.rpc.proxy.v1.StatsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_proxy_v1_StatsResponse(buffer_arg) {
  return infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// ProxyService is the service that provides proxy functionality.
var StatsServiceService = exports.StatsServiceService = {
  // Stats returns the stats for the given proxy.
stats: {
    path: '/infrastructure.rpc.proxy.v1.StatsService/Stats',
    requestStream: false,
    responseStream: false,
    requestType: infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest,
    responseType: infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse,
    requestSerialize: serialize_infrastructure_rpc_proxy_v1_StatsRequest,
    requestDeserialize: deserialize_infrastructure_rpc_proxy_v1_StatsRequest,
    responseSerialize: serialize_infrastructure_rpc_proxy_v1_StatsResponse,
    responseDeserialize: deserialize_infrastructure_rpc_proxy_v1_StatsResponse,
  },
};

exports.StatsServiceClient = grpc.makeGenericClientConstructor(StatsServiceService);
