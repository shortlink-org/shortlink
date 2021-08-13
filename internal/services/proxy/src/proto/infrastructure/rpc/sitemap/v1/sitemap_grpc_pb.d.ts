// package: infrastructure.rpc.sitemap.v1
// file: infrastructure/rpc/sitemap/v1/sitemap.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as infrastructure_rpc_sitemap_v1_sitemap_pb from "../../../../infrastructure/rpc/sitemap/v1/sitemap_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

interface ISitemapServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    parse: ISitemapServiceService_IParse;
}

interface ISitemapServiceService_IParse extends grpc.MethodDefinition<infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest, google_protobuf_empty_pb.Empty> {
    path: "/infrastructure.rpc.sitemap.v1.SitemapService/Parse";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}

export const SitemapServiceService: ISitemapServiceService;

export interface ISitemapServiceServer extends grpc.UntypedServiceImplementation {
    parse: grpc.handleUnaryCall<infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest, google_protobuf_empty_pb.Empty>;
}

export interface ISitemapServiceClient {
    parse(request: infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    parse(request: infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    parse(request: infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}

export class SitemapServiceClient extends grpc.Client implements ISitemapServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public parse(request: infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public parse(request: infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public parse(request: infrastructure_rpc_sitemap_v1_sitemap_pb.ParseRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}
