/**
 * Event types for structured logging
 * Used to type-check event names in log metadata
 */
export type LogEvent =
  | `usecase.${string}`
  | `http.${string}`
  | `connect.${string}`
  | string; // Allow custom events

/**
 * Helper type for structured event logs
 * Ensures that meta contains an 'event' field for structured logging
 *
 * @template TEvent - Event name (e.g., "usecase.start", "http.error")
 * @template TMeta - Additional metadata fields
 */
export type EventLogMeta<
  TEvent extends LogEvent = LogEvent,
  TMeta = Record<string, unknown>
> = TMeta & {
  event: TEvent;
  error?: Error;
};

/**
 * Logger interface for abstraction from concrete implementation
 * Allows easy replacement of Winston with another logging library
 *
 * @template TMeta - Type for metadata object. Defaults to `any` for backward compatibility.
 *                   Error objects can be passed directly as meta or in meta.error and will be automatically serialized.
 */
export interface ILogger<TMeta = any> {
  /**
   * Logs an informational message
   * @param message - Log message
   * @param meta - Metadata object or Error instance (Error will be automatically serialized)
   */
  info(message: string, meta?: (TMeta & { error?: Error }) | Error): void;

  /**
   * Logs a warning
   * @param message - Log message
   * @param meta - Metadata object or Error instance (Error will be automatically serialized)
   */
  warn(message: string, meta?: (TMeta & { error?: Error }) | Error): void;

  /**
   * Logs an error
   * @param message - Log message
   * @param meta - Metadata object or Error instance (Error will be automatically serialized to OTEL-compatible format)
   */
  error(message: string, meta?: (TMeta & { error?: Error }) | Error): void;

  /**
   * Logs a debug message
   * @param message - Log message
   * @param meta - Metadata object or Error instance (Error will be automatically serialized)
   */
  debug(message: string, meta?: (TMeta & { error?: Error }) | Error): void;

  /**
   * Logs an HTTP request (special method for HTTP logging)
   * @param message - Log message
   * @param meta - Metadata object or Error instance (Error will be automatically serialized)
   */
  http(message: string, meta?: (TMeta & { error?: Error }) | Error): void;

  /**
   * Logs a structured event
   * Type-safe method for logging events with required 'event' field
   *
   * @param event - Event name (e.g., "usecase.start", "http.error")
   * @param meta - Event metadata (must include 'event' field matching the event parameter)
   * @param level - Log level (defaults to 'info')
   *
   * @example
   * logger.event("usecase.start", {
   *   event: "usecase.start",
   *   useCase: "GetLink",
   *   durationMs: 100
   * });
   */
  event<TEvent extends LogEvent>(
    event: TEvent,
    meta: EventLogMeta<TEvent, TMeta>,
    level?: "info" | "warn" | "error" | "debug"
  ): void;
}
