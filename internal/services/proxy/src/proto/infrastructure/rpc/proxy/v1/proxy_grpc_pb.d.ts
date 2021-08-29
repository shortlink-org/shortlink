// package: infrastructure.rpc.proxy.v1
// file: infrastructure/rpc/proxy/v1/proxy.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as infrastructure_rpc_proxy_v1_proxy_pb from "../../../../infrastructure/rpc/proxy/v1/proxy_pb";
import * as domain_proxy_v1_proxy_pb from "../../../../domain/proxy/v1/proxy_pb";

interface IStatsServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    stats: IStatsServiceService_IStats;
}

interface IStatsServiceService_IStats extends grpc.MethodDefinition<infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest, infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse> {
    path: "/infrastructure.rpc.proxy.v1.StatsService/Stats";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse>;
}

export const StatsServiceService: IStatsServiceService;

export interface IStatsServiceServer extends grpc.UntypedServiceImplementation {
    stats: grpc.handleUnaryCall<infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest, infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse>;
}

export interface IStatsServiceClient {
    stats(request: infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    stats(request: infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    stats(request: infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
}

export class StatsServiceClient extends grpc.Client implements IStatsServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public stats(request: infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    public stats(request: infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
    public stats(request: infrastructure_rpc_proxy_v1_proxy_pb.StatsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_proxy_v1_proxy_pb.StatsResponse) => void): grpc.ClientUnaryCall;
}
