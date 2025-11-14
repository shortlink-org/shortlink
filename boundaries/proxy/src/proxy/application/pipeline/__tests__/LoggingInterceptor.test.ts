import { describe, it, expect, beforeEach, vi } from "vitest";
import { LoggingInterceptor } from "../LoggingInterceptor.js";
import { UseCaseExecutionContext } from "../IUseCaseInterceptor.js";
import { ILogger } from "../../../../infrastructure/logging/ILogger.js";

describe("LoggingInterceptor", () => {
  let interceptor: LoggingInterceptor;
  let mockLogger: ILogger;

  beforeEach(() => {
    mockLogger = {
      debug: vi.fn(),
      info: vi.fn(),
      warn: vi.fn(),
      error: vi.fn(),
    } as unknown as ILogger;

    interceptor = new LoggingInterceptor(mockLogger);
  });

  describe("before", () => {
    it("should log debug message with use case name and request", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: { hash: "test-hash" },
        metadata: new Map(),
      };

      // Act
      const result = interceptor.before(context);

      // Assert
      expect(result).toEqual(context.request);
      expect(mockLogger.debug).toHaveBeenCalledWith(
        "[UseCase] Starting: TestUseCase",
        expect.objectContaining({
          useCase: "TestUseCase",
          request: { hash: "test-hash" },
        })
      );
    });

    it("should sanitize sensitive fields in request", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {
          hash: "test-hash",
          password: "secret123",
          token: "abc123",
          apiKey: "key123",
        },
        metadata: new Map(),
      };

      // Act
      interceptor.before(context);

      // Assert
      expect(mockLogger.debug).toHaveBeenCalledWith(
        expect.any(String),
        expect.objectContaining({
          request: expect.objectContaining({
            password: "***REDACTED***",
            token: "***REDACTED***",
            apiKey: "***REDACTED***",
          }),
        })
      );
    });
  });

  describe("after", () => {
    it("should log info message with duration on success", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {},
        response: { link: {} },
        duration: 150,
        metadata: new Map(),
      };

      // Act
      interceptor.after(context);

      // Assert
      expect(mockLogger.info).toHaveBeenCalledWith(
        "[UseCase] Completed: TestUseCase",
        expect.objectContaining({
          useCase: "TestUseCase",
          duration: "150ms",
          success: true,
        })
      );
    });

    it("should handle zero duration", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {},
        response: {},
        duration: 0,
        metadata: new Map(),
      };

      // Act
      interceptor.after(context);

      // Assert
      expect(mockLogger.info).toHaveBeenCalledWith(
        expect.any(String),
        expect.objectContaining({
          duration: "0ms",
        })
      );
    });
  });

  describe("onError", () => {
    it("should log error message with error details", () => {
      // Arrange
      const error = new Error("Test error");
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {},
        error,
        duration: 50,
        metadata: new Map(),
      };

      // Act
      interceptor.onError(context);

      // Assert
      expect(mockLogger.error).toHaveBeenCalledWith(
        "[UseCase] Failed: TestUseCase",
        error,
        expect.objectContaining({
          useCase: "TestUseCase",
          duration: "50ms",
          success: false,
        })
      );
    });
  });

  describe("finally", () => {
    it("should not throw when called", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {},
        metadata: new Map(),
      };

      // Act & Assert
      expect(() => interceptor.finally(context)).not.toThrow();
    });
  });
});

