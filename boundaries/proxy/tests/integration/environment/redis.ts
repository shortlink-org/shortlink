import {
  RedisContainer,
  type StartedRedisContainer,
} from "@testcontainers/redis";

/**
 * Creates and starts a Redis container for integration tests.
 * Uses Testcontainers for isolated test environments.
 *
 * @returns Started Redis container with connection URI
 */
export async function createRedis(): Promise<StartedRedisContainer> {
  const container = await new RedisContainer("redis:7.4-alpine").start();

  return container;
}
