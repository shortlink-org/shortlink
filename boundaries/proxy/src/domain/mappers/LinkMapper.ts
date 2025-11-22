import { Link } from "../entities/Link.js";
import { Hash } from "../entities/Hash.js";
import { Link as LinkProto, LinkSchema } from "../../infrastructure/proto/infrastructure/rpc/link/v1/link_pb.js";
import { create } from "@bufbuild/protobuf";
import { Timestamp, TimestampSchema } from "@bufbuild/protobuf/wkt";

/**
 * Маппер для преобразования между доменной сущностью Link и protobuf Link
 * Изолирует преобразование между слоями
 */
export class LinkMapper {
  /**
   * Преобразует protobuf Link в доменную сущность
   */
  static toDomain(protoLink: LinkProto): Link {
    const hash = new Hash(protoLink.hash);
    const url = protoLink.url;

    // Преобразуем protobuf Timestamp в Date
    const createdAt = protoLink.createdAt
      ? new Date(Number(protoLink.createdAt.seconds) * 1000 + Number(protoLink.createdAt.nanos) / 1000000)
      : new Date();
    const updatedAt = protoLink.updatedAt
      ? new Date(Number(protoLink.updatedAt.seconds) * 1000 + Number(protoLink.updatedAt.nanos) / 1000000)
      : new Date();

    return new Link(hash, url, createdAt, updatedAt);
  }

  /**
   * Преобразует доменную сущность Link в protobuf Link
   */
  static toProto(domainLink: Link): LinkProto {
    const protoLink = create(LinkSchema, {
      hash: domainLink.hash.value,
      url: domainLink.url,
    });

    // Преобразуем Date в protobuf Timestamp
    if (domainLink.createdAt) {
      protoLink.createdAt = this.createTimestamp(domainLink.createdAt);
    }
    if (domainLink.updatedAt) {
      protoLink.updatedAt = this.createTimestamp(domainLink.updatedAt);
    }

    return protoLink;
  }

  /**
   * Преобразует массив protobuf Links в массив доменных сущностей
   */
  static toDomainArray(protoLinks: LinkProto[]): Link[] {
    return protoLinks.map((protoLink) => LinkMapper.toDomain(protoLink));
  }

  /**
   * Преобразует массив доменных сущностей Links в массив protobuf Links
   */
  static toProtoArray(domainLinks: Link[]): LinkProto[] {
    return domainLinks.map((domainLink) => LinkMapper.toProto(domainLink));
  }

  /**
   * Создает protobuf Timestamp из Date
   */
  private static createTimestamp(date: Date): LinkProto["createdAt"] {
    const seconds = BigInt(Math.floor(date.getTime() / 1000));
    const nanos = (date.getTime() % 1000) * 1000000;
    return create(TimestampSchema, { seconds, nanos });
  }
}

