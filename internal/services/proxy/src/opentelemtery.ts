const {NodeTracerProvider} = require('@opentelemetry/sdk-trace-node');
const {registerInstrumentations} = require('@opentelemetry/instrumentation');
const {HttpInstrumentation} = require('@opentelemetry/instrumentation-http');
const {ExpressInstrumentation} = require('@opentelemetry/instrumentation-express');

const provider = new NodeTracerProvider();
provider.register();

registerInstrumentations({
    instrumentations: [
        // Express instrumentation expects HTTP layer to be instrumented
        new HttpInstrumentation(),
        new ExpressInstrumentation(),
    ],
});