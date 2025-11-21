import type { FastifyInstance } from "fastify";
import { createDIContainer, createRequestScope, type ContainerDependencies } from "../../../container/index.js";
import { buildServer } from "../../../proxy/infrastructure/http/fastify/server.js";
import type { AwilixContainer } from "awilix";
import { asValue } from "awilix";

/**
 * Creates a Fastify application for integration tests.
 * Uses Awilix container with ability to override dependencies.
 */
export async function createTestServer(
  overrides?: Partial<ContainerDependencies>
): Promise<FastifyInstance> {
  // Create DI container
  const container = createDIContainer();

  // Override dependencies if provided
  if (overrides) {
    Object.entries(overrides).forEach(([key, value]) => {
      container.register({
        [key]: asValue(value),
      } as any);
    });
  }

  // Build Fastify server
  const server = await buildServer(container);

  return server;
}

/**
 * Helper to create a mock container for testing
 */
export function createMockContainer(
  overrides: Partial<ContainerDependencies>
): AwilixContainer<ContainerDependencies> {
  const container = createDIContainer();

  Object.entries(overrides).forEach(([key, value]) => {
    container.register({
      [key]: asValue(value),
    } as any);
  });

  return container;
}
