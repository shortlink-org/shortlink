declare module "node:async_context" {
  export class AsyncContext<T = unknown> {
    run<TResult>(store: T, callback: () => TResult): TResult;
    getStore(): T | undefined;
  }
}
