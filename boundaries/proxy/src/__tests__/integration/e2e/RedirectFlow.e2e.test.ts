import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import request from "supertest";
import { Container } from "inversify";
import * as express from "express";
import { createTestServer } from "../helpers/testServer.js";
import { LinkApplicationService } from "../../../proxy/application/services/LinkApplicationService.js";
import { GetLinkByHashUseCase } from "../../../proxy/application/use-cases/GetLinkByHashUseCase.js";
import { PublishEventUseCase } from "../../../proxy/application/use-cases/PublishEventUseCase.js";
import { ILinkRepository } from "../../../proxy/domain/repositories/ILinkRepository.js";
import { IEventPublisher } from "../../../proxy/application/use-cases/PublishEventUseCase.js";
import { ILinkServiceAdapter } from "../../../proxy/infrastructure/adapters/ILinkServiceAdapter.js";
import { LinkServiceRepository } from "../../../proxy/infrastructure/repositories/LinkServiceRepository.js";
import { Hash } from "../../../proxy/domain/entities/Hash.js";
import { Link } from "../../../proxy/domain/entities/Link.js";
import TYPES from "../../../types.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import {
  UseCasePipeline,
  LoggingInterceptor,
  MetricsInterceptor,
} from "../../../proxy/application/pipeline/index.js";
import { ILinkCache } from "../../../proxy/infrastructure/cache/RedisLinkCache.js";

/**
 * End-to-end интеграционные тесты для полного flow
 * Тестируют весь путь: Controller → Application Service → Use Cases → Repository → Adapter
 */
