# 10. logger format

Date: 2022-12-21

## Status

Accepted

## Context

We use logging to track the performance of our services. We want to develop a standard approach for logging all of our
services.

## Decision

We selected JSON format for logging. We will use the following format:

```json
{
  "level": "info",
  "msg": "message",
  "timestamp": "2006-01-02T15:04:05.999999999Z07:00",
  "caller": "main.go:42",
  "trace_id": "trace_id"
}
```

#### Application logs

To record application logs, use [span events](https://uptrace.dev/opentelemetry/distributed-tracing.html#event).
You must set the event name to `log` and
use [semantic attributes](https://uptrace.dev/opentelemetry/distributed-tracing.html#attributes) to record the
context:

- `level` - to record the log severity. Must be one of `DEBUG`, `INFO`, `WARN`, `ERROR` and `FATAL`.
- `msg` - to record the log message.
- `timestamp` - to record the log timestamp (_RFC3339Nano_).
- `caller` - to record the caller function name.
- `trace_id` - to record the trace ID (_hex-encoded_).

## Consequences

We use next material for made our logger format:

- [Structured Logging](https://uptrace.dev/opentelemetry/structured-logging.html)
