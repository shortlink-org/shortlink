// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var infrastructure_rpc_cqrs_link_v1_link_query_pb = require('../../../../../infrastructure/rpc/cqrs/link/v1/link_query_pb.js');
var domain_link_cqrs_v1_link_pb = require('../../../../../domain/link_cqrs/v1/link_pb.js');

function serialize_infrastructure_rpc_cqrs_link_v1_GetRequest(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.GetRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_GetRequest(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_cqrs_link_v1_GetResponse(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.GetResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_GetResponse(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var LinkQueryServiceService = exports.LinkQueryServiceService = {
  get: {
    path: '/infrastructure.rpc.cqrs.link.v1.LinkQueryService/Get',
    requestStream: false,
    responseStream: false,
    requestType: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest,
    responseType: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse,
    requestSerialize: serialize_infrastructure_rpc_cqrs_link_v1_GetRequest,
    requestDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_GetRequest,
    responseSerialize: serialize_infrastructure_rpc_cqrs_link_v1_GetResponse,
    responseDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_GetResponse,
  },
};

exports.LinkQueryServiceClient = grpc.makeGenericClientConstructor(LinkQueryServiceService);
