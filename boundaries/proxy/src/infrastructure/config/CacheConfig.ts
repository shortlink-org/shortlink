import { ConfigReader } from "./ConfigReader.js";

/**
 * Конфигурация кэширования
 * Децентрализованный подход - модуль сам определяет свои дефолты
 */
export class CacheConfig {
  /**
   * Включено ли кэширование
   */
  readonly enabled: boolean;

  /**
   * URL для подключения к Redis
   */
  readonly redisUrl: string;

  /**
   * TTL для положительных результатов (в секундах)
   */
  readonly ttlPositive: number;

  /**
   * TTL для отрицательных результатов (в секундах)
   */
  readonly ttlNegative: number;

  /**
   * Префикс для ключей кэша
   */
  readonly keyPrefix: string;

  constructor() {
    this.enabled = ConfigReader.boolean("CACHE_ENABLED", false);
    this.redisUrl = ConfigReader.string(
      "REDIS_URL",
      "redis://localhost:6379"
    );
    this.ttlPositive = ConfigReader.number("CACHE_TTL_POSITIVE", 3600); // 1 hour
    this.ttlNegative = ConfigReader.number("CACHE_TTL_NEGATIVE", 300); // 5 minutes
    this.keyPrefix = ConfigReader.string("CACHE_KEY_PREFIX", "shortlink:proxy");
  }
}

