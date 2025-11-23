import {
  diag,
  DiagConsoleLogger,
  DiagLogLevel,
  context,
  trace,
} from "@opentelemetry/api";
import { NodeSDK, logs } from "@opentelemetry/sdk-node";
import { getNodeAutoInstrumentations } from "@opentelemetry/auto-instrumentations-node";
import { OTLPTraceExporter } from "@opentelemetry/exporter-trace-otlp-grpc";
import { OTLPMetricExporter } from "@opentelemetry/exporter-metrics-otlp-grpc";
import { OTLPLogExporter } from "@opentelemetry/exporter-logs-otlp-grpc";
import { PrometheusExporter } from "@opentelemetry/exporter-prometheus";
import { PeriodicExportingMetricReader } from "@opentelemetry/sdk-metrics";
import { resourceFromAttributes } from "@opentelemetry/resources";
import {
  ATTR_SERVICE_NAME,
  ATTR_SERVICE_VERSION,
  ATTR_HTTP_ROUTE,
  ATTR_USER_AGENT_ORIGINAL,
} from "@opentelemetry/semantic-conventions";
import { WinstonInstrumentation } from "@opentelemetry/instrumentation-winston";
import { RuntimeNodeInstrumentation } from "@opentelemetry/instrumentation-runtime-node";
import { KafkaJsInstrumentation } from "@opentelemetry/instrumentation-kafkajs";
import type { SpanExporter } from "@opentelemetry/sdk-trace-base";
import type { FastifyInstance, FastifyReply, FastifyRequest } from "fastify";

const DEFAULT_SERVICE_NAME = "proxy-service";
const DEFAULT_COLLECTOR_ENDPOINT = "grpc://grafana-tempo.grafana:4317";
const DEFAULT_DIAG_LOG_LEVEL = DiagLogLevel.WARN;

/**
 * Global Prometheus exporter instance for /metrics
 */
let prometheusExporter: PrometheusExporter | null = null;
/**
 * Global NodeSDK instance for graceful shutdown
 */
let sdkInstance: NodeSDK | null = null;
let shutdownRegistered = false;

/**
 * Configure OTEL diag logger + suppress noisy warnings
 * Note: NodeSDK will also set a logger, so we only configure if not already set
 */
function configureDiagnostics(): void {
  if (!process.env.OTEL_LOG_LEVEL) {
    process.env.OTEL_LOG_LEVEL = "warn";
  }

  // Only set logger if not already set (to avoid "Current logger will overwrite" warning)
  // NodeSDK will set its own logger, so we skip setting it here
  // diag.setLogger(new DiagConsoleLogger(), DEFAULT_DIAG_LOG_LEVEL);

  const originalWarn = console.warn;
  console.warn = (...args: unknown[]) => {
    const message = String(args[0] ?? "");
    if (
      message.includes(
        "has been loaded before @opentelemetry/instrumentation"
      ) ||
      message.includes("@opentelemetry/winston-transport is not available") ||
      message.includes("Current logger will overwrite")
    ) {
      return;
    }
    originalWarn.apply(console, args);
  };
}

/**
 * Map SERVICE_NAME → OTEL_SERVICE_NAME if needed
 */
function normalizeServiceNameEnv(): void {
  if (!process.env.OTEL_SERVICE_NAME && process.env.SERVICE_NAME) {
    process.env.OTEL_SERVICE_NAME = process.env.SERVICE_NAME;
  }
}

/**
 * Setup SIGTERM/SIGINT graceful shutdown for the SDK
 */
function registerSdkShutdownHook(sdk: NodeSDK): void {
  if (shutdownRegistered) return;
  shutdownRegistered = true;

  const handler = async (signal: NodeJS.Signals) => {
    console.log(
      `[Telemetry] Received ${signal}, shutting down OpenTelemetry SDK...`
    );
    try {
      await sdk.shutdown();
      console.log("[Telemetry] OpenTelemetry SDK shutdown complete");
    } catch (err) {
      console.error("[Telemetry] Error during OpenTelemetry SDK shutdown", err);
    } finally {
      // Let the process exit naturally (don't call process.exit())
    }
  };

  process.once("SIGTERM", handler);
  process.once("SIGINT", handler);
}

/**
 * Initialize OpenTelemetry SDK:
 *  - Traces → OTLP gRPC
 *  - Metrics → Prometheus + OTLP gRPC
 *  - Logs → OTLP gRPC
 *  - Runtime metrics auto-instrumentation
 *  - Winston auto log injection + sending to OTLP Logs SDK
 *
 * Returns PrometheusExporter for /metrics endpoint.
 */
