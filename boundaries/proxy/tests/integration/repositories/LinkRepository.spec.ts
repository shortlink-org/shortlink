import {
  describe,
  it,
  expect,
  beforeAll,
  afterAll,
  beforeEach,
  vi,
  type Mocked,
} from "vitest";
import { BaseTestEnvironment } from "../environment/BaseTestEnvironment.js";
import type { ILinkRepository } from "../../../src/proxy/domain/repositories/ILinkRepository.js";
import { Link } from "../../../src/proxy/domain/entities/Link.js";
import { Hash } from "../../../src/proxy/domain/entities/Hash.js";
import { overrideDI } from "../helpers/overrideDI.js";
import type { ILinkServiceAdapter } from "../../../src/proxy/infrastructure/adapters/ILinkServiceAdapter.js";
import type { ILinkCache } from "../../../src/proxy/infrastructure/cache/RedisLinkCache.js";

/**
 * Integration tests for LinkRepository.
 * Demonstrates repository testing pattern with Testcontainers and DI overrides.
 *
 * This test shows:
 * - BaseTestEnvironment usage
 * - DI override mechanism
 * - Repository interaction testing
 * - Domain entity mapping verification
 */
describe("LinkRepository Integration", () => {
  let env: BaseTestEnvironment;
  let repository: ILinkRepository;
  let mockAdapter: Mocked<ILinkServiceAdapter>;
  let mockCache: Mocked<ILinkCache>;

  beforeAll(async () => {
    // Start test environment (Redis container)
    env = new BaseTestEnvironment();
    await env.start();
  }, 60000);

  afterAll(async () => {
    // Stop all containers and cleanup
    await env.stop();
  }, 30000);

  beforeEach(() => {
    // Create strictly typed mocks for adapter and cache
    mockAdapter = {
      getLinkByHash: vi.fn<ILinkServiceAdapter["getLinkByHash"]>(),
    } satisfies Mocked<ILinkServiceAdapter>;

    mockCache = {
      get: vi.fn<ILinkCache["get"]>().mockResolvedValue(undefined), // Cache miss by default
      setPositive: vi
        .fn<ILinkCache["setPositive"]>()
        .mockResolvedValue(undefined),
      setNegative: vi
        .fn<ILinkCache["setNegative"]>()
        .mockResolvedValue(undefined),
      clear: vi.fn<ILinkCache["clear"]>().mockResolvedValue(undefined),
    } satisfies Mocked<ILinkCache>;

    // Override DI with mocks
    const container = env.getContainer();
    overrideDI(container, {
      linkServiceAdapter: mockAdapter,
      linkCache: mockCache,
    });

    // Resolve repository from container
    repository = container.resolve<ILinkRepository>("linkRepository");
  });

  describe("findByHash", () => {
    it("should find link by hash and return domain entity", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");

      // Mock adapter to return link
      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result).toBeInstanceOf(Link);
      expect(result?.hash.value).toBe("testhash123");
      expect(result?.url).toBe("https://example.com");
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
    });

    it("should return null when link not found", async () => {
      // Arrange
      const hash = new Hash("nonexistent");

      // Mock adapter to return null
      mockAdapter.getLinkByHash.mockResolvedValue(null);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).toBeNull();
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
    });

    it("should use cache when available", async () => {
      // Arrange
      const hash = new Hash("cachedhash");
      const cachedLink = new Link(hash, "https://cached.example.com");

      // Mock cache to return cached link
      mockCache.get.mockResolvedValue(cachedLink);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).not.toBeNull();
      expect(result?.url).toBe("https://cached.example.com");
      // Adapter should not be called when cache hit
      expect(mockAdapter.getLinkByHash).not.toHaveBeenCalled();
    });

    it("should handle negative cache correctly", async () => {
      // Arrange
      const hash = new Hash("negcache");

      // Mock cache to return null (negative cache)
      mockCache.get.mockResolvedValue(null);

      // Act
      const result = await repository.findByHash(hash);

      // Assert
      expect(result).toBeNull();
      // Adapter should not be called when negative cache hit
      expect(mockAdapter.getLinkByHash).not.toHaveBeenCalled();
    });
  });

  describe("exists", () => {
    it("should return true when link exists", async () => {
      // Arrange
      const hash = new Hash("existing");
      const link = new Link(hash, "https://example.com");

      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      const exists = await repository.exists(hash);

      // Assert
      expect(exists).toBe(true);
    });

    it("should return false when link does not exist", async () => {
      // Arrange
      const hash = new Hash("nonexistent");

      (mockAdapter.getLinkByHash as any).mockResolvedValue(null);

      // Act
      const exists = await repository.exists(hash);

      // Assert
      expect(exists).toBe(false);
    });
  });

  describe("domain entity mapping", () => {
    it("should preserve domain entity invariants", async () => {
      // Arrange
      const hash = new Hash("validhash");
      const link = new Link(
        hash,
        "https://example.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );

      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      const result = await repository.findByHash(hash);

      // Assert - verify domain entity properties are preserved
      expect(result).not.toBeNull();
      expect(result?.hash.equals(hash)).toBe(true);
      expect(result?.isValidForRedirect()).toBe(true);
      expect(result?.createdAt).toBeInstanceOf(Date);
      expect(result?.updatedAt).toBeInstanceOf(Date);
    });
  });
});
