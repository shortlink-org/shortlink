## Observability

This document describes how to configure observability for your application.

### Standard Metrics

All services implement standard metrics according to [ADR-0018](../ADR/decisions/0018-service-metrics.md). See [Observability Metrics](./observability-metrics.md) for detailed documentation on:

- Build information metrics (version, commit, build time)
- SLA/SLO/SLI metrics (availability, latency, error rate)
- Cache metrics (hit/miss ratios)
- Basic service metrics (RPS, response time, error percentage)

### Docs

- [Standard Observability Metrics](./observability-metrics.md)
- [OpenTelemetry Distributed Tracing](https://uptrace.dev/opentelemetry/distributed-tracing.html)
- [OpenTelemetry Metrics](https://uptrace.dev/opentelemetry/metrics.html)