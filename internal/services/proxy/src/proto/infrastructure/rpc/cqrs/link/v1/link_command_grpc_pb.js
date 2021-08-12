// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var infrastructure_rpc_cqrs_link_v1_link_command_pb = require('../../../../../infrastructure/rpc/cqrs/link/v1/link_command_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var domain_link_v1_link_pb = require('../../../../../domain/link/v1/link_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_cqrs_link_v1_AddRequest(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_command_pb.AddRequest)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.AddRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_AddRequest(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_command_pb.AddRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_cqrs_link_v1_AddResponse(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_command_pb.AddResponse)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.AddResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_AddResponse(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_command_pb.AddResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_cqrs_link_v1_DeleteRequest(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_command_pb.DeleteRequest)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.DeleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_DeleteRequest(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_command_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_cqrs_link_v1_UpdateRequest(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_command_pb.UpdateRequest)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.UpdateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_UpdateRequest(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_command_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_cqrs_link_v1_UpdateResponse(arg) {
  if (!(arg instanceof infrastructure_rpc_cqrs_link_v1_link_command_pb.UpdateResponse)) {
    throw new Error('Expected argument of type infrastructure.rpc.cqrs.link.v1.UpdateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_cqrs_link_v1_UpdateResponse(buffer_arg) {
  return infrastructure_rpc_cqrs_link_v1_link_command_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var LinkCommandServiceService = exports.LinkCommandServiceService = {
  add: {
    path: '/infrastructure.rpc.cqrs.link.v1.LinkCommandService/Add',
    requestStream: false,
    responseStream: false,
    requestType: infrastructure_rpc_cqrs_link_v1_link_command_pb.AddRequest,
    responseType: infrastructure_rpc_cqrs_link_v1_link_command_pb.AddResponse,
    requestSerialize: serialize_infrastructure_rpc_cqrs_link_v1_AddRequest,
    requestDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_AddRequest,
    responseSerialize: serialize_infrastructure_rpc_cqrs_link_v1_AddResponse,
    responseDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_AddResponse,
  },
  update: {
    path: '/infrastructure.rpc.cqrs.link.v1.LinkCommandService/Update',
    requestStream: false,
    responseStream: false,
    requestType: infrastructure_rpc_cqrs_link_v1_link_command_pb.UpdateRequest,
    responseType: infrastructure_rpc_cqrs_link_v1_link_command_pb.UpdateResponse,
    requestSerialize: serialize_infrastructure_rpc_cqrs_link_v1_UpdateRequest,
    requestDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_UpdateRequest,
    responseSerialize: serialize_infrastructure_rpc_cqrs_link_v1_UpdateResponse,
    responseDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_UpdateResponse,
  },
  delete: {
    path: '/infrastructure.rpc.cqrs.link.v1.LinkCommandService/Delete',
    requestStream: false,
    responseStream: false,
    requestType: infrastructure_rpc_cqrs_link_v1_link_command_pb.DeleteRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_infrastructure_rpc_cqrs_link_v1_DeleteRequest,
    requestDeserialize: deserialize_infrastructure_rpc_cqrs_link_v1_DeleteRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.LinkCommandServiceClient = grpc.makeGenericClientConstructor(LinkCommandServiceService);
