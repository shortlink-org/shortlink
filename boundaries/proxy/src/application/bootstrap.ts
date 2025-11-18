import "reflect-metadata";
import * as dotenv from "dotenv";

// 1. Load environment variables FIRST (до логгера, до конфигов)
dotenv.config();

import http from "http";
import log from "../logger.js";
import container from "../inversify.config.js";
import { initializeTelemetry } from "../infrastructure/telemetry.js";
import { initializeProfiling } from "../infrastructure/profiling.js";
import { configureHealthChecks } from "../infrastructure/health.js";
import { createServer } from "../interfaces/http/server.js";
import { AppConfig } from "../infrastructure/config/index.js";
import { logPermissions } from "../infrastructure/permissions.js";
import TYPES from "../types.js";
import { IMessageBus } from "../proxy/domain/interfaces/IMessageBus.js";
import { IEventPublisher } from "../proxy/application/use-cases/PublishEventUseCase.js";
import { StandardMetrics } from "../proxy/infrastructure/metrics/StandardMetrics.js";

// Build information - injected at build time
const version = process.env.APP_VERSION || "dev";
const commit = process.env.APP_COMMIT || "none";
const buildTime = process.env.APP_BUILD_TIME || "unknown";

// Improved type guard (короче, безопаснее)
type ExchangeInitializer = {
  initializeExchanges: () => Promise<void>;
};
function isExchangeInitializer(
  candidate: unknown
): candidate is ExchangeInitializer {
  if (
    typeof candidate !== "object" ||
    candidate === null ||
    !("initializeExchanges" in candidate)
  ) {
    return false;
  }

  return (
    typeof (candidate as { initializeExchanges: unknown })
      .initializeExchanges === "function"
  );
}

function bootstrap(): http.Server {
  // 0. Validate Permissions API
  if (process.env.NODE_ENV === "production") {
    logPermissions();
  }

  // 1. Telemetry & Profiling
  initializeTelemetry();

  // Initialize standard metrics and set build info (ADR-0014)
  const standardMetrics = new StandardMetrics();
  standardMetrics.setBuildInfo(version, commit, buildTime);
  log.info(`[Bootstrap] Service initialized - version: ${version}, commit: ${commit}, build_time: ${buildTime}`);

  initializeProfiling().catch((err) => {
    console.error("[Bootstrap] Failed to initialize profiling:", err);
  });

  // 2. Message Bus (fire-and-forget)
  const messageBus = container.get<IMessageBus>(
    TYPES.INFRASTRUCTURE.MessageBus
  );
  const eventPublisher = container.get<IEventPublisher>(
    TYPES.INFRASTRUCTURE.EventPublisher
  );

  messageBus
    .connect()
    .then(async () => {
      if (isExchangeInitializer(eventPublisher)) {
        await eventPublisher.initializeExchanges();
      }
    })
    .catch((err) => {
      log.error("[Bootstrap] Failed to connect Message Bus:", err);
    });

  // 3. Load configuration
  const appConfig = new AppConfig();

  // 4. Create app
  const app = createServer(container);

  // 5. Create HTTP server
  const server = http.createServer(app);

  // 6. Health checks
  configureHealthChecks(server);

  // 7. Start server — with error listener
  server.on("error", (err) => {
    log.error("[Bootstrap] HTTP server failed to start:", err);
    process.exit(1);
  });

  server.listen(appConfig.port, () => {
    log.info(`App running on ${appConfig.port}`);
  });

  return server;
}

// Start app
let server: http.Server;
try {
  server = bootstrap();
} catch (error) {
  log.error("Failed to start application", error);
  process.exit(1);
}

export default server;
