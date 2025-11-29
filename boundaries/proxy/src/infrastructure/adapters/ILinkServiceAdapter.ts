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
   * @param userId - optional user_id from Kratos session, or "anonymous" if not authenticated
   * @returns Promise с доменной сущностью Link
   * @throws LinkNotFoundError если ссылка не найдена или доступ запрещен
   */
  getLinkByHash(hash: Hash, userId?: string | null): Promise<Link>;
}
