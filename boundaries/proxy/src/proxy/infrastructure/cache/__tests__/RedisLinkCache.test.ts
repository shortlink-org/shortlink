import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { RedisLinkCache } from "../RedisLinkCache.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import { ILogger } from "../../../../infrastructure/logging/ILogger.js";
import { CacheConfig } from "../../../../infrastructure/config/CacheConfig.js";
import Redis from "ioredis";

// Mock ioredis
vi.mock("ioredis", () => {
  return {
    default: vi.fn().mockImplementation(() => ({
      get: vi.fn(),
      setex: vi.fn(),
      del: vi.fn(),
      quit: vi.fn(),
      connect: vi.fn().mockResolvedValue(undefined),
      on: vi.fn(),
      status: "ready",
    })),
  };
});

describe("RedisLinkCache", () => {
  let cache: RedisLinkCache;
  let mockLogger: ILogger;
  let mockConfig: CacheConfig;
  let mockRedis: any;

  beforeEach(() => {
    // Mock logger
    mockLogger = {
      info: vi.fn(),
      warn: vi.fn(),
      error: vi.fn(),
      debug: vi.fn(),
      http: vi.fn(),
    };

    // Mock config
    mockConfig = {
      enabled: true,
      redisUrl: "redis://localhost:6379",
      ttlPositive: 3600,
      ttlNegative: 300,
      keyPrefix: "shortlink:proxy",
    } as CacheConfig;

    // Create Redis mock instance
    mockRedis = {
      get: vi.fn(),
      setex: vi.fn(),
      del: vi.fn(),
      quit: vi.fn(),
      connect: vi.fn().mockResolvedValue(undefined),
      on: vi.fn(),
      status: "ready",
    };

    // Mock Redis constructor
    vi.mocked(Redis).mockImplementation(() => mockRedis);

    cache = new RedisLinkCache(mockLogger, mockConfig);
  });

  afterEach(() => {
    vi.clearAllMocks();
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
      const cachedData = JSON.stringify({
        hash: { value: hash.value },
        url: link.url,
        createdAt: link.createdAt.toISOString(),
        updatedAt: link.updatedAt.toISOString(),
      });

      mockRedis.get.mockResolvedValue(cachedData);

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result).not.toBeUndefined();
      expect(result?.hash.value).toBe(hash.value);
      expect(result?.url).toBe(link.url);
      expect(mockRedis.get).toHaveBeenCalledWith(
        "shortlink:proxy:hash:abc123"
      );
    });

    it("should return null for negative cache", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      mockRedis.get.mockResolvedValue("NEGATIVE");

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeNull();
      expect(mockRedis.get).toHaveBeenCalledWith(
        "shortlink:proxy:hash:nonexistent"
      );
    });

    it("should return undefined for cache miss", async () => {
      // Arrange
      const hash = new Hash("miss");
      mockRedis.get.mockResolvedValue(null);

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeUndefined();
      expect(mockRedis.get).toHaveBeenCalledWith(
        "shortlink:proxy:hash:miss"
      );
    });

    it("should return undefined when Redis is not available", async () => {
      // Arrange
      const hash = new Hash("test");
      mockRedis.status = "end";

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeUndefined();
      expect(mockRedis.get).not.toHaveBeenCalled();
    });

    it("should handle Redis errors gracefully", async () => {
      // Arrange
      const hash = new Hash("error");
      mockRedis.get.mockRejectedValue(new Error("Redis error"));

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeUndefined();
      expect(mockLogger.error).toHaveBeenCalled();
    });

    it("should clear corrupted cache entry", async () => {
      // Arrange
      const hash = new Hash("corrupted");
      mockRedis.get.mockResolvedValue("invalid json");
      mockRedis.del.mockResolvedValue(1);

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeUndefined();
      expect(mockRedis.del).toHaveBeenCalledWith(
        "shortlink:proxy:hash:corrupted"
      );
      expect(mockLogger.error).toHaveBeenCalled();
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
      mockRedis.setex.mockResolvedValue("OK");

      // Act
      await cache.setPositive(hash, link);

      // Assert
      expect(mockRedis.setex).toHaveBeenCalledWith(
        "shortlink:proxy:hash:abc123",
        3600,
        expect.stringContaining('"url":"https://example.com"')
      );
    });

    it("should skip when Redis is not available", async () => {
      // Arrange
      const hash = new Hash("test");
      const link = new Link(hash, "https://example.com");
      mockRedis.status = "end";

      // Act
      await cache.setPositive(hash, link);

      // Assert
      expect(mockRedis.setex).not.toHaveBeenCalled();
    });

    it("should handle Redis errors gracefully", async () => {
      // Arrange
      const hash = new Hash("error");
      const link = new Link(hash, "https://example.com");
      mockRedis.setex.mockRejectedValue(new Error("Redis error"));

      // Act
      await cache.setPositive(hash, link);

      // Assert
      expect(mockLogger.error).toHaveBeenCalled();
    });
  });

  describe("setNegative", () => {
    it("should save negative cache", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      mockRedis.setex.mockResolvedValue("OK");

      // Act
      await cache.setNegative(hash);

      // Assert
      expect(mockRedis.setex).toHaveBeenCalledWith(
        "shortlink:proxy:hash:nonexistent",
        300,
        "NEGATIVE"
      );
    });

    it("should skip when Redis is not available", async () => {
      // Arrange
      const hash = new Hash("test");
      mockRedis.status = "end";

      // Act
      await cache.setNegative(hash);

      // Assert
      expect(mockRedis.setex).not.toHaveBeenCalled();
    });

    it("should handle Redis errors gracefully", async () => {
      // Arrange
      const hash = new Hash("error");
      mockRedis.setex.mockRejectedValue(new Error("Redis error"));

      // Act
      await cache.setNegative(hash);

      // Assert
      expect(mockLogger.error).toHaveBeenCalled();
    });
  });

  describe("clear", () => {
    it("should delete cache entry", async () => {
      // Arrange
      const hash = new Hash("abc123");
      mockRedis.del.mockResolvedValue(1);

      // Act
      await cache.clear(hash);

      // Assert
      expect(mockRedis.del).toHaveBeenCalledWith(
        "shortlink:proxy:hash:abc123"
      );
    });

    it("should skip when Redis is not available", async () => {
      // Arrange
      const hash = new Hash("test");
      mockRedis.status = "end";

      // Act
      await cache.clear(hash);

      // Assert
      expect(mockRedis.del).not.toHaveBeenCalled();
    });

    it("should handle Redis errors gracefully", async () => {
      // Arrange
      const hash = new Hash("error");
      mockRedis.del.mockRejectedValue(new Error("Redis error"));

      // Act
      await cache.clear(hash);

      // Assert
      expect(mockLogger.error).toHaveBeenCalled();
    });
  });

  describe("when cache is disabled", () => {
    beforeEach(() => {
      const disabledConfig = {
        ...mockConfig,
        enabled: false,
      } as CacheConfig;
      cache = new RedisLinkCache(mockLogger, disabledConfig);
    });

    it("should return undefined for get", async () => {
      // Arrange
      const hash = new Hash("test");

      // Act
      const result = await cache.get(hash);

      // Assert
      expect(result).toBeUndefined();
      expect(mockRedis.get).not.toHaveBeenCalled();
    });

    it("should skip setPositive", async () => {
      // Arrange
      const hash = new Hash("test");
      const link = new Link(hash, "https://example.com");

      // Act
      await cache.setPositive(hash, link);

      // Assert
      expect(mockRedis.setex).not.toHaveBeenCalled();
    });
  });
});

