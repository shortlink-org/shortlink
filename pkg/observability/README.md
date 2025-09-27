## Observability

This package provides observability primitives for the application:

- `monitoring` - provides a Prometheus metrics registry and a HTTP handler for exposing metrics
- `tracing` - provides a Tracing provider for OpenTelemetry with Go 1.25 FlightRecorder support
- `logging` - provides a structured logger

### Tracing Package

The `tracing` package offers comprehensive observability capabilities including:

- **OpenTelemetry Integration**: Full OTLP tracing support with automatic span collection
- **Go 1.25 FlightRecorder**: Continuous, low-overhead tracing with rolling buffer for perfect tracing
- **Automatic Trace Capture**: Captures traces on errors, panics, and user-defined signals
- **Environment Configuration**: Configurable via environment variables or code
- **Middleware Support**: Easy integration with existing application code

#### FlightRecorder Features

- **Perfect Tracing**: Always-available trace data for post-mortem analysis
- **Low Overhead**: In-memory ring buffer with configurable retention (default: 1 minute, 3MB)
- **Production Ready**: Thread-safe operations with comprehensive error handling
- **Signal Support**: Manual trace capture via SIGUSR1/SIGUSR2 signals
- **Health Monitoring**: Built-in health check functionality

For detailed FlightRecorder documentation, see [`FLIGHT_RECORDER_IMPLEMENTATION.md`](../../FLIGHT_RECORDER_IMPLEMENTATION.md).

### References

- [uptrace](https://uptrace.dev/opentelemetry/) - more articles and tips
