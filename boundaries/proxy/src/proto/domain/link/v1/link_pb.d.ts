// package: domain.link.v1
// file: domain/link/v1/link.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_field_mask_pb from "google-protobuf/google/protobuf/field_mask_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Link extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: Link): Link.AsObject;

    static serializeBinaryToWriter(message: Link, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): Link;

    static deserializeBinaryFromReader(message: Link, reader: jspb.BinaryReader): Link;

    hasFieldMask(): boolean;

    clearFieldMask(): void;

    getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;

    setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): Link;

    getUrl(): string;

    setUrl(value: string): Link;

    getHash(): string;

    setHash(value: string): Link;

    getDescribe(): string;

    setDescribe(value: string): Link;

    hasCreatedAt(): boolean;

    clearCreatedAt(): void;

    getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;

    setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Link;

    hasUpdatedAt(): boolean;

    clearUpdatedAt(): void;

    getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;

    setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Link;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): Link.AsObject;
}

export namespace Link {
    export type AsObject = {
        fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
        url: string,
        hash: string,
        describe: string,
        createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
}

export class Links extends jspb.Message {
    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: Links): Links.AsObject;

    static serializeBinaryToWriter(message: Links, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): Links;

    static deserializeBinaryFromReader(message: Links, reader: jspb.BinaryReader): Links;

    clearLinkList(): void;

    getLinkList(): Array<Link>;

    setLinkList(value: Array<Link>): Links;

    addLink(value?: Link, index?: number): Link;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): Links.AsObject;
}

export namespace Links {
    export type AsObject = {
        linkList: Array<Link.AsObject>,
    }
}

export enum LinkEvent {
    LINK_EVENT_UNSPECIFIED = 0,
    LINK_EVENT_ADD = 1,
    LINK_EVENT_GET = 2,
    LINK_EVENT_LIST = 3,
    LINK_EVENT_UPDATE = 4,
    LINK_EVENT_DELETE = 5,
}
