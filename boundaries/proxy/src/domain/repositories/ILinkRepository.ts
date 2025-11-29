import { Link } from "../entities/Link.js";
import { Hash } from "../entities/Hash.js";

/**
 * Интерфейс репозитория для работы со ссылками
 * Работает с domain entities, не с protobuf или Prisma моделями
 */
export interface ILinkRepository {
  /**
   * Находит ссылку по хешу
   * @param hash - хеш ссылки
   * @param userId - optional user_id from Kratos session, undefined if not authenticated (treated as anonymous)
   * @returns Promise с доменной сущностью Link
   * @throws LinkNotFoundError если ссылка не найдена или доступ запрещен
   */
  findByHash(hash: Hash, userId?: string | null): Promise<Link>;

  /**
   * Сохраняет ссылку
   * @param link - доменная сущность ссылки
   * @returns Promise с сохраненной ссылкой
   */
  save(link: Link): Promise<Link>;

  /**
   * Проверяет существование ссылки по хешу
   * @param hash - хеш ссылки
   * @returns Promise<boolean> - true если ссылка существует
   */
  exists(hash: Hash): Promise<boolean>;
}
