import { Span } from "@opentelemetry/api";

/**
 * Интерфейс для трейсинга gRPC вызовов
 * Абстрагирует создание и управление OpenTelemetry spans
 */
export interface IGrpcTracing {
  /**
   * Создает span для gRPC вызова
   * @param method - Название gRPC метода
   * @param operation - Название операции (например, 'getLinkByHash')
   * @returns Span для трейсинга
   */
  startSpan(method: string, operation: string): Span;

  /**
   * Завершает span с успешным результатом
   * @param span - Span для завершения
   */
  endSpan(span: Span): void;

  /**
   * Завершает span с ошибкой
   * @param span - Span для завершения
   * @param error - Ошибка
   */
  endSpanWithError(span: Span, error: Error): void;
}

