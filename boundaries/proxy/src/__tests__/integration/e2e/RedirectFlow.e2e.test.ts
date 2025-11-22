import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import type { FastifyInstance } from "fastify";
import { createTestServer } from "../helpers/testServer.js";
import { LinkApplicationService } from "../../../application/services/LinkApplicationService.js";
import { GetLinkByHashUseCase } from "../../../application/use-cases/GetLinkByHashUseCase.js";
import { PublishEventUseCase } from "../../../application/use-cases/PublishEventUseCase.js";
import { LinkServiceRepository } from "../../../infrastructure/repositories/LinkServiceRepository.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import type { ILogger } from "../../../infrastructure/logging/ILogger.js";
import {
  UseCasePipeline,
  LoggingInterceptor,
  MetricsInterceptor,
} from "../../../application/pipeline/index.js";
import type { ILinkCache } from "../../../infrastructure/cache/RedisLinkCache.js";
import type { IEventPublisher } from "../../../application/use-cases/PublishEventUseCase.js";
import type { ILinkServiceAdapter } from "../../../infrastructure/adapters/ILinkServiceAdapter.js";
import type { DomainEvent } from "../../../domain/events/index.js";

/**
 * End-to-end интеграционные тесты для полного flow
 * Тестируют весь путь: Controller → Application Service → Use Cases → Repository → Adapter
 */
