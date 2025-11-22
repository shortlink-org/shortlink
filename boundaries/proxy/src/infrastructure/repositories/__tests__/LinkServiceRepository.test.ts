import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { LinkServiceRepository } from "../LinkServiceRepository.js";
import { ILinkServiceAdapter } from "../../adapters/ILinkServiceAdapter.js";
import { ILinkCache } from "../../cache/RedisLinkCache.js";
import { ILogger } from "../../logging/ILogger.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";

describe("LinkServiceRepository", () => {
  let repository: LinkServiceRepository;
  let mockAdapter: {
    getLinkByHash: ReturnType<typeof vi.fn>;
  };
  let mockCache: {
    get: ReturnType<typeof vi.fn>;
    setPositive: ReturnType<typeof vi.fn>;
    setNegative: ReturnType<typeof vi.fn>;
  };
  let mockLogger: ILogger;

  beforeEach(() => {
    // Создаем мок адаптера
    mockAdapter = {
      getLinkByHash: vi.fn(),
    } as any;

    // Создаем мок кэша
    mockCache = {
      get: vi.fn().mockResolvedValue(undefined),
      setPositive: vi.fn().mockResolvedValue(undefined),
      setNegative: vi.fn().mockResolvedValue(undefined),
    } as any;

    // Создаем мок логгера
    mockLogger = {
      info: vi.fn(),
      warn: vi.fn(),
      error: vi.fn(),
      debug: vi.fn(),
      http: vi.fn(),
    };

    repository = new LinkServiceRepository(
      mockAdapter as any,
      mockCache as any,
      mockLogger
    );
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  describe("findByHash", () => {
    it("should return cached Link when cache hit", async () => {
      // Arrange
      const hash = new Hash("abc123");
      const link = new Link(
        hash,
        "https://example.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );

      mockCache.get.mockResolvedValue(link);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result).toEqual(link);
      expect(mockCache.get).toHaveBeenCalledWith(hash);
      expect(mockAdapter.getLinkByHash).not.toHaveBeenCalled();
    });

    it("should return null when negative cache hit", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      mockCache.get.mockResolvedValue(null); // negative cache

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).toBeNull();
      expect(mockCache.get).toHaveBeenCalledWith(hash);
      expect(mockAdapter.getLinkByHash).not.toHaveBeenCalled();
    });

    it("should fetch from adapter on cache miss and cache result", async () => {
      // Arrange
      const hash = new Hash("abc123");
      const link = new Link(
        hash,
        "https://example.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );

      mockCache.get.mockResolvedValue(undefined); // cache miss
      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result).toEqual(link);
      expect(mockCache.get).toHaveBeenCalledWith(hash);
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
      expect(mockCache.setPositive).toHaveBeenCalledWith(hash, link);
    });

    it("should cache negative result when link not found", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      mockCache.get.mockResolvedValue(undefined); // cache miss
      mockAdapter.getLinkByHash.mockResolvedValue(null);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).toBeNull();
      expect(mockCache.get).toHaveBeenCalledWith(hash);
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
      expect(mockCache.setNegative).toHaveBeenCalledWith(hash);
    });

    it("should return Link when found", async () => {
      // Arrange
      const hash = new Hash("abc123");
      const link = new Link(
        hash,
        "https://example.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );

      mockCache.get.mockResolvedValue(undefined); // cache miss
      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result).toEqual(link);
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledTimes(1);
    });

    it("should return null when link not found", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      mockCache.get.mockResolvedValue(undefined); // cache miss
      mockAdapter.getLinkByHash.mockResolvedValue(null);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).toBeNull();
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
    });

    it("should handle adapter errors", async () => {
      // Arrange
      const hash = new Hash("error");
      const adapterError = new Error("Service unavailable");
      mockAdapter.getLinkByHash.mockRejectedValue(adapterError);

      // Act & Assert
      await expect(repository.findByHash(hash)).rejects.toThrow(
        "Service unavailable"
      );
    });
  });

  describe("save", () => {
    it("should throw error as LinkServiceRepository is read-only", async () => {
      // Arrange
      const hash = new Hash("test");
      const link = new Link(hash, "https://example.com");

      // Act & Assert
      await expect(repository.save(link)).rejects.toThrow(
        "LinkServiceRepository does not support saving links"
      );
    });
  });

  describe("exists", () => {
    it("should return true when link exists", async () => {
      // Arrange
      const hash = new Hash("exists");
      const link = new Link(hash, "https://example.com");
      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      const result = await repository.exists(hash);

      // Assert
      expect(result).toBe(true);
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
    });

    it("should return false when link does not exist", async () => {
      // Arrange
      const hash = new Hash("notexists");
      mockAdapter.getLinkByHash.mockResolvedValue(null);

      // Act
      const result = await repository.exists(hash);

      // Assert
      expect(result).toBe(false);
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
    });
  });
});
