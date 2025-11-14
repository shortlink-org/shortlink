import * as express from "express";
import { DomainError } from "../../../proxy/domain/exceptions/index.js";
import { InvalidHashError } from "../../../proxy/domain/exceptions/index.js";
import { LinkNotFoundError } from "../../../proxy/domain/exceptions/index.js";
import {
  ApplicationError,
  ValidationError,
  InfrastructureError,
  ExternalServiceError,
} from "../../../proxy/application/exceptions/index.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";

/**
 * Маппинг доменных и прикладных ошибок в HTTP ответы
 * Преобразует типизированные ошибки в стандартизированные HTTP ответы
 */
export class ErrorMapper {
  constructor(private readonly logger: ILogger) {}

  /**
   * Преобразует ошибку в HTTP ответ
   */
  mapToHttpResponse(
    error: unknown,
    req: express.Request,
    res: express.Response
  ): void {
    // Domain Errors
    if (error instanceof InvalidHashError) {
      this.sendError(res, 400, {
        code: "INVALID_HASH",
        message: error.message,
        field: "hash",
      });
      return;
    }

    if (error instanceof LinkNotFoundError) {
      this.sendError(res, 404, {
        code: "LINK_NOT_FOUND",
        message: error.message,
      });
      return;
    }

    if (error instanceof DomainError) {
      this.sendError(res, 400, {
        code: "DOMAIN_ERROR",
        message: error.message,
      });
      return;
    }

    // Application Errors
    if (error instanceof ValidationError) {
      this.sendError(res, error.statusCode, {
        code: error.code,
        message: error.message,
        field: error.field,
        details: error.details,
      });
      return;
    }

    if (error instanceof ExternalServiceError) {
      this.logger.warn("External service error", {
        service: error.service,
        statusCode: error.statusCode,
        message: error.message,
        path: req.path,
      });
      this.sendError(res, error.statusCode, {
        code: error.code,
        message: error.message,
        service: error.service,
      });
      return;
    }

    if (error instanceof InfrastructureError) {
      this.logger.error("Infrastructure error", error.originalError || error, {
        service: error.service,
        path: req.path,
        method: req.method,
      });
      this.sendError(res, error.statusCode, {
        code: error.code,
        message: error.message,
        service: error.service,
      });
      return;
    }

    if (error instanceof ApplicationError) {
      this.sendError(res, error.statusCode, {
        code: error.code,
        message: error.message,
      });
      return;
    }

    // Generic Error
    if (error instanceof Error) {
      this.logger.error("Unhandled error", error, {
        path: req.path,
        method: req.method,
      });
      this.sendError(res, 500, {
        code: "INTERNAL_SERVER_ERROR",
        message: "An unexpected error occurred",
      });
      return;
    }

    // Unknown error type
    this.logger.error("Unknown error type", error, {
      path: req.path,
      method: req.method,
    });
    this.sendError(res, 500, {
      code: "INTERNAL_SERVER_ERROR",
      message: "An unexpected error occurred",
    });
  }

  /**
   * Отправляет стандартизированный HTTP ответ с ошибкой
   */
  private sendError(
    res: express.Response,
    statusCode: number,
    body: {
      code: string;
      message: string;
      field?: string;
      details?: Record<string, unknown>;
    }
  ): void {
    res.status(statusCode).json({
      error: {
        ...body,
        timestamp: new Date().toISOString(),
      },
    });
  }
}

