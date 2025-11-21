import type { FastifyInstance } from "fastify";
import type { AwilixContainer } from "awilix";
import type { ContainerDependencies } from "../container/index.js";
import type { IMessageBus } from "../proxy/domain/interfaces/IMessageBus.js";

const READY_ENDPOINT = "/ready";

/**
 * Health check handler for Kubernetes readiness/liveness probes.
 *
 * @returns Promise resolving to true if healthy
 */
async function onHealthCheck(): Promise<boolean> {
  // TODO: Add actual health checks (e.g., database connection, external services)
  return true;
}

/**
 * Configure graceful shutdown and health checks for Fastify server.
 *
 * Sets up:
 * - Health check endpoint at /ready
 * - Message Bus disconnection on shutdown
 *
 * @param server - Fastify server instance
 * @param container - DI container for accessing dependencies
 */
export function configureHealthChecks(
  server: any, // Fastify server's internal HTTP server
  container?: AwilixContainer<ContainerDependencies>
): void {
  // Register health check endpoint on Fastify
  // Note: server is the Node.js HTTP server from fastify.server
  const fastify = (server as any).fastify as FastifyInstance | undefined;
  
  if (fastify) {
    // Register /ready endpoint
    fastify.get(READY_ENDPOINT, {
      schema: {
        description: "Health check endpoint for Kubernetes readiness/liveness probes",
        tags: ["Health"],
        response: {
          200: {
            description: "Service is ready",
            type: "string",
            example: "OK",
          },
          503: {
            description: "Service is not ready (e.g., during graceful shutdown)",
            type: "string",
            example: "Service Unavailable",
          },
        },
      },
    }, async (request, reply) => {
      const healthy = await onHealthCheck();
      if (healthy) {
        return reply.code(200).send("OK");
      }
      return reply.code(503).send("Service Unavailable");
    });

    // Graceful shutdown handler
    fastify.addHook("onClose", async () => {
      if (container) {
        try {
          const messageBus = container.resolve<IMessageBus>("messageBus");
          if (messageBus.isConnected()) {
            await messageBus.disconnect();
          }
        } catch (error) {
          console.error("[Shutdown] Failed to disconnect Message Bus:", error);
        }
      }
    });
  }
}

