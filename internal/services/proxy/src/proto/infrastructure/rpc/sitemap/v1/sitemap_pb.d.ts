// package: infrastructure.rpc.sitemap.v1
// file: infrastructure/rpc/sitemap/v1/sitemap.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class ParseRequest extends jspb.Message {
    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: ParseRequest): ParseRequest.AsObject;

    static serializeBinaryToWriter(message: ParseRequest, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): ParseRequest;

    static deserializeBinaryFromReader(message: ParseRequest, reader: jspb.BinaryReader): ParseRequest;

    getUrl(): string;

    setUrl(value: string): ParseRequest;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): ParseRequest.AsObject;
}

export namespace ParseRequest {
    export type AsObject = {
        url: string,
    }
}
