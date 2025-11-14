import { ApplicationError } from "./ApplicationError.js";

/**
 * Ошибка валидации входных данных
 * Используется для ошибок валидации HTTP DTOs
 */
export class ValidationError extends ApplicationError {
  constructor(
    message: string,
    public readonly field?: string,
    public readonly details?: Record<string, unknown>
  ) {
    super(message, "VALIDATION_ERROR", 400);
    this.name = "ValidationError";
    Object.setPrototypeOf(this, ValidationError.prototype);
  }
}

