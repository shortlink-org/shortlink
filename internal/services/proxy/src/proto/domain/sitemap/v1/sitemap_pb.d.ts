// package: domain.sitemap.v1
// file: domain/sitemap/v1/sitemap.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as tagger_tagger_pb from "../../../tagger/tagger_pb";

export class Sitemap extends jspb.Message { 
    getLocation(): string;
    setLocation(value: string): Sitemap;
    getLastModified(): string;
    setLastModified(value: string): Sitemap;

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
        location: string,
        lastModified: string,
    }
}
