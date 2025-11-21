import type { FastifyRequest, FastifyReply } from "fastify";
import { getPrometheusExporter } from "../../../../../infrastructure/telemetry.js";

/**
 * HTTP Controller for Prometheus metrics.
 * Provides /metrics endpoint for Prometheus scraping.
 */
export class MetricsController {
  /**
   * Handles metrics request
   * GET /metrics
   *
   * @param request - Fastify request
   * @param reply - Fastify reply object
   */
  async getMetrics(
    request: FastifyRequest,
    reply: FastifyReply
  ): Promise<void> {
    const exporter = getPrometheusExporter();

    if (!exporter) {
      reply.code(503).send("# Prometheus exporter not initialized\n");
      return;
    }

    try {
      // Fastify-compatible handler for Prometheus metrics
      // Convert Fastify request/reply to Express-like interface
      const expressReq = {
        url: request.url,
        method: request.method,
        headers: request.headers,
        connection: {
          remoteAddress: request.ip,
        },
      } as any;

      const expressRes = {
        statusCode: 200,
        headers: {},
        setHeader: (name: string, value: string) => {
          reply.header(name, value);
        },
        end: (chunk: string) => {
          reply.send(chunk);
        },
      } as any;

      exporter.getMetricsRequestHandler(expressReq, expressRes);
    } catch (error) {
      reply.code(500).send(`# Error collecting metrics: ${error}\n`);
    }
  }
}

