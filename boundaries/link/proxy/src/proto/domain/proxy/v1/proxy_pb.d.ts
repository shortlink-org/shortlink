// package: domain.proxy.v1
// file: domain/proxy/v1/proxy.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_field_mask_pb from "google-protobuf/google/protobuf/field_mask_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Stats extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: Stats): Stats.AsObject;

    static serializeBinaryToWriter(message: Stats, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): Stats;

    static deserializeBinaryFromReader(message: Stats, reader: jspb.BinaryReader): Stats;

    hasFieldMask(): boolean;

    clearFieldMask(): void;

    getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;

    setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): Stats;

    getHash(): string;

    setHash(value: string): Stats;

    getCountRedirect(): number;

    setCountRedirect(value: number): Stats;

    hasUpdatedAt(): boolean;

    clearUpdatedAt(): void;

    getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;

    setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Stats;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): Stats.AsObject;
}

export namespace Stats {
    export type AsObject = {
        fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
        hash: string,
        countRedirect: number,
        updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
}
