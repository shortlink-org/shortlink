import { Container } from "inversify";
import TYPES from "./types.js";

// Domain Services
import { LinkDomainService } from "./proxy/domain/services/LinkDomainService.js";

// Application Layer
import { GetLinkByHashUseCase } from "./proxy/application/use-cases/GetLinkByHashUseCase.js";
import { PublishEventUseCase } from "./proxy/application/use-cases/PublishEventUseCase.js";
import { LinkApplicationService } from "./proxy/application/services/LinkApplicationService.js";

// Application Layer - Pipeline
import {
  UseCasePipeline,
  LoggingInterceptor,
  MetricsInterceptor,
  AuthorizationInterceptor,
  DefaultAuthorizationChecker,
  IAuthorizationChecker,
} from "./proxy/application/pipeline/index.js";

// Infrastructure Layer - Repositories
import { LinkServiceRepository } from "./proxy/infrastructure/repositories/LinkServiceRepository.js";
import { ILinkRepository } from "./proxy/domain/repositories/ILinkRepository.js";

// Infrastructure Layer - Adapters & ACL
import { LinkServiceACL } from "./proxy/infrastructure/anti-corruption/LinkServiceACL.js";
import {
  ILinkServiceAdapter,
  LinkServiceConnectAdapter,
} from "./proxy/infrastructure/adapters/index.js";

// Infrastructure Layer - Messaging
import { AMQPEventPublisher } from "./proxy/infrastructure/messaging/AMQPEventPublisher.js";
import { IEventPublisher } from "./proxy/application/use-cases/PublishEventUseCase.js";

// Infrastructure Layer - Logging
import { ILogger } from "./infrastructure/logging/ILogger.js";
import { WinstonLogger } from "./infrastructure/logging/WinstonLogger.js";

// Infrastructure Layer - Configuration
import { ExternalServicesConfig } from "./infrastructure/config/ExternalServicesConfig.js";
import { CacheConfig } from "./infrastructure/config/CacheConfig.js";

// Infrastructure Layer - Messaging
import { RabbitMQMessageBus } from "./proxy/infrastructure/messaging/RabbitMQMessageBus.js";
import { IMessageBus } from "./proxy/domain/interfaces/IMessageBus.js";

// Infrastructure Layer - Cache
import {
  ILinkCache,
  RedisLinkCache,
} from "./proxy/infrastructure/cache/RedisLinkCache.js";


// Infrastructure Layer - Metrics & Tracing
import {
  IGrpcMetrics,
  OpenTelemetryGrpcMetrics,
} from "./proxy/infrastructure/metrics/index.js";
import {
  IGrpcTracing,
  OpenTelemetryGrpcTracing,
} from "./proxy/infrastructure/tracing/index.js";

const container = new Container();

// ============================================================================
// DOMAIN BINDINGS
// ============================================================================

// Domain Services
container
  .bind<LinkDomainService>(TYPES.DOMAIN.LinkDomainService)
  .to(LinkDomainService)
  .inSingletonScope();

// ============================================================================
// APPLICATION BINDINGS
// ============================================================================

// Use Cases
container
  .bind<GetLinkByHashUseCase>(TYPES.APPLICATION.GetLinkByHashUseCase)
  .to(GetLinkByHashUseCase)
  .inSingletonScope();

container
  .bind<PublishEventUseCase>(TYPES.APPLICATION.PublishEventUseCase)
  .to(PublishEventUseCase)
  .inSingletonScope();

// Application Services
container
  .bind<LinkApplicationService>(TYPES.APPLICATION.LinkApplicationService)
  .to(LinkApplicationService)
  .inSingletonScope();

// Pipeline
container
  .bind<UseCasePipeline>(TYPES.APPLICATION.UseCasePipeline)
  .to(UseCasePipeline)
  .inSingletonScope();

// Pipeline Interceptors
container
  .bind<LoggingInterceptor>(TYPES.APPLICATION.LoggingInterceptor)
  .to(LoggingInterceptor)
  .inSingletonScope();

container
  .bind<MetricsInterceptor>(TYPES.APPLICATION.MetricsInterceptor)
  .to(MetricsInterceptor)
  .inSingletonScope();

container
  .bind<AuthorizationInterceptor>(TYPES.APPLICATION.AuthorizationInterceptor)
  .to(AuthorizationInterceptor)
  .inSingletonScope();

// Authorization Checker
container
  .bind<IAuthorizationChecker>(TYPES.INFRASTRUCTURE.AuthorizationChecker)
  .to(DefaultAuthorizationChecker)
  .inSingletonScope();

// ============================================================================
// INFRASTRUCTURE BINDINGS
// ============================================================================

// Repositories
// LinkServiceRepository требует LinkServiceAdapter и ACL
container
  .bind<ILinkRepository>(TYPES.REPOSITORY.LinkRepository)
  .to(LinkServiceRepository)
  .inSingletonScope();

// Anti-corruption Layer
container
  .bind<LinkServiceACL>(Symbol.for("LinkServiceACL"))
  .to(LinkServiceACL)
  .inSingletonScope();

// Adapters - используем только gRPC/Connect адаптер
container
  .bind<ILinkServiceAdapter>(TYPES.INFRASTRUCTURE.LinkServiceAdapter)
  .to(LinkServiceConnectAdapter)
  .inSingletonScope();

// Messaging
container
  .bind<IEventPublisher>(TYPES.INFRASTRUCTURE.EventPublisher)
  .to(AMQPEventPublisher)
  .inSingletonScope();

// Logging
container
  .bind<ILogger>(TYPES.INFRASTRUCTURE.Logger)
  .to(WinstonLogger)
  .inSingletonScope();

// Configuration
container
  .bind<ExternalServicesConfig>(Symbol.for("ExternalServicesConfig"))
  .to(ExternalServicesConfig)
  .inSingletonScope();

container
  .bind<CacheConfig>(Symbol.for("CacheConfig"))
  .to(CacheConfig)
  .inSingletonScope();

// Cache
container
  .bind<ILinkCache>(TYPES.INFRASTRUCTURE.LinkCache)
  .to(RedisLinkCache)
  .inSingletonScope();

// Message Bus
container
  .bind<IMessageBus>(TYPES.INFRASTRUCTURE.MessageBus)
  .to(RabbitMQMessageBus)
  .inSingletonScope();


// gRPC Metrics
container
  .bind<IGrpcMetrics>(TYPES.INFRASTRUCTURE.GrpcMetrics)
  .to(OpenTelemetryGrpcMetrics)
  .inSingletonScope();

// gRPC Tracing
container
  .bind<IGrpcTracing>(TYPES.INFRASTRUCTURE.GrpcTracing)
  .to(OpenTelemetryGrpcTracing)
  .inSingletonScope();

export default container;
