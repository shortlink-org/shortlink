import type { AwilixContainer } from "awilix";
import type { StartedRedisContainer } from "@testcontainers/redis";
import { createRedis } from "./redis.js";
import {
  createDIContainer,
  type ContainerDependencies,
} from "../../../src/container/index.js";
import { asValue } from "awilix";
import { CacheConfig } from "../../../src/infrastructure/config/CacheConfig.js";

/**
 * Base test environment for integration tests.
 * Manages Testcontainers (Redis) and DI container setup.
 * Follows Clean Architecture principles - domain tests do NOT use containers.
 */
export class BaseTestEnvironment {
  redis: StartedRedisContainer | null = null;
  container: AwilixContainer<ContainerDependencies> | null = null;

  /**
   * Starts all containers and creates DI container with test configuration.
   * Registers container connection URIs into DI as config values.
   */
  async start(): Promise<void> {
    // Start Redis container
    this.redis = await createRedis();

    // Create DI container
    this.container = createDIContainer();

    // Register container connection URIs as config overrides
    const redisUri = this.redis.getConnectionUrl();

    // Override cache config with test Redis URL
    // Create a test CacheConfig instance that matches the interface
    const cacheConfig = Object.create(CacheConfig.prototype) as CacheConfig;
    Object.assign(cacheConfig, {
      enabled: true,
      redisUrl: redisUri,
      ttlPositive: 60,
      ttlNegative: 30,
      keyPrefix: "test:shortlink:proxy",
    });

    this.container.register({
      cacheConfig: asValue(cacheConfig),
    });

    // Store connection URI for test access
    // This can be used to create direct connections if needed
    (this.container as any).__testRedisUri = redisUri;
  }

  /**
   * Stops all containers and disposes DI container.
   */
  async stop(): Promise<void> {
    // Stop Redis container
    if (this.redis) {
      await this.redis.stop();
      this.redis = null;
    }

    // Dispose DI container
    if (this.container) {
      await this.container.dispose();
      this.container = null;
    }
  }

  /**
   * Gets Redis connection URI.
   * @returns Connection URI string
   */
  getRedisUri(): string {
    if (!this.redis) {
      throw new Error("Redis container is not started. Call start() first.");
    }
    return this.redis.getConnectionUrl();
  }

  /**
   * Gets the DI container.
   * @returns Awilix container instance
   */
  getContainer(): AwilixContainer<ContainerDependencies> {
    if (!this.container) {
      throw new Error("DI container is not initialized. Call start() first.");
    }
    return this.container;
  }
}
