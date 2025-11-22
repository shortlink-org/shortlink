import { ApplicationError } from "./ApplicationError.js";

/**
 * Ошибка инфраструктуры (БД, внешние сервисы, сеть)
 * Используется для ошибок, возникающих в Infrastructure Layer
 */
export class InfrastructureError extends ApplicationError {
  constructor(
    message: string,
    public readonly service?: string,
    public readonly originalError?: Error,
    code: string = "INFRASTRUCTURE_ERROR",
    statusCode: number = 503
  ) {
    super(message, code, statusCode);
    this.name = "InfrastructureError";
    Object.setPrototypeOf(this, InfrastructureError.prototype);
  }
}

/**
 * Ошибка внешнего сервиса
 */
export class ExternalServiceError extends InfrastructureError {
  constructor(
    message: string,
    service: string,
    statusCode?: number,
    originalError?: Error
  ) {
    super(
      message,
      service,
      originalError,
      "EXTERNAL_SERVICE_ERROR",
      statusCode ?? 503
    );
    this.name = "ExternalServiceError";
    Object.setPrototypeOf(this, ExternalServiceError.prototype);
  }
}


