// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var infrastructure_rpc_sitemap_sitemap_pb = require('../../../infrastructure/rpc/sitemap/sitemap_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var domain_sitemap_v1_sitemap_pb = require('../../../domain/sitemap/v1/sitemap_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_infrastructure_rpc_sitemap_v1_ParseRequest(arg) {
  if (!(arg instanceof infrastructure_rpc_sitemap_sitemap_pb.ParseRequest)) {
    throw new Error('Expected argument of type infrastructure.rpc.sitemap.v1.ParseRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_infrastructure_rpc_sitemap_v1_ParseRequest(buffer_arg) {
  return infrastructure_rpc_sitemap_sitemap_pb.ParseRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var SitemapServiceService = exports.SitemapServiceService = {
  parse: {
    path: '/infrastructure.rpc.sitemap.v1.SitemapService/Parse',
    requestStream: false,
    responseStream: false,
    requestType: infrastructure_rpc_sitemap_sitemap_pb.ParseRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_infrastructure_rpc_sitemap_v1_ParseRequest,
    requestDeserialize: deserialize_infrastructure_rpc_sitemap_v1_ParseRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.SitemapServiceClient = grpc.makeGenericClientConstructor(SitemapServiceService);
