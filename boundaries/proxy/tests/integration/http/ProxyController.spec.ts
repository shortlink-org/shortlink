import {
  describe,
  it,
  expect,
  beforeAll,
  afterAll,
  beforeEach,
  afterEach,
  vi,
  type Mocked,
} from "vitest";
import type { FastifyInstance } from "fastify";
import { BaseTestEnvironment } from "../environment/BaseTestEnvironment.js";
import { buildTestServer } from "../helpers/buildTestServer.js";
import { overrideDI } from "../helpers/overrideDI.js";
import { Link } from "../../../src/proxy/domain/entities/Link.js";
import { Hash } from "../../../src/proxy/domain/entities/Hash.js";
import type { ILinkServiceAdapter } from "../../../src/proxy/infrastructure/adapters/ILinkServiceAdapter.js";
import type { ILinkCache } from "../../../src/proxy/infrastructure/cache/RedisLinkCache.js";

/**
 * Integration tests for ProxyController HTTP routes.
 * Tests full HTTP request/response flow with Fastify test server.
 *
 * This test demonstrates:
 * - BaseTestEnvironment usage
 * - Fastify test server setup
 * - HTTP route testing with server.inject()
 * - DI override for mocks
 * - Repository interaction verification
 */
describe("ProxyController Integration", () => {
  let env: BaseTestEnvironment;
  let server: FastifyInstance;
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

  beforeEach(async () => {
    // Create strictly typed mocks
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

    // Build test server
    server = await buildTestServer(env);
  });

  afterEach(async () => {
    // Close server after each test
    await server.close();
    vi.clearAllMocks();
  });

  describe("GET /s/:hash", () => {
    it("should redirect to original URL when link found", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");

      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      const response = await server.inject({
        method: "GET",
        url: "/s/testhash123",
      });

      // Assert
      expect(response.statusCode).toBe(301);
      expect(response.headers.location).toBe("https://example.com");
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
    });

    it("should return 404 when link not found", async () => {
      // Arrange
      const hash = new Hash("nonexistent");

      mockAdapter.getLinkByHash.mockResolvedValue(null);

      // Act
      const response = await server.inject({
        method: "GET",
        url: "/s/nonexistent",
      });

      // Assert
      expect(response.statusCode).toBe(404);
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(hash);
    });

    it("should use cache when available", async () => {
      // Arrange
      const hash = new Hash("cachedhash");
      const cachedLink = new Link(hash, "https://cached.example.com");

      mockCache.get.mockResolvedValue(cachedLink);

      // Act
      const response = await server.inject({
        method: "GET",
        url: "/s/cachedhash",
      });

      // Assert
      expect(response.statusCode).toBe(301);
      expect(response.headers.location).toBe("https://cached.example.com");
      // Adapter should not be called when cache hit
      expect(mockAdapter.getLinkByHash).not.toHaveBeenCalled();
    });

    it("should handle invalid hash format", async () => {
      // Act - invalid hash with special characters
      const response = await server.inject({
        method: "GET",
        url: "/s/invalid-hash!",
      });

      // Assert - should return 400 or 404 depending on validation
      expect([400, 404]).toContain(response.statusCode);
    });

    it("should handle multiple redirects correctly", async () => {
      // Arrange
      const hash1 = new Hash("hash1");
      const link1 = new Link(hash1, "https://example1.com");
      const hash2 = new Hash("hash2");
      const link2 = new Link(hash2, "https://example2.com");

      mockAdapter.getLinkByHash
        .mockResolvedValueOnce(link1)
        .mockResolvedValueOnce(link2);

      // Act
      const response1 = await server.inject({
        method: "GET",
        url: "/s/hash1",
      });
      const response2 = await server.inject({
        method: "GET",
        url: "/s/hash2",
      });

      // Assert
      expect(response1.statusCode).toBe(301);
      expect(response1.headers.location).toBe("https://example1.com");
      expect(response2.statusCode).toBe(301);
      expect(response2.headers.location).toBe("https://example2.com");
    });
  });

  describe("repository interactions", () => {
    it("should call repository with correct hash", async () => {
      // Arrange
      const hash = new Hash("testhash");
      const link = new Link(hash, "https://example.com");

      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      await server.inject({
        method: "GET",
        url: "/s/testhash",
      });

      // Assert - verify repository was called with domain entity
      expect(mockAdapter.getLinkByHash).toHaveBeenCalledWith(
        expect.objectContaining({
          value: "testhash",
        })
      );
    });

    it("should cache result after successful lookup", async () => {
      // Arrange
      const hash = new Hash("cacheme");
      const link = new Link(hash, "https://example.com");

      mockAdapter.getLinkByHash.mockResolvedValue(link);

      // Act
      await server.inject({
        method: "GET",
        url: "/s/cacheme",
      });

      // Assert - verify cache was updated
      expect(mockCache.setPositive).toHaveBeenCalledWith(hash, link);
    });
  });
});
