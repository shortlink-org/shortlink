import { AsyncContext } from "node:async_context";

export type RequestContextSeed = Record<string, unknown>;

type RequestContextStore = Map<string, unknown>;

/**
 * RequestContext wraps AsyncContext to expose a clean abstraction for request-scoped state
 * (trace ids, metadata, etc.) without leaking framework-specific details.
 */
export class RequestContext {
  private static readonly TRACE_ID_KEY = "trace_id";
  private static readonly storage = new AsyncContext<RequestContextStore>();

  private static getStore(): RequestContextStore | undefined {
    return (
      this.storage as unknown as {
        getStore: () => RequestContextStore | undefined;
      }
    ).getStore();
  }

  /**
   * Creates a new store seeded with the provided values.
   */
  static seed(seed: RequestContextSeed = {}): RequestContextStore {
    const store = new Map<string, unknown>();

    for (const [key, value] of Object.entries(seed)) {
      store.set(key, value);
    }

    return store;
  }

  /**
   * Runs the provided callback within a new AsyncContext scope.
   */
  static run<T>(store: RequestContextStore, callback: () => T): T {
    return this.storage.run(store, callback);
  }

  /**
   * Retrieves a value from the request-scoped store.
   */
  static get<T = unknown>(key: string): T | undefined {
    return this.getStore()?.get(key) as T | undefined;
  }

  /**
   * Sets a value on the current request-scoped store.
   */
  static set(key: string, value: unknown): void {
    const store = this.getStore();

    if (!store) {
      throw new Error(
        "RequestContext store is not available. Did you register the middleware?"
      );
    }

    store.set(key, value);
  }

  static getTraceId(): string | undefined {
    return this.get<string>(this.TRACE_ID_KEY);
  }
}
