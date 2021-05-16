// package: infrastructure.rpc.v1
// file: src/proto/infrastructure/rpc/v1/proxy.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as src_proto_infrastructure_rpc_v1_proxy_pb from "../../../../../src/proto/infrastructure/rpc/v1/proxy_pb";
import * as src_proto_domain_proxy_v1_proxy_pb from "../../../../../src/proto/domain/proxy/v1/proxy_pb";

interface IStatsService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    stats: IStatsService_IStats;
}

interface IStatsService_IStats extends grpc.MethodDefinition<src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse> {
    path: "/infrastructure.rpc.v1.Stats/Stats";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest>;
    requestDeserialize: grpc.deserialize<src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest>;
    responseSerialize: grpc.serialize<src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse>;
    responseDeserialize: grpc.deserialize<src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse>;
}

export const StatsService: IStatsService;

export interface IStatsServer extends grpc.UntypedServiceImplementation {
    stats: grpc.handleUnaryCall<src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse>;
}

export interface IStatsClient {
    stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
}

export class StatsClient extends grpc.Client implements IStatsClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    public stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    public stats(request: src_proto_infrastructure_rpc_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: src_proto_infrastructure_rpc_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
}
