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
import type { SpanExporter } from "@opentelemetry/sdk-trace-base";
import type { FastifyInstance, FastifyReply, FastifyRequest } from "fastify";
import Pyroscope from "@pyroscope/nodejs";

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
 * Map SERVICE_NAME → OTEL_SERVICE_NAME if нужно
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
      // Let the process exit naturally (не вызываем process.exit())
    }
  };

  process.once("SIGTERM", handler);
  process.once("SIGINT", handler);
}

/**
 * Initialize Pyroscope profiler (pprof → Pyroscope server).
 *
 * This is NOT OTLP profiles yet, а прямой push в Pyroscope.
 * Pyroscope уже потом клеится с остальной телеметрией в Grafana.
 */
function initializePyroscopeProfiling(serviceName: string): void {
  const serverAddress = process.env.PYROSCOPE_SERVER_ADDRESS;
  if (!serverAddress) {
    return;
  }

  try {
    Pyroscope.init({
      serverAddress,
      appName: serviceName,
      // Немного тегов для удобной фильтрации
      tags: {
        service: serviceName,
        env: process.env.DEPLOY_ENV || process.env.NODE_ENV || "unknown",
      },
      // Пример включения CPU-time для wall профилей:
      // wall: { collectCpuTime: true },
    });

    Pyroscope.start();
    console.log("[Telemetry] Pyroscope profiling enabled");
  } catch (err) {
    console.warn("[Telemetry] Failed to initialize Pyroscope profiler", err);
  }
}

/**
 * Initialize OpenTelemetry SDK:
 *  - Traces → OTLP gRPC
 *  - Metrics → Prometheus + OTLP gRPC
 *  - Logs → OTLP gRPC
 *  - Runtime metrics auto-instrumentation
 *  - Winston auto log injection + sending в OTLP Logs SDK
 *  - Pyroscope pprof (если задан PYROSCOPE_SERVER_ADDRESS)
 *
 * Возвращает PrometheusExporter для /metrics endpoint.
 */
export function initializeTelemetry(): PrometheusExporter | null {
  if (sdkInstance) {
    // Уже инициализировано (например, в тестах / dev hot reload)
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

  // OTLP metrics (push) — в Alloy/Collector
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
    // добавляем trace_id/span_id/trace_flags в логи
    disableLogCorrelation: false,
    // и шлём сами логи в OTEL Logs SDK (через @opentelemetry/winston-transport)
    disableLogSending: false,
    // можно дописать поля через hook
    logHook: (span, record) => {
      record["service.name"] = serviceName;
      if (span) {
        const spanContext = span.spanContext();
        record["span.name"] = spanContext?.spanId || "unknown";
      }
    },
  });

  const runtimeInstrumentation = new RuntimeNodeInstrumentation({
    // пример: можно включать/отключать отдельные метрики
    enabled: true,
  });

  const instrumentations = [
    getNodeAutoInstrumentations({
      // При желании можно детюнить конкретные auto-instrumentations
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
  initializePyroscopeProfiling(serviceName);

  console.log(
    "[Telemetry] OpenTelemetry SDK initialized (traces+metrics+logs+profiling)"
  );

  return prometheusExporter;
}

/**
 * Explicit SDK shutdown helper (например, для тестов)
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
 * Fastify hook для семантического обогащения спанов:
 * - http.route (если вдруг auto-instr не проставил)
 * - client.address / user_agent
 *
 * Используешь так:
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
