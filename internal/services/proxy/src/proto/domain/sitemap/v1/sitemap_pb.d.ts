// package: domain.sitemap.v1
// file: domain/sitemap/v1/sitemap.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as tagger_tagger_pb from "../../../tagger/tagger_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Url extends jspb.Message { 
    getLoc(): string;
    setLoc(value: string): Url;
    getLastMod(): string;
    setLastMod(value: string): Url;
    getChangeFreq(): string;
    setChangeFreq(value: string): Url;
    getPriority(): number;
    setPriority(value: number): Url;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Url.AsObject;
    static toObject(includeInstance: boolean, msg: Url): Url.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Url, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Url;
    static deserializeBinaryFromReader(message: Url, reader: jspb.BinaryReader): Url;
}

export namespace Url {
    export type AsObject = {
        loc: string,
        lastMod: string,
        changeFreq: string,
        priority: number,
    }
}

export class Sitemap extends jspb.Message { 
    clearUrlList(): void;
    getUrlList(): Array<Url>;
    setUrlList(value: Array<Url>): Sitemap;
    addUrl(value?: Url, index?: number): Url;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Sitemap.AsObject;
    static toObject(includeInstance: boolean, msg: Sitemap): Sitemap.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Sitemap, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Sitemap;
    static deserializeBinaryFromReader(message: Sitemap, reader: jspb.BinaryReader): Sitemap;
}

export namespace Sitemap {
    export type AsObject = {
        urlList: Array<Url.AsObject>,
    }
}
