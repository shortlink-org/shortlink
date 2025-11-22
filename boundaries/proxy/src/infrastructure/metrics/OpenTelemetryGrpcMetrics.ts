import { metrics, Meter, Counter, Histogram } from "@opentelemetry/api";
import { IGrpcMetrics } from "./IGrpcMetrics.js";

/**
 * Реализация метрик gRPC вызовов через OpenTelemetry Metrics API
 * Использует OpenTelemetry для сбора метрик, которые можно экспортировать в Prometheus/Grafana
 */
export class OpenTelemetryGrpcMetrics implements IGrpcMetrics {
  private readonly meter: Meter;
  private readonly requestCounter: Counter;
  private readonly durationHistogram: Histogram;
  private readonly errorCounter: Counter;

  constructor() {
    // Получаем глобальный Meter из OpenTelemetry
    this.meter = metrics.getMeter("proxy-service", "1.0.0");

    // Создаем счетчик запросов
    this.requestCounter = this.meter.createCounter("grpc_requests_total", {
      description: "Total number of gRPC requests",
      unit: "1",
    });

    // Создаем гистограмму времени выполнения
    this.durationHistogram = this.meter.createHistogram("grpc_request_duration_ms", {
      description: "Duration of gRPC requests in milliseconds",
      unit: "ms",
    });

    // Создаем счетчик ошибок
    this.errorCounter = this.meter.createCounter("grpc_errors_total", {
      description: "Total number of gRPC errors",
      unit: "1",
    });
  }

  recordRequest(method: string, status: "success" | "error" | "not_found"): void {
    this.requestCounter.add(1, {
      method,
      status,
      service: "link-service",
    });
  }

  recordDuration(method: string, durationMs: number): void {
    this.durationHistogram.record(durationMs, {
      method,
      service: "link-service",
    });
  }

  recordError(method: string, errorCode: number): void {
    this.errorCounter.add(1, {
      method,
      error_code: errorCode.toString(),
      service: "link-service",
    });
  }
}

