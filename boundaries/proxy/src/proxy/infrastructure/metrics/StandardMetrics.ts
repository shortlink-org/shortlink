import { metrics, ValueType } from "@opentelemetry/api";

/**
 * Standard metrics implementation according to ADR-0018
 * Provides consistent metrics across all services
 */
export class StandardMetrics {
  private readonly meter = metrics.getMeter("proxy-service", "1.0.0");

  // Build Info Metric (ADR-0014)
  private readonly buildInfo = this.meter.createGauge("proxy_application_build_info", {
    description: "Build information about the proxy service including version, commit, and build time",
    valueType: ValueType.INT,
  });

  // SLA/SLO/SLI Metrics (ADR-0018)
  private readonly serviceAvailabilityRatio = this.meter.createGauge(
    "proxy_application_service_availability_ratio",
    {
      description: "Ratio of uptime to total time for SLA tracking",
      valueType: ValueType.DOUBLE,
    }
  );

  private readonly httpRequestDuration = this.meter.createHistogram(
    "proxy_application_http_request_duration_seconds",
    {
      description: "HTTP request response times for SLO tracking",
      unit: "s",
      valueType: ValueType.DOUBLE,
    }
  );

  private readonly errorRatePerMinute = this.meter.createCounter("proxy_application_error_rate_per_minute", {
    description: "Number of errors per minute for system reliability tracking",
    unit: "1",
    valueType: ValueType.INT,
  });

  // Cache Metrics (ADR-0018)
  private readonly cacheHitTotal = this.meter.createCounter("proxy_cache_hit_total", {
    description: "Total number of cache hits",
    unit: "1",
    valueType: ValueType.INT,
  });

  private readonly cacheMissTotal = this.meter.createCounter("proxy_cache_miss_total", {
    description: "Total number of cache misses",
    unit: "1",
    valueType: ValueType.INT,
  });

  // Basic Service Metrics (ADR-0018)
  private readonly requestsPerSecond = this.meter.createCounter("proxy_application_requests_per_second", {
    description: "Service request load per second",
    unit: "1",
    valueType: ValueType.INT,
  });

  private readonly responseTimeSeconds = this.meter.createHistogram("proxy_application_response_time_seconds", {
    description: "Service response time",
    unit: "s",
    valueType: ValueType.DOUBLE,
  });

  private readonly errorRatePercentage = this.meter.createGauge("proxy_application_error_rate_percentage", {
    description: "Percentage of error requests",
    valueType: ValueType.DOUBLE,
  });

  /**
   * Set build information metric
   */
  setBuildInfo(version: string, commit: string, buildTime: string): void {
    this.buildInfo.record(1, {
      version,
      commit,
      build_time: buildTime,
    });
  }

  /**
   * Set service availability ratio
   */
  setServiceAvailability(ratio: number): void {
    this.serviceAvailabilityRatio.record(ratio);
  }

  /**
   * Record HTTP request duration
   */
  recordHttpRequestDuration(method: string, endpoint: string, status: string, duration: number): void {
    this.httpRequestDuration.record(duration, {
      method,
      endpoint,
      status,
    });
  }

  /**
   * Increment error rate counter
   */
  incrementErrorRate(errorType: string, operation: string): void {
    this.errorRatePerMinute.add(1, {
      type: errorType,
      operation,
    });
  }

  /**
   * Record cache hit
   */
  recordCacheHit(cacheType: string): void {
    this.cacheHitTotal.add(1, {
      cache_type: cacheType,
    });
  }

  /**
   * Record cache miss
   */
  recordCacheMiss(cacheType: string): void {
    this.cacheMissTotal.add(1, {
      cache_type: cacheType,
    });
  }

  /**
   * Increment requests per second counter
   */
  incrementRequestsPerSecond(operation: string): void {
    this.requestsPerSecond.add(1, {
      operation,
    });
  }

  /**
   * Record response time
   */
  recordResponseTime(operation: string, duration: number): void {
    this.responseTimeSeconds.record(duration, {
      operation,
    });
  }

  /**
   * Set error rate percentage
   */
  setErrorRatePercentage(operation: string, percentage: number): void {
    this.errorRatePercentage.record(percentage, {
      operation,
    });
  }
}
