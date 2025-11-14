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
