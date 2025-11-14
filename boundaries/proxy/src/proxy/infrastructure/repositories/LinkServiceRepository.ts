import { metrics } from "@opentelemetry/api";
import type { Counter, Meter } from "@opentelemetry/api";
import { injectable, inject } from "inversify";
import { Link } from "../../domain/entities/Link.js";
import { Hash } from "../../domain/entities/Hash.js";
import { ILinkRepository } from "../../domain/repositories/ILinkRepository.js";
import { ILinkServiceAdapter } from "../adapters/ILinkServiceAdapter.js";
import { ILinkCache } from "../cache/RedisLinkCache.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import TYPES from "../../../types.js";

/**
 * Реализация ILinkRepository через внешний Link Service
 * Использует адаптер для получения ссылок из внешнего сервиса
 * Работает с domain entities, не с protobuf или HTTP моделями
 * Интегрирован с Redis кэшем для оптимизации производительности
 */
@injectable()
export class LinkServiceRepository implements ILinkRepository {
  private readonly meter: Meter;
  private readonly linkServiceErrorCounter: Counter;

  constructor(
    @inject(TYPES.INFRASTRUCTURE.LinkServiceAdapter)
    private readonly linkServiceAdapter: ILinkServiceAdapter,
    @inject(TYPES.INFRASTRUCTURE.LinkCache)
    private readonly cache: ILinkCache,
    @inject(TYPES.INFRASTRUCTURE.Logger) private readonly logger: ILogger
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

  async findByHash(hash: Hash): Promise<Link | null> {
    // Проверяем кэш сначала
    const cached = await this.cache.get(hash);

    // Если найден положительный результат в кэше
    if (cached !== undefined && cached !== null) {
      this.logger.debug("Cache hit - returning cached link", {
        hash: hash.value,
      });
      return cached;
    }

    // Если отрицательный кэш - возвращаем null без обращения к адаптеру
    if (cached === null) {
      this.logger.debug("Cache hit - negative, returning null", {
        hash: hash.value,
      });
      return null;
    }

    // Кэш miss - обращаемся к адаптеру
    this.logger.debug("Cache miss - fetching from adapter", {
      hash: hash.value,
    });

    try {
      const link = await this.linkServiceAdapter.getLinkByHash(hash);

      // Сохраняем результат в кэш
      if (link !== null) {
        await this.cache.setPositive(hash, link);
      } else {
        await this.cache.setNegative(hash);
      }

      return link;
    } catch (error) {
      // При ошибке адаптера не сохраняем в кэш, но логируем
      this.logger.error("Adapter error in findByHash", error, {
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
    const link = await this.findByHash(hash);
    return link !== null;
  }
}
