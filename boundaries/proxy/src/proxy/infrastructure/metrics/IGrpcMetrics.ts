/**
 * Интерфейс для метрик gRPC вызовов
 * Абстрагирует реализацию метрик для тестирования и замены
 */
export interface IGrpcMetrics {
  /**
   * Увеличивает счетчик запросов
   * @param method - Название gRPC метода
   * @param status - Статус ответа ('success' | 'error' | 'not_found')
   */
  recordRequest(method: string, status: "success" | "error" | "not_found"): void;

  /**
   * Записывает время выполнения запроса
   * @param method - Название gRPC метода
   * @param durationMs - Время выполнения в миллисекундах
   */
  recordDuration(method: string, durationMs: number): void;

  /**
   * Увеличивает счетчик ошибок
   * @param method - Название gRPC метода
   * @param errorCode - Код ошибки gRPC
   */
  recordError(method: string, errorCode: number): void;
}

