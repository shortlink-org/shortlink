import type { Interceptor } from "@connectrpc/connect";
import { ILogger } from "../../../../logging/ILogger.js";

/**
 * Конфигурация для retry interceptor
 */
export interface RetryConfig {
  /**
   * Максимальное количество попыток (включая первую)
   */
  maxAttempts: number;

  /**
   * Начальная задержка перед повторной попыткой в миллисекундах
   */
  initialDelayMs: number;

  /**
   * Максимальная задержка в миллисекундах (для exponential backoff)
   */
  maxDelayMs: number;

  /**
   * Множитель для exponential backoff
   */
  backoffMultiplier: number;

  /**
   * HTTP статус коды, при которых нужно повторять запрос
   * По умолчанию: 429 (Too Many Requests), 500-599 (Server Errors)
   */
  retryableStatusCodes?: number[];
}

const DEFAULT_RETRY_CONFIG: RetryConfig = {
  maxAttempts: 3,
  initialDelayMs: 100,
  maxDelayMs: 5000,
  backoffMultiplier: 2,
  retryableStatusCodes: [429, 500, 502, 503, 504],
};

/**
 * Проверяет, является ли ошибка временной и требует ли повторной попытки
 */
function isRetryableError(error: any, retryableStatusCodes: number[]): boolean {
  // Проверяем HTTP статус код
  if (error?.status && retryableStatusCodes.includes(error.status)) {
    return true;
  }

  // Проверяем Connect код ошибки
  // Connect использует gRPC коды: UNAVAILABLE (14), DEADLINE_EXCEEDED (4), RESOURCE_EXHAUSTED (8)
  if (error?.code === 14 || error?.code === 4 || error?.code === 8) {
    return true;
  }

  // Проверяем строковые коды Connect
  if (
    error?.code === "UNAVAILABLE" ||
    error?.code === "DEADLINE_EXCEEDED" ||
    error?.code === "RESOURCE_EXHAUSTED"
  ) {
    return true;
  }

  // Проверяем сетевые ошибки
  if (error?.name === "NetworkError" || error?.message?.includes("ECONNREFUSED")) {
    return true;
  }

  return false;
}

/**
 * Вычисляет задержку для exponential backoff
 */
function calculateDelay(attempt: number, config: RetryConfig): number {
  const delay = config.initialDelayMs * Math.pow(config.backoffMultiplier, attempt - 1);
  return Math.min(delay, config.maxDelayMs);
}

/**
 * Создает задержку на указанное количество миллисекунд
 */
function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/**
 * Connect interceptor для повторных попыток при временных ошибках
 * Использует exponential backoff для задержек между попытками
 */
export function createRetryInterceptor(
  logger: ILogger,
  config: Partial<RetryConfig> = {}
): Interceptor {
  const retryConfig = { ...DEFAULT_RETRY_CONFIG, ...config };

  return (next) => async (req) => {
    let attempt = 0;
    let lastError: any;

    while (attempt < retryConfig.maxAttempts) {
      attempt++;

      try {
        const response = await next(req);
        return response;
      } catch (error: any) {
        lastError = error;

        // Если это последняя попытка или ошибка не требует повторной попытки - пробрасываем ошибку
        if (
          attempt >= retryConfig.maxAttempts ||
          !isRetryableError(error, retryConfig.retryableStatusCodes!)
        ) {
          if (attempt > 1) {
            logger.warn("Connect request failed after retries", {
              method: (req as any).method?.name || "unknown",
              service: (req as any).method?.service?.typeName || "unknown",
              attempts: attempt,
              error: error?.message,
            });
          }
          throw error;
        }

        // Вычисляем задержку для следующей попытки
        const delay = calculateDelay(attempt, retryConfig);

        logger.debug("Connect request retry", {
          method: (req as any).method?.name || "unknown",
          service: (req as any).method?.service?.typeName || "unknown",
          attempt,
          maxAttempts: retryConfig.maxAttempts,
          delay,
          error: error?.message,
        });

        // Ждем перед следующей попыткой
        await sleep(delay);
      }
    }

    // Это не должно произойти, но на всякий случай
    throw lastError || new Error("Connect request failed after all retries");
  };
}

