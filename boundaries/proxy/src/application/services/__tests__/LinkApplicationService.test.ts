import { describe, it, expect, beforeEach, vi } from "vitest";
import { LinkApplicationService, HandleRedirectRequest } from "../LinkApplicationService.js";
import { GetLinkByHashUseCase } from "../../use-cases/GetLinkByHashUseCase.js";
import { PublishEventUseCase } from "../../use-cases/PublishEventUseCase.js";
import { UseCasePipeline } from "../../pipeline/UseCasePipeline.js";
import { LoggingInterceptor } from "../../pipeline/LoggingInterceptor.js";
import { MetricsInterceptor } from "../../pipeline/MetricsInterceptor.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import { LinkNotFoundError } from "../../../domain/exceptions/index.js";
import { LinkEvents } from "../../../domain/events/index.js";
import { ok, err } from "neverthrow";

describe("LinkApplicationService", () => {
  let service: LinkApplicationService;
  let mockGetLinkByHashUseCase: GetLinkByHashUseCase;
  let mockPublishEventUseCase: PublishEventUseCase;
  let mockLogger: ILogger;
  let mockPipeline: UseCasePipeline;
  let mockLoggingInterceptor: LoggingInterceptor;
  let mockMetricsInterceptor: MetricsInterceptor;

  beforeEach(() => {
    mockGetLinkByHashUseCase = {
      execute: vi.fn(),
    } as unknown as GetLinkByHashUseCase;

    mockPublishEventUseCase = {
      execute: vi.fn(),
    } as unknown as PublishEventUseCase;

    mockLogger = {
      debug: vi.fn(),
      info: vi.fn(),
      warn: vi.fn(),
      error: vi.fn(),
    } as unknown as ILogger;

    mockPipeline = {
      execute: vi.fn(),
    } as unknown as UseCasePipeline;

    mockLoggingInterceptor = {} as LoggingInterceptor;
    mockMetricsInterceptor = {} as MetricsInterceptor;

    service = new LinkApplicationService(
      mockGetLinkByHashUseCase,
      mockPublishEventUseCase,
      mockLogger,
      mockPipeline,
      mockLoggingInterceptor,
      mockMetricsInterceptor
    );
  });

  describe("getByHash", () => {
    it("should call GetLinkByHashUseCase with correct request", async () => {
      // Arrange
      const hash = "testhash123";
      const mockResponse = {
        link: new Link(new Hash("testhash123"), "https://example.com"),
      };
      vi.mocked(mockGetLinkByHashUseCase.execute).mockResolvedValue(mockResponse);

      // Act
      const result = await service.getByHash(hash);

      // Assert
      expect(mockGetLinkByHashUseCase.execute).toHaveBeenCalledWith({
        hash,
      });
      expect(result).toEqual(mockResponse);
    });
  });

  describe("handleRedirect", () => {
    it("should successfully handle redirect and publish event", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");
      const request: HandleRedirectRequest = { hash };

      const mockLinkResponse = { link };
      vi.mocked(mockGetLinkByHashUseCase.execute).mockResolvedValue(mockLinkResponse);
      vi.mocked(mockPipeline.execute).mockResolvedValue(ok({}));

      // Act
      const result = await service.handleRedirect(request);

      // Assert
      expect(result.isOk()).toBe(true);
      if (result.isOk()) {
        expect(result.value.link).toEqual(link);
      }
      expect(mockPipeline.execute).toHaveBeenCalledWith(
        mockPublishEventUseCase,
        expect.objectContaining({
          event: expect.objectContaining({
            type: "LinkRedirected",
            hash,
            link,
          }),
        }),
        expect.any(Array)
      );
    });

    it("should return error when link not found", async () => {
      // Arrange
      const hash = new Hash("nonexistent123");
      const request: HandleRedirectRequest = { hash };
      const notFoundError = new LinkNotFoundError(hash);

      vi.mocked(mockGetLinkByHashUseCase.execute).mockRejectedValue(notFoundError);

      // Act
      const result = await service.handleRedirect(request);

      // Assert
      expect(result.isErr()).toBe(true);
      if (result.isErr()) {
        expect(result.error).toBeInstanceOf(LinkNotFoundError);
      }
      expect(mockPipeline.execute).not.toHaveBeenCalled();
    });

    it("should log warning when event publishing fails but continue", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");
      const request: HandleRedirectRequest = { hash };

      const mockLinkResponse = { link };
      vi.mocked(mockGetLinkByHashUseCase.execute).mockResolvedValue(mockLinkResponse);
      vi.mocked(mockPipeline.execute).mockRejectedValue(new Error("Publish failed"));

      // Act
      const result = await service.handleRedirect(request);

      // Assert
      expect(result.isOk()).toBe(true);
      expect(mockLogger.warn).toHaveBeenCalledWith(
        "Failed to publish redirect event",
        expect.objectContaining({
          hash: hash.value,
          error: expect.any(Error),
        })
      );
    });

    it("should throw unexpected errors", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const request: HandleRedirectRequest = { hash };
      const unexpectedError = new Error("Unexpected error");

      vi.mocked(mockGetLinkByHashUseCase.execute).mockRejectedValue(unexpectedError);

      // Act & Assert
      await expect(service.handleRedirect(request)).rejects.toThrow("Unexpected error");
    });
  });
});