export function initializeTelemetry(): PrometheusExporter | null {
  if (sdkInstance) {
    // Already initialized (e.g., in tests / dev hot reload)
    return prometheusExporter;
  }

  configureDiagnostics();
  normalizeServiceNameEnv();

  const collectorEndpoint =
    process.env.OTEL_EXPORTER_OTLP_ENDPOINT ?? DEFAULT_COLLECTOR_ENDPOINT;

  const serviceName =
    process.env.OTEL_SERVICE_NAME ??
    process.env.SERVICE_NAME ??
    DEFAULT_SERVICE_NAME;

  // --- Traces: OTLP gRPC exporter ---
  const traceExporter = new OTLPTraceExporter({
    url: collectorEndpoint,
  }) as unknown as SpanExporter;

  // --- Metrics: Prometheus pull + OTLP push ---
  // Prometheus endpoint (pull)
  const prometheusOptions = {
    port: Number(process.env.METRICS_PORT ?? 9464),
    host: process.env.METRICS_HOST,
    endpoint: process.env.METRICS_PATH ?? "/metrics",
  };
  prometheusExporter = new PrometheusExporter(prometheusOptions);

  // OTLP metrics (push) — to Alloy/Collector
  const otlpMetricExporter = new OTLPMetricExporter({
    url: collectorEndpoint,
  });

  const otlpMetricReader = new PeriodicExportingMetricReader({
    exporter: otlpMetricExporter,
    exportIntervalMillis: Number(
      process.env.OTEL_METRICS_EXPORT_INTERVAL ?? 10000
    ),
  });

  // --- Logs: OTLP gRPC exporter ---
  const otlpLogExporter = new OTLPLogExporter({
    url: collectorEndpoint,
  });

  const logRecordProcessor = new logs.BatchLogRecordProcessor(otlpLogExporter);

  // --- Common resource (service.name, env, version, etc.) ---
  const resource = resourceFromAttributes({
    [ATTR_SERVICE_NAME]: serviceName,
    "deployment.environment":
      process.env.DEPLOY_ENV ?? process.env.NODE_ENV ?? "development",
    [ATTR_SERVICE_VERSION]:
      process.env.SERVICE_VERSION ??
      process.env.npm_package_version ??
      "unknown",
  });

  // --- Instrumentations ---
  const winstonInstrumentation = new WinstonInstrumentation({
    // Add trace_id/span_id/trace_flags to logs
    disableLogCorrelation: false,
    // Send logs to OTEL Logs SDK (via @opentelemetry/winston-transport)
    disableLogSending: false,
    // Can add additional fields via hook
    logHook: (span, record) => {
      record["service.name"] = serviceName;
      if (span) {
        const spanContext = span.spanContext();
        record["span.name"] = spanContext?.spanId || "unknown";
      }
    },
  });

  const runtimeInstrumentation = new RuntimeNodeInstrumentation({
    // Example: can enable/disable individual metrics
    enabled: true,
  });

  const kafkaInstrumentation = new KafkaJsInstrumentation({
    // Enable instrumentation for producer and consumer
    enabled: true,
    // Can configure operation filtering if needed
    // consumerHook: (span, info) => { ... },
    // producerHook: (span, info) => { ... },
  });

  const instrumentations = [
    getNodeAutoInstrumentations({
      // Can tune specific auto-instrumentations if desired
      "@opentelemetry/instrumentation-http": { enabled: true },
      "@opentelemetry/instrumentation-fastify": { enabled: true },
      "@opentelemetry/instrumentation-winston": {
        enabled: true,
        disableLogCorrelation: false,
        disableLogSending: false,
      },
    }),
    winstonInstrumentation,
    runtimeInstrumentation,
    kafkaInstrumentation,
  ];

  // --- SDK init ---
  const sdk = new NodeSDK({
    resource,
    traceExporter,
    metricReaders: [prometheusExporter, otlpMetricReader],
    logRecordProcessors: [logRecordProcessor],
    instrumentations,
  });

  sdk.start();

  sdkInstance = sdk;
  registerSdkShutdownHook(sdk);

  console.log(
    "[Telemetry] OpenTelemetry SDK initialized (traces+metrics+logs)"
  );

  return prometheusExporter;
}

/**
 * Explicit SDK shutdown helper (e.g., for tests)
 */
export async function shutdownTelemetry(): Promise<void> {
  if (!sdkInstance) return;
  try {
    await sdkInstance.shutdown();
    console.log("[Telemetry] SDK shutdown via shutdownTelemetry()");
  } catch (err) {
    console.error("[Telemetry] Error in shutdownTelemetry()", err);
  } finally {
    sdkInstance = null;
  }
}

/**
 * Get Prometheus exporter instance for /metrics endpoint.
 * Returns null if telemetry not initialized.
 */
export function getPrometheusExporter(): PrometheusExporter | null {
  return prometheusExporter;
}

/**
 * Fastify hook for semantic span enrichment:
 * - http.route (in case auto-instr didn't set it)
 * - client.address / user_agent
 *
 * Usage:
 *   fastify.register(semanticEnrichmentPlugin);
 */
export async function semanticEnrichmentPlugin(
  fastify: FastifyInstance
): Promise<void> {
  fastify.addHook(
    "onRequest",
    (req: FastifyRequest, _reply: FastifyReply, done) => {
      const span = trace.getSpan(context.active());
      if (span) {
        const route =
          // Fastify v4
          (req as any).routerPath ||
          (req.routeOptions && req.routeOptions.url) ||
          req.raw.url;

        span.setAttributes({
          [ATTR_HTTP_ROUTE]: route,
          "net.peer.ip": req.ip,
          [ATTR_USER_AGENT_ORIGINAL]: req.headers["user-agent"] ?? "",
        });
      }
      done();
    }
  );
}
