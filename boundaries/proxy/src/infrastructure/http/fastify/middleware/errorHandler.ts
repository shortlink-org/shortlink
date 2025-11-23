import type { FastifyError, FastifyRequest, FastifyReply } from "fastify";
import type { ILogger } from "../../../logging/ILogger.js";
import { ErrorMapper } from "./ErrorMapper.js";

/**
 * Fastify error handler middleware.
 * Centralized error handling for all routes with structured logging.
 * Automatically logs all HTTP errors with request context.
 *
 * @param logger - Logger instance
 * @returns Error handler function
 */
export function errorHandler(logger: ILogger) {
  const errorMapper = new ErrorMapper(logger);

  return async (
    error: FastifyError | Error,
    request: FastifyRequest,
    reply: FastifyReply
  ): Promise<void> => {
    const startTime = (request as any).startTime || Date.now();
    const duration = Date.now() - startTime;

    // Map error to HTTP response
    const { statusCode, payload } = errorMapper.mapToHttpResponse(
      error,
      request
    );

    // Log HTTP error with full context
    // error object will be automatically serialized by WinstonLogger to { name, message, stack }
    logger.error("http.error", {
      event: "http.error",
      method: request.method,
      url: request.url,
      path: (request as any).routerPath || request.url,
      statusCode,
      durationMs: duration,
      errorCode: payload.error.code,
      requestId: (request as any).id,
      userAgent: request.headers["user-agent"],
      ip: request.ip,
      error: error instanceof Error ? error : new Error(String(error)),
    });

    reply.status(statusCode).send(payload);
  };
}
