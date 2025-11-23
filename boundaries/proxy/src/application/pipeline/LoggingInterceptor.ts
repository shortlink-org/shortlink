import {
  IUseCaseInterceptor,
  UseCaseExecutionContext,
} from "./IUseCaseInterceptor.js";
import { ILogger } from "../../infrastructure/logging/ILogger.js";

/**
 * Interceptor for automatic logging of Use Case execution
 * Logs entry, exit, errors, and execution time
 * Uses improved logger with OTEL-compatible error serialization
 */
export class LoggingInterceptor<TRequest = any, TResponse = any>
  implements IUseCaseInterceptor<TRequest, TResponse>
{
  constructor(private readonly logger: ILogger) {}

  before(context: UseCaseExecutionContext<TRequest, TResponse>): TRequest {
    this.logger.debug(`UseCase started: ${context.useCaseName}`, {
      event: "usecase.start",
      useCase: context.useCaseName,
      request: this.sanitizeRequest(context.request),
    });
    return context.request;
  }

  after(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    const duration = context.duration || 0;
    this.logger.info(`UseCase completed: ${context.useCaseName}`, {
      event: "usecase.success",
      useCase: context.useCaseName,
      durationMs: duration,
      success: true,
    });
  }

  onError(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    const duration = context.duration || 0;
    this.logger.error(`UseCase failed: ${context.useCaseName}`, {
      event: "usecase.error",
      useCase: context.useCaseName,
      durationMs: duration,
      success: false,
      error: context.error,
    });
  }

  finally(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    // Additional cleanup if needed
  }

  /**
   * List of sensitive field names that should be redacted
   */
  private readonly sensitiveFields = [
    "password",
    "token",
    "secret",
    "apiKey",
    "api_key",
    "accessToken",
    "refreshToken",
    "authorization",
    "auth",
    "credentials",
    "privateKey",
    "private_key",
    "sessionId",
    "session_id",
  ];

  /**
   * Checks if a field name is sensitive and should be redacted
   */
  private isSensitiveField(fieldName: string): boolean {
    const lowerFieldName = fieldName.toLowerCase();
    return this.sensitiveFields.some((sensitive) =>
      lowerFieldName.includes(sensitive.toLowerCase())
    );
  }

  /**
   * Checks if value is a built-in JavaScript object that should not be recursively processed
   */
  private isBuiltInObject(value: any): boolean {
    return (
      value instanceof Date ||
      value instanceof URL ||
      value instanceof RegExp ||
      value instanceof Map ||
      value instanceof Set ||
      (typeof Buffer !== "undefined" && Buffer.isBuffer(value)) ||
      value instanceof Uint8Array ||
      value instanceof ArrayBuffer
    );
  }

  /**
   * Recursively sanitizes request by removing sensitive data for logging
   * Prevents logging of passwords, tokens, secrets, and API keys at any nesting level
   * Handles nested objects, arrays, and built-in types safely
   */
  private sanitizeRequest(request: any, visited = new WeakSet()): any {
    // Handle null, undefined, primitives
    if (request === null || request === undefined) {
      return request;
    }

    if (typeof request !== "object") {
      return request;
    }

    // Handle built-in objects (don't recurse into them)
    if (this.isBuiltInObject(request)) {
      return request;
    }

    // Handle arrays - recursively sanitize each element
    if (Array.isArray(request)) {
      return request.map((item) => this.sanitizeRequest(item, visited));
    }

    // Handle circular references
    if (visited.has(request)) {
      return "[Circular Reference]";
    }

    visited.add(request);

    try {
      const sanitized: any = {};

      // Process all properties recursively
      for (const [key, value] of Object.entries(request)) {
        // Check if field name is sensitive (case-insensitive, partial match)
        if (this.isSensitiveField(key)) {
          sanitized[key] = "***REDACTED***";
        } else if (value === null || value === undefined) {
          sanitized[key] = value;
        } else if (typeof value === "object") {
          // Recursively sanitize nested objects
          sanitized[key] = this.sanitizeRequest(value, visited);
        } else {
          // Primitive values - keep as is
          sanitized[key] = value;
        }
      }

      return sanitized;
    } catch (err) {
      // If serialization fails (e.g., symbol properties), return safe string
      return "[Object with non-enumerable properties]";
    }
  }
}
