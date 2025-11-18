# Standard Observability Metrics

This document describes the standard metrics implemented across all services according to ADR-0018 (Standard Metrics for Services).

## Overview

All services now expose standard Prometheus metrics that follow consistent naming conventions and provide insights into service health, performance, and resource utilization.

## Build Information Metric (ADR-0014)

Every service exposes a `build_info` metric that contains version, commit SHA, and build time information.

### Format

```
<service>_application_build_info{version="<version>",commit="<commit>",build_time="<timestamp>"} 1
```

### Examples

- `link_application_build_info{version="1.2.3",commit="abc123",build_time="2025-11-18T14:00:00Z"} 1`
- `metadata_application_build_info{version="1.2.3",commit="abc123",build_time="2025-11-18T14:00:00Z"} 1`
- `proxy_application_build_info{version="1.2.3",commit="abc123",build_time="2025-11-18T14:00:00Z"} 1`

## SLA/SLO/SLI Metrics

### Service Availability Ratio

Tracks the ratio of uptime to total time for SLA monitoring.

```
<service>_application_service_availability_ratio <ratio>
```

### HTTP Request Duration

Histogram metric tracking HTTP request response times for SLO monitoring.

```
<service>_application_http_request_duration_seconds{method="<method>",endpoint="<endpoint>",status="<status>"} <duration>
```

### Error Rate Per Minute

Counter tracking the number of errors per minute for system reliability.

```
<service>_application_error_rate_per_minute{type="<error_type>",operation="<operation>"} <count>
```

## Cache Metrics

### Cache Hit Total

Total number of cache hits.

```
<service>_cache_hit_total{cache_type="<type>"} <count>
```

### Cache Miss Total

Total number of cache misses.

```
<service>_cache_miss_total{cache_type="<type>"} <count>
```

### Cache Hit Ratio

Calculated metric: `cache_hit_total / (cache_hit_total + cache_miss_total)`

## Basic Service Metrics

### Requests Per Second (RPS)

Tracks service request load per second.

```
<service>_application_requests_per_second{operation="<operation>"} <count>
```

### Response Time

Histogram of service response times.

```
<service>_application_response_time_seconds{operation="<operation>"} <duration>
```

### Error Rate Percentage

Percentage of error requests.

```
<service>_application_error_rate_percentage{operation="<operation>"} <percentage>
```

## Implementation

### Go Services

Go services (Link, Metadata, BFF) use the Prometheus client library with `promauto` for automatic registration:

```go
import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var buildInfo = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "link",
		Subsystem: "application",
		Name:      "build_info",
		Help:      "Build information including version, commit, and build time",
	},
	[]string{"version", "commit", "build_time"},
)

// Set build info at startup
func SetBuildInfo(version, commit, buildTime string) {
	buildInfo.WithLabelValues(version, commit, buildTime).Set(1)
}
```

### TypeScript/Node.js Services

The Proxy service uses OpenTelemetry Metrics API:

```typescript
import { metrics } from "@opentelemetry/api";

export class StandardMetrics {
  private readonly meter = metrics.getMeter("proxy-service", "1.0.0");

  private readonly buildInfo = this.meter.createGauge(
    "proxy_application_build_info",
    {
      description: "Build information including version, commit, and build time",
    }
  );

  setBuildInfo(version: string, commit: string, buildTime: string): void {
    this.buildInfo.record(1, { version, commit, build_time: buildTime });
  }
}
```

## Querying Metrics

### PromQL Examples

**Get service versions:**
```promql
{__name__=~".*_application_build_info"}
```

**Calculate cache hit ratio:**
```promql
sum(rate(link_cache_hit_total[5m])) / 
(sum(rate(link_cache_hit_total[5m])) + sum(rate(link_cache_miss_total[5m])))
```

**Track error rate:**
```promql
rate(link_application_error_rate_per_minute[5m])
```

**95th percentile response time:**
```promql
histogram_quantile(0.95, rate(link_application_http_request_duration_seconds_bucket[5m]))
```

## Best Practices

1. **Consistent Naming**: Follow the namespace_subsystem_name pattern
2. **Use Labels**: Add dimensions via labels for better filtering
3. **Document Metrics**: Include help text for all metrics
4. **Monitor Cardinality**: Avoid high-cardinality label values
5. **Set Appropriate Buckets**: Use suitable histogram buckets for duration metrics

## References

- [ADR-0014: Standardizing Observability Tools](../ADR/decisions/0014-observability.md)
- [ADR-0018: Standard Metrics for Services](../ADR/decisions/0018-service-metrics.md)
- [ADR-0023: Prometheus Metrics Naming](../ADR/decisions/0023-naming-prometheus-metrics.md)
- [Prometheus Best Practices](https://prometheus.io/docs/practices/naming/)
