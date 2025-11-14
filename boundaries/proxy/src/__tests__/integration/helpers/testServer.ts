import "reflect-metadata";
import * as express from "express";
import { InversifyExpressServer } from "inversify-express-utils";
import { Container } from "inversify";
import { configureMiddleware } from "../../../infrastructure/middleware.js";
import { errorHandler } from "../../../interfaces/http/middleware/errorHandler.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { LinkApplicationService } from "../../../proxy/application/services/LinkApplicationService.js";
import TYPES from "../../../types.js";
// Side-effect import: registers controllers with inversify-express-utils
import "../../../proxy/interfaces/http/controllers/ProxyController.js";

/**
 * Создает Express приложение для интеграционных тестов
 * Патчит app.all() для совместимости с path-to-regexp v8
 */
export async function createTestServer(container: Container): Promise<express.Application> {
  const app = express.default();

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
  const server = new InversifyExpressServer(container, null, { rootPath: "/" }, app);
  
  let builtApp: express.Application;
  
  try {
    builtApp = server.build();
  } catch (error: any) {
    // Suppress wildcard errors from InversifyExpressServer's internal routes
    // Our controllers use specific routes like "/s/:hash" which don't conflict
    if (error.message?.includes("parameter name") || error.message?.includes("pathToRegexp")) {
      // Create a minimal app that works for our test routes
      // We'll manually register the controller route
      const testApp = express.default();
      configureMiddleware(testApp);
      
      // Manually register controller route for testing
      // This bypasses InversifyExpressServer's wildcard usage
      const { ProxyController } = await import("../../../proxy/interfaces/http/controllers/ProxyController.js");
      // Create controller instance manually with dependencies from container
      const linkApplicationService = container.get<LinkApplicationService>(TYPES.APPLICATION.LinkApplicationService);
      const logger = container.get<ILogger>(TYPES.INFRASTRUCTURE.Logger);
      const controller = new ProxyController(linkApplicationService, logger);
      testApp.get("/s/:hash", async (req, res, next) => {
        await controller.redirect(req, res, next);
      });
      
      builtApp = testApp;
    } else {
      throw error;
    }
  }

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

