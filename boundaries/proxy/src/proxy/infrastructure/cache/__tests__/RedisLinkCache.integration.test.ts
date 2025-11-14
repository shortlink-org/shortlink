import { describe, it, expect, beforeAll, afterAll } from "vitest";
import { RedisLinkCache } from "../RedisLinkCache.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import { ILogger } from "../../../../infrastructure/logging/ILogger.js";
import { CacheConfig } from "../../../../infrastructure/config/CacheConfig.js";
import Redis from "ioredis";
import { RedisContainer } from "@testcontainers/redis";

/**
 * Integration tests for RedisLinkCache
 * Используют Testcontainers для запуска временного Redis
 *
 * Требования:
 * - Запущенный Docker (для Testcontainers)
 * - Или можно указать REDIS_URL для использования внешнего Redis
 */
describe("RedisLinkCache Integration", () => {
  let cache: RedisLinkCache;
  let logger: ILogger;
  let config: CacheConfig;
  let testRedis: Redis;
  let redisContainer: RedisContainer | null = null;
  let isReady = false;
  let failureReason = "Redis Testcontainer is not available";

  beforeAll(async () => {
    // Setup logger
    logger = {
      info: () => {},
      warn: () => {},
      error: () => {},
      debug: () => {},
      http: () => {},
    };

    try {
      const redisUrl = await (async () => {
        if (process.env.REDIS_URL) {
          return process.env.REDIS_URL;
        }

        redisContainer = await new RedisContainer("redis:7.4-alpine").start();
        return redisContainer.getConnectionUrl();
      })();

      // Setup config - создаем объект конфигурации для тестов
      config = {
        enabled: true,
        redisUrl,
        ttlPositive: 60, // 1 minute for tests
        ttlNegative: 30, // 30 seconds for tests
        keyPrefix: "test:shortlink:proxy",
      } as CacheConfig;

      // Create cache instance
      cache = new RedisLinkCache(logger, config);

      // Create test Redis client for cleanup
      testRedis = new Redis(config.redisUrl);
      await testRedis.ping();
      isReady = true;
    } catch (error) {
      failureReason = `Redis Testcontainer setup failed: ${
        error instanceof Error ? error.message : String(error)
      }`;
      console.error(`[RedisLinkCache.integration] ${failureReason}`);
      throw error instanceof Error ? error : new Error(String(error));
    }
  });

  afterAll(async () => {
    if (!isReady) {
      return;
    }

    // Cleanup test keys
    const keys = await testRedis.keys(`${config.keyPrefix}:*`);
    if (keys.length > 0) {
      await testRedis.del(...keys);
    }
    await testRedis.quit();
    await cache.disconnect();
    if (redisContainer) {
      await redisContainer.stop();
    }
  });

  const skipIfNotReady = (): boolean => {
    if (!isReady) {
      console.error(`[RedisLinkCache.integration] ${failureReason}`);
      throw new Error(failureReason);
      return true;
    }
    return false;
  };

  describe("cache hit scenarios", () => {
    it("should return cached Link on cache hit", async () => {
      if (skipIfNotReady()) return;

      // Arrange
      const hash = new Hash("cachehit123");
      const link = new Link(
        hash,
        "https://example.com/cached",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );

      // Set in cache
      await cache.setPositive(hash, link);

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result).not.toBeUndefined();
      expect(result?.hash.value).toBe(hash.value);
      expect(result?.url).toBe(link.url);
    });

    it("should return null for negative cache hit", async () => {
      if (skipIfNotReady()) return;

      // Arrange
      const hash = new Hash("negcache123");

      // Set negative cache
      await cache.setNegative(hash);

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeNull();
    });
  });

  describe("cache miss scenarios", () => {
    it("should return undefined on cache miss", async () => {
      if (skipIfNotReady()) return;

      // Arrange
      const hash = new Hash("cachemiss123");

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeUndefined();
    });
  });

  describe("cache operations", () => {
    it("should store and retrieve Link correctly", async () => {
      if (skipIfNotReady()) return;

      // Arrange
      const hash = new Hash("storetest123");
      const link = new Link(
        hash,
        "https://example.com/stored",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );

      // Act - store
      await cache.setPositive(hash, link);

      // Act - retrieve
      const result = await cache.get(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result?.hash.value).toBe(hash.value);
      expect(result?.url).toBe(link.url);
    });

    it("should clear cache entry", async () => {
      if (skipIfNotReady()) return;

      // Arrange
      const hash = new Hash("cleartest123");
      const link = new Link(hash, "https://example.com/clear");

      await cache.setPositive(hash, link);
      expect(await cache.get(hash)).not.toBeUndefined();

      // Act
      await cache.clear(hash);

      // Assert
      const result = await cache.get(hash);
      expect(result).toBeUndefined();
    });
  });

  describe("graceful degradation", () => {
    it("should handle Redis down gracefully", async () => {
      if (skipIfNotReady()) return;

      // Arrange - use invalid Redis URL
      const invalidConfig = {
        enabled: true,
        redisUrl: "redis://localhost:9999", // Invalid port
        ttlPositive: 60,
        ttlNegative: 30,
        keyPrefix: "test:shortlink:proxy",
      } as CacheConfig;

      const cacheWithInvalidRedis = new RedisLinkCache(logger, invalidConfig);
      const hash = new Hash("graceful123");

      // Wait a bit for connection attempt
      await new Promise((resolve) => setTimeout(resolve, 500));

      // Act - should not throw
      const result = await cacheWithInvalidRedis.get(hash);

      // Assert - should return undefined, not throw
      expect(result).toBeUndefined();

      // Cleanup
      await cacheWithInvalidRedis.disconnect();
    });
  });
});
