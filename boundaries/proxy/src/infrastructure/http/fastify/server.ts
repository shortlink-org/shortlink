import type { FastifyInstance, FastifyPluginOptions } from "fastify";
import Fastify from "fastify";
import helmet from "@fastify/helmet";
import formbody from "@fastify/formbody";
import swagger from "@fastify/swagger";
import swaggerUi from "@fastify/swagger-ui";
import { context, trace } from "@opentelemetry/api";
import type { AwilixContainer } from "awilix";
import type { ContainerDependencies } from "../../../di/container.js";
import { createRequestScope } from "../../../di/container.js";
import { registerRoutes } from "./routes/index.js";
import { errorHandler } from "./middleware/errorHandler.js";
import type { ILogger } from "../../logging/ILogger.js";

/**
 * Fastify server configuration options
 */
export interface FastifyServerConfig {
  port?: number;
  host?: string;
  logger?: boolean;
  rootPath?: string;
}

/**
 * Builds and configures a Fastify server instance.
 * This is the HTTP adapter layer in Clean Architecture.
 *
 * @param container - Awilix dependency injection container
 * @param config - Server configuration options
 * @returns Configured Fastify instance
 */
export async function buildServer(
  container: AwilixContainer<ContainerDependencies>,
  config: FastifyServerConfig = {}
): Promise<FastifyInstance> {
  const logger = container.resolve<ILogger>("logger");
  const appConfig = container.resolve("appConfig");

  const port = config.port ?? appConfig.port;
  const host = config.host ?? "0.0.0.0";

  // Create Fastify instance
  const fastify = Fastify({
    logger: config.logger ?? false,
    requestIdLogLabel: "reqId",
    disableRequestLogging: false,
  });

  // Register plugins
  await fastify.register(helmet, {
    contentSecurityPolicy: false, // Disable CSP for development
  });

  await fastify.register(formbody);

  // Swagger/OpenAPI
  await fastify.register(swagger, {
    openapi: {
      info: {
        title: "Proxy Service API",
        description: "Proxy service for redirect to original URL",
        version: "1.0.0",
      },
      servers: [
        {
          url: `http://${host}:${port}`,
          description: "Development server",
        },
      ],
    },
  });

  await fastify.register(swaggerUi, {
    routePrefix: "/docs",
    uiConfig: {
      docExpansion: "list",
      deepLinking: false,
    },
    staticCSP: true,
    transformSpecificationClone: true,
  });

  // Request scoping for dependency injection
  // Each request gets its own container scope
  fastify.addHook("onRequest", async (request: any) => {
    const requestScope = createRequestScope(container);
    request.container = requestScope;
  });

  // Attach trace identifier to every outgoing response
  fastify.addHook("onSend", async (request, reply, payload) => {
    const span = trace.getSpan(context.active());
    const traceId = span?.spanContext()?.traceId;

    if (traceId) {
      reply.header("trace-id", traceId);
    }

    return payload;
  });

  // Cleanup request scope after request completes
  fastify.addHook("onResponse", async (request: any) => {
    const scope = request.container as
      | AwilixContainer<ContainerDependencies>
      | undefined;
    if (scope) {
      await scope.dispose();
    }
  });

  // Register error handler (must be registered before routes)
  fastify.setErrorHandler(errorHandler(logger));

  // Register routes
  await fastify.register(registerRoutes, {
    container,
  } as FastifyPluginOptions & {
    container: AwilixContainer<ContainerDependencies>;
  });

  return fastify;
}
