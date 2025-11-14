import { Link } from "../../domain/entities/Link.js";
import { Hash } from "../../domain/entities/Hash.js";
import { LinkMapper } from "../../domain/mappers/LinkMapper.js";
import { Link as LinkProto } from "../../../proto/infrastructure/rpc/link/v1/link_pb.js";

/**
 * Anti-corruption Layer для Link Service
 * Защищает домен от изменений внешнего API
 * Преобразует внешние модели (protobuf, HTTP JSON) в domain entities
 */
export class LinkServiceACL {
  /**
   * Преобразует protobuf Link в доменную сущность
   *
   * @param protoLink - protobuf Link из внешнего сервиса
   * @returns доменная сущность Link
   */
  toDomainEntityFromProto(protoLink: LinkProto): Link {
    return LinkMapper.toDomain(protoLink);
  }

  /**
   * Преобразует HTTP JSON ответ в доменную сущность
   * Защищает домен от изменений в HTTP API
   *
   * @param httpResponse - HTTP JSON ответ от внешнего сервиса
   * @returns доменная сущность Link
   */
  toDomainEntityFromHttp(httpResponse: {
    hash: string;
    url: string;
    createdAt?: string | Date;
    updatedAt?: string | Date;
  }): Link {
    const hash = new Hash(httpResponse.hash);
    const url = httpResponse.url;

    // Преобразуем строки в Date, если нужно
    const createdAt = httpResponse.createdAt
      ? typeof httpResponse.createdAt === "string"
        ? new Date(httpResponse.createdAt)
        : httpResponse.createdAt
      : new Date();

    const updatedAt = httpResponse.updatedAt
      ? typeof httpResponse.updatedAt === "string"
        ? new Date(httpResponse.updatedAt)
        : httpResponse.updatedAt
      : new Date();

    return new Link(hash, url, createdAt, updatedAt);
  }

  /**
   * Преобразует доменную сущность в protobuf (если нужно для запросов)
   *
   * @param domainLink - доменная сущность Link
   * @returns protobuf Link
   */
  toProto(domainLink: Link): LinkProto {
    return LinkMapper.toProto(domainLink);
  }
}