describe("Redirect Flow E2E Integration Tests", () => {
  let app: FastifyInstance;
  let publishMock: ReturnType<typeof vi.fn<(event: DomainEvent) => Promise<void>>>;
  let getLinkByHashMock: ReturnType<typeof vi.fn<(hash: Hash) => Promise<Link | null>>>;
  let cacheGetMock: ReturnType<typeof vi.fn<(hash: Hash) => Promise<Link | null | undefined>>>;
  let cacheSetPositiveMock: ReturnType<typeof vi.fn<(hash: Hash, link: Link) => Promise<void>>>;
  let cacheSetNegativeMock: ReturnType<typeof vi.fn<(hash: Hash) => Promise<void>>>;
  let cacheClearMock: ReturnType<typeof vi.fn<(hash: Hash) => Promise<void>>>;
  let loggerInfoMock: ReturnType<typeof vi.fn<(message: string, meta?: any) => void>>;
  let loggerWarnMock: ReturnType<typeof vi.fn<(message: string, meta?: any) => void>>;
  let loggerErrorMock: ReturnType<typeof vi.fn<(message: string, error?: any, meta?: any) => void>>;
  let loggerDebugMock: ReturnType<typeof vi.fn<(message: string, meta?: any) => void>>;
  let loggerHttpMock: ReturnType<typeof vi.fn<(message: string) => void>>;
  let mockEventPublisher: IEventPublisher;
  let mockLinkServiceAdapter: ILinkServiceAdapter;
  let mockLinkCache: ILinkCache;
  let mockLogger: ILogger;

  beforeEach(async () => {
    // Mock Event Publisher
    publishMock = vi.fn<(event: DomainEvent) => Promise<void>>().mockResolvedValue(undefined);
    mockEventPublisher = {
      publish: publishMock,
    };

    // Mock Adapter
    getLinkByHashMock = vi.fn<(hash: Hash) => Promise<Link | null>>();
    mockLinkServiceAdapter = {
      getLinkByHash: getLinkByHashMock,
    } as any;

    cacheGetMock = vi.fn<(hash: Hash) => Promise<Link | null | undefined>>().mockResolvedValue(undefined);
    cacheSetPositiveMock = vi.fn<(hash: Hash, link: Link) => Promise<void>>().mockResolvedValue(undefined);
    cacheSetNegativeMock = vi.fn<(hash: Hash) => Promise<void>>().mockResolvedValue(undefined);
    cacheClearMock = vi.fn<(hash: Hash) => Promise<void>>().mockResolvedValue(undefined);
    mockLinkCache = {
      get: cacheGetMock,
      setPositive: cacheSetPositiveMock,
      setNegative: cacheSetNegativeMock,
      clear: cacheClearMock,
    } as any;

    // Mock Logger
    loggerInfoMock = vi.fn<(message: string, meta?: any) => void>();
    loggerWarnMock = vi.fn<(message: string, meta?: any) => void>();
    loggerErrorMock = vi.fn<(message: string, error?: any, meta?: any) => void>();
    loggerDebugMock = vi.fn<(message: string, meta?: any) => void>();
    loggerHttpMock = vi.fn<(message: string) => void>();
    mockLogger = {
      info: loggerInfoMock,
      warn: loggerWarnMock,
      error: loggerErrorMock,
      debug: loggerDebugMock,
      http: loggerHttpMock,
    } as any;

    // Create Fastify server with mocked dependencies
    // Real services: LinkServiceRepository, GetLinkByHashUseCase, PublishEventUseCase, LinkApplicationService
    app = await createTestServer({
      linkServiceAdapter: mockLinkServiceAdapter,
      linkCache: mockLinkCache,
      eventPublisher: mockEventPublisher,
      logger: mockLogger,
      // Real implementations will be used (from container)
      // They will use the mocked dependencies above
    } as any);
  });

  afterEach(async () => {
    vi.clearAllMocks();
    await app.close();
  });

  describe("Full redirect flow", () => {
    it("should complete full redirect flow successfully", async () => {
      // Arrange
      const hash = new Hash("testhash");
      const link = new Link(hash, "https://example.com");

      // Настраиваем моки для полного flow
      // LinkServiceRepository использует LinkServiceAdapter, поэтому мокаем адаптер
      getLinkByHashMock.mockResolvedValue(link);

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/testhash",
      });

      // Assert - проверяем весь flow
      // 1. Controller получил запрос
      expect(response.statusCode).toBe(301);
      expect(response.headers.location).toBe("https://example.com");

      // 2. LinkRepository был вызван через LinkServiceRepository → LinkServiceAdapter
      expect(getLinkByHashMock).toHaveBeenCalledWith(
        expect.objectContaining({ value: "testhash" })
      );

      // 3. Event Publisher был вызван для публикации события
      expect(publishMock).toHaveBeenCalled();

      // 4. Статистика собирается через eBPF, не требует записи в БД
    });

    it("should handle link not found in full flow", async () => {
      // Arrange
      // LinkServiceRepository использует LinkServiceAdapter
      getLinkByHashMock.mockResolvedValue(null);

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/nonexistent",
      });

      // Assert
      expect(response.statusCode).toBe(404);
      // В E2E тесте LinkServiceRepository использует LinkServiceAdapter
      // Проверяем, что адаптер был вызван
      expect(getLinkByHashMock).toHaveBeenCalled();
    });

    it("should handle adapter errors gracefully", async () => {
      // Arrange
      getLinkByHashMock.mockRejectedValue(new Error("Adapter error"));

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/testhash",
      });

      // Assert
      // Ошибка адаптера должна быть обработана error handler middleware
      expect(response.statusCode).toBeGreaterThanOrEqual(400);
      expect(getLinkByHashMock).toHaveBeenCalled();
    });

    it("should continue redirect even if event publishing fails", async () => {
      // Arrange
      const link = new Link(new Hash("testhash"), "https://example.com");

      // LinkServiceRepository использует LinkServiceAdapter
      getLinkByHashMock.mockResolvedValue(link);
      publishMock.mockRejectedValue(new Error("Event publishing error"));

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/testhash",
      });

      // Assert - редирект должен выполниться, даже если событие не опубликовано
      expect(response.statusCode).toBe(301);
      expect(response.headers.location).toBe("https://example.com");
      // Логгер должен зафиксировать предупреждение
      expect(loggerWarnMock).toHaveBeenCalled();
    });
  });
});
