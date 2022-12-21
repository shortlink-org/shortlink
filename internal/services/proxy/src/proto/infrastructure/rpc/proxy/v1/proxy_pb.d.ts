// package: infrastructure.rpc.proxy.v1
// file: infrastructure/rpc/proxy/v1/proxy.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as domain_proxy_v1_proxy_pb from "../../../../domain/proxy/v1/proxy_pb";

export class StatsRequest extends jspb.Message {
    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: StatsRequest): StatsRequest.AsObject;

    static serializeBinaryToWriter(message: StatsRequest, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): StatsRequest;

    static deserializeBinaryFromReader(message: StatsRequest, reader: jspb.BinaryReader): StatsRequest;

    getHash(): string;

    setHash(value: string): StatsRequest;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): StatsRequest.AsObject;
}

export namespace StatsRequest {
    export type AsObject = {
        hash: string,
    }
}

export class StatsResponse extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: StatsResponse): StatsResponse.AsObject;

    static serializeBinaryToWriter(message: StatsResponse, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): StatsResponse;

    static deserializeBinaryFromReader(message: StatsResponse, reader: jspb.BinaryReader): StatsResponse;

    hasStats(): boolean;

    clearStats(): void;

    getStats(): domain_proxy_v1_proxy_pb.Stats | undefined;

    setStats(value?: domain_proxy_v1_proxy_pb.Stats): StatsResponse;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): StatsResponse.AsObject;
}

export namespace StatsResponse {
    export type AsObject = {
        stats?: domain_proxy_v1_proxy_pb.Stats.AsObject,
    }
}