describe("Redirect Flow E2E Integration Tests", () => {
  let app: express.Application;
  let container: Container;
  let publishMock: ReturnType<typeof vi.fn>;
  let getLinkByHashMock: ReturnType<typeof vi.fn>;
  let cacheGetMock: ReturnType<typeof vi.fn>;
  let cacheSetPositiveMock: ReturnType<typeof vi.fn>;
  let cacheSetNegativeMock: ReturnType<typeof vi.fn>;
  let cacheClearMock: ReturnType<typeof vi.fn>;
  let loggerInfoMock: ReturnType<typeof vi.fn>;
  let loggerWarnMock: ReturnType<typeof vi.fn>;
  let loggerErrorMock: ReturnType<typeof vi.fn>;
  let loggerDebugMock: ReturnType<typeof vi.fn>;
  let loggerHttpMock: ReturnType<typeof vi.fn>;
  let mockEventPublisher: IEventPublisher;
  let mockLinkServiceAdapter: ILinkServiceAdapter;
  let mockLinkCache: ILinkCache;
  let mockLogger: ILogger;

  beforeEach(async () => {
    // Создаем новый контейнер для каждого теста
    container = new Container();

    // Мокаем Event Publisher
    publishMock = vi.fn().mockResolvedValue(undefined);
    mockEventPublisher = {
      publish: publishMock,
    };

    // Мокаем Adapter
    getLinkByHashMock = vi.fn();
    mockLinkServiceAdapter = {
      getLinkByHash: getLinkByHashMock,
    };

    cacheGetMock = vi.fn().mockResolvedValue(undefined);
    cacheSetPositiveMock = vi.fn().mockResolvedValue(undefined);
    cacheSetNegativeMock = vi.fn().mockResolvedValue(undefined);
    cacheClearMock = vi.fn().mockResolvedValue(undefined);
    mockLinkCache = {
      get: cacheGetMock,
      setPositive: cacheSetPositiveMock,
      setNegative: cacheSetNegativeMock,
      clear: cacheClearMock,
    };

    // Мокаем Logger
    loggerInfoMock = vi.fn();
    loggerWarnMock = vi.fn();
    loggerErrorMock = vi.fn();
    loggerDebugMock = vi.fn();
    loggerHttpMock = vi.fn();
    mockLogger = {
      info: loggerInfoMock,
      warn: loggerWarnMock,
      error: loggerErrorMock,
      debug: loggerDebugMock,
      http: loggerHttpMock,
    };

    // Биндим реальный LinkServiceRepository, который использует адаптер
    // Это позволяет проверить полный flow: Repository → Adapter
    container
      .bind<ILinkRepository>(TYPES.REPOSITORY.LinkRepository)
      .to(LinkServiceRepository)
      .inSingletonScope();

    container
      .bind<IEventPublisher>(TYPES.INFRASTRUCTURE.EventPublisher)
      .toConstantValue(mockEventPublisher);

    container
      .bind<ILinkServiceAdapter>(TYPES.INFRASTRUCTURE.LinkServiceAdapter)
      .toConstantValue(mockLinkServiceAdapter);

    container
      .bind<ILinkCache>(TYPES.INFRASTRUCTURE.LinkCache)
      .toConstantValue(mockLinkCache);

    container
      .bind<ILogger>(TYPES.INFRASTRUCTURE.Logger)
      .toConstantValue(mockLogger);

    // Биндим реальные Use Cases и Application Service
    container
      .bind<GetLinkByHashUseCase>(TYPES.APPLICATION.GetLinkByHashUseCase)
      .to(GetLinkByHashUseCase)
      .inSingletonScope();

    container
      .bind<PublishEventUseCase>(TYPES.APPLICATION.PublishEventUseCase)
      .to(PublishEventUseCase)
      .inSingletonScope();

    container
      .bind<LinkApplicationService>(TYPES.APPLICATION.LinkApplicationService)
      .to(LinkApplicationService)
      .inSingletonScope();

    // Биндим Pipeline и Interceptors
    container
      .bind<UseCasePipeline>(TYPES.APPLICATION.UseCasePipeline)
      .to(UseCasePipeline)
      .inSingletonScope();

    container
      .bind<LoggingInterceptor>(TYPES.APPLICATION.LoggingInterceptor)
      .to(LoggingInterceptor)
      .inSingletonScope();

    container
      .bind<MetricsInterceptor>(TYPES.APPLICATION.MetricsInterceptor)
      .to(MetricsInterceptor)
      .inSingletonScope();

    // Создаем Express приложение с тестовым контейнером
    app = await createTestServer(container);
  });

  afterEach(() => {
    vi.clearAllMocks();
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
      const response = await request(app).get("/s/testhash");

      // Assert - проверяем весь flow
      // 1. Controller получил запрос
      expect(response.status).toBe(301);
      expect(response.headers.location).toBe("https://example.com");

      // 2. Application Service был вызван
      // (проверяем через моки репозиториев)

      // 3. LinkRepository был вызван через LinkServiceRepository → LinkServiceAdapter
      expect(getLinkByHashMock).toHaveBeenCalledWith(
        expect.objectContaining({ value: "testhash" })
      );

      // 4. Event Publisher был вызван для публикации события
      expect(publishMock).toHaveBeenCalled();

      // 5. Статистика собирается через eBPF, не требует записи в БД
    });

    it("should handle link not found in full flow", async () => {
      // Arrange
      // LinkServiceRepository использует LinkServiceAdapter
      getLinkByHashMock.mockResolvedValue(null);

      // Act
      const response = await request(app).get("/s/nonexistent");

      // Assert
      expect(response.status).toBe(404);
      // В E2E тесте LinkServiceRepository использует LinkServiceAdapter
      // Проверяем, что адаптер был вызван
      expect(getLinkByHashMock).toHaveBeenCalled();
    });

    it("should handle adapter errors gracefully", async () => {
      // Arrange
      getLinkByHashMock.mockRejectedValue(new Error("Adapter error"));

      // Act
      const response = await request(app).get("/s/testhash");

      // Assert
      // Ошибка адаптера должна быть обработана error handler middleware
      expect(response.status).toBeGreaterThanOrEqual(400);
      expect(getLinkByHashMock).toHaveBeenCalled();
    });

    it("should continue redirect even if event publishing fails", async () => {
      // Arrange
      const link = new Link(new Hash("testhash"), "https://example.com");

      // LinkServiceRepository использует LinkServiceAdapter
      getLinkByHashMock.mockResolvedValue(link);
      publishMock.mockRejectedValue(new Error("Event publishing error"));

      // Act
      const response = await request(app).get("/s/testhash");

      // Assert - редирект должен выполниться, даже если событие не опубликовано
      expect(response.status).toBe(301);
      expect(response.headers.location).toBe("https://example.com");
      // Логгер должен зафиксировать предупреждение
      expect(loggerWarnMock).toHaveBeenCalled();
    });
  });
});
