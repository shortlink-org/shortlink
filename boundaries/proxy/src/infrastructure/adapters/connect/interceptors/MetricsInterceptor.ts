import type { Interceptor } from "@connectrpc/connect";
import { IGrpcMetrics } from "../../../metrics/IGrpcMetrics.js";

/**
 * Connect interceptor для записи метрик
 * Записывает метрики для каждого Connect вызова
 */
export function createMetricsInterceptor(metrics: IGrpcMetrics): Interceptor {
  return (next) => async (req) => {
    const startTime = Date.now();
    const method = (req as any).method?.name || "unknown";

    try {
      const response = await next(req);

      // Записываем метрики для успешного запроса
      const duration = Date.now() - startTime;
      metrics.recordRequest(method, "success");
      metrics.recordDuration(method, duration);

      return response;
    } catch (error: any) {
      // Записываем метрики для ошибки
      const duration = Date.now() - startTime;
      metrics.recordRequest(method, "error");
      metrics.recordDuration(method, duration);

      // Определяем код ошибки
      let errorCode = -1;
      if (typeof error?.code === "number") {
        errorCode = error.code;
      } else if (error?.status) {
        errorCode = error.status;
      }

      metrics.recordError(method, errorCode);

      // Проверяем, является ли это NOT_FOUND
      if (
        error?.code === "NOT_FOUND" ||
        error?.code === 5 ||
        error?.status === 404
      ) {
        metrics.recordRequest(method, "not_found");
      }

      throw error;
    }
  };
}

