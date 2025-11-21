import * as dotenv from "dotenv";

// 1. Load environment variables FIRST (before logger, before configs)
dotenv.config();

import { createDIContainer } from "./container/index.js";
import { buildServer } from "./proxy/infrastructure/http/fastify/server.js";
import { initializeTelemetry } from "./infrastructure/telemetry.js";
import { initializeProfiling } from "./infrastructure/profiling.js";
import { configureHealthChecks } from "./infrastructure/health.js";
import { AppConfig } from "./infrastructure/config/index.js";
import { logPermissions } from "./infrastructure/permissions.js";
import type { IMessageBus } from "./proxy/domain/interfaces/IMessageBus.js";
import type { IEventPublisher } from "./proxy/application/use-cases/PublishEventUseCase.js";
import type { FastifyInstance } from "fastify";
import type { AwilixContainer } from "awilix";
import type { ContainerDependencies } from "./container/index.js";

// Improved type guard
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

async function bootstrap(): Promise<FastifyInstance> {
  // 0. Validate Permissions API (production only)
  if (process.env.NODE_ENV === "production") {
    logPermissions();
  }

  // 1. Telemetry & Profiling
  initializeTelemetry();
  await initializeProfiling().catch((err: unknown) => {
    console.error("[Bootstrap] Failed to initialize profiling:", err);
  });

  // 2. Create DI container (Composition Root)
  const container = createDIContainer();
  const logger = container.resolve("logger");

  // 3. Message Bus (fire-and-forget)
  const messageBus = container.resolve<IMessageBus>("messageBus");
  const eventPublisher = container.resolve<IEventPublisher>("eventPublisher");

  messageBus
    .connect()
    .then(async () => {
      if (isExchangeInitializer(eventPublisher)) {
        await eventPublisher.initializeExchanges();
      }
    })
    .catch((err) => {
      logger.error("[Bootstrap] Failed to connect Message Bus:", err);
    });

  // 4. Build Fastify server
  const app = await buildServer(container);

  // 5. Health checks
  const appConfig = container.resolve<AppConfig>("appConfig");
  configureHealthChecks(app.server, container);

  // 6. Start server
  try {
    await app.listen({
      port: appConfig.port,
      host: "0.0.0.0",
    });

    logger.info(`App running on ${appConfig.port}`);
  } catch (err) {
    logger.error("[Bootstrap] HTTP server failed to start:", err);
    process.exit(1);
  }

  return app;
}

// Start app
let app: FastifyInstance | undefined;

bootstrap()
  .then((instance) => {
    app = instance;
  })
  .catch((error) => {
    console.error("Failed to start application", error);
    process.exit(1);
  });

// Graceful shutdown
async function gracefulShutdown() {
  if (app) {
    await app.close();
  }
  process.exit(0);
}

process.on("SIGTERM", gracefulShutdown);
process.on("SIGINT", gracefulShutdown);

export default app;
