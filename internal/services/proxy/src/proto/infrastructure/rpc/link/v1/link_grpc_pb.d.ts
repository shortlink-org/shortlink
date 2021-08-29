// package: infrastructure.rpc.link.v1
// file: infrastructure/rpc/link/v1/link.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as infrastructure_rpc_link_v1_link_pb from "../../../../infrastructure/rpc/link/v1/link_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as domain_link_v1_link_pb from "../../../../domain/link/v1/link_pb";

interface ILinkServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    get: ILinkServiceService_IGet;
    list: ILinkServiceService_IList;
    add: ILinkServiceService_IAdd;
    update: ILinkServiceService_IUpdate;
    delete: ILinkServiceService_IDelete;
}

interface ILinkServiceService_IGet extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_pb.GetRequest, infrastructure_rpc_link_v1_link_pb.GetResponse> {
    path: "/infrastructure.rpc.link.v1.LinkService/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.GetRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.GetRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.GetResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.GetResponse>;
}
interface ILinkServiceService_IList extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_pb.ListRequest, infrastructure_rpc_link_v1_link_pb.ListResponse> {
    path: "/infrastructure.rpc.link.v1.LinkService/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.ListRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.ListRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.ListResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.ListResponse>;
}
interface ILinkServiceService_IAdd extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_pb.AddRequest, infrastructure_rpc_link_v1_link_pb.AddResponse> {
    path: "/infrastructure.rpc.link.v1.LinkService/Add";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.AddRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.AddRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.AddResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.AddResponse>;
}
interface ILinkServiceService_IUpdate extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_pb.UpdateRequest, infrastructure_rpc_link_v1_link_pb.UpdateResponse> {
    path: "/infrastructure.rpc.link.v1.LinkService/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.UpdateRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.UpdateRequest>;
    responseSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.UpdateResponse>;
    responseDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.UpdateResponse>;
}
interface ILinkServiceService_IDelete extends grpc.MethodDefinition<infrastructure_rpc_link_v1_link_pb.DeleteRequest, google_protobuf_empty_pb.Empty> {
    path: "/infrastructure.rpc.link.v1.LinkService/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<infrastructure_rpc_link_v1_link_pb.DeleteRequest>;
    requestDeserialize: grpc.deserialize<infrastructure_rpc_link_v1_link_pb.DeleteRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}

export const LinkServiceService: ILinkServiceService;

export interface ILinkServiceServer extends grpc.UntypedServiceImplementation {
    get: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_pb.GetRequest, infrastructure_rpc_link_v1_link_pb.GetResponse>;
    list: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_pb.ListRequest, infrastructure_rpc_link_v1_link_pb.ListResponse>;
    add: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_pb.AddRequest, infrastructure_rpc_link_v1_link_pb.AddResponse>;
    update: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_pb.UpdateRequest, infrastructure_rpc_link_v1_link_pb.UpdateResponse>;
    delete: grpc.handleUnaryCall<infrastructure_rpc_link_v1_link_pb.DeleteRequest, google_protobuf_empty_pb.Empty>;
}

export interface ILinkServiceClient {
    get(request: infrastructure_rpc_link_v1_link_pb.GetRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.GetResponse) => void): grpc.ClientUnaryCall;
    get(request: infrastructure_rpc_link_v1_link_pb.GetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.GetResponse) => void): grpc.ClientUnaryCall;
    get(request: infrastructure_rpc_link_v1_link_pb.GetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.GetResponse) => void): grpc.ClientUnaryCall;
    list(request: infrastructure_rpc_link_v1_link_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.ListResponse) => void): grpc.ClientUnaryCall;
    list(request: infrastructure_rpc_link_v1_link_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.ListResponse) => void): grpc.ClientUnaryCall;
    list(request: infrastructure_rpc_link_v1_link_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.ListResponse) => void): grpc.ClientUnaryCall;
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

export class LinkServiceClient extends grpc.Client implements ILinkServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public get(request: infrastructure_rpc_link_v1_link_pb.GetRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public get(request: infrastructure_rpc_link_v1_link_pb.GetRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public get(request: infrastructure_rpc_link_v1_link_pb.GetRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.GetResponse) => void): grpc.ClientUnaryCall;
    public list(request: infrastructure_rpc_link_v1_link_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public list(request: infrastructure_rpc_link_v1_link_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public list(request: infrastructure_rpc_link_v1_link_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: infrastructure_rpc_link_v1_link_pb.ListResponse) => void): grpc.ClientUnaryCall;
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
