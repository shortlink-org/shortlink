import type { FastifyInstance } from "fastify";
import { buildServer } from "../../../src/proxy/infrastructure/http/fastify/server.js";
import { BaseTestEnvironment } from "../environment/BaseTestEnvironment.js";

/**
 * Builds a Fastify test server with the test environment's DI container.
 * Server is ready for use with server.inject() for integration tests.
 *
 * @param env - Base test environment with initialized containers and DI
 * @returns Fastify instance ready for testing
 *
 * @example
 * ```ts
 * const env = new BaseTestEnvironment();
 * await env.start();
 * const server = await buildTestServer(env);
 *
 * const response = await server.inject({
 *   method: "GET",
 *   url: "/s/testhash",
 * });
 * ```
 */
export async function buildTestServer(
  env: BaseTestEnvironment
): Promise<FastifyInstance> {
  const container = env.getContainer();
  const server = await buildServer(container);
  await server.ready();
  return server;
}
