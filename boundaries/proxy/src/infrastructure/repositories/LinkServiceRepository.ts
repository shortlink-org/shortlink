import { metrics } from "@opentelemetry/api";
import type { Counter, Meter } from "@opentelemetry/api";
import { Link } from "../../domain/entities/Link.js";
import { Hash } from "../../domain/entities/Hash.js";
import { LinkNotFoundError } from "../../domain/exceptions/index.js";
import { ILinkRepository } from "../../domain/repositories/ILinkRepository.js";
import { ILinkServiceAdapter } from "../adapters/ILinkServiceAdapter.js";
import { ILinkCache } from "../cache/RedisLinkCache.js";
import { ILogger } from "../logging/ILogger.js";

/**
 * Реализация ILinkRepository через внешний Link Service
 * Использует адаптер для получения ссылок из внешнего сервиса
 * Работает с domain entities, не с protobuf или HTTP моделями
 * Интегрирован с Redis кэшем для оптимизации производительности
 */
export class LinkServiceRepository implements ILinkRepository {
  private readonly meter: Meter;
  private readonly linkServiceErrorCounter: Counter;

  constructor(
    private readonly linkServiceAdapter: ILinkServiceAdapter,
    private readonly linkCache: ILinkCache,
    private readonly logger: ILogger
  ) {
    this.meter = metrics.getMeter("proxy-service", "1.0.0");
    this.linkServiceErrorCounter = this.meter.createCounter(
      "linkservice_errors_total",
      {
        description: "Total number of Link Service adapter errors",
        unit: "1",
      }
    );
  }

  async findByHash(hash: Hash, userId?: string | null): Promise<Link> {
    // Note: Cache doesn't consider userId, so private links might be cached incorrectly
    // For now, we skip cache when userId is provided and not "anonymous" (private link access)
    // TODO: Consider cache key that includes userId for private links

    if (userId && userId !== "anonymous") {
      // For private links, skip cache and go directly to adapter
      this.logger.debug("Cache bypass - fetching private link from adapter", {
        hash: hash.value,
        hasUserId: !!userId,
      });

      try {
        return await this.linkServiceAdapter.getLinkByHash(hash, userId);
      } catch (error) {
        this.logger.error("Adapter error in findByHash (private link)", {
          error: error instanceof Error ? error : new Error(String(error)),
          hash: hash.value,
        });
        this.linkServiceErrorCounter.add(1, {
          error_name: error instanceof Error ? error.name : "UnknownError",
        });
        throw error;
      }
    }

    // For public links, use cache
    const cached = await this.linkCache.get(hash);

    // Если найден положительный результат в кэше
    if (cached !== undefined && cached !== null) {
      this.logger.debug("Cache hit - returning cached link", {
        hash: hash.value,
      });
      return cached;
    }

    // Если отрицательный кэш - бросаем LinkNotFoundError
    if (cached === null) {
      this.logger.debug("Cache hit - negative, throwing LinkNotFoundError", {
        hash: hash.value,
      });
      throw new LinkNotFoundError(hash);
    }

    // Кэш miss - обращаемся к адаптеру
    this.logger.debug("Cache miss - fetching from adapter", {
      hash: hash.value,
    });

    try {
      const link = await this.linkServiceAdapter.getLinkByHash(hash);

      // Сохраняем результат в кэш
      await this.linkCache.setPositive(hash, link);

      return link;
    } catch (error) {
      // Если LinkNotFoundError - сохраняем отрицательный кэш
      if (error instanceof LinkNotFoundError) {
        await this.linkCache.setNegative(hash);
      }

      // Логируем и пробрасываем ошибку
      this.logger.error("Adapter error in findByHash", {
        error: error instanceof Error ? error : new Error(String(error)),
        hash: hash.value,
      });
      this.linkServiceErrorCounter.add(1, {
        error_name: error instanceof Error ? error.name : "UnknownError",
      });
      throw error;
    }
  }

  async save(link: Link): Promise<Link> {
    // Link Service Repository только читает данные из внешнего сервиса
    // Сохранение ссылок не поддерживается (это ответственность Link Service)
    throw new Error(
      "LinkServiceRepository does not support saving links. Links are managed by Link Service."
    );
  }

  async exists(hash: Hash): Promise<boolean> {
    try {
      await this.findByHash(hash);
      return true;
    } catch (error) {
      if (error instanceof LinkNotFoundError) {
        return false;
      }
      // Re-throw other errors
      throw error;
    }
  }
}
