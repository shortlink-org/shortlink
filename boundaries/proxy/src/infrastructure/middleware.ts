import * as express from "express";
import * as bodyParser from "body-parser";
import helmet from "helmet";
import morgan from "morgan";
import log from "../logger.js";

const MORGAN_FORMAT = ":method :url :status :res[content-length] - :response-time ms";

/**
 * Configure and apply Express middleware stack.
 *
 * Applies:
 * - Body parser (URL encoded and JSON)
 * - Helmet security headers
 * - Morgan HTTP request logging
 *
 * @param app - Express application instance
 */
export function configureMiddleware(app: express.Application): void {
  // Body parsing middleware
  app.use(
    bodyParser.urlencoded({
      extended: true,
    })
  );
  app.use(bodyParser.json());

  // Security headers
  app.use(helmet());

  // HTTP request logging
  const morganMiddleware = morgan(MORGAN_FORMAT, {
    stream: {
      write: (message: string) => log.http(message.trim()),
    },
  });
  app.use(morganMiddleware);
}

