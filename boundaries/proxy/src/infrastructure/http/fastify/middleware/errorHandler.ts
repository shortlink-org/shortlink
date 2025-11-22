import type { FastifyError, FastifyRequest, FastifyReply } from "fastify";
import type { ILogger } from "../../../logging/ILogger.js";
import { ErrorMapper } from "./ErrorMapper.js";

/**
 * Fastify error handler middleware.
 * Centralized error handling for all routes.
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
    // Map error to HTTP response
    const { statusCode, payload } = errorMapper.mapToHttpResponse(error, request);

    reply.status(statusCode).send(payload);
  };
}

