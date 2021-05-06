// package: infrastructure.rpc.v1
// file: src/proto/infrastructure/rpc/v1/proxy.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as src_proto_infrastructure_rpc_v1_proxy_pb from "../../../../../src/proto/infrastructure/rpc/v1/proxy_pb";
import * as src_proto_domain_proxy_v1_proxy_pb from "../../../../../src/proto/domain/proxy/v1/proxy_pb";

interface IProxyService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    stats: IProxyService_IStats;
}

interface IProxyService_IStats extends grpc.MethodDefinition<src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse> {
    path: "/infrastructure.rpc.v1.Proxy/Stats";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest>;
    requestDeserialize: grpc.deserialize<src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest>;
    responseSerialize: grpc.serialize<src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse>;
    responseDeserialize: grpc.deserialize<src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse>;
}

export const ProxyService: IProxyService;

export interface IProxyServer extends grpc.UntypedServiceImplementation {
    stats: grpc.handleUnaryCall<src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse>;
}

export interface IProxyClient {
    stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
}

export class ProxyClient extends grpc.Client implements IProxyClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    public stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    public stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
}
