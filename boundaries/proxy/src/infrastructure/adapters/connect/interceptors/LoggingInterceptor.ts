import type { Interceptor } from "@connectrpc/connect";
import { ILogger } from "../../../logging/ILogger.js";

/**
 * Connect interceptor for logging requests and responses
 * Logs information about each Connect call
 */
export function createLoggingInterceptor(logger: ILogger): Interceptor {
  return (next) => async (req) => {
    const startTime = Date.now();
    // In Connect 2.x req.method may have different structure
    // Use safe property access
    const method = (req as any).method?.name || "unknown";
    const service = (req as any).method?.service?.typeName || "unknown";

    // Log request start
    logger.debug("connect.request.started", {
      event: "connect.request.started",
      service,
      method,
      url: (req as any).url || "unknown",
    });

    try {
      const response = await next(req);

      // Log successful response
      const duration = Date.now() - startTime;
      logger.debug("connect.request.completed", {
        event: "connect.request.completed",
        service,
        method,
        durationMs: duration,
        status: "success",
      });

      return response;
    } catch (error: any) {
      // Log error
      // error object will be automatically serialized by WinstonLogger to { name, message, stack }
      const duration = Date.now() - startTime;
      logger.error("connect.request.failed", {
        event: "connect.request.failed",
        service,
        method,
        durationMs: duration,
        status: "error",
        errorCode: error?.code,
        error: error instanceof Error ? error : new Error(String(error)),
      });

      throw error;
    }
  };
}
