import { createContainer, asClass, asValue, asFunction, InjectionMode, AwilixContainer } from "awilix";
import type { FastifyInstance } from "fastify";

// Domain Services
import { LinkDomainService } from "../proxy/domain/services/LinkDomainService.js";

// Application Layer
import { GetLinkByHashUseCase } from "../proxy/application/use-cases/GetLinkByHashUseCase.js";
import { PublishEventUseCase } from "../proxy/application/use-cases/PublishEventUseCase.js";
import { LinkApplicationService } from "../proxy/application/services/LinkApplicationService.js";

// Application Layer - Pipeline
import {
  UseCasePipeline,
  LoggingInterceptor,
  MetricsInterceptor,
  AuthorizationInterceptor,
  DefaultAuthorizationChecker,
  type IAuthorizationChecker,
} from "../proxy/application/pipeline/index.js";

// Infrastructure Layer - Repositories
import { LinkServiceRepository } from "../proxy/infrastructure/repositories/LinkServiceRepository.js";
import type { ILinkRepository } from "../proxy/domain/repositories/ILinkRepository.js";

// Infrastructure Layer - Adapters & ACL
import { LinkServiceACL } from "../proxy/infrastructure/anti-corruption/LinkServiceACL.js";
import {
  type ILinkServiceAdapter,
  LinkServiceConnectAdapter,
} from "../proxy/infrastructure/adapters/index.js";

// Infrastructure Layer - Messaging
import { AMQPEventPublisher } from "../proxy/infrastructure/messaging/AMQPEventPublisher.js";
import type { IEventPublisher } from "../proxy/application/use-cases/PublishEventUseCase.js";

// Infrastructure Layer - Logging
import type { ILogger } from "../infrastructure/logging/ILogger.js";
import { WinstonLogger } from "../infrastructure/logging/WinstonLogger.js";

// Infrastructure Layer - Configuration
import { ExternalServicesConfig } from "../infrastructure/config/ExternalServicesConfig.js";
import { CacheConfig } from "../infrastructure/config/CacheConfig.js";
import { AppConfig } from "../infrastructure/config/index.js";

// Infrastructure Layer - Messaging
import { RabbitMQMessageBus } from "../proxy/infrastructure/messaging/RabbitMQMessageBus.js";
import type { IMessageBus } from "../proxy/domain/interfaces/IMessageBus.js";

// Infrastructure Layer - Cache
import {
  type ILinkCache,
  RedisLinkCache,
} from "../proxy/infrastructure/cache/RedisLinkCache.js";

// Infrastructure Layer - Metrics & Tracing
import {
  type IGrpcMetrics,
  OpenTelemetryGrpcMetrics,
} from "../proxy/infrastructure/metrics/index.js";
import {
  type IGrpcTracing,
  OpenTelemetryGrpcTracing,
} from "../proxy/infrastructure/tracing/index.js";

// Controllers
import { ProxyController } from "../proxy/infrastructure/http/fastify/controllers/ProxyController.js";
import { MetricsController } from "../proxy/infrastructure/http/fastify/controllers/MetricsController.js";

/**
 * Container dependencies interface
 */
export interface ContainerDependencies {
  // Config
  appConfig: AppConfig;
  externalServicesConfig: ExternalServicesConfig;
  cacheConfig: CacheConfig;

  // Domain
  linkDomainService: LinkDomainService;

  // Application
  getLinkByHashUseCase: GetLinkByHashUseCase;
  publishEventUseCase: PublishEventUseCase;
  linkApplicationService: LinkApplicationService;
  useCasePipeline: UseCasePipeline;
  loggingInterceptor: LoggingInterceptor;
  metricsInterceptor: MetricsInterceptor;
  authorizationInterceptor: AuthorizationInterceptor;
  authorizationChecker: IAuthorizationChecker;

  // Infrastructure
  logger: ILogger;
  linkRepository: ILinkRepository;
  linkServiceACL: LinkServiceACL;
  linkServiceAdapter: ILinkServiceAdapter;
  eventPublisher: IEventPublisher;
  messageBus: IMessageBus;
  linkCache: ILinkCache;
  grpcMetrics: IGrpcMetrics;
  grpcTracing: IGrpcTracing;

