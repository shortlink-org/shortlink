import { IUseCaseInterceptor, UseCaseExecutionContext } from "./IUseCaseInterceptor.js";
import { metrics, Meter, Counter, Histogram } from "@opentelemetry/api";

/**
 * Интерцептор для сбора метрик выполнения Use Cases
 * Собирает метрики: количество запросов, время выполнения, ошибки
 */
export class MetricsInterceptor<TRequest = any, TResponse = any>
  implements IUseCaseInterceptor<TRequest, TResponse>
{
  private readonly meter: Meter;
  private readonly requestCounter: Counter;
  private readonly durationHistogram: Histogram;
  private readonly errorCounter: Counter;

  constructor() {
    // Получаем глобальный Meter из OpenTelemetry
    this.meter = metrics.getMeter("proxy-service", "1.0.0");

    // Создаем счетчик запросов
    this.requestCounter = this.meter.createCounter("usecase_requests_total", {
      description: "Total number of Use Case requests",
      unit: "1",
    });

    // Создаем гистограмму времени выполнения
    this.durationHistogram = this.meter.createHistogram(
      "usecase_duration_ms",
      {
        description: "Duration of Use Case execution in milliseconds",
        unit: "ms",
      }
    );

    // Создаем счетчик ошибок
    this.errorCounter = this.meter.createCounter("usecase_errors_total", {
      description: "Total number of Use Case errors",
      unit: "1",
    });
  }

  before(context: UseCaseExecutionContext<TRequest, TResponse>): TRequest {
    // Метрики записываются после выполнения
    return context.request;
  }

  after(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    const duration = context.duration || 0;

    // Записываем метрики успешного выполнения
    this.requestCounter.add(1, {
      usecase: context.useCaseName,
      status: "success",
    });

    this.durationHistogram.record(duration, {
      usecase: context.useCaseName,
    });
  }

  onError(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    const duration = context.duration || 0;

    // Записываем метрики ошибки
    this.requestCounter.add(1, {
      usecase: context.useCaseName,
      status: "error",
    });

    this.errorCounter.add(1, {
      usecase: context.useCaseName,
      error_type: context.error?.constructor.name || "Unknown",
    });

    this.durationHistogram.record(duration, {
      usecase: context.useCaseName,
    });
  }
}

