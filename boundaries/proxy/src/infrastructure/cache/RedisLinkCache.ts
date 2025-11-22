import Redis from "ioredis";
import { metrics, trace, context } from "@opentelemetry/api";
import type { Meter, Counter, Histogram } from "@opentelemetry/api";
import { Link } from "../../domain/entities/Link.js";
import { Hash } from "../../domain/entities/Hash.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { CacheConfig } from "../../../infrastructure/config/CacheConfig.js";

/**
 * Интерфейс для кэша ссылок
 */
export interface ILinkCache {
  /**
   * Получает ссылку из кэша
   * @param hash - хеш ссылки
   * @returns Link если найдена, null если отрицательный кэш, undefined если не найдено
   */
  get(hash: Hash): Promise<Link | null | undefined>;

  /**
   * Сохраняет положительный результат в кэш
   * @param hash - хеш ссылки
   * @param link - ссылка для сохранения
   */
  setPositive(hash: Hash, link: Link): Promise<void>;

  /**
   * Сохраняет отрицательный результат в кэш (ссылка не найдена)
   * @param hash - хеш ссылки
   */
  setNegative(hash: Hash): Promise<void>;

  /**
   * Очищает кэш для конкретного хеша
   * @param hash - хеш ссылки
   */
  clear(hash: Hash): Promise<void>;
}

/**
 * Реализация кэша ссылок на базе Redis
 * Использует ioredis для работы с Redis
 * Gracefully деградирует при недоступности Redis - не ломает работу приложения
 */
export class RedisLinkCache implements ILinkCache {
  private readonly redis: Redis | null = null;
  private readonly enabled: boolean;
  private readonly keyPrefix: string;
  private readonly ttlPositive: number;
  private readonly ttlNegative: number;

  // Prometheus metrics
  private readonly meter: Meter;
  private readonly requestCounter: Counter;
  private readonly durationHistogram: Histogram;
  private readonly errorCounter: Counter;
  private readonly cacheHitCounter: Counter;
  private readonly cacheMissCounter: Counter;

  constructor(
    private readonly logger: ILogger,
    private readonly cacheConfig: CacheConfig
  ) {
    this.enabled = cacheConfig.enabled;
    this.keyPrefix = cacheConfig.keyPrefix;
    this.ttlPositive = cacheConfig.ttlPositive;
    this.ttlNegative = cacheConfig.ttlNegative;

    // Initialize OpenTelemetry metrics
    this.meter = metrics.getMeter("proxy-service", "1.0.0");

    // Cache requests counter (hit/miss/error)
    this.requestCounter = this.meter.createCounter("cache_requests_total", {
      description: "Total number of cache requests",
      unit: "1",
    });

    // Cache operation duration histogram
    this.durationHistogram = this.meter.createHistogram("cache_duration_ms", {
      description: "Duration of cache operations in milliseconds",
      unit: "ms",
    });

    // Cache errors counter
    this.errorCounter = this.meter.createCounter("cache_errors_total", {
      description: "Total number of cache errors",
      unit: "1",
    });
    this.cacheHitCounter = this.meter.createCounter("cache_hits_total", {
      description: "Total number of cache hits",
      unit: "1",
    });
    this.cacheMissCounter = this.meter.createCounter("cache_misses_total", {
      description: "Total number of cache misses",
      unit: "1",
    });

    if (this.enabled) {
      try {
        this.redis = new Redis(cacheConfig.redisUrl, {
          retryStrategy: (times) => {
            // Экспоненциальная задержка с максимумом 3 секунды
            const delay = Math.min(times * 50, 3000);
            this.logger.debug("Redis retry attempt", { times, delay });
            return delay;
          },
          maxRetriesPerRequest: 3,
          enableReadyCheck: true,
          lazyConnect: true,
        });

        // Обработка ошибок подключения
        this.redis.on("error", (error) => {
          this.logger.error("Redis connection error", error, {
            redisUrl: cacheConfig.redisUrl,
          });
        });

        this.redis.on("connect", () => {
          this.logger.debug("Redis connected", { redisUrl: cacheConfig.redisUrl });
        });

        this.redis.on("close", () => {
          this.logger.warn("Redis connection closed");
        });

        // Подключаемся асинхронно, не блокируя конструктор
        this.redis.connect().catch((error) => {
          this.logger.error("Failed to connect to Redis", error, {
            redisUrl: cacheConfig.redisUrl,
          });
        });
      } catch (error) {
        this.logger.error("Failed to initialize Redis", error, {
          redisUrl: cacheConfig.redisUrl,
        });
      }
    } else {
      this.logger.debug("Redis cache disabled");
    }
  }

