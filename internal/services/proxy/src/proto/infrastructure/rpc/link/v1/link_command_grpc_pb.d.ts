// package: infrastructure.rpc.link.v1
// file: infrastructure/rpc/link/v1/link_command.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as infrastructure_rpc_link_v1_link_command_pb from "../../../../infrastructure/rpc/link/v1/link_command_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as domain_link_v1_link_pb from "../../../../domain/link/v1/link_pb";
import * as infrastructure_rpc_link_v1_link_pb from "../../../../infrastructure/rpc/link/v1/link_pb";

interface ILinkCommandServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    add: ILinkCommandServiceService_IAdd;
    update: ILinkCommandServiceService_IUpdate;
    delete: ILinkCommandServiceService_IDelete;
}

interface ILinkCommandServiceService_IAdd extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_pb.AddRequest, infrastructure_rpc_link_v1_link_pb.AddResponse> {
    path: "/infrastructure.rpc.link.v1.LinkCommandService/Add";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.AddRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.AddRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.AddResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.AddResponse>;
}
interface ILinkCommandServiceService_IUpdate extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_pb.UpdateRequest, infrastructure_rpc_link_v1_link_pb.UpdateResponse> {
    path: "/infrastructure.rpc.link.v1.LinkCommandService/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.UpdateRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.UpdateRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.UpdateResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.UpdateResponse>;
}
interface ILinkCommandServiceService_IDelete extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_pb.DeleteRequest, google_protobuf_empty_pb.Empty> {
    path: "/infrastructure.rpc.link.v1.LinkCommandService/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.DeleteRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.DeleteRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}

export const LinkCommandServiceService: ILinkCommandServiceService;

export interface ILinkCommandServiceServer extends grpc.UntypedServiceImplementation {
    add: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_pb.AddRequest, infrastructure_rpc_link_v1_link_pb.AddResponse>;
    update: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_pb.UpdateRequest, infrastructure_rpc_link_v1_link_pb.UpdateResponse>;
    delete: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_pb.DeleteRequest, google_protobuf_empty_pb.Empty>;
}

export interface ILinkCommandServiceClient {
    add(request: infrastructure_rpc_link_v1_link_pb.AddRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.AddResponse) => void): grpc.ClientUnaryCall;
    add(request: infrastructure_rpc_link_v1_link_pb.AddRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.AddResponse) => void): grpc.ClientUnaryCall;
    add(request: infrastructure_rpc_link_v1_link_pb.AddRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.AddResponse) => void): grpc.ClientUnaryCall;
    update(request: infrastructure_rpc_link_v1_link_pb.UpdateRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    update(request: infrastructure_rpc_link_v1_link_pb.UpdateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    update(request: infrastructure_rpc_link_v1_link_pb.UpdateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    delete(request: infrastructure_rpc_link_v1_link_pb.DeleteRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    delete(request: infrastructure_rpc_link_v1_link_pb.DeleteRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    delete(request: infrastructure_rpc_link_v1_link_pb.DeleteRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}

export class LinkCommandServiceClient extends grpc.Client implements ILinkCommandServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public add(request: infrastructure_rpc_link_v1_link_pb.AddRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.AddResponse) => void): grpc.ClientUnaryCall;
    public add(request: infrastructure_rpc_link_v1_link_pb.AddRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.AddResponse) => void): grpc.ClientUnaryCall;
    public add(request: infrastructure_rpc_link_v1_link_pb.AddRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.AddResponse) => void): grpc.ClientUnaryCall;
    public update(request: infrastructure_rpc_link_v1_link_pb.UpdateRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public update(request: infrastructure_rpc_link_v1_link_pb.UpdateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public update(request: infrastructure_rpc_link_v1_link_pb.UpdateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public delete(request: infrastructure_rpc_link_v1_link_pb.DeleteRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public delete(request: infrastructure_rpc_link_v1_link_pb.DeleteRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public delete(request: infrastructure_rpc_link_v1_link_pb.DeleteRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}
