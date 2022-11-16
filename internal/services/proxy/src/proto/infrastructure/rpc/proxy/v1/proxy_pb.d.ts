// package: infrastructure.rpc.proxy.v1
// file: infrastructure/rpc/proxy/v1/proxy.proto

import * as jspb from "google-protobuf";
import * as domain_proxy_v1_proxy_pb from "../../../../domain/proxy/v1/proxy_pb";

export class StatsRequest extends jspb.Message {
  getHash(): string;
  setHash(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StatsRequest): StatsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StatsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatsRequest;
  static deserializeBinaryFromReader(message: StatsRequest, reader: jspb.BinaryReader): StatsRequest;
}

export namespace StatsRequest {
  export type AsObject = {
    hash: string,
  }
}

export class StatsResponse extends jspb.Message {
  hasStats(): boolean;
  clearStats(): void;
  getStats(): domain_proxy_v1_proxy_pb.Stats | undefined;
  setStats(value?: domain_proxy_v1_proxy_pb.Stats): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StatsResponse): StatsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StatsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatsResponse;
  static deserializeBinaryFromReader(message: StatsResponse, reader: jspb.BinaryReader): StatsResponse;
}

export namespace StatsResponse {
  export type AsObject = {
    stats?: domain_proxy_v1_proxy_pb.Stats.AsObject,
  }
}

