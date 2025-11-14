import * as express from "express";
import { z } from "zod";
import { ValidationError } from "../../../proxy/application/exceptions/index.js";

/**
 * Middleware для валидации HTTP запросов с использованием Zod схем
 * @param schema - Zod схема для валидации
 * @param source - Источник данных ('params', 'query', 'body')
 */
export function validateRequest<T extends z.ZodTypeAny>(
  schema: T,
  source: "params" | "query" | "body" = "body"
): express.RequestHandler {
  return (req: express.Request, res: express.Response, next: express.NextFunction) => {
    try {
      const data = source === "params" ? req.params : source === "query" ? req.query : req.body;
      const validated = schema.parse(data);
      
      // Сохраняем валидированные данные обратно в request
      if (source === "params") {
        Object.assign(req.params, validated);
      } else if (source === "query") {
        Object.assign(req.query, validated);
      } else {
        req.body = validated;
      }
      
      next();
    } catch (error) {
      if (error instanceof z.ZodError) {
        const firstError = error.issues[0];
        const field = firstError?.path.join(".") || source;
        const validationError = new ValidationError(
          firstError?.message || "Validation failed",
          field,
          { issues: error.issues }
        );
        return next(validationError);
      }
      return next(new ValidationError("Invalid request data"));
    }
  };
}

