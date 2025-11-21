import type { AwilixContainer, Resolver } from "awilix";
import { asValue, asClass, asFunction } from "awilix";
import type { ContainerDependencies } from "../../../src/container/index.js";

/**
 * Overrides dependencies in the DI container for testing.
 * Automatically detects and wraps values with appropriate Awilix resolver.
 *
 * Supports:
 * - Plain values/mocks → asValue()
 * - Class constructors → asClass()
 * - Factory functions → asFunction()
 * - Already wrapped resolvers → used as-is
 *
 * @param container - Awilix container to override
 * @param overrides - Partial container dependencies to override
 *
 * @example
 * ```ts
 * overrideDI(container, {
 *   linkRepository: mockRepo,
 *   redisClient: fakeRedis,
 * });
 * ```
 */
export function overrideDI(
  container: AwilixContainer<ContainerDependencies>,
  overrides: Partial<ContainerDependencies>
): void {
  const registrations: Record<string, Resolver<unknown>> = {};

  for (const [key, value] of Object.entries(overrides)) {
    // Skip undefined — explicit undefined should not override the container
    if (value === undefined) continue;

    // --- Detect already wrapped Awilix resolver ---
    // Check if value is already a Resolver (has resolve method and is an object)
    const isResolver =
      value &&
      typeof value === "object" &&
      "resolve" in (value as any) &&
      typeof (value as any).resolve === "function" &&
      // Additional check: Awilix resolvers have specific properties
      ("lifetime" in (value as any) || "inject" in (value as any));

    if (isResolver) {
      // Already wrapped resolver - use as-is
      registrations[key] = value as Resolver<unknown>;
      continue;
    }

    // --- Detect class constructor ---
    // Class has prototype with constructor pointing to itself
    // Exclude arrow functions (no prototype) and regular functions (prototype.constructor is Function)
    const isClass =
      typeof value === "function" &&
      value.prototype &&
      value.prototype.constructor === value &&
      value !== Function.prototype &&
      // Simple check: class constructors have their own prototype
      Object.getPrototypeOf(value.prototype) === Object.prototype;

    if (isClass) {
      registrations[key] = asClass(
        value as new (...args: any[]) => any
      ).singleton();
      continue;
    }

    // --- Function factories (arrow functions, vi.fn mocks, etc.) ---
    if (typeof value === "function") {
      registrations[key] = asFunction(
        value as (...args: any[]) => any
      ).singleton();
      continue;
    }

    // --- Default: plain values, mocks, objects, instances ---
    registrations[key] = asValue(value);
  }

  container.register(registrations as any);
}
