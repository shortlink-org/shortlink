import type { FastifyInstance } from "fastify";
import type { AwilixContainer } from "awilix";
import type { ContainerDependencies } from "../di/container.js";
import type { IMessageBus } from "../domain/interfaces/IMessageBus.js";

const READY_ENDPOINT = "/ready";
const LIVE_ENDPOINT = "/live";

/**
 * Server readiness state manager for graceful shutdown
 */
class ServerState {
  private isReady: boolean = true;
  private isShuttingDown: boolean = false;

  /**
   * Mark server as ready to accept traffic
   */
  setReady(ready: boolean): void {
    this.isReady = ready;
  }

  /**
   * Check if server is ready to accept traffic
   */
  getReady(): boolean {
    return this.isReady && !this.isShuttingDown;
  }

  /**
   * Mark server as shutting down
   */
  setShuttingDown(shuttingDown: boolean): void {
    this.isShuttingDown = shuttingDown;
    if (shuttingDown) {
      this.isReady = false;
    }
  }

  /**
   * Check if server is shutting down
   */
  getShuttingDown(): boolean {
    return this.isShuttingDown;
  }
}

// Global server state instance
const serverState = new ServerState();

/**
 * Default health check timeout in milliseconds
 * If health check takes longer, it's considered failed
 */
const DEFAULT_HEALTH_CHECK_TIMEOUT_MS = 2000; // 2 seconds

/**
 * Maximum allowed timeout value (24 hours) to prevent overflow
 */
const MAX_TIMEOUT_MS = 24 * 60 * 60 * 1000;

/**
 * Safely parses a timeout value from environment variable
 * Ensures the value is positive, valid, and within reasonable bounds
 *
 * @param envValue - Environment variable value
 * @param defaultValue - Default value to use if parsing fails
 * @returns Valid timeout in milliseconds (always positive)
 */
function parseTimeoutMs(
  envValue: string | undefined,
  defaultValue: number
): number {
  if (!envValue) {
    return defaultValue;
  }

  const parsed = parseInt(envValue, 10);

  // Check for NaN, negative values, or values that are too large
  if (isNaN(parsed) || parsed < 0 || parsed > MAX_TIMEOUT_MS) {
    console.warn(
      `[HealthCheck] Invalid timeout value "${envValue}", using default ${defaultValue}ms`
    );
    return defaultValue;
  }

  return parsed;
}

/**
 * Creates a timeout promise that rejects after specified time
 */
function createTimeout(timeoutMs: number): Promise<never> {
  // Ensure timeout is always positive (safety check)
  const safeTimeout = Math.max(0, timeoutMs);
  return new Promise((_, reject) => {
    setTimeout(() => {
      reject(new Error(`Health check timeout after ${safeTimeout}ms`));
    }, safeTimeout);
  });
}

/**
 * Wraps a health check function with a timeout
 * If the check doesn't complete within the timeout, returns false (not ready)
 *
 * @param healthCheck - The health check function to wrap
 * @param timeoutMs - Timeout in milliseconds
 * @returns Promise resolving to health check result or false on timeout
 */
async function withTimeout(
  healthCheck: () => Promise<boolean>,
  timeoutMs: number
): Promise<boolean> {
  try {
    return await Promise.race([healthCheck(), createTimeout(timeoutMs)]);
  } catch (error) {
    // On timeout or error, log and return false (not ready)
    console.error("[HealthCheck] Health check failed or timed out:", error);
    return false;
  }
}

/**
 * Health check handler for Kubernetes readiness probes.
 * Checks if server is ready to accept traffic.
 * Includes timeout protection to prevent hanging.
 *
 * @returns Promise resolving to true if ready
 */
