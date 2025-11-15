import compression from "compression";
import express, {
  type Express,
  type RequestHandler,
  json,
  urlencoded,
} from "express";
import helmet from "helmet";
import { performance } from "node:perf_hooks";
import { constants as zlibConstants } from "node:zlib";

import { RequestContext } from "../observability/RequestContext.js";
import { errorHandler } from "./middlewares/errorHandler.js";
import { notFoundHandler } from "./middlewares/notFoundHandler.js";
import { requestContextMiddleware } from "./middlewares/requestContextMiddleware.js";

type BuildExpressAppOptions = {
  jsonLimit?: string;
  urlEncodedLimit?: string;
  strictRouting?: boolean;
};

const DEFAULT_JSON_LIMIT = process.env.HTTP_JSON_BODY_LIMIT ?? "1mb";
const DEFAULT_URL_ENCODED_LIMIT =
  process.env.HTTP_URLENCODED_BODY_LIMIT ?? "1mb";

const BROTLI_ENABLED = typeof zlibConstants.BROTLI_PARAM_QUALITY === "number";

const configureApplicationDefaults = (
  app: Express,
  strictRouting?: boolean
) => {
  app.set("trust proxy", true);
  app.disable("x-powered-by");
  app.set("json spaces", 0);

  if (strictRouting) {
    app.enable("strict routing");
  } else {
    app.disable("strict routing");
  }

  app.set("etag", "strong");
};

type CompressionOptionsWithBrotli = compression.CompressionOptions & {
  brotli?: {
    enabled: boolean;
    params: Record<number, number>;
  };
};

const registerMiddlewares = (app: Express, options: BuildExpressAppOptions) => {
  const jsonLimit = options.jsonLimit ?? DEFAULT_JSON_LIMIT;
  const urlEncodedLimit = options.urlEncodedLimit ?? DEFAULT_URL_ENCODED_LIMIT;

  app.use(requestContextMiddleware);
  app.use(
    helmet({
      contentSecurityPolicy: false,
    })
  );

  const compressionOptions: CompressionOptionsWithBrotli = {
    threshold: 0,
    filter: compression.filter,
  };

  if (BROTLI_ENABLED) {
    compressionOptions.brotli = {
      enabled: true,
      params: {
        [zlibConstants.BROTLI_PARAM_QUALITY]: 4,
      },
    };
  }

  app.use(compression(compressionOptions) as unknown as RequestHandler);

  app.use(json({ limit: jsonLimit, strict: true }));
  app.use(
    urlencoded({
      limit: urlEncodedLimit,
      extended: true,
    })
  );

  app.use((req, res, next) => {
    const start =
      RequestContext.get<number>("request_start_time") ?? performance.now();

    res.on("finish", () => {
      const duration = performance.now() - start;

      const logEntry = {
        trace_id: RequestContext.getTraceId() ?? "unknown",
        method: req.method,
        path: req.originalUrl ?? req.url,
        status_code: res.statusCode,
        content_length: res.getHeader("content-length") ?? undefined,
        duration_ms: Number(duration.toFixed(4)),
      };

      console.info(JSON.stringify(logEntry));
    });

    next();
  });
};

const registerRoutes = (app: Express) => {
  app.get("/healthz", (_req, res) => {
    res.status(200).json({ status: "ok" });
  });
};

const registerFallbacks = (app: Express) => {
  app.use(notFoundHandler);
  app.use(errorHandler);
};

export const buildExpressApp = (
  options: BuildExpressAppOptions = {}
): Express => {
  const app = express();

  configureApplicationDefaults(app, options.strictRouting);
  registerMiddlewares(app, options);
  registerRoutes(app);
  registerFallbacks(app);

  return app;
};
