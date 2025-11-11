import { diag, DiagConsoleLogger, DiagLogLevel } from "@opentelemetry/api";
import { NodeSDK } from "@opentelemetry/sdk-node";
import { getNodeAutoInstrumentations } from "@opentelemetry/auto-instrumentations-node";
import { OTLPTraceExporter } from "@opentelemetry/exporter-trace-otlp-grpc";
import { resourceFromAttributes } from "@opentelemetry/resources";
import type { SpanExporter } from "@opentelemetry/sdk-trace-base";

// For troubleshooting, set the log level to DiagLogLevel.DEBUG
diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.INFO);

// Map legacy SERVICE_NAME to standard OTEL_SERVICE_NAME if provided
if (!process.env.OTEL_SERVICE_NAME && process.env.SERVICE_NAME) {
  process.env.OTEL_SERVICE_NAME = process.env.SERVICE_NAME;
}

const collectorEndpoint =
  process.env.OTEL_EXPORTER_OTLP_ENDPOINT ??
  "grpc://grafana-tempo.grafana:4317";

const traceExporter = new OTLPTraceExporter({
  url: collectorEndpoint,
}) as unknown as SpanExporter;

const sdk = new NodeSDK({
  resource: resourceFromAttributes({
    "service.name":
      process.env.OTEL_SERVICE_NAME ??
      process.env.SERVICE_NAME ??
      "proxy-service",
  }),
  traceExporter,
  instrumentations: [getNodeAutoInstrumentations()],
});

sdk.start();
