import {
  createContainer,
  asClass,
  asFunction,
  InjectionMode,
  AwilixContainer,
} from "awilix";

// Type imports
import type { AppConfig } from "../infrastructure/config/index.js";
import type { ExternalServicesConfig } from "../infrastructure/config/ExternalServicesConfig.js";
import type { CacheConfig } from "../infrastructure/config/CacheConfig.js";
import type { LinkDomainService } from "../domain/services/LinkDomainService.js";
import type { GetLinkByHashUseCase } from "../application/use-cases/GetLinkByHashUseCase.js";
import type { PublishEventUseCase } from "../application/use-cases/PublishEventUseCase.js";
import type { LinkApplicationService } from "../application/services/LinkApplicationService.js";
import type {
  UseCasePipeline,
  LoggingInterceptor,
  MetricsInterceptor,
  AuthorizationInterceptor,
  IAuthorizationChecker,
} from "../application/pipeline/index.js";
import type { ILinkRepository } from "../domain/repositories/ILinkRepository.js";
import type { LinkServiceACL } from "../infrastructure/anti-corruption/LinkServiceACL.js";
import type { ILinkServiceAdapter } from "../infrastructure/adapters/index.js";
import type { IEventPublisher } from "../application/use-cases/PublishEventUseCase.js";
import type { ILogger } from "../infrastructure/logging/ILogger.js";
import type { IMessageBus } from "../domain/interfaces/IMessageBus.js";
import type { ILinkCache } from "../infrastructure/cache/RedisLinkCache.js";
import type { IGrpcMetrics } from "../infrastructure/metrics/index.js";
import type { IGrpcTracing } from "../infrastructure/tracing/index.js";
import type { ProxyController } from "../infrastructure/http/fastify/controllers/ProxyController.js";
import type { MetricsController } from "../infrastructure/http/fastify/controllers/MetricsController.js";

// Registries
import { CONFIG, INFRA, DOMAIN, APP, CONTROLLERS } from "./index.js";

// Special cases (require .inject())
import { WinstonLogger } from "../infrastructure/logging/WinstonLogger.js";
import { RabbitMQMessageBus } from "../infrastructure/messaging/RabbitMQMessageBus.js";
import { KafkaMessageBus } from "../infrastructure/messaging/KafkaMessageBus.js";
import { RedisLinkCache } from "../infrastructure/cache/RedisLinkCache.js";
import { AMQPEventPublisher } from "../infrastructure/messaging/AMQPEventPublisher.js";
import { KafkaEventPublisher } from "../infrastructure/messaging/KafkaEventPublisher.js";
import { ConfigReader } from "../infrastructure/config/ConfigReader.js";

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
    // Use CLASSIC mode so controller method access doesn't trigger Awilix property resolution.
    injectionMode: InjectionMode.CLASSIC,
  });

  // ============================================================================
  // AUTO-REGISTRATION
  // ============================================================================
  // Register CONFIG (asValue/asClass already configured)
  container.register(CONFIG);

  // Register LOGGING (needed early)
  container.register({
    logger: asClass(WinstonLogger).singleton(),
  });

  // Register INFRA classes (auto-register as singleton classes)
  // Exclude eventPublisher - it needs explicit injection
  Object.entries(INFRA).forEach(([name, clazz]) => {
    if (name !== "eventPublisher") {
      container.register(name, asClass(clazz as any).singleton());
    }
  });

  // Register DOMAIN classes
  Object.entries(DOMAIN).forEach(([name, clazz]) => {
    container.register(name, asClass(clazz as any).singleton());
  });

  // Register APP classes
  Object.entries(APP).forEach(([name, clazz]) => {
    container.register(name, asClass(clazz as any).singleton());
  });

  // Register CONTROLLERS classes
  // ProxyController needs explicit injection to avoid PROXY mode issues
  Object.entries(CONTROLLERS).forEach(([name, clazz]) => {
    if (name === "proxyController") {
      container.register(
        "proxyController",
        asClass(clazz as any)
          .inject(() => ({
            linkApplicationService: container.resolve("linkApplicationService"),
            logger: container.resolve("logger"),
          }))
          .singleton()
      );
    } else {
      container.register(name, asClass(clazz as any).singleton());
    }
  });

  // ============================================================================
  // SPECIAL CASES (require .inject())
  // ============================================================================
  // Cache - depends on logger, cacheConfig
  container.register(
    "linkCache",
    asClass(RedisLinkCache)
      .inject(() => ({
        logger: container.resolve("logger"),
        cacheConfig: container.resolve("cacheConfig"),
      }))
      .singleton()
  );

  // Message Bus - depends on logger only
  // Use asFunction to avoid PROXY mode auto-resolution issues with amqplib's debug dependency
  // Selects implementation based on MQ_TYPE environment variable
  container.register(
    "messageBus",
    asFunction(() => {
      const logger = container.resolve<ILogger>("logger");
      const mqType = ConfigReader.string("MQ_TYPE", "rabbitmq").toLowerCase();

      if (mqType === "kafka") {
        return new KafkaMessageBus(logger);
      } else {
        return new RabbitMQMessageBus(logger);
      }
    }).singleton()
  );

  // Event Publisher - depends on messageBus and logger
  // Use asFunction to ensure dependencies are properly injected
  // Selects implementation based on MQ_TYPE environment variable
  container.register(
    "eventPublisher",
    asFunction(() => {
      const messageBus = container.resolve<IMessageBus>("messageBus");
      const logger = container.resolve<ILogger>("logger");
      const mqType = ConfigReader.string("MQ_TYPE", "rabbitmq").toLowerCase();

      if (mqType === "kafka") {
        return new KafkaEventPublisher(messageBus, logger);
      } else {
        return new AMQPEventPublisher(messageBus, logger);
      }
    }).singleton()
  );

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
