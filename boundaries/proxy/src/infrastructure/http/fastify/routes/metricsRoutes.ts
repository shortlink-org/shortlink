import type { FastifyInstance, FastifyPluginOptions } from "fastify";
import type { AwilixContainer } from "awilix";
import type { ContainerDependencies } from "../../../../di/container.js";

/**
 * Fastify JSON Schema for metrics endpoint
 */
const metricsSchema = {
  description: "Prometheus metrics endpoint",
  tags: ["Metrics"],
  response: {
    200: {
      description: "Metrics in Prometheus format",
      type: "string",
      contentMediaType: "text/plain",
    },
    503: {
      description: "Prometheus exporter not initialized",
      type: "string",
    },
    500: {
      description: "Error collecting metrics",
      type: "string",
    },
  },
};

/**
 * Registers metrics routes
 */
export async function registerMetricsRoutes(
  fastify: FastifyInstance,
  opts: FastifyPluginOptions & { container: AwilixContainer<ContainerDependencies> }
): Promise<void> {
  const container = opts.container;
  const metricsController = container.resolve("metricsController");

  // GET /metrics - Prometheus metrics
  fastify.get(
    "/metrics",
    {
      schema: metricsSchema,
    },
    async (request, reply) => {
      // Get controller from request-scoped container
      const requestContainer = (request as any).container as AwilixContainer<ContainerDependencies>;
      const controller = requestContainer.resolve("metricsController");

      await controller.getMetrics(request, reply);
    }
  );
}

