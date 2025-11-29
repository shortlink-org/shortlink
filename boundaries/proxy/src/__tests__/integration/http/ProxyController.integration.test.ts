import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import type { FastifyInstance } from "fastify";
import { createTestServer } from "../helpers/testServer.js";
import { LinkApplicationService } from "../../../application/services/LinkApplicationService.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import { LinkNotFoundError } from "../../../domain/exceptions/index.js";
import { Result, ok, err } from "neverthrow";
import type { ILogger } from "../../../infrastructure/logging/ILogger.js";

/**
 * Integration tests for HTTP endpoints
 * Tests real Fastify server with full middleware stack
 */
describe("ProxyController Integration Tests", () => {
  let app: FastifyInstance;
  let mockLinkApplicationService: {
    handleRedirect: ReturnType<typeof vi.fn>;
  };
  let mockLogger: {
    info: ReturnType<typeof vi.fn>;
    warn: ReturnType<typeof vi.fn>;
    error: ReturnType<typeof vi.fn>;
    debug: ReturnType<typeof vi.fn>;
    http: ReturnType<typeof vi.fn>;
  };

  beforeEach(async () => {
    // Mock LinkApplicationService
    mockLinkApplicationService = {
      handleRedirect: vi.fn(),
    } as any;

    // Mock Logger
    mockLogger = {
      info: vi.fn(),
      warn: vi.fn(),
      error: vi.fn(),
      debug: vi.fn(),
      http: vi.fn(),
    } as any;

    // Create Fastify server with mocked dependencies
    app = await createTestServer({
      linkApplicationService: mockLinkApplicationService as any,
      logger: mockLogger as any,
    } as any);
  });

  afterEach(async () => {
    vi.clearAllMocks();
    await app.close();
  });

  describe("GET /s/:hash", () => {
    it("should redirect to original URL when link is found", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");
      mockLinkApplicationService.handleRedirect.mockResolvedValue(ok({ link }));

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/testhash123",
      });

      // Assert
      expect(response.statusCode).toBe(301);
      expect(response.headers.location).toBe("https://example.com");
      expect(mockLinkApplicationService.handleRedirect).toHaveBeenCalledWith({
        hash: expect.any(Hash),
        userId: expect.any(String),
      });
    });

    it("should return 404 when link is not found", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      mockLinkApplicationService.handleRedirect.mockResolvedValue(
        err(new LinkNotFoundError(hash))
      );

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/nonexistent",
      });

      // Assert
      expect(response.statusCode).toBe(404);
      expect(response.json()).toHaveProperty("error.code", "LINK_NOT_FOUND");
      expect(response.json()).toHaveProperty("error.message");
      expect(mockLinkApplicationService.handleRedirect).toHaveBeenCalledWith({
        hash: expect.any(Hash),
        userId: expect.any(String),
      });
    });

    it("should return 400 when hash is invalid", async () => {
      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/invalid-hash!",
      });

      // Assert
      expect(response.statusCode).toBe(400);
      expect(response.json()).toHaveProperty("error.code");
      expect(response.json()).toHaveProperty("error.message");
      expect(mockLinkApplicationService.handleRedirect).not.toHaveBeenCalled();
    });

    it("should return 404 when hash is empty", async () => {
      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/",
      });

      // Assert
      // Fastify will return 404 for unmatched route pattern (empty hash after /s/)
      expect([400, 404]).toContain(response.statusCode);
    });

    it("should handle application service errors", async () => {
      // Arrange
      const hash = new Hash("testhash");
      mockLinkApplicationService.handleRedirect.mockRejectedValue(
        new Error("Internal server error")
      );

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/testhash",
      });

      // Assert
      expect(response.statusCode).toBe(500);
      expect(response.json()).toHaveProperty("error.code");
      expect(response.json()).toHaveProperty("error.message");
    });
  });
});
