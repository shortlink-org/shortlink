import { Link } from "../../domain/entities/Link.js";
import { Hash } from "../../domain/entities/Hash.js";

/**
 * Интерфейс адаптера для получения ссылок из внешнего Link Service
 * Абстрагирует способ получения данных (HTTP, gRPC)
 */
export interface ILinkServiceAdapter {
  /**
   * Получает ссылку по хешу из внешнего сервиса
   *
   * @param hash - хеш ссылки
   * @returns Promise с доменной сущностью Link или null если не найдена
   */
  getLinkByHash(hash: Hash): Promise<Link | null>;
}

