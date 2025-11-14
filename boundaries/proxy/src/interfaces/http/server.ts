import "reflect-metadata";
import * as express from "express";
import { InversifyExpressServer } from "inversify-express-utils";
import { Container } from "inversify";
import { configureMiddleware } from "../../infrastructure/middleware.js";
import { errorHandler } from "./middleware/errorHandler.js";
import { ILogger } from "../../infrastructure/logging/ILogger.js";
import TYPES from "../../types.js";
// Side-effect import: registers controllers with inversify-express-utils
import "../../proxy/interfaces/http/controllers/ProxyController.js";
import "./controllers/MetricsController.js";

const DEFAULT_ROOT_PATH = "/";

/**
 * Server configuration options.
 */
export interface ServerConfig {
  rootPath?: string;
}

/**
 * Build and configure Express application with Inversify container.
 *
 * Sets up:
 * - Express application instance
 * - Middleware stack
 * - InversifyExpressServer integration
 * - Controller registration (via side-effect imports)
 *
 * Note: Patches app.all() to convert "*" to ":splat*" for path-to-regexp v8 compatibility
 * before InversifyExpressServer.build() since InversifyExpressServer uses app.all("*") internally
 *
 * @param container - Inversify dependency injection container
 * @param config - Server configuration options
 * @param existingApp - Optional existing Express app (for testing/mocking)
 * @returns Configured Express application
 */
export function createServer(
  container: Container,
  config: ServerConfig = {},
  existingApp?: express.Application
): express.Application {
  const app = existingApp ?? express.default();
  const rootPath = config.rootPath ?? DEFAULT_ROOT_PATH;

  // Apply middleware
  configureMiddleware(app);

  // Patch app.all() to convert "*" to ":splat*" for path-to-regexp v8 compatibility
  // InversifyExpressServer uses app.all("*") internally, which needs conversion
  const originalAll = app.all.bind(app);
  (app as any).all = function (path: any, ...handlers: any[]) {
    // Convert "*" to ":splat*" for path-to-regexp v8 compatibility
    if (typeof path === "string" && path === "*") {
      path = ":splat*";
    }
    return originalAll(path, ...handlers);
  };

  // Build InversifyExpressServer
  const server = new InversifyExpressServer(container, null, { rootPath }, app);

  const builtApp = server.build();

  // Add error handling middleware (must be last)
  try {
    const logger = container.get<ILogger>(TYPES.INFRASTRUCTURE.Logger);
    builtApp.use(errorHandler(logger));
  } catch (error) {
    // Fallback if logger is not available
    builtApp.use(errorHandler());
  }

  return builtApp;
}
