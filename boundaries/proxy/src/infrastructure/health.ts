import http from "http";
import { createTerminus } from "@godaddy/terminus";
import container from "../inversify.config.js";
import TYPES from "../types.js";
import { IMessageBus } from "../proxy/domain/interfaces/IMessageBus.js";

const DEFAULT_SIGNAL = "SIGINT";
const READY_ENDPOINT = "/ready";

/**
 * Health check handler for Kubernetes readiness/liveness probes.
 *
 * @swagger
 * /ready:
 *   get:
 *     summary: Health check endpoint для Kubernetes
 *     description: |
 *       Проверка готовности сервиса для Kubernetes readiness/liveness probes.
 *       Возвращает 200 OK если сервис готов к обработке запросов.
 *     tags:
 *       - Health
 *     responses:
 *       '200':
 *         description: Сервис готов к работе
 *         content:
 *           text/plain:
 *             schema:
 *               type: string
 *               example: OK
 *       '503':
 *         description: Сервис не готов (например, при graceful shutdown)
 *         content:
 *           text/plain:
 *             schema:
 *               type: string
 *               example: Service Unavailable
 *
 * @returns Promise resolving to true if healthy
 */
async function onHealthCheck(): Promise<boolean> {
  // TODO: Add actual health checks (e.g., database connection, external services)
  return true;
}

/**
 * Configure graceful shutdown and health checks for HTTP server.
 *
 * Sets up:
 * - Graceful shutdown on SIGINT
 * - Health check endpoint at /ready
 * - Message Bus disconnection on shutdown
 *
 * @param server - HTTP server instance
 */
export function configureHealthChecks(server: http.Server): void {
  createTerminus(server, {
    signal: DEFAULT_SIGNAL,
    healthChecks: {
      [READY_ENDPOINT]: onHealthCheck,
    },
    onShutdown: async () => {
      // Gracefully disconnect Message Bus on shutdown
      try {
        const messageBus = container.get<IMessageBus>(TYPES.INFRASTRUCTURE.MessageBus);
        if (messageBus.isConnected()) {
          await messageBus.disconnect();
        }
      } catch (error) {
        console.error("[Shutdown] Failed to disconnect Message Bus:", error);
      }
    },
  });
}

