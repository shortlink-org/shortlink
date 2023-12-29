# 18. Standard Metrics for Services using Prometheus

Date: 2023-12-29

## Status

Accepted

## Context

The drive for standardizing metrics in our services is informed by the need for clear, consistent, 
and actionable data that can guide operational decisions and provide insights into service performance. 
Prometheus, with its robust monitoring capabilities, offers a suitable platform for this endeavor. 
This decision also aligns with our commitment to maintaining high service availability and reliability and is influenced 
by our previous ADR on Prometheus Metrics Naming [(ADR #23)](./0023-naming-prometheus-metrics.md).

## Decision

We have decided to adopt standard metrics for our services using Prometheus.

### Enhancements to Prometheus Metric Examples

#### SLA/SLO/SLI Metrics

- **`service_availability_ratio`**: A gauge metric representing the ratio of uptime to the total time, aligning with our SLA uptime guarantees.
- **`http_request_duration_seconds`**: A histogram metric measuring the request response times, aiding in tracking our SLOs related to response time.
- **`error_rate_per_minute`**: A counter metric tracking the number of errors per minute, an SLI for system reliability.

#### Cache Metrics

- **`cache_hit_total`** and **`cache_miss_total`**: Counter metrics for the total number of cache hits and misses, respectively.
- **`cache_hit_ratio`**: A calculated ratio from `cache_hit_total` and `cache_miss_total` to determine the effectiveness of the cache.
- **`hot_key_access_frequency`**: A gauge metric indicating the frequency of access for hot keys.

#### Basic Service Metrics
- 
- **`requests_per_second` (RPS)**: Counter metric measuring service request load per second.
- **`transactions_per_second` (TPS)** & **`queries_per_second` (QPS)**: Counter metrics for transactions and database queries per second, respectively.
- **`response_time_seconds`**: Histogram metric tracking service response time.
- **`error_rate_percentage`**: Gauge metric for percentage of error requests.

#### Resource Utilization Metrics
- 
- **`network_traffic_bytes`**: Counter metric for inbound and outbound network traffic.
- **`cpu_usage_percentage`**: Gauge metric for CPU utilization.
- **`ram_usage_bytes`**: Gauge metric for RAM usage.
- **`disk_usage_bytes`**: Gauge metric for disk space usage (HDD/SSD).

#### Additional Metrics

- **`queue_size`**: Gauge metric for the size of each critical queue.
- **`process_count`** & **`thread_count`**: Gauge metrics for monitoring the number of processes and threads.

### Reference to ADR #23: Prometheus Metrics Naming

Consistent with [ADR #23](./0023-naming-prometheus-metrics.md), all metrics will follow the prescribed naming conventions and utilize labels for additional dimensions.
This will enhance clarity, ease of understanding, and consistency in metric categorization.

## Consequences

This standardization will enable systematic monitoring and improvement of service performance. 
However, challenges include ensuring accuracy in distributed systems and avoiding over-reliance on quantitative metrics. 
These will be mitigated through continuous monitoring strategy refinement and periodic metric reviews.
