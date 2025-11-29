import type { FastifyRequest } from "fastify";
import { ConnectError, Code } from "@connectrpc/connect";
import { DomainError } from "../../../../domain/exceptions/index.js";
import { InvalidHashError } from "../../../../domain/exceptions/index.js";
import { LinkNotFoundError } from "../../../../domain/exceptions/index.js";
import {
  ApplicationError,
  ValidationError,
  InfrastructureError,
  ExternalServiceError,
} from "../../../../application/exceptions/index.js";
import type { ILogger } from "../../../logging/ILogger.js";

/**
 * Error response structure
 */
export interface ErrorResponse {
  statusCode: number;
  payload: {
    error: {
      code: string;
      message: string;
      field?: string;
      details?: Record<string, unknown>;
      service?: string;
      timestamp: string;
    };
  };
}

/**
 * Maps domain and application errors to HTTP responses.
 * Transforms typed errors into standardized HTTP responses for Fastify.
 */
export class ErrorMapper {
  constructor(private readonly logger: ILogger) {}

  /**
   * Maps error to HTTP response
   */
  mapToHttpResponse(error: unknown, request?: FastifyRequest): ErrorResponse {
    // Handle Connect/gRPC errors
    // According to ADR 42: PermissionDenied should return 404 Not Found
    // Also handle InvalidArgument with "permission denied" message (legacy error mapping issue)
    if (
      error instanceof ConnectError &&
      (error.code === Code.PermissionDenied ||
        (error.code === Code.InvalidArgument &&
          error.message.toLowerCase().includes("permission denied")))
    ) {
      return {
        statusCode: 404,
        payload: {
          error: {
            code: "LINK_NOT_FOUND",
            message: "Link not found",
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    // Domain Errors
    if (error instanceof InvalidHashError) {
      return {
        statusCode: 400,
        payload: {
          error: {
            code: "INVALID_HASH",
            message: error.message,
            field: "hash",
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    if (error instanceof LinkNotFoundError) {
      return {
        statusCode: 404,
        payload: {
          error: {
            code: "LINK_NOT_FOUND",
            message: error.message,
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    if (error instanceof DomainError) {
      return {
        statusCode: 400,
        payload: {
          error: {
            code: "DOMAIN_ERROR",
            message: error.message,
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    // Application Errors
    if (error instanceof ValidationError) {
      return {
        statusCode: error.statusCode,
        payload: {
          error: {
            code: error.code,
            message: error.message,
            field: error.field,
            details: error.details,
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    if (error instanceof ExternalServiceError) {
      this.logger.warn("http.error.external_service", {
        event: "http.error.external_service",
        service: error.service,
        statusCode: error.statusCode,
        message: error.message,
        path: request?.url,
        method: request?.method,
        error: error.originalError || error,
      });
      const statusCode = error.statusCode ?? 503;
      return {
        statusCode,
        payload: {
          error: {
            code: error.code,
            message: error.message,
            service: error.service,
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    if (error instanceof InfrastructureError) {
      this.logger.error("http.error.infrastructure", {
        event: "http.error.infrastructure",
        service: error.service,
        path: request?.url,
        method: request?.method,
        statusCode: error.statusCode,
        errorCode: error.code,
        error: error.originalError || error,
      });
      return {
        statusCode: error.statusCode,
        payload: {
          error: {
            code: error.code,
            message: error.message,
            service: error.service,
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    if (error instanceof ApplicationError) {
      return {
        statusCode: error.statusCode,
        payload: {
          error: {
            code: error.code,
            message: error.message,
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    // Generic Error
    if (error instanceof Error) {
      this.logger.error("http.error.unhandled", {
        event: "http.error.unhandled",
        path: request?.url,
        method: request?.method,
        statusCode: 500,
        error: error,
      });
      return {
        statusCode: 500,
        payload: {
          error: {
            code: "INTERNAL_SERVER_ERROR",
            message: "An unexpected error occurred",
            timestamp: new Date().toISOString(),
          },
        },
      };
    }

    // Unknown error type
    this.logger.error("http.error.unknown", {
      event: "http.error.unknown",
      path: request?.url,
      method: request?.method,
      statusCode: 500,
      error: error instanceof Error ? error : new Error(String(error)),
      errorType: typeof error,
      errorString: String(error),
    });
    return {
      statusCode: 500,
      payload: {
        error: {
          code: "INTERNAL_SERVER_ERROR",
          message: "An unexpected error occurred",
          timestamp: new Date().toISOString(),
        },
      },
    };
  }
}
