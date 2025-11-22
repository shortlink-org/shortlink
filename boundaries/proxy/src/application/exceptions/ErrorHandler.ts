import { DomainError } from "../../domain/exceptions/DomainError.js";
import { LinkNotFoundError } from "../../domain/exceptions/LinkNotFoundError.js";
import { InvalidHashError } from "../../domain/exceptions/InvalidHashError.js";
import { ApplicationError, ValidationError } from "./index.js";
import { ILogger } from "../../infrastructure/logging/ILogger.js";

/**
 * Утилита для централизованной обработки ошибок в Use Cases
 * Обеспечивает консистентную обработку ошибок и логирование
 */
export class ErrorHandler {
  constructor(private readonly logger?: ILogger) {}

  /**
   * Обрабатывает ошибку, выбрасываемую в Use Case
   * Логирует ошибку и пробрасывает её дальше для обработки в ErrorMapper
   *
   * @param error - ошибка для обработки
   * @param context - контекст выполнения (название Use Case, параметры запроса)
   */
  handleUseCaseError(
    error: unknown,
    context?: {
      useCaseName?: string;
      request?: unknown;
    }
  ): never {
    // Логируем ошибку с контекстом
    if (this.logger) {
      if (error instanceof DomainError || error instanceof ApplicationError) {
        // Доменные и прикладные ошибки логируем как warning (ожидаемые ошибки)
        this.logger.warn(`Use Case error: ${context?.useCaseName || "Unknown"}`, {
          error: error.message,
          code: error instanceof ApplicationError ? error.code : error.constructor.name,
          context,
        });
      } else {
        // Неожиданные ошибки логируем как error
        this.logger.error(`Unexpected error in Use Case: ${context?.useCaseName || "Unknown"}`, error, {
          context,
        });
      }
    }

    // Пробрасываем ошибку дальше для обработки в ErrorMapper
    throw error;
  }

  /**
   * Оборачивает выполнение Use Case с обработкой ошибок
   *
   * @param useCaseName - название Use Case для логирования
   * @param execute - функция выполнения Use Case
   * @param request - запрос для логирования
   */
  async executeWithErrorHandling<TRequest, TResponse>(
    useCaseName: string,
    execute: (request: TRequest) => Promise<TResponse>,
    request: TRequest
  ): Promise<TResponse> {
    try {
      return await execute(request);
    } catch (error) {
      this.handleUseCaseError(error, {
        useCaseName,
        request,
      });
    }
  }

  /**
   * Проверяет, является ли ошибка ожидаемой (Domain или Application ошибкой)
   */
  isExpectedError(error: unknown): error is DomainError | ApplicationError {
    return error instanceof DomainError || error instanceof ApplicationError;
  }

  /**
   * Проверяет, является ли ошибка критической (неожиданной)
   */
  isCriticalError(error: unknown): boolean {
    return !this.isExpectedError(error);
  }
}

