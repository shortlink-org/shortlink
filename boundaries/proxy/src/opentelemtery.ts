import * as opentelemetry from "@opentelemetry/sdk-node";
import { diag, DiagConsoleLogger, DiagLogLevel } from '@opentelemetry/api';
import { ConsoleSpanExporter } from "@opentelemetry/sdk-trace-base";

// For troubleshooting, set the log level to DiagLogLevel.DEBUG
diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.INFO);

// Map legacy SERVICE_NAME to standard OTEL_SERVICE_NAME if provided
if (!process.env.OTEL_SERVICE_NAME && process.env.SERVICE_NAME) {
  process.env.OTEL_SERVICE_NAME = process.env.SERVICE_NAME;
}

const sdk = new opentelemetry.NodeSDK({
  traceExporter: new ConsoleSpanExporter(),
  instrumentations: [],
});

sdk.start()
