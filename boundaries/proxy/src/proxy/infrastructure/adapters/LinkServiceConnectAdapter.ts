import { injectable, inject } from "inversify";
import { createConnectTransport } from "@connectrpc/connect-node";
import type { Transport } from "@connectrpc/connect";
import { LinkServiceACL } from "../anti-corruption/LinkServiceACL.js";
import { ILinkServiceAdapter } from "./ILinkServiceAdapter.js";
import { Link } from "../../domain/entities/Link.js";
import { Hash } from "../../domain/entities/Hash.js";
import { ExternalServicesConfig } from "../../../infrastructure/config/ExternalServicesConfig.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { IGrpcMetrics } from "../metrics/IGrpcMetrics.js";
import { IGrpcTracing } from "../tracing/IGrpcTracing.js";
import TYPES from "../../../types.js";
import {
  createLoggingInterceptor,
  createRetryInterceptor,
  createMetricsInterceptor,
  createTracingInterceptor,
} from "./connect/interceptors/index.js";
// Импортируем типы из сгенерированных proto файлов
import {
  GetRequest,
  GetRequestSchema,
  GetResponse,
  GetResponseSchema,
} from "../../../proto/infrastructure/rpc/link/v1/link_pb.js";
import { create, fromBinary, toBinary } from "@bufbuild/protobuf";

/**
 * Connect адаптер для получения ссылок из Link Service
 * Использует @connectrpc/connect через unary transport API
 * Работает с типами из @bufbuild/protobuf без генерации connect-es файлов
 */
@injectable()
export class LinkServiceConnectAdapter implements ILinkServiceAdapter {
  private readonly transport: Transport;

  constructor(
    @inject(Symbol.for("LinkServiceACL")) private readonly acl: LinkServiceACL,
    @inject(Symbol.for("ExternalServicesConfig"))
    private readonly config: ExternalServicesConfig,
    @inject(TYPES.INFRASTRUCTURE.Logger) private readonly logger: ILogger,
    @inject(TYPES.INFRASTRUCTURE.GrpcMetrics)
    private readonly metrics: IGrpcMetrics,
    @inject(TYPES.INFRASTRUCTURE.GrpcTracing)
    private readonly tracing: IGrpcTracing
  ) {
    // Создаем Connect interceptors
    const interceptors = [
      // Порядок важен: сначала трейсинг, потом метрики, потом retry, потом логирование
      createTracingInterceptor(this.tracing),
      createMetricsInterceptor(this.metrics),
      createRetryInterceptor(this.logger, {
        maxAttempts: this.config.retryCount,
        initialDelayMs: 100,
        maxDelayMs: 5000,
        backoffMultiplier: 2,
      }),
      createLoggingInterceptor(this.logger),
    ];

    // Создаем Connect transport один раз при инициализации с interceptors
    this.transport = createConnectTransport({
      baseUrl: `http://${this.config.linkServiceGrpcUrl}`,
      httpVersion: "2",
      interceptors,
    });
  }

  async getLinkByHash(hash: Hash): Promise<Link | null> {
    const hashValue = hash.value;

    // Создаем GetRequest используя create() из @bufbuild/protobuf
    const request = create(GetRequestSchema, {
      hash: hashValue,
    });

    // Выполняем Connect вызов через unary transport API
    // Полностью типобезопасно через @bufbuild/protobuf типы
    // Connect 2.x сигнатура: unary(method, signal, timeoutMs, header, message, contextValues)
    // Interceptors автоматически обработают логирование, метрики, трейсинг и retry
    const methodDesc = {
      service: {
        typeName: "infrastructure.rpc.link.v1.LinkService",
      },
      method: {
        name: "Get",
        kind: "unary" as const,
        I: GetRequestSchema,
        O: GetResponseSchema,
      },
    } as any; // Временный any для обхода проблем с типизацией Connect 2.x

    const signal = AbortSignal.timeout(this.config.requestTimeout);
    const timeoutMs = this.config.requestTimeout;

    // Преобразуем GetRequest в бинарный формат для Connect
    const requestBinary = toBinary(GetRequestSchema, request);

    try {
      const response = await this.transport.unary(
        methodDesc,
        signal,
        timeoutMs,
        undefined, // headers (опционально)
        requestBinary as any, // Connect ожидает Record, но принимает Uint8Array
        undefined // contextValues (опционально)
      );

      // response - это UnaryResponse, который содержит message в бинарном формате
      // Преобразуем бинарный ответ обратно в GetResponse
      const responseBinary = (response as any).message || response;
      const getResponse = fromBinary(GetResponseSchema, responseBinary as Uint8Array);

      if (!getResponse || !getResponse.link) {
        // NOT_FOUND обрабатывается interceptors (метрики и трейсинг)
        return null;
      }

      // Преобразуем protobuf Link в доменную сущность через ACL
      // Используем Link из domain/link/v1/link_pb (который уже использует @bufbuild/protobuf)
      const protoLink = getResponse.link;
      const domainLink = this.acl.toDomainEntityFromProto(protoLink);

      return domainLink;
    } catch (error: any) {
      // Обрабатываем Connect ошибки
      // Interceptors уже обработали логирование, метрики и трейсинг
      // Здесь только бизнес-логика обработки ошибок

      // Проверяем NOT_FOUND
      if (
        error?.code === "NOT_FOUND" ||
        error?.code === 5 ||
        error?.status === 404
      ) {
        return null;
      }

      // Все остальные ошибки пробрасываем дальше
      // Interceptors уже залогировали и записали метрики
      throw error;
    }
  }
}

