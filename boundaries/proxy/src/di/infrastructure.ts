import { OpenTelemetryGrpcMetrics } from "../infrastructure/metrics/index.js";
import { OpenTelemetryGrpcTracing } from "../infrastructure/tracing/index.js";
import { LinkServiceACL } from "../infrastructure/anti-corruption/LinkServiceACL.js";
import { LinkServiceConnectAdapter } from "../infrastructure/adapters/index.js";
import { AMQPEventPublisher } from "../infrastructure/messaging/AMQPEventPublisher.js";
import { LinkServiceRepository } from "../infrastructure/repositories/LinkServiceRepository.js";
import { KratosSessionExtractor } from "../infrastructure/auth/index.js";

export const INFRA = {
  grpcMetrics: OpenTelemetryGrpcMetrics,
  grpcTracing: OpenTelemetryGrpcTracing,
  linkServiceACL: LinkServiceACL,
  linkServiceAdapter: LinkServiceConnectAdapter,
  eventPublisher: AMQPEventPublisher,
  linkRepository: LinkServiceRepository,
  kratosSessionExtractor: KratosSessionExtractor,
};