  /**
   * Генерирует ключ кэша для хеша
   */
  private getCacheKey(hash: Hash): string {
    return `${this.keyPrefix}:hash:${hash.value}`;
  }

  /**
   * Проверяет доступность Redis
   */
  private isAvailable(): boolean {
    if (!this.enabled || this.redis === null) {
      return false;
    }
    // Проверяем статус подключения
    const status = this.redis.status;
    return status === "ready" || status === "connect";
  }

  private setCacheSpanAttributes(
    hash: Hash,
    result: "HIT" | "MISS" | "NEGATIVE"
  ): void {
    const span = trace.getSpan(context.active());
    if (!span) {
      return;
    }
    span.setAttribute("cache.result", result);
    span.setAttribute("url.hash", hash.value);
  }

  async get(hash: Hash): Promise<Link | null | undefined> {
    const startTime = Date.now();

    if (!this.isAvailable()) {
      this.logger.debug("Cache miss - Redis not available", {
        hash: hash.value,
      });
      this.requestCounter.add(1, {
        operation: "get",
        type: "miss",
        reason: "unavailable",
      });
      this.cacheMissCounter.add(1, {
        reason: "unavailable",
      });
      this.setCacheSpanAttributes(hash, "MISS");
      return undefined;
    }

    try {
      const key = this.getCacheKey(hash);
      const cached = await this.redis!.get(key);
      const duration = Date.now() - startTime;
      this.durationHistogram.record(duration, {
        operation: "get",
      });

      if (cached === null) {
        this.logger.debug("Cache miss", { hash: hash.value });
        this.requestCounter.add(1, {
          operation: "get",
          type: "miss",
        });
        this.cacheMissCounter.add(1, {
          reason: "not_cached",
        });
        this.setCacheSpanAttributes(hash, "MISS");
        return undefined;
      }

      // Проверяем отрицательный кэш (специальное значение)
      if (cached === "NEGATIVE") {
        this.logger.debug("Cache hit - negative", { hash: hash.value });
        this.requestCounter.add(1, {
          operation: "get",
          type: "hit",
          result: "negative",
        });
        this.cacheHitCounter.add(1, {
          result: "NEGATIVE",
        });
        this.setCacheSpanAttributes(hash, "NEGATIVE");
        return null;
      }

      // Десериализуем Link из JSON
      try {
        const linkData = JSON.parse(cached);
        const link = new Link(
          new Hash(linkData.hash.value),
          linkData.url,
          new Date(linkData.createdAt),
          new Date(linkData.updatedAt)
        );
        this.logger.debug("Cache hit - positive", { hash: hash.value });
        this.requestCounter.add(1, {
          operation: "get",
          type: "hit",
          result: "positive",
        });
        this.cacheHitCounter.add(1, {
          result: "HIT",
        });
        this.setCacheSpanAttributes(hash, "HIT");
        return link;
      } catch (parseError) {
        this.logger.error(
          "Failed to parse cached link",
          parseError,
          { hash: hash.value, cached }
        );
        this.errorCounter.add(1, {
          operation: "get",
          error_type: "parse_error",
        });
        this.cacheMissCounter.add(1, {
          reason: "parse_error",
        });
        this.setCacheSpanAttributes(hash, "MISS");
        // Очищаем поврежденный кэш
        await this.clear(hash);
        return undefined;
      }
    } catch (error) {
      const duration = Date.now() - startTime;
      this.durationHistogram.record(duration, {
        operation: "get",
      });
      this.logger.error("Cache get error", error, { hash: hash.value });
      this.errorCounter.add(1, {
        operation: "get",
        error_type: "redis_error",
      });
      this.requestCounter.add(1, {
        operation: "get",
        type: "error",
      });
      this.cacheMissCounter.add(1, {
        reason: "redis_error",
      });
      this.setCacheSpanAttributes(hash, "MISS");
      // Graceful degradation - возвращаем undefined вместо ошибки
      return undefined;
    }
  }

