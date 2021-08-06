// package: infrastructure.rpc.link.v1
// file: infrastructure/rpc/link/v1/link_query.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as infrastructure_rpc_link_v1_link_query_pb from "../../../../infrastructure/rpc/link/v1/link_query_pb";
import * as domain_link_v1_link_pb from "../../../../domain/link/v1/link_pb";

interface ILinkQueryServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    get: ILinkQueryServiceService_IGet;
    list: ILinkQueryServiceService_IList;
}

interface ILinkQueryServiceService_IGet extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_query_pb.GetRequest, infrastructure_rpc_link_v1_link_query_pb.GetResponse> {
    path: "/infrastructure.rpc.link.v1.LinkQueryService/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_query_pb.GetRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_query_pb.GetRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_query_pb.GetResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_query_pb.GetResponse>;
}
interface ILinkQueryServiceService_IList extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_query_pb.ListRequest, infrastructure_rpc_link_v1_link_query_pb.ListResponse> {
    path: "/infrastructure.rpc.link.v1.LinkQueryService/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_query_pb.ListRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_query_pb.ListRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_query_pb.ListResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_query_pb.ListResponse>;
}

export const LinkQueryServiceService: ILinkQueryServiceService;

export interface ILinkQueryServiceServer extends grpc.UntypedServiceImplementation {
    get: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_query_pb.GetRequest, infrastructure_rpc_link_v1_link_query_pb.GetResponse>;
    list: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_query_pb.ListRequest, infrastructure_rpc_link_v1_link_query_pb.ListResponse>;
}

export interface ILinkQueryServiceClient {
    get(request: infrastructure_rpc_link_v1_link_query_pb.GetRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    get(request: infrastructure_rpc_link_v1_link_query_pb.GetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    get(request: infrastructure_rpc_link_v1_link_query_pb.GetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    list(request: infrastructure_rpc_link_v1_link_query_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.ListResponse) => void): grpc.ClientUnaryCall;
    list(request: infrastructure_rpc_link_v1_link_query_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.ListResponse) => void): grpc.ClientUnaryCall;
    list(request: infrastructure_rpc_link_v1_link_query_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.ListResponse) => void): grpc.ClientUnaryCall;
}

export class LinkQueryServiceClient extends grpc.Client implements ILinkQueryServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public get(request: infrastructure_rpc_link_v1_link_query_pb.GetRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public get(request: infrastructure_rpc_link_v1_link_query_pb.GetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public get(request: infrastructure_rpc_link_v1_link_query_pb.GetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public list(request: infrastructure_rpc_link_v1_link_query_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public list(request: infrastructure_rpc_link_v1_link_query_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public list(request: infrastructure_rpc_link_v1_link_query_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_query_pb.ListResponse) => void): grpc.ClientUnaryCall;
}
