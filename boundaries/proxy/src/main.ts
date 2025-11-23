// CRITICAL: Import telemetry-init FIRST to set up OpenTelemetry instrumentation
// before any modules (amqplib, ioredis, winston) are loaded
import "./telemetry-init.js";

// Now safe to import other modules
import { createDIContainer } from "./di/container.js";
import { buildServer } from "./infrastructure/http/fastify/server.js";
import { initializeProfiling } from "./infrastructure/profiling.js";
import { configureHealthChecks } from "./infrastructure/health.js";
import { AppConfig } from "./infrastructure/config/index.js";
import { logPermissions } from "./infrastructure/permissions.js";
import type { IMessageBus } from "./domain/interfaces/IMessageBus.js";
import type { IEventPublisher } from "./application/use-cases/PublishEventUseCase.js";
import type { FastifyInstance } from "fastify";
import type { AwilixContainer } from "awilix";
import type { ContainerDependencies } from "./di/container.js";

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

async function bootstrap(): Promise<{
  app: FastifyInstance;
  gracefulShutdownHandler: () => Promise<void>;
}> {
  // 0. Validate Permissions API (production only)
  if (process.env.NODE_ENV === "production") {
    logPermissions();
  }

  // 1. Profiling (Telemetry already initialized at module load time)
  await initializeProfiling().catch((err: unknown) => {
    console.error("[Bootstrap] Failed to initialize profiling:", err);
  });

  // 2. Create DI container (Composition Root)
  const container = createDIContainer();
  const logger = container.resolve("logger");

  // 3. Message Bus - must connect before starting server
  const messageBus = container.resolve<IMessageBus>("messageBus");
  const eventPublisher = container.resolve<IEventPublisher>("eventPublisher");

  try {
    await messageBus.connect();

    // Initialize exchanges if event publisher supports it
    if (isExchangeInitializer(eventPublisher)) {
      await eventPublisher.initializeExchanges();
    }

    logger.info("[Bootstrap] Message Bus connected and exchanges initialized");
  } catch (err) {
    logger.error("[Bootstrap] Failed to connect Message Bus:", err);
    logger.error(
      "[Bootstrap] Application cannot start without Message Bus. Exiting..."
    );
    process.exit(1);
  }

  // 4. Build Fastify server
  const app = await buildServer(container);

  // 5. Health checks and graceful shutdown handler
  const appConfig = container.resolve<AppConfig>("appConfig");
  const gracefulShutdownHandler = configureHealthChecks(app, container);

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

  return { app, gracefulShutdownHandler };
}

// Start app
let app: FastifyInstance | undefined;
let gracefulShutdownHandler: (() => Promise<void>) | undefined;

bootstrap()
  .then(({ app: instance, gracefulShutdownHandler: handler }) => {
    app = instance;
    gracefulShutdownHandler = handler;
  })
  .catch((error) => {
    console.error("Failed to start application", error);
    process.exit(1);
  });

// Graceful shutdown
async function gracefulShutdown() {
  if (gracefulShutdownHandler) {
    await gracefulShutdownHandler();
  } else if (app) {
    // Fallback if handler not available
    await app.close();
  }
  process.exit(0);
}

process.on("SIGTERM", gracefulShutdown);
process.on("SIGINT", gracefulShutdown);

export default app;
