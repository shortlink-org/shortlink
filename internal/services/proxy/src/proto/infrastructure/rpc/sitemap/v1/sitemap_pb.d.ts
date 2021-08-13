// package: infrastructure.rpc.sitemap.v1
// file: infrastructure/rpc/sitemap/v1/sitemap.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class ParseRequest extends jspb.Message { 
    getUrl(): string;
    setUrl(value: string): ParseRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ParseRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ParseRequest): ParseRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ParseRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ParseRequest;
    static deserializeBinaryFromReader(message: ParseRequest, reader: jspb.BinaryReader): ParseRequest;
}

export namespace ParseRequest {
    export type AsObject = {
        url: string,
    }
}
