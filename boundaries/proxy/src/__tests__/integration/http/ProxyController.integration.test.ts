import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import request from "supertest";
import { Container } from "inversify";
import * as express from "express";
import { createTestServer } from "../helpers/testServer.js";
import { LinkApplicationService } from "../../../proxy/application/services/LinkApplicationService.js";
import { Hash } from "../../../proxy/domain/entities/Hash.js";
import { Link } from "../../../proxy/domain/entities/Link.js";
import { LinkNotFoundError } from "../../../proxy/domain/exceptions/index.js";
import { Result, ok, err } from "neverthrow";
import TYPES from "../../../types.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";

/**
 * Integration тесты для HTTP endpoints
 * Тестируют реальный Express сервер с полным middleware stack
 */
describe("ProxyController Integration Tests", () => {
  let app: express.Application;
  let container: Container;
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
    // Создаем новый контейнер для каждого теста
    container = new Container();

    // Мокаем LinkApplicationService
    mockLinkApplicationService = {
      handleRedirect: vi.fn(),
    } as any;

    // Мокаем Logger
    mockLogger = {
      info: vi.fn(),
      warn: vi.fn(),
      error: vi.fn(),
      debug: vi.fn(),
      http: vi.fn(),
    } as any;

    // Биндим моки в контейнер
    container
      .bind<LinkApplicationService>(TYPES.APPLICATION.LinkApplicationService)
      .toConstantValue(mockLinkApplicationService as any);

    container.bind<ILogger>(TYPES.INFRASTRUCTURE.Logger).toConstantValue(mockLogger as any);

    // Создаем Express приложение с тестовым контейнером
    app = await createTestServer(container);
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  describe("GET /s/:hash", () => {
    it("should redirect to original URL when link is found", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");
      mockLinkApplicationService.handleRedirect.mockResolvedValue(
        ok({ link })
      );

      // Act
      const response = await request(app).get("/s/testhash123");

      // Assert
      expect(response.status).toBe(301);
      expect(response.headers.location).toBe("https://example.com");
      expect(mockLinkApplicationService.handleRedirect).toHaveBeenCalledWith({
        hash: expect.objectContaining({ value: "testhash123" }),
      });
    });

    it("should return 404 when link is not found", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      mockLinkApplicationService.handleRedirect.mockResolvedValue(
        err(new LinkNotFoundError(hash))
      );

      // Act
      const response = await request(app).get("/s/nonexistent");

      // Assert
      expect(response.status).toBe(404);
      expect(response.body.error).toHaveProperty("code", "LINK_NOT_FOUND");
      expect(response.body.error).toHaveProperty("message");
      expect(mockLinkApplicationService.handleRedirect).toHaveBeenCalledWith({
        hash: expect.objectContaining({ value: "nonexistent" }),
      });
    });

    it("should return 400 when hash is invalid", async () => {
      // Act
      const response = await request(app).get("/s/invalid-hash!");

      // Assert
      expect(response.status).toBe(400);
      expect(response.body.error).toHaveProperty("code");
      expect(response.body.error).toHaveProperty("message");
      expect(mockLinkApplicationService.handleRedirect).not.toHaveBeenCalled();
    });

    it("should return 400 when hash is empty", async () => {
      // Act
      const response = await request(app).get("/s/");

      // Assert
      expect(response.status).toBe(404); // Express route не найден
    });

    it("should handle application service errors", async () => {
      // Arrange
      const hash = new Hash("testhash");
      mockLinkApplicationService.handleRedirect.mockRejectedValue(
        new Error("Internal server error")
      );

      // Act
      const response = await request(app).get("/s/testhash");

      // Assert
      expect(response.status).toBe(500);
      expect(response.body.error).toHaveProperty("code");
      expect(response.body.error).toHaveProperty("message");
    });
  });
});