  async setPositive(hash: Hash, link: Link): Promise<void> {
    const startTime = Date.now();

    if (!this.isAvailable()) {
      this.logger.debug("Cache set skipped - Redis not available", {
        hash: hash.value,
      });
      return;
    }

    try {
      const key = this.getCacheKey(hash);
      const linkData = {
        hash: { value: link.hash.value },
        url: link.url,
        createdAt: link.createdAt.toISOString(),
        updatedAt: link.updatedAt.toISOString(),
      };
      const serialized = JSON.stringify(linkData);

      await this.redis!.setex(key, this.ttlPositive, serialized);
      const duration = Date.now() - startTime;
      this.durationHistogram.record(duration, {
        operation: "set",
        type: "positive",
      });
      this.logger.debug("Cache set - positive", {
        hash: hash.value,
        ttl: this.ttlPositive,
      });
    } catch (error) {
      const duration = Date.now() - startTime;
      this.durationHistogram.record(duration, {
        operation: "set",
        type: "positive",
      });
      this.logger.error("Cache set error", error, { hash: hash.value });
      this.errorCounter.add(1, {
        operation: "set",
        type: "positive",
        error_type: "redis_error",
      });
      // Graceful degradation - не выбрасываем ошибку
    }
  }

  async setNegative(hash: Hash): Promise<void> {
    const startTime = Date.now();

    if (!this.isAvailable()) {
      this.logger.debug("Cache set skipped - Redis not available", {
        hash: hash.value,
      });
      return;
    }

    try {
      const key = this.getCacheKey(hash);
      // Используем специальное значение для отрицательного кэша
      await this.redis!.setex(key, this.ttlNegative, "NEGATIVE");
      const duration = Date.now() - startTime;
      this.durationHistogram.record(duration, {
        operation: "set",
        type: "negative",
      });
      this.logger.debug("Cache set - negative", {
        hash: hash.value,
        ttl: this.ttlNegative,
      });
    } catch (error) {
      const duration = Date.now() - startTime;
      this.durationHistogram.record(duration, {
        operation: "set",
        type: "negative",
      });
      this.logger.error("Cache set error", error, { hash: hash.value });
      this.errorCounter.add(1, {
        operation: "set",
        type: "negative",
        error_type: "redis_error",
      });
      // Graceful degradation - не выбрасываем ошибку
    }
  }

  async clear(hash: Hash): Promise<void> {
    const startTime = Date.now();

    if (!this.isAvailable()) {
      this.logger.debug("Cache clear skipped - Redis not available", {
        hash: hash.value,
      });
      return;
    }

    try {
      const key = this.getCacheKey(hash);
      await this.redis!.del(key);
      const duration = Date.now() - startTime;
      this.durationHistogram.record(duration, {
        operation: "clear",
      });
      this.logger.debug("Cache cleared", { hash: hash.value });
    } catch (error) {
      const duration = Date.now() - startTime;
      this.durationHistogram.record(duration, {
        operation: "clear",
      });
      this.logger.error("Cache clear error", error, { hash: hash.value });
      this.errorCounter.add(1, {
        operation: "clear",
        error_type: "redis_error",
      });
      // Graceful degradation - не выбрасываем ошибку
    }
  }

  /**
   * Закрывает соединение с Redis
   * Должен быть вызван при завершении работы приложения
   */
  async disconnect(): Promise<void> {
    if (this.redis) {
      try {
        await this.redis.quit();
        this.logger.debug("Redis disconnected");
      } catch (error) {
        this.logger.error("Error disconnecting Redis", error);
      }
    }
  }
}

