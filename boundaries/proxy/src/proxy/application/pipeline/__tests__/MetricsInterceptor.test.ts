import { describe, it, expect, beforeEach, vi } from "vitest";
import { MetricsInterceptor } from "../MetricsInterceptor.js";
import { UseCaseExecutionContext } from "../IUseCaseInterceptor.js";
import * as otel from "@opentelemetry/api";

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

// Mock OpenTelemetry
vi.mock("@opentelemetry/api", () => {
  const mockCounter = {
    add: vi.fn(),
  };

  const mockHistogram = {
    record: vi.fn(),
  };

  const mockMeter = {
    createCounter: vi.fn().mockReturnValue(mockCounter),
    createHistogram: vi.fn().mockReturnValue(mockHistogram),
  };

  return {
    metrics: {
      getMeter: vi.fn().mockReturnValue(mockMeter),
    },
  };
});

describe("MetricsInterceptor", () => {
  let interceptor: MetricsInterceptor;
  let mockCounter: any;
  let mockHistogram: any;

  beforeEach(() => {
    vi.clearAllMocks();
    interceptor = new MetricsInterceptor();

    // Get mocked meter and instruments
    const mockMeter = otel.metrics.getMeter("proxy-service", "1.0.0");
    mockCounter = mockMeter.createCounter("usecase_requests_total");
    mockHistogram = mockMeter.createHistogram("usecase_duration_ms");
  });

  describe("before", () => {
    it("should return request unchanged", () => {
      // Arrange
      const context = createContext({
        request: { hash: "test-hash" },
      });

      // Act
      const result = interceptor.before(context);

      // Assert
      expect(result).toEqual(context.request);
    });
  });

  describe("after", () => {
    it("should record success metrics", () => {
      // Arrange
      const context = createContext({
        duration: 150,
      });

      // Act
      interceptor.after(context);

      // Assert
      expect(mockCounter.add).toHaveBeenCalledWith(1, {
        usecase: "TestUseCase",
        status: "success",
      });
      expect(mockHistogram.record).toHaveBeenCalledWith(150, {
        usecase: "TestUseCase",
      });
    });

    it("should handle zero duration", () => {
      // Arrange
      const context = createContext({
        duration: 0,
      });

      // Act
      interceptor.after(context);

      // Assert
      expect(mockHistogram.record).toHaveBeenCalledWith(0, {
        usecase: "TestUseCase",
      });
    });
  });

  describe("onError", () => {
    it("should record error metrics", () => {
      // Arrange
      const error = new Error("Test error");
      const context = createContext({
        error,
        duration: 50,
      });

      // Act
      interceptor.onError(context);

      // Assert
      expect(mockCounter.add).toHaveBeenCalledWith(1, {
        usecase: "TestUseCase",
        status: "error",
      });
      expect(mockCounter.add).toHaveBeenCalledWith(1, {
        usecase: "TestUseCase",
        error_type: "Error",
      });
      expect(mockHistogram.record).toHaveBeenCalledWith(50, {
        usecase: "TestUseCase",
      });
    });

    it("should handle unknown error type", () => {
      // Arrange
      const context = createContext({
        error: undefined,
        duration: 50,
      });

      // Act
      interceptor.onError(context);

      // Assert
      expect(mockCounter.add).toHaveBeenCalledWith(1, {
        usecase: "TestUseCase",
        error_type: "Unknown",
      });
    });
  });
});

