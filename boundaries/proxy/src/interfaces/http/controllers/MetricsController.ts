import * as express from "express";
import { controller, httpGet, request, response } from "inversify-express-utils";
import { getPrometheusExporter } from "../../../infrastructure/telemetry.js";

/**
 * HTTP Controller для Prometheus метрик
 * Предоставляет endpoint /metrics для scraping Prometheus
 */
@controller("/metrics")
export class MetricsController {
  /**
   * @swagger
   * /metrics:
   *   get:
   *     summary: Prometheus metrics endpoint
   *     description: |
   *       Endpoint для Prometheus scraping метрик OpenTelemetry.
   *       Возвращает метрики в формате Prometheus (text/plain).
   *     tags:
   *       - Metrics
   *     responses:
   *       '200':
   *         description: Метрики в формате Prometheus
   *         content:
   *           text/plain:
   *             schema:
   *               type: string
   *               example: |
   *                 # HELP grpc_requests_total Total number of gRPC requests
   *                 # TYPE grpc_requests_total counter
   *                 grpc_requests_total{method="GetLinkByHash",service="link-service",status="success"} 42
   */
  @httpGet("/")
  public async getMetrics(
    @request() req: express.Request,
    @response() res: express.Response
  ): Promise<void> {
    const exporter = getPrometheusExporter();

    if (!exporter) {
      res.status(503).send("# Prometheus exporter not initialized\n");
      return;
    }

    try {
      exporter.getMetricsRequestHandler(req, res);
    } catch (error) {
      res.status(500).send(`# Error collecting metrics: ${error}\n`);
    }
  }
}

