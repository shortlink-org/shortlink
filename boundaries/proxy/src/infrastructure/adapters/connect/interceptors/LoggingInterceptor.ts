import type { Interceptor } from "@connectrpc/connect";
import { ILogger } from "../../../logging/ILogger.js";

/**
 * Connect interceptor для логирования запросов и ответов
 * Логирует информацию о каждом Connect вызове
 */
export function createLoggingInterceptor(logger: ILogger): Interceptor {
  return (next) => async (req) => {
    const startTime = Date.now();
    // В Connect 2.x req.method может иметь разную структуру
    // Используем безопасный доступ к свойствам
    const method = (req as any).method?.name || "unknown";
    const service = (req as any).method?.service?.typeName || "unknown";

    // Логируем начало запроса
    logger.debug("Connect request started", {
      service,
      method,
      url: (req as any).url || "unknown",
    });

    try {
      const response = await next(req);

      // Логируем успешный ответ
      const duration = Date.now() - startTime;
      logger.debug("Connect request completed", {
        service,
        method,
        duration,
        status: "success",
      });

      return response;
    } catch (error: any) {
      // Логируем ошибку
      const duration = Date.now() - startTime;
      logger.error("Connect request failed", error, {
        service,
        method,
        duration,
        status: "error",
        errorCode: error?.code,
        errorMessage: error?.message,
      });

      throw error;
    }
  };
}

