import type { RequestHandler } from "express";
import { randomUUID } from "node:crypto";
import { performance } from "node:perf_hooks";

import { RequestContext } from "../../../observability/RequestContext";

/**
 * Bootstraps a per-request AsyncContext scope, adding trace metadata that downstream
 * handlers can access without passing arguments through the call stack.
 */
export const requestContextMiddleware: RequestHandler = (req, _res, next) => {
  const traceId = randomUUID();
  const store = RequestContext.seed({
    trace_id: traceId,
    request_start_time: performance.now(),
    method: req.method,
    path: req.originalUrl ?? req.url,
  });

  RequestContext.run(store, () => {
    next();
  });
};

