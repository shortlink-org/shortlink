import { describe, it, expect, beforeEach, vi } from "vitest";
import {
  AuthorizationInterceptor,
  AuthorizationError,
  DefaultAuthorizationChecker,
  IAuthorizationChecker,
} from "../AuthorizationInterceptor.js";
import { UseCaseExecutionContext } from "../IUseCaseInterceptor.js";
import { ILogger } from "../../../../infrastructure/logging/ILogger.js";

describe("AuthorizationInterceptor", () => {
  let interceptor: AuthorizationInterceptor;
  let mockLogger: ILogger;
  let mockAuthorizationChecker: IAuthorizationChecker;

  beforeEach(() => {
    mockLogger = {
      debug: vi.fn(),
      info: vi.fn(),
      warn: vi.fn(),
      error: vi.fn(),
    } as unknown as ILogger;

    mockAuthorizationChecker = {
      isAuthorized: vi.fn().mockResolvedValue(true),
    } as unknown as IAuthorizationChecker;

    interceptor = new AuthorizationInterceptor(
      mockLogger,
      mockAuthorizationChecker
    );
  });

  describe("before", () => {
    it("should allow request when authorized", async () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: { hash: "test-hash" },
        metadata: new Map(),
      };

      vi.mocked(mockAuthorizationChecker.isAuthorized).mockResolvedValue(true);

      // Act
      const result = await interceptor.before(context);

      // Assert
      expect(result).toEqual(context.request);
      expect(mockAuthorizationChecker.isAuthorized).toHaveBeenCalledWith(
        "TestUseCase",
        context.request,
        context
      );
    });

    it("should throw AuthorizationError when not authorized", async () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: { hash: "test-hash" },
        metadata: new Map(),
      };

      vi.mocked(mockAuthorizationChecker.isAuthorized).mockResolvedValue(false);

      // Act & Assert
      await expect(interceptor.before(context)).rejects.toThrow(
        AuthorizationError
      );
      expect(mockLogger.warn).toHaveBeenCalledWith(
        "[Authorization] Access denied for: TestUseCase",
        expect.objectContaining({
          useCase: "TestUseCase",
        })
      );
    });

    it("should handle synchronous authorization check", async () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {},
        metadata: new Map(),
      };

      const syncChecker: IAuthorizationChecker = {
        isAuthorized: vi.fn().mockReturnValue(true),
      };

      const syncInterceptor = new AuthorizationInterceptor(
        mockLogger,
        syncChecker
      );

      // Act
      const result = await syncInterceptor.before(context);

      // Assert
      expect(result).toEqual(context.request);
    });
  });

  describe("after", () => {
    it("should not throw when called", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {},
        response: {},
        metadata: new Map(),
      };

      // Act & Assert
      expect(() => interceptor.after(context)).not.toThrow();
    });
  });

  describe("onError", () => {
    it("should log warning for AuthorizationError", () => {
      // Arrange
      const error = new AuthorizationError("Access denied");
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {},
        error,
        metadata: new Map(),
      };

      // Act
      interceptor.onError(context);

      // Assert
      expect(mockLogger.warn).toHaveBeenCalledWith(
        "[Authorization] Access denied",
        expect.objectContaining({
          useCase: "TestUseCase",
        })
      );
    });

    it("should not log for non-authorization errors", () => {
      // Arrange
      const error = new Error("Other error");
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "TestUseCase",
        request: {},
        error,
        metadata: new Map(),
      };

      // Act
      interceptor.onError(context);

      // Assert
      expect(mockLogger.warn).not.toHaveBeenCalled();
    });
  });
});

describe("DefaultAuthorizationChecker", () => {
  let checker: DefaultAuthorizationChecker;

  beforeEach(() => {
    checker = new DefaultAuthorizationChecker();
  });

  describe("isAuthorized", () => {
    it("should allow public use cases", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "GetLinkByHashUseCase",
        request: {},
        metadata: new Map(),
      };

      // Act
      const result = checker.isAuthorized(
        "GetLinkByHashUseCase",
        {},
        context
      );

      // Assert
      expect(result).toBe(true);
    });

    it("should allow use cases when user is in context", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "PrivateUseCase",
        request: {},
        metadata: new Map([["user", { id: "123" }]]),
      };

      // Act
      const result = checker.isAuthorized("PrivateUseCase", {}, context);

      // Assert
      expect(result).toBe(true);
    });

    it("should deny use cases when user is not in context", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "PrivateUseCase",
        request: {},
        metadata: new Map(),
      };

      // Act
      const result = checker.isAuthorized("PrivateUseCase", {}, context);

      // Assert
      expect(result).toBe(false);
    });

    it("should deny when user is null", () => {
      // Arrange
      const context: UseCaseExecutionContext<any, any> = {
        useCaseName: "PrivateUseCase",
        request: {},
        metadata: new Map([["user", null]]),
      };

      // Act
      const result = checker.isAuthorized("PrivateUseCase", {}, context);

      // Assert
      expect(result).toBe(false);
    });
  });
});

