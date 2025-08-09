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

function serialize_infrastructure_rpc_cqrs_link_v1_ListRequest(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_query_pb.ListRequest)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.ListRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_ListRequest(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_query_pb.ListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_cqrs_link_v1_ListResponse(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_query_pb.ListResponse)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.ListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_ListResponse(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_query_pb.ListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// LinkQueryService is the service that provides the query methods for the Link aggregate.
var LinkQueryServiceService = exports.LinkQueryServiceService = {
  // Get returns a LinkView for the given hash.
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
  // List returns a LinksView for the given filter.
list: {
    path: '/infrastructure.rpc.cqrs.link.v1.LinkQueryService/List',
    requestStream: false,
    responseStream: false,
    requestType: infrastructure_rpc_cqrs_link_v1_link_query_pb.ListRequest,
    responseType: infrastructure_rpc_cqrs_link_v1_link_query_pb.ListResponse,
    requestSerialize: serialize_infrastructure_rpc_cqrs_link_v1_ListRequest,
    requestDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_ListRequest,
    responseSerialize: serialize_infrastructure_rpc_cqrs_link_v1_ListResponse,
    responseDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_ListResponse,
  },
};

exports.LinkQueryServiceClient = grpc.makeGenericClientConstructor(LinkQueryServiceService);
