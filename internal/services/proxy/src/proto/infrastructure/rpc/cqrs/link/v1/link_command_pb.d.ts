// package: infrastructure.rpc.cqrs.link.v1
// file: infrastructure/rpc/cqrs/link/v1/link_command.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as domain_link_v1_link_pb from "../../../../../domain/link/v1/link_pb";

export class AddRequest extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: AddRequest): AddRequest.AsObject;

    static serializeBinaryToWriter(message: AddRequest, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): AddRequest;

    static deserializeBinaryFromReader(message: AddRequest, reader: jspb.BinaryReader): AddRequest;

    hasLink(): boolean;

    clearLink(): void;

    getLink(): domain_link_v1_link_pb.Link | undefined;

    setLink(value?: domain_link_v1_link_pb.Link): AddRequest;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): AddRequest.AsObject;
}

export namespace AddRequest {
    export type AsObject = {
        link?: domain_link_v1_link_pb.Link.AsObject,
    }
}

export class AddResponse extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: AddResponse): AddResponse.AsObject;

    static serializeBinaryToWriter(message: AddResponse, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): AddResponse;

    static deserializeBinaryFromReader(message: AddResponse, reader: jspb.BinaryReader): AddResponse;

    hasLink(): boolean;

    clearLink(): void;

    getLink(): domain_link_v1_link_pb.Link | undefined;

    setLink(value?: domain_link_v1_link_pb.Link): AddResponse;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): AddResponse.AsObject;
}

export namespace AddResponse {
    export type AsObject = {
        link?: domain_link_v1_link_pb.Link.AsObject,
    }
}

export class UpdateRequest extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;

    static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): UpdateRequest;

    static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;

    hasLink(): boolean;

    clearLink(): void;

    getLink(): domain_link_v1_link_pb.Link | undefined;

    setLink(value?: domain_link_v1_link_pb.Link): UpdateRequest;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): UpdateRequest.AsObject;
}

export namespace UpdateRequest {
    export type AsObject = {
        link?: domain_link_v1_link_pb.Link.AsObject,
    }
}

export class UpdateResponse extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: UpdateResponse): UpdateResponse.AsObject;

    static serializeBinaryToWriter(message: UpdateResponse, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): UpdateResponse;

    static deserializeBinaryFromReader(message: UpdateResponse, reader: jspb.BinaryReader): UpdateResponse;

    hasLink(): boolean;

    clearLink(): void;

    getLink(): domain_link_v1_link_pb.Link | undefined;

    setLink(value?: domain_link_v1_link_pb.Link): UpdateResponse;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): UpdateResponse.AsObject;
}

export namespace UpdateResponse {
    export type AsObject = {
        link?: domain_link_v1_link_pb.Link.AsObject,
    }
}

export class DeleteRequest extends jspb.Message {
    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;

    static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): DeleteRequest;

    static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;

    getHash(): string;

    setHash(value: string): DeleteRequest;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): DeleteRequest.AsObject;
}

export namespace DeleteRequest {
    export type AsObject = {
        hash: string,
    }
}
