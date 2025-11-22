import type { FastifyInstance, FastifyPluginOptions } from "fastify";
import type { AwilixContainer } from "awilix";
import type { ContainerDependencies } from "../../../../di/container.js";
import { registerProxyRoutes } from "./proxyRoutes.js";
import { registerMetricsRoutes } from "./metricsRoutes.js";

/**
 * Registers all routes for the Fastify application.
 * This is the routes composition point following Clean Architecture.
 *
 * @param fastify - Fastify instance
 * @param opts - Plugin options including container
 */
export async function registerRoutes(
  fastify: FastifyInstance,
  opts: FastifyPluginOptions & {
    container: AwilixContainer<ContainerDependencies>;
  }
): Promise<void> {
  // Register all route modules
  await fastify.register(registerProxyRoutes, opts);
  await fastify.register(registerMetricsRoutes, opts);
}
