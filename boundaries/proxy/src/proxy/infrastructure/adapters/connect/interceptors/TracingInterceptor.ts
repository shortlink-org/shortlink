import type { Interceptor } from "@connectrpc/connect";
import { IGrpcTracing } from "../../../tracing/IGrpcTracing.js";

/**
 * Connect interceptor для трейсинга через OpenTelemetry
 * Создает spans для каждого Connect вызова
 */
export function createTracingInterceptor(tracing: IGrpcTracing): Interceptor {
  return (next) => async (req) => {
    const method = (req as any).method?.name || "unknown";
    const serviceName = (req as any).method?.service?.typeName || "unknown";
    const operation = `${serviceName}.${method}`;
    const span = tracing.startSpan(method, operation);

    // Устанавливаем атрибуты span
    span.setAttribute("rpc.service", serviceName);
    span.setAttribute("rpc.method", method);
    span.setAttribute("rpc.system", "connect");

    try {
      const response = await next(req);

      // Завершаем span с успешным результатом
      span.setAttribute("rpc.grpc.status_code", 0); // OK
      tracing.endSpan(span);

      return response;
    } catch (error: any) {
      // Определяем код ошибки для span
      let statusCode = -1;
      if (typeof error?.code === "number") {
        statusCode = error.code;
      } else if (error?.status) {
        statusCode = error.status;
      } else if (error?.code === "NOT_FOUND" || error?.code === 5) {
        statusCode = 5; // NOT_FOUND
      }

      span.setAttribute("rpc.grpc.status_code", statusCode);
      span.setAttribute("error", true);
      span.setAttribute("error.message", error?.message || String(error));

      // Завершаем span с ошибкой
      if (error instanceof Error) {
        tracing.endSpanWithError(span, error);
      } else {
        tracing.endSpan(span);
      }

      throw error;
    }
  };
}