async function onReadinessCheck(): Promise<boolean> {
  const timeoutMs = parseTimeoutMs(
    process.env.HEALTH_CHECK_TIMEOUT_MS,
    DEFAULT_HEALTH_CHECK_TIMEOUT_MS
  );

  // Wrap the actual health check with timeout
  return await withTimeout(async () => {
    // Check server state first (fast, no external calls)
    const isReady = serverState.getReady();
    if (!isReady) {
      console.warn("[HealthCheck] Server not ready:", {
        isReady: serverState.getReady(),
        isShuttingDown: serverState.getShuttingDown(),
      });
      return false;
    }

    // TODO: Add actual health checks (e.g., database connection, external services)
    // These should be wrapped in individual timeouts if they call external services
    // Example:
    // const dbHealthy = await withTimeout(() => checkDatabase(), 1000);
    // if (!dbHealthy) return false;

    return true;
  }, timeoutMs);
}

/**
 * Configure graceful shutdown and health checks for Fastify server.
 *
 * Sets up:
 * - Liveness endpoint at /live (always OK if process is alive)
 * - Readiness endpoint at /ready (503 during shutdown)
 * - Message Bus disconnection on shutdown
 *
 * @param fastify - Fastify server instance
 * @param container - DI container for accessing dependencies
 * @returns Function to initiate graceful shutdown
 */
export function configureHealthChecks(
  fastify: FastifyInstance,
  container?: AwilixContainer<ContainerDependencies>
): () => Promise<void> {
  // Register /live endpoint (liveness probe)
  // Always returns OK if process is alive
  fastify.get(
    LIVE_ENDPOINT,
    {
      schema: {
        description:
          "Liveness probe endpoint for Kubernetes. Always returns OK if process is alive.",
        tags: ["Health"],
        response: {
          200: {
            description: "Process is alive",
            type: "string",
            example: "OK",
          },
        },
      },
    },
    async (request, reply) => {
      return reply.code(200).send("OK");
    }
  );

  // Register /ready endpoint (readiness probe)
  // Returns 503 during graceful shutdown
  fastify.get(
    READY_ENDPOINT,
    {
      schema: {
        description:
          "Readiness probe endpoint for Kubernetes. Returns 503 during graceful shutdown.",
        tags: ["Health"],
        response: {
          200: {
            description: "Service is ready to accept traffic",
            type: "string",
            example: "OK",
          },
          503: {
            description:
              "Service is not ready (e.g., during graceful shutdown)",
            type: "string",
            example: "Service Unavailable",
          },
        },
      },
    },
    async (request, reply) => {
      try {
        const ready = await onReadinessCheck();
        if (ready) {
          return reply.code(200).send("OK");
        }
        return reply.code(503).send("Service Unavailable");
      } catch (error) {
        // If health check throws (shouldn't happen due to timeout wrapper, but safety first)
        console.error(
          "[HealthCheck] Unexpected error in readiness check:",
          error
        );
        return reply.code(503).send("Service Unavailable");
      }
    }
  );

  // Graceful shutdown handler
  fastify.addHook("onClose", async () => {
    if (container) {
      try {
        const messageBus = container.resolve<IMessageBus>("messageBus");
        if (messageBus.isConnected()) {
          await messageBus.disconnect();
        }
      } catch (error) {
        console.error("[Shutdown] Failed to disconnect Message Bus:", error);
      }
    }
  });

  /**
   * Initiate graceful shutdown
   * 1. Mark server as not ready (503 on /ready)
   * 2. Wait for traffic to drain
   * 3. Close resources
   */
  return async (): Promise<void> => {
    console.log("[Shutdown] Initiating graceful shutdown...");

    // Step 1: Mark as not ready - K8S will stop sending traffic
    serverState.setShuttingDown(true);
    console.log("[Shutdown] Server marked as not ready (503 on /ready)");

    // Step 2: Wait for traffic to drain
    // Kubernetes typically takes a few seconds to update endpoints
    const drainTime = parseTimeoutMs(process.env.SHUTDOWN_DRAIN_TIME_MS, 5000);
    console.log(`[Shutdown] Waiting ${drainTime}ms for traffic to drain...`);
    await new Promise((resolve) => setTimeout(resolve, drainTime));

    // Step 3: Close Fastify server (triggers onClose hook)
    console.log("[Shutdown] Closing server...");
    await fastify.close();
    console.log("[Shutdown] Server closed");
  };
}

/**
 * Get server state for external access
 */
export function getServerState(): ServerState {
  return serverState;
}