  // Controllers
  proxyController: ProxyController;
  metricsController: MetricsController;
}

/**
 * Creates and configures the Awilix dependency injection container.
 * This is the Composition Root following Clean Architecture principles.
 *
 * @returns Configured Awilix container with all dependencies
 */
export function createDIContainer(): AwilixContainer<ContainerDependencies> {
  const container = createContainer<ContainerDependencies>({
    injectionMode: InjectionMode.PROXY,
  });

  // ============================================================================
  // CONFIG (asValue - no dependencies)
  // ============================================================================
  container.register({
    appConfig: asValue(new AppConfig()),
    externalServicesConfig: asClass(ExternalServicesConfig).singleton(),
    cacheConfig: asClass(CacheConfig).singleton(),
  });

  // ============================================================================
  // LOGGING (needed early)
  // ============================================================================
  container.register({
    logger: asClass(WinstonLogger).singleton(),
  });

  // ============================================================================
  // INFRASTRUCTURE - CORE
  // ============================================================================
  container.register({
    // Metrics & Tracing (no dependencies)
    grpcMetrics: asClass(OpenTelemetryGrpcMetrics).singleton(),
    grpcTracing: asClass(OpenTelemetryGrpcTracing).singleton(),

    // Anti-corruption Layer (no dependencies)
    linkServiceACL: asClass(LinkServiceACL).singleton(),

    // Adapters - depends on ACL, config, logger, metrics, tracing
    // Parameter names: linkServiceACL, externalServicesConfig, logger, grpcMetrics, grpcTracing
    linkServiceAdapter: asClass(LinkServiceConnectAdapter).singleton(),

    // Cache - depends on logger, cacheConfig
    // Parameter names: logger, cacheConfig
    linkCache: asClass(RedisLinkCache)
      .inject(() => ({
        logger: container.resolve("logger"),
        cacheConfig: container.resolve("cacheConfig"),
      }))
      .singleton(),

    // Message Bus - depends on logger only
    // Parameter names: logger
    messageBus: asClass(RabbitMQMessageBus).singleton(),

    // Event Publisher - depends on messageBus, logger
    // Parameter names: messageBus, logger
    eventPublisher: asClass(AMQPEventPublisher).singleton(),

    // Repository - depends on adapter, cache, logger
    // Parameter names: linkServiceAdapter, linkCache, logger
    linkRepository: asClass(LinkServiceRepository).singleton(),
  });

  // ============================================================================
  // DOMAIN SERVICES
  // ============================================================================
  container.register({
    linkDomainService: asClass(LinkDomainService).singleton(),
  });

  // ============================================================================
  // APPLICATION LAYER
  // ============================================================================
  container.register({
    // Authorization
    authorizationChecker: asClass(DefaultAuthorizationChecker).singleton(),

    // Pipeline Interceptors
    loggingInterceptor: asClass(LoggingInterceptor).singleton(),
    metricsInterceptor: asClass(MetricsInterceptor).singleton(),
    authorizationInterceptor: asClass(AuthorizationInterceptor).singleton(),

    // Pipeline
    useCasePipeline: asClass(UseCasePipeline).singleton(),

    // Use Cases
    getLinkByHashUseCase: asClass(GetLinkByHashUseCase).singleton(),
    publishEventUseCase: asClass(PublishEventUseCase).singleton(),

    // Application Services
    linkApplicationService: asClass(LinkApplicationService).singleton(),
  });

  // ============================================================================
  // CONTROLLERS (HTTP Layer)
  // ============================================================================
  container.register({
    proxyController: asClass(ProxyController).singleton(),
    metricsController: asClass(MetricsController).singleton(),
  });

  return container;
}

/**
 * Creates a request-scoped container for per-request isolation.
 * Each request gets its own container that can override dependencies if needed.
 *
 * @param parentContainer - Parent container to create scope from
 * @returns Request-scoped container
 */
export function createRequestScope(
  parentContainer: AwilixContainer<ContainerDependencies>
): AwilixContainer<ContainerDependencies> {
  return parentContainer.createScope();
}

