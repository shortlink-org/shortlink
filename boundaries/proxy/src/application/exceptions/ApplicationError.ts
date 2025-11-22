/**
 * Базовый класс для всех ошибок Application Layer
 * Используется для типизации и централизованной обработки ошибок приложения
 */
export class ApplicationError extends Error {
  constructor(
    message: string,
    public readonly code: string = "APPLICATION_ERROR",
    public readonly statusCode: number = 500
  ) {
    super(message);
    this.name = "ApplicationError";
    Object.setPrototypeOf(this, ApplicationError.prototype);
  }
}

