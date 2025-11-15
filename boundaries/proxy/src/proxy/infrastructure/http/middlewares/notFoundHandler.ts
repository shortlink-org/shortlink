import type { RequestHandler } from "express";

import { RequestContext } from "../../observability/RequestContext.js";

export const notFoundHandler: RequestHandler = (req, res) => {
  const traceId = RequestContext.getTraceId();

  res.status(404).json({
    error: {
      trace_id: traceId ?? "unknown",
      message: `Route ${req.method} ${
        req.originalUrl ?? req.url
      } was not found`,
    },
  });
};
