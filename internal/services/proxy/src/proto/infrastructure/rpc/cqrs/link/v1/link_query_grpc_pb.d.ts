// package: infrastructure.rpc.cqrs.link.v1
// file: infrastructure/rpc/cqrs/link/v1/link_query.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as infrastructure_rpc_cqrs_link_v1_link_query_pb from "../../../../../infrastructure/rpc/cqrs/link/v1/link_query_pb";
import * as domain_link_cqrs_v1_link_pb from "../../../../../domain/link_cqrs/v1/link_pb";

interface ILinkQueryServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    get: ILinkQueryServiceService_IGet;
}

interface ILinkQueryServiceService_IGet extends grpc.MethodDefinition<infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest, infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse> {
    path: "/infrastructure.rpc.cqrs.link.v1.LinkQueryService/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse>;
}

export const LinkQueryServiceService: ILinkQueryServiceService;

export interface ILinkQueryServiceServer extends grpc.UntypedServiceImplementation {
    get: grpc.handleUnaryCall<infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest, infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse>;
}

export interface ILinkQueryServiceClient {
    get(request: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    get(request: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    get(request: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
}

export class LinkQueryServiceClient extends grpc.Client implements ILinkQueryServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public get(request: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public get(request: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public get(request: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_cqrs_link_v1_link_query_pb.GetResponse) => void): grpc.ClientUnaryCall;
}
