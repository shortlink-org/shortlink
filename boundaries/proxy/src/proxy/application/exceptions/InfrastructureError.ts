import { ApplicationError } from "./ApplicationError.js";

/**
 * Ошибка инфраструктуры (БД, внешние сервисы, сеть)
 * Используется для ошибок, возникающих в Infrastructure Layer
 */
export class InfrastructureError extends ApplicationError {
  constructor(
    message: string,
    public readonly service?: string,
    public readonly originalError?: Error
  ) {
    super(message, "INFRASTRUCTURE_ERROR", 503);
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
    public readonly statusCode?: number,
    originalError?: Error
  ) {
    super(message, service, originalError);
    this.code = "EXTERNAL_SERVICE_ERROR";
    this.statusCode = statusCode || 503;
    this.name = "ExternalServiceError";
    Object.setPrototypeOf(this, ExternalServiceError.prototype);
  }
}


