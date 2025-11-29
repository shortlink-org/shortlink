import {
  describe,
  it,
  expect,
  beforeAll,
  afterAll,
  beforeEach,
  afterEach,
} from "vitest";
import { RedisLinkCache } from "../RedisLinkCache.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import { ILogger } from "../../logging/ILogger.js";
import { CacheConfig } from "../../config/CacheConfig.js";
import Redis from "ioredis";
import {
  RedisContainer,
  type StartedRedisContainer,
} from "@testcontainers/redis";

/**
 * Unit tests for RedisLinkCache using Testcontainers
 * Uses real Redis through Testcontainers instead of mocks
 * This provides more realistic testing
 */
describe("RedisLinkCache", () => {
  let cache: RedisLinkCache;
  let logger: ILogger;
  let config: CacheConfig;
  let testRedis: Redis;
  let redisContainer: StartedRedisContainer | null = null;

  beforeAll(async () => {
    // Setup logger
    logger = {
      info: () => {},
      warn: () => {},
      error: () => {},
      debug: () => {},
      http: () => {},
      event: () => {},
    };

    // Start Redis container (or use REDIS_URL if provided)
    const redisUrl = process.env.REDIS_URL
      ? process.env.REDIS_URL
      : await (async () => {
          redisContainer = await new RedisContainer("redis:7.4-alpine").start();
          return redisContainer.getConnectionUrl();
        })();

    // Setup config
    config = {
      enabled: true,
      redisUrl,
      ttlPositive: 3600,
      ttlNegative: 300,
      keyPrefix: "test:shortlink:proxy",
    } as CacheConfig;

    // Create test Redis client for cleanup
    testRedis = new Redis(config.redisUrl);
    await testRedis.ping();
  }, 30000);

  afterAll(async () => {
    if (testRedis) {
      await testRedis.quit();
    }
    if (cache) {
      await cache.disconnect();
    }
    if (redisContainer) {
      await redisContainer.stop();
    }
  }, 30000);

  beforeEach(async () => {
    // Cleanup all test keys before each test
    const keys = await testRedis.keys(`${config.keyPrefix}:*`);
    if (keys.length > 0) {
      await testRedis.del(...keys);
    }

    // Create fresh cache instance for each test
    cache = new RedisLinkCache(logger, config);

    // Wait a bit for cache to connect
    await new Promise((resolve) => setTimeout(resolve, 100));
  });

  afterEach(async () => {
    if (cache) {
      // Cleanup all test keys after each test
      const keys = await testRedis.keys(`${config.keyPrefix}:*`);
      if (keys.length > 0) {
        await testRedis.del(...keys);
      }
    }
  });

  describe("get", () => {
    it("should return Link when cached", async () => {
      // Arrange
      const hash = new Hash("abc123");
      const link = new Link(
        hash,
        "https://example.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );

      await cache.setPositive(hash, link);

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result).not.toBeUndefined();
      expect(result?.hash.value).toBe(hash.value);
      expect(result?.url).toBe(link.url);
    });

    it("should return null for negative cache", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      await cache.setNegative(hash);

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeNull();
    });

    it("should return undefined for cache miss", async () => {
      // Arrange
      const hash = new Hash("miss");

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeUndefined();
    });

    it("should return undefined when Redis is not available", async () => {
      // Arrange - create cache with invalid Redis URL
      const invalidConfig = {
        enabled: true,
        redisUrl: "redis://localhost:9999", // Invalid port
        ttlPositive: 3600,
        ttlNegative: 300,
        keyPrefix: "test:shortlink:proxy",
      } as CacheConfig;

      const cacheWithInvalidRedis = new RedisLinkCache(logger, invalidConfig);
      const hash = new Hash("test");

      // Wait a bit for connection attempt
      await new Promise((resolve) => setTimeout(resolve, 500));

      // Act
      const result = await cacheWithInvalidRedis.get(hash);

      // Assert
      expect(result).toBeUndefined();

      // Cleanup
      await cacheWithInvalidRedis.disconnect();
    });

    it("should handle Redis errors gracefully", async () => {
      // Arrange - use invalid Redis URL to simulate connection error
      const invalidConfig = {
        enabled: true,
        redisUrl: "redis://localhost:9999", // Invalid port
        ttlPositive: 3600,
        ttlNegative: 300,
        keyPrefix: "test:shortlink:proxy",
      } as CacheConfig;

      const cacheWithInvalidRedis = new RedisLinkCache(logger, invalidConfig);
      const hash = new Hash("error");

      // Wait a bit for connection attempt
      await new Promise((resolve) => setTimeout(resolve, 500));

      // Act - should not throw, should return undefined gracefully
      const result = await cacheWithInvalidRedis.get(hash);

      // Assert
      expect(result).toBeUndefined();

      // Cleanup
      await cacheWithInvalidRedis.disconnect();
    });

    it("should clear corrupted cache entry", async () => {
      // Arrange - manually set corrupted JSON
      const hash = new Hash("corrupted");
      const key = `${config.keyPrefix}:hash:${hash.value}`;
      await testRedis.set(key, "invalid json");

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeUndefined();

      // Verify corrupted entry was cleared
      const exists = await testRedis.exists(key);
      expect(exists).toBe(0);
    });
  });

  describe("setPositive", () => {
    it("should save Link to cache", async () => {
      // Arrange
      const hash = new Hash("abc123");
      const link = new Link(
        hash,
        "https://example.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );

      // Act
      await cache.setPositive(hash, link);

      // Assert - verify it was saved by reading it back
      const result = await cache.get(hash);
      expect(result).not.toBeNull();
      expect(result?.hash.value).toBe(hash.value);
      expect(result?.url).toBe(link.url);
    });

    it("should skip when Redis is not available", async () => {
      // Arrange - create cache with invalid Redis URL
      const invalidConfig = {
        enabled: true,
        redisUrl: "redis://localhost:9999",
        ttlPositive: 3600,
        ttlNegative: 300,
        keyPrefix: "test:shortlink:proxy",
      } as CacheConfig;

      const cacheWithInvalidRedis = new RedisLinkCache(logger, invalidConfig);
      const hash = new Hash("test");
      const link = new Link(hash, "https://example.com");

      // Wait a bit for connection attempt
      await new Promise((resolve) => setTimeout(resolve, 500));

      // Act - should not throw
      await cacheWithInvalidRedis.setPositive(hash, link);

      // Assert - should not have saved
      const result = await cacheWithInvalidRedis.get(hash);
      expect(result).toBeUndefined();

      // Cleanup
      await cacheWithInvalidRedis.disconnect();
    });
  });

  describe("setNegative", () => {
    it("should save negative cache", async () => {
      // Arrange
      const hash = new Hash("nonexistent");

      // Act
      await cache.setNegative(hash);

      // Assert - verify it was saved by reading it back
      const result = await cache.get(hash);
      expect(result).toBeNull();
    });

    it("should skip when Redis is not available", async () => {
      // Arrange - create cache with invalid Redis URL
      const invalidConfig = {
        enabled: true,
        redisUrl: "redis://localhost:9999",
        ttlPositive: 3600,
        ttlNegative: 300,
        keyPrefix: "test:shortlink:proxy",
      } as CacheConfig;

      const cacheWithInvalidRedis = new RedisLinkCache(logger, invalidConfig);
      const hash = new Hash("test");

      // Wait a bit for connection attempt
      await new Promise((resolve) => setTimeout(resolve, 500));

      // Act - should not throw
      await cacheWithInvalidRedis.setNegative(hash);

      // Assert - should not have saved
      const result = await cacheWithInvalidRedis.get(hash);
      expect(result).toBeUndefined();

      // Cleanup
      await cacheWithInvalidRedis.disconnect();
    });
  });

  describe("clear", () => {
    it("should delete cache entry", async () => {
      // Arrange
      const hash = new Hash("abc123");
      const link = new Link(hash, "https://example.com");
      await cache.setPositive(hash, link);
      expect(await cache.get(hash)).not.toBeUndefined();

      // Act
      await cache.clear(hash);

      // Assert
      const result = await cache.get(hash);
      expect(result).toBeUndefined();
    });

    it("should skip when Redis is not available", async () => {
      // Arrange - create cache with invalid Redis URL
      const invalidConfig = {
        enabled: true,
        redisUrl: "redis://localhost:9999",
        ttlPositive: 3600,
        ttlNegative: 300,
        keyPrefix: "test:shortlink:proxy",
      } as CacheConfig;

      const cacheWithInvalidRedis = new RedisLinkCache(logger, invalidConfig);
      const hash = new Hash("test");

      // Wait a bit for connection attempt
      await new Promise((resolve) => setTimeout(resolve, 500));

      // Act - should not throw
      await cacheWithInvalidRedis.clear(hash);

      // Cleanup
      await cacheWithInvalidRedis.disconnect();
    });
  });

  describe("when cache is disabled", () => {
    it("should return undefined for get", async () => {
      // Arrange
      const disabledConfig = {
        ...config,
        enabled: false,
      } as CacheConfig;
      const disabledCache = new RedisLinkCache(logger, disabledConfig);
      const hash = new Hash("test");

      // Act
      const result = await disabledCache.get(hash);

      // Assert
      expect(result).toBeUndefined();

      // Cleanup
      await disabledCache.disconnect();
    });

    it("should skip setPositive", async () => {
      // Arrange
      const disabledConfig = {
        ...config,
        enabled: false,
      } as CacheConfig;
      const disabledCache = new RedisLinkCache(logger, disabledConfig);
      const hash = new Hash("test");
      const link = new Link(hash, "https://example.com");

      // Act
      await disabledCache.setPositive(hash, link);

      // Assert - verify nothing was saved
      const result = await disabledCache.get(hash);
      expect(result).toBeUndefined();

      // Also verify directly in Redis
      const key = `${disabledConfig.keyPrefix}:hash:${hash.value}`;
      const exists = await testRedis.exists(key);
      expect(exists).toBe(0);

      // Cleanup
      await disabledCache.disconnect();
    });
  });
});
