import * as opentelemetry from "@opentelemetry/sdk-node";
import { diag, DiagConsoleLogger, DiagLogLevel } from '@opentelemetry/api';
import { Resource } from "@opentelemetry/resources";
import { SemanticResourceAttributes } from "@opentelemetry/semantic-conventions";
import { NodeTracerProvider } from "@opentelemetry/sdk-trace-node";
import { registerInstrumentations } from "@opentelemetry/instrumentation";
import { ConsoleSpanExporter, BatchSpanProcessor } from "@opentelemetry/sdk-trace-base";

// For troubleshooting, set the log level to DiagLogLevel.DEBUG
diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.INFO);

const sdk = new opentelemetry.NodeSDK({
  traceExporter: new opentelemetry.tracing.ConsoleSpanExporter(),
  instrumentations: [],
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: process.env.SERVICE_NAME,
  }),
});

sdk.start()
