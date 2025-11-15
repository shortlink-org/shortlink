import { diag, DiagConsoleLogger, DiagLogLevel } from "@opentelemetry/api";
import { NodeSDK } from "@opentelemetry/sdk-node";
import { getNodeAutoInstrumentations } from "@opentelemetry/auto-instrumentations-node";
import { OTLPTraceExporter } from "@opentelemetry/exporter-trace-otlp-grpc";
import { PrometheusExporter } from "@opentelemetry/exporter-prometheus";
import { resourceFromAttributes } from "@opentelemetry/resources";
import type { SpanExporter } from "@opentelemetry/sdk-trace-base";

const DEFAULT_SERVICE_NAME = "proxy-service";
const DEFAULT_COLLECTOR_ENDPOINT = "grpc://grafana-tempo.grafana:4317";
const DEFAULT_DIAG_LOG_LEVEL = DiagLogLevel.INFO;

/**
 * Prometheus exporter instance for metrics endpoint
 */
let prometheusExporter: PrometheusExporter | null = null;

/**
 * Initialize OpenTelemetry SDK for distributed tracing and metrics.
 *
 * Configures:
 * - Service name from OTEL_SERVICE_NAME or SERVICE_NAME env vars
 * - OTLP trace exporter to collector endpoint
 * - Prometheus metrics exporter for /metrics endpoint
 * - Auto-instrumentations for Node.js
 *
 * Side effects: Starts the SDK and sets up global instrumentation.
 *
 * @returns Prometheus exporter instance for /metrics endpoint
 */
export function initializeTelemetry(): PrometheusExporter | null {
  // Configure diagnostic logger
  diag.setLogger(new DiagConsoleLogger(), DEFAULT_DIAG_LOG_LEVEL);

  // Map legacy SERVICE_NAME to standard OTEL_SERVICE_NAME if provided
  if (!process.env.OTEL_SERVICE_NAME && process.env.SERVICE_NAME) {
    process.env.OTEL_SERVICE_NAME = process.env.SERVICE_NAME;
  }

  const collectorEndpoint =
    process.env.OTEL_EXPORTER_OTLP_ENDPOINT ?? DEFAULT_COLLECTOR_ENDPOINT;

  const traceExporter = new OTLPTraceExporter({
    url: collectorEndpoint,
  }) as unknown as SpanExporter;

  const serviceName =
    process.env.OTEL_SERVICE_NAME ??
    process.env.SERVICE_NAME ??
    DEFAULT_SERVICE_NAME;

  // Configure Prometheus metrics exporter
  // PrometheusExporter exposes an HTTP server that Prometheus scrapes
  const metricsPort = Number(process.env.METRICS_PORT ?? 9464);
  const metricsHost = process.env.METRICS_HOST;
  prometheusExporter = new PrometheusExporter({
    port: metricsPort,
    host: metricsHost,
  });

  const sdk = new NodeSDK({
    resource: resourceFromAttributes({
      "service.name": serviceName,
    }),
    traceExporter,
    metricReader: prometheusExporter,
    instrumentations: [getNodeAutoInstrumentations()],
  });

  sdk.start();

  console.log("[Telemetry] OpenTelemetry SDK initialized with Prometheus metrics exporter");

  return prometheusExporter;
}

/**
 * Get Prometheus exporter instance for /metrics endpoint
 * Returns null if telemetry is not initialized
 */
export function getPrometheusExporter(): PrometheusExporter | null {
  return prometheusExporter;
}

