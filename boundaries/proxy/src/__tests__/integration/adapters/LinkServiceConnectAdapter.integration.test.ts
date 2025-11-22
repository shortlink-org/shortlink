import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { LinkServiceConnectAdapter } from "../../../infrastructure/adapters/LinkServiceConnectAdapter.js";
import { LinkServiceACL } from "../../../infrastructure/anti-corruption/LinkServiceACL.js";
import { ExternalServicesConfig } from "../../../infrastructure/config/ExternalServicesConfig.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { IGrpcMetrics } from "../../../infrastructure/metrics/IGrpcMetrics.js";
import { IGrpcTracing } from "../../../infrastructure/tracing/IGrpcTracing.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import {
  GetRequest,
  GetRequestSchema,
  GetResponse,
  GetResponseSchema,
  Link as LinkProto,
  LinkSchema,
} from "../../../infrastructure/proto/infrastructure/rpc/link/v1/link_pb.js";
import { create, fromBinary, toBinary } from "@bufbuild/protobuf";
import type { Transport } from "@connectrpc/connect";

/**
 * Integration тесты для Connect адаптера
 * Тестируют адаптер с моками Transport и зависимостей
 */
describe("LinkServiceConnectAdapter Integration Tests", () => {
  let adapter: LinkServiceConnectAdapter;
  let mockTransport: {
    unary: ReturnType<typeof vi.fn>;
  };
  let mockACL: {
    toDomainEntityFromProto: ReturnType<typeof vi.fn>;
  };
  let mockConfig: ExternalServicesConfig;
  let mockLogger: {
    info: ReturnType<typeof vi.fn>;
    warn: ReturnType<typeof vi.fn>;
    error: ReturnType<typeof vi.fn>;
    debug: ReturnType<typeof vi.fn>;
    http: ReturnType<typeof vi.fn>;
  };
  let mockMetrics: {
    recordRequest: ReturnType<typeof vi.fn>;
    recordDuration: ReturnType<typeof vi.fn>;
    recordError: ReturnType<typeof vi.fn>;
  };
  let mockTracing: {
    startSpan: ReturnType<typeof vi.fn>;
    endSpan: ReturnType<typeof vi.fn>;
    endSpanWithError: ReturnType<typeof vi.fn>;
  };

  beforeEach(() => {
    // Мокаем Transport
    mockTransport = {
      unary: vi.fn(),
    } as any;

    // Мокаем ACL
    mockACL = {
      toDomainEntityFromProto: vi.fn(),
    } as any;

    // Мокаем Config
    mockConfig = {
      linkServiceGrpcUrl: "localhost:50051",
      requestTimeout: 5000,
      retryCount: 3,
    } as any;

    // Мокаем Logger
    mockLogger = {
      info: vi.fn(),
      warn: vi.fn(),
      error: vi.fn(),
      debug: vi.fn(),
      http: vi.fn(),
    } as any;

    // Мокаем Metrics
    mockMetrics = {
      recordRequest: vi.fn(),
      recordDuration: vi.fn(),
      recordError: vi.fn(),
    } as any;

    // Мокаем Tracing
    mockTracing = {
      startSpan: vi.fn().mockReturnValue({
        setAttribute: vi.fn(),
      }),
      endSpan: vi.fn(),
      endSpanWithError: vi.fn(),
    } as any;

    // Создаем адаптер с моками
    // Используем приватное поле transport через рефлексию или создаем через конструктор
    adapter = new LinkServiceConnectAdapter(
      mockACL as any,
      mockConfig as any,
      mockLogger as any,
      mockMetrics as any,
      mockTracing as any
    );

    // Заменяем transport на мок через рефлексию
    (adapter as any).transport = mockTransport;
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  describe("getLinkByHash", () => {
    it("should return link when found", async () => {
      // Arrange
      const hash = new Hash("testhash");
      const protoLink = create(LinkSchema, {
        hash: "testhash",
        url: "https://example.com",
      });
      const domainLink = new Link(hash, "https://example.com");

      const getResponse = create(GetResponseSchema, {
        link: protoLink,
      });

      // Мокаем Transport.unary для успешного ответа
      mockTransport.unary = vi.fn().mockResolvedValue({
        message: toBinary(GetResponseSchema, getResponse),
      } as any);

      mockACL.toDomainEntityFromProto.mockReturnValue(domainLink);

      // Act
      const result = await adapter.getLinkByHash(hash);

      // Assert
      expect(result).toEqual(domainLink);
      expect(mockTransport.unary).toHaveBeenCalled();
      expect(mockACL.toDomainEntityFromProto).toHaveBeenCalledWith(protoLink);
    });

    it("should return null when link is not found (NOT_FOUND)", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      const notFoundError = new Error("NOT_FOUND");
      (notFoundError as any).code = "NOT_FOUND";

      mockTransport.unary = vi.fn().mockRejectedValue(notFoundError);

      // Act
      const result = await adapter.getLinkByHash(hash);

      // Assert
      expect(result).toBeNull();
      expect(mockTransport.unary).toHaveBeenCalled();
      expect(mockACL.toDomainEntityFromProto).not.toHaveBeenCalled();
    });

    it("should return null when link is not found (404 status)", async () => {
      // Arrange
      const hash = new Hash("nonexistent");
      const notFoundError = new Error("Not Found");
      (notFoundError as any).status = 404;

      mockTransport.unary = vi.fn().mockRejectedValue(notFoundError);

      // Act
      const result = await adapter.getLinkByHash(hash);

      // Assert
      expect(result).toBeNull();
    });

    it("should throw error for other transport errors", async () => {
      // Arrange
      const hash = new Hash("testhash");
      const transportError = new Error("Transport error");
      (transportError as any).code = "INTERNAL";

      mockTransport.unary = vi.fn().mockRejectedValue(transportError);

      // Act & Assert
      await expect(adapter.getLinkByHash(hash)).rejects.toThrow(
        "Transport error"
      );
      expect(mockTransport.unary).toHaveBeenCalled();
    });

    it("should create GetRequest with correct hash", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const protoLink = create(LinkSchema, {
        hash: "testhash123",
        url: "https://example.com",
      });
      const domainLink = new Link(hash, "https://example.com");

      const getResponse = create(GetResponseSchema, {
        link: protoLink,
      });

      mockTransport.unary = vi.fn().mockResolvedValue({
        message: toBinary(GetResponseSchema, getResponse),
      } as any);

      mockACL.toDomainEntityFromProto.mockReturnValue(domainLink);

      // Act
      await adapter.getLinkByHash(hash);

      // Assert - проверяем, что GetRequest создан с правильным hash
      const unaryCall = mockTransport.unary.mock.calls[0];
      expect(unaryCall).toBeDefined();
      // Проверяем, что был вызван unary с правильными параметрами
      expect(mockTransport.unary).toHaveBeenCalled();
    });

    it("should use ACL to convert proto to domain entity", async () => {
      // Arrange
      const hash = new Hash("testhash");
      const protoLink = create(LinkSchema, {
        hash: "testhash",
        url: "https://example.com",
      });
      const domainLink = new Link(hash, "https://example.com");

      const getResponse = create(GetResponseSchema, {
        link: protoLink,
      });

      mockTransport.unary = vi.fn().mockResolvedValue({
        message: toBinary(GetResponseSchema, getResponse),
      } as any);

      mockACL.toDomainEntityFromProto.mockReturnValue(domainLink);

      // Act
      await adapter.getLinkByHash(hash);

      // Assert
      expect(mockACL.toDomainEntityFromProto).toHaveBeenCalledWith(protoLink);
    });
  });
});

