import type { ErrorRequestHandler } from "express";

import { RequestContext } from "../../observability/RequestContext.js";

const isProduction = process.env.NODE_ENV === "production";

const resolveStatus = (error: unknown): number => {
  if (typeof error === "object" && error !== null) {
    const maybeError = error as { statusCode?: number; status?: number };

    if (typeof maybeError.statusCode === "number") {
      return maybeError.statusCode;
    }

    if (typeof maybeError.status === "number") {
      return maybeError.status;
    }
  }

  return 500;
};

export const errorHandler: ErrorRequestHandler = (err, req, res, next) => {
  if (res.headersSent) {
    next(err);
    return;
  }

  const traceId = RequestContext.getTraceId();
  const spanId = RequestContext.get<string>("span_id");
  const timestamp = new Date().toISOString();
  const method = req.method;
  const path = req.originalUrl ?? req.url;

  const logPayload = {
    trace_id: traceId ?? "unknown",
    span_id: spanId ?? null,
    timestamp,
    method,
    path,
    error: {
      name: err?.name ?? "Error",
      message: err?.message ?? "Internal Server Error",
      stack: isProduction ? undefined : err?.stack,
    },
  };

  console.error(JSON.stringify(logPayload));

  const statusCode = resolveStatus(err);

  res.status(statusCode).json({
    error: {
      trace_id: traceId ?? "unknown",
      message: isProduction
        ? "Internal Server Error"
        : err?.message ?? "Internal Server Error",
    },
  });
};
