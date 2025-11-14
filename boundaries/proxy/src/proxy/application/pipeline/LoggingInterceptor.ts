import { injectable, inject } from "inversify";
import { IUseCaseInterceptor, UseCaseExecutionContext } from "./IUseCaseInterceptor.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import TYPES from "../../../types.js";

/**
 * Интерцептор для логирования выполнения Use Cases
 * Логирует вход, выход, ошибки и время выполнения
 */
@injectable()
export class LoggingInterceptor<TRequest = any, TResponse = any>
  implements IUseCaseInterceptor<TRequest, TResponse>
{
  constructor(
    @inject(TYPES.INFRASTRUCTURE.Logger) private readonly logger: ILogger
  ) {}

  before(context: UseCaseExecutionContext<TRequest, TResponse>): TRequest {
    this.logger.debug(`[UseCase] Starting: ${context.useCaseName}`, {
      useCase: context.useCaseName,
      request: this.sanitizeRequest(context.request),
    });
    return context.request;
  }

  after(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    const duration = context.duration || 0;
    this.logger.info(`[UseCase] Completed: ${context.useCaseName}`, {
      useCase: context.useCaseName,
      duration: `${duration}ms`,
      success: true,
    });
  }

  onError(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    const duration = context.duration || 0;
    this.logger.error(
      `[UseCase] Failed: ${context.useCaseName}`,
      context.error,
      {
        useCase: context.useCaseName,
        duration: `${duration}ms`,
        success: false,
      }
    );
  }

  finally(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    // Дополнительная очистка при необходимости
  }

  /**
   * Очищает запрос от чувствительных данных для логирования
   */
  private sanitizeRequest(request: any): any {
    if (!request || typeof request !== "object") {
      return request;
    }

    const sanitized = { ...request };
    const sensitiveFields = ["password", "token", "secret", "apiKey"];

    for (const field of sensitiveFields) {
      if (field in sanitized) {
        sanitized[field] = "***REDACTED***";
      }
    }

    return sanitized;
  }
}

