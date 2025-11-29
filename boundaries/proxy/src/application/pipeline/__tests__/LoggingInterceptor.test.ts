import { describe, it, expect, beforeEach, vi } from "vitest";
import { LoggingInterceptor } from "../LoggingInterceptor.js";
import { UseCaseExecutionContext } from "../IUseCaseInterceptor.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";

const createContext = <TRequest = any, TResponse = any>(
  overrides: Partial<UseCaseExecutionContext<TRequest, TResponse>> = {}
): UseCaseExecutionContext<TRequest, TResponse> => {
  const base: UseCaseExecutionContext<TRequest, TResponse> = {
    useCaseName: "TestUseCase",
    request: {} as TRequest,
    metadata: new Map<string, unknown>(),
    startTime: Date.now(),
  };

  return {
    ...base,
    ...overrides,
    metadata: overrides.metadata ?? base.metadata,
  };
};

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
      const context = createContext({
        request: { hash: "test-hash" },
      });

      // Act
      const result = interceptor.before(context);

      // Assert
      expect(result).toEqual(context.request);
      expect(mockLogger.debug).toHaveBeenCalledWith(
        "UseCase started: TestUseCase",
        expect.objectContaining({
          event: "usecase.start",
          useCase: "TestUseCase",
          request: { hash: "test-hash" },
        })
      );
    });

    it("should sanitize sensitive fields in request", () => {
      // Arrange
      const context = createContext({
        request: {
          hash: "test-hash",
          password: "secret123",
          token: "abc123",
          apiKey: "key123",
        },
      });

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
      const context = createContext({
        duration: 150,
      });

      // Act
      interceptor.after(context);

      // Assert
      expect(mockLogger.info).toHaveBeenCalledWith(
        "UseCase completed: TestUseCase",
        expect.objectContaining({
          event: "usecase.success",
          useCase: "TestUseCase",
          durationMs: 150,
          success: true,
        })
      );
    });

    it("should handle zero duration", () => {
      // Arrange
      const context = createContext({
        duration: 0,
      });

      // Act
      interceptor.after(context);

      // Assert
      expect(mockLogger.info).toHaveBeenCalledWith(
        "UseCase completed: TestUseCase",
        expect.objectContaining({
          event: "usecase.success",
          durationMs: 0,
          success: true,
        })
      );
    });
  });

  describe("onError", () => {
    it("should log error message with error details", () => {
      // Arrange
      const error = new Error("Test error");
      const context = createContext({
        error,
        duration: 50,
      });

      // Act
      interceptor.onError(context);

      // Assert
      expect(mockLogger.error).toHaveBeenCalledWith(
        "UseCase failed: TestUseCase",
        expect.objectContaining({
          event: "usecase.error",
          useCase: "TestUseCase",
          durationMs: 50,
          success: false,
          error: error,
        })
      );
    });
  });

  describe("finally", () => {
    it("should not throw when called", () => {
      // Arrange
      const context = createContext();

      // Act & Assert
      expect(() => interceptor.finally(context)).not.toThrow();
    });
  });
});
