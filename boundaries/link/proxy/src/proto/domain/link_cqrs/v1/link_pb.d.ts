// package: domain.link_cqrs.v1
// file: domain/link_cqrs/v1/link.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_field_mask_pb from "google-protobuf/google/protobuf/field_mask_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class LinkView extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: LinkView): LinkView.AsObject;

    static serializeBinaryToWriter(message: LinkView, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): LinkView;

    static deserializeBinaryFromReader(message: LinkView, reader: jspb.BinaryReader): LinkView;

    hasFieldMask(): boolean;

    clearFieldMask(): void;

    getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;

    setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): LinkView;

    getUrl(): string;

    setUrl(value: string): LinkView;

    getHash(): string;

    setHash(value: string): LinkView;

    getDescribe(): string;

    setDescribe(value: string): LinkView;

    getImageUrl(): string;

    setImageUrl(value: string): LinkView;

    getMetaDescription(): string;

    setMetaDescription(value: string): LinkView;

    getMetaKeywords(): string;

    setMetaKeywords(value: string): LinkView;

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
}

export namespace LinkView {
    export type AsObject = {
        fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
        url: string,
        hash: string,
        describe: string,
        imageUrl: string,
        metaDescription: string,
        metaKeywords: string,
        createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
}

export class LinksView extends jspb.Message {
    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: LinksView): LinksView.AsObject;

    static serializeBinaryToWriter(message: LinksView, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): LinksView;

    static deserializeBinaryFromReader(message: LinksView, reader: jspb.BinaryReader): LinksView;

    clearLinksList(): void;

    getLinksList(): Array<LinkView>;

    setLinksList(value: Array<LinkView>): LinksView;

    addLinks(value?: LinkView, index?: number): LinkView;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): LinksView.AsObject;
}

export namespace LinksView {
    export type AsObject = {
        linksList: Array<LinkView.AsObject>,
    }
}
