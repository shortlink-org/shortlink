import { describe, it, expect, vi, beforeEach } from "vitest";
import { Request, Response, NextFunction } from "express";
import { z } from "zod";
import { validateRequest } from "../validationMiddleware.js";
import { ValidationError } from "../../../../proxy/application/exceptions/index.js";

describe("validateRequest", () => {
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;
  let mockNext: NextFunction;

  beforeEach(() => {
    mockRequest = {
      params: {},
      query: {},
      body: {},
    };
    mockResponse = {
      status: vi.fn().mockReturnThis(),
      json: vi.fn().mockReturnThis(),
    };
    mockNext = vi.fn();
  });

  describe("body validation", () => {
    it("should validate and pass valid body data", () => {
      // Arrange
      const schema = z.object({
        hash: z.string(),
        url: z.string().url(),
      });
      mockRequest.body = { hash: "test-hash", url: "https://example.com" };

      const middleware = validateRequest(schema, "body");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith();
      expect(mockRequest.body).toEqual({ hash: "test-hash", url: "https://example.com" });
    });

    it("should reject invalid body data", () => {
      // Arrange
      const schema = z.object({
        hash: z.string(),
        url: z.string().url(),
      });
      mockRequest.body = { hash: "test-hash", url: "invalid-url" };

      const middleware = validateRequest(schema, "body");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith(expect.any(ValidationError));
    });

    it("should update request body with validated data", () => {
      // Arrange
      const schema = z.object({
        hash: z.string(),
      });
      mockRequest.body = { hash: "test-hash", extra: "field" };

      const middleware = validateRequest(schema, "body");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockRequest.body).toEqual({ hash: "test-hash" });
    });
  });

  describe("params validation", () => {
    it("should validate and pass valid params", () => {
      // Arrange
      const schema = z.object({
        hash: z.string(),
      });
      mockRequest.params = { hash: "test-hash" };

      const middleware = validateRequest(schema, "params");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith();
      expect(mockRequest.params).toEqual({ hash: "test-hash" });
    });

    it("should reject invalid params", () => {
      // Arrange
      const schema = z.object({
        hash: z.string().min(5),
      });
      mockRequest.params = { hash: "abc" };

      const middleware = validateRequest(schema, "params");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith(expect.any(ValidationError));
    });
  });

  describe("query validation", () => {
    it("should validate and pass valid query parameters", () => {
      // Arrange
      const schema = z.object({
        page: z.string().transform(Number),
      });
      mockRequest.query = { page: "1" };

      const middleware = validateRequest(schema, "query");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith();
    });

    it("should reject invalid query parameters", () => {
      // Arrange
      const schema = z.object({
        page: z.string().regex(/^\d+$/),
      });
      mockRequest.query = { page: "invalid" };

      const middleware = validateRequest(schema, "query");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith(expect.any(ValidationError));
    });
  });

  describe("error handling", () => {
    it("should create ValidationError with field path", () => {
      // Arrange
      const schema = z.object({
        user: z.object({
          email: z.string().email(),
        }),
      });
      mockRequest.body = { user: { email: "invalid" } };

      const middleware = validateRequest(schema, "body");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith(
        expect.objectContaining({
          field: "user.email",
        })
      );
    });

    it("should handle non-ZodError exceptions", () => {
      // Arrange
      const schema = z.object({
        hash: z.string(),
      });
      mockRequest.body = null;

      const middleware = validateRequest(schema, "body");

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith(expect.any(ValidationError));
    });
  });

  describe("default source", () => {
    it("should default to body validation", () => {
      // Arrange
      const schema = z.object({
        hash: z.string(),
      });
      mockRequest.body = { hash: "test-hash" };

      const middleware = validateRequest(schema);

      // Act
      middleware(mockRequest as Request, mockResponse as Response, mockNext);

      // Assert
      expect(mockNext).toHaveBeenCalledWith();
    });
  });
});

