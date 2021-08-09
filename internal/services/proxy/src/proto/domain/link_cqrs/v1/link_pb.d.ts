// package: domain.link_cqrs.v1
// file: domain/link_cqrs/v1/link.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class LinkView extends jspb.Message { 
    getUrl(): string;
    setUrl(value: string): LinkView;
    getHash(): string;
    setHash(value: string): LinkView;
    getDescribe(): string;
    setDescribe(value: string): LinkView;
    getImageurl(): string;
    setImageurl(value: string): LinkView;
    getMetadescription(): string;
    setMetadescription(value: string): LinkView;
    getMetakeywords(): string;
    setMetakeywords(value: string): LinkView;

    hasCreatedAt(): boolean;
    clearCreatedAt(): void;
    getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): LinkView;

    hasUpdatedAt(): boolean;
    clearUpdatedAt(): void;
    getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): LinkView;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LinkView.AsObject;
    static toObject(includeInstance: boolean, msg: LinkView): LinkView.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LinkView, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LinkView;
    static deserializeBinaryFromReader(message: LinkView, reader: jspb.BinaryReader): LinkView;
}

export namespace LinkView {
    export type AsObject = {
        url: string,
        hash: string,
        describe: string,
        imageurl: string,
        metadescription: string,
        metakeywords: string,
        createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
}

export class LinksView extends jspb.Message { 
    clearLinkList(): void;
    getLinkList(): Array<LinkView>;
    setLinkList(value: Array<LinkView>): LinksView;
    addLink(value?: LinkView, index?: number): LinkView;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LinksView.AsObject;
    static toObject(includeInstance: boolean, msg: LinksView): LinksView.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LinksView, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LinksView;
    static deserializeBinaryFromReader(message: LinksView, reader: jspb.BinaryReader): LinksView;
}

export namespace LinksView {
    export type AsObject = {
        linkList: Array<LinkView.AsObject>,
    }
}
