import winston from "winston";
import { trace } from "@opentelemetry/api";
import { ILogger, LogEvent, EventLogMeta } from "./ILogger.js";

/**
 * ILogger implementation based on Winston
 * Encapsulates Winston logic in the infrastructure layer
 * Automatically adds trace_id and span_id from OpenTelemetry context
 *
 * @template TMeta - Type for metadata object. Defaults to `any` for backward compatibility.
 */
export class WinstonLogger<TMeta = any> implements ILogger<TMeta> {
  private logger: winston.Logger;

  constructor() {
    this.logger = winston.createLogger({
      level: process.env.LOG_LEVEL || "info",
      format: winston.format.combine(
        winston.format.timestamp(),
        winston.format.errors({ stack: true }),
        winston.format.json()
      ),
      transports: [
        new winston.transports.Console({
          // Output pure JSON format (no colorize/simple override)
        }),
      ],
    });
  }

  /**
   * Extracts trace_id and span_id from the current OpenTelemetry context
   */
  private getTraceContext(): { trace_id?: string; span_id?: string } {
    const activeSpan = trace.getActiveSpan();
    if (!activeSpan) {
      return {};
    }

    const spanContext = activeSpan.spanContext();
    if (!spanContext || !spanContext.traceId || !spanContext.spanId) {
      return {};
    }

    return {
      trace_id: spanContext.traceId,
      span_id: spanContext.spanId,
    };
  }

  /**
   * Merges metadata with trace context
   * Trace context is always applied last and cannot be overwritten by metadata
   */
  private enrichMeta(meta?: any): any {
    const traceContext = this.getTraceContext();
    // Trace context always on top - meta can never overwrite trace_id/span_id
    return { ...meta, ...traceContext };
  }

  /**
   * Serializes Error object to OTEL-compatible JSON format
   * Uses standard fields: error.message, error.stack, error.name
   */
  private serializeError(error: Error): {
    message: string;
    stack?: string;
    name: string;
  } {
    return {
      name: error.name,
      message: error.message,
      stack: error.stack,
    };
  }

  /**
   * Checks if value is a built-in JavaScript object that should not be recursively processed
   * These objects have symbol properties and can cause errors during serialization
   */
  private isBuiltInObject(value: any): boolean {
    return (
      value instanceof Date ||
      value instanceof URL ||
      value instanceof RegExp ||
      value instanceof Map ||
      value instanceof Set ||
      (typeof Buffer !== "undefined" && Buffer.isBuffer(value)) ||
      value instanceof Uint8Array ||
      value instanceof ArrayBuffer
    );
  }

  /**
   * Serializes built-in objects to safe string representation
   */
  private serializeBuiltInObject(value: any): string {
    if (value instanceof Date) {
      return value.toISOString();
    }
    if (value instanceof URL) {
      return value.toString();
    }
    if (value instanceof RegExp) {
      return value.toString();
    }
    if (value instanceof Map) {
      return `Map(${value.size})`;
    }
    if (value instanceof Set) {
      return `Set(${value.size})`;
    }
    if (typeof Buffer !== "undefined" && Buffer.isBuffer(value)) {
      return `<Buffer ${value.length} bytes>`;
    }
    if (value instanceof Uint8Array) {
      return `<Uint8Array ${value.length} bytes>`;
    }
    if (value instanceof ArrayBuffer) {
      return `<ArrayBuffer ${value.byteLength} bytes>`;
    }
    return String(value);
  }

  /**
   * Processes Error objects in meta, serializing them to OTEL-compatible format
   * Unified serialization layer for all errors
   * Safely handles built-in objects (Date, URL, RegExp, etc.) to prevent serialization errors
   */
  private processErrorInMeta(meta?: any): any {
    if (!meta) return meta;

    // If meta itself is an Error, serialize it
    if (meta instanceof Error) {
      return {
        error: this.serializeError(meta),
      };
    }

    // If meta is a built-in object, serialize it safely
    if (this.isBuiltInObject(meta)) {
      return this.serializeBuiltInObject(meta);
    }

    // If meta is an array, process each element
    if (Array.isArray(meta)) {
      return meta.map((item) => this.processErrorInMeta(item));
    }

    // If meta is an object, check for Error objects in its properties
    const processed: any = { ...meta };

    // Check all meta properties and serialize found Error objects
    // Use try-catch to handle objects with symbol properties safely
    try {
      for (const [key, value] of Object.entries(meta)) {
        if (value instanceof Error) {
          processed[key] = this.serializeError(value);
        } else if (this.isBuiltInObject(value)) {
          processed[key] = this.serializeBuiltInObject(value);
        } else if (
          typeof value === "object" &&
          value !== null &&
          !Array.isArray(value)
        ) {
          // Recursively process nested objects (only plain objects)
          processed[key] = this.processErrorInMeta(value);
        }
      }
    } catch (err) {
      // If Object.entries fails (e.g., due to symbol properties), fallback to string representation
      return String(meta);
    }

    return processed;
  }

  info(message: string, meta?: any): void {
    const processedMeta = this.processErrorInMeta(meta);
    this.logger.info(message, this.enrichMeta(processedMeta));
  }

  warn(message: string, meta?: any): void {
    const processedMeta = this.processErrorInMeta(meta);
    this.logger.warn(message, this.enrichMeta(processedMeta));
  }

  error(message: string, meta?: any): void {
    const processedMeta = this.processErrorInMeta(meta);
    this.logger.error(message, this.enrichMeta(processedMeta));
  }

  debug(message: string, meta?: any): void {
    const processedMeta = this.processErrorInMeta(meta);
    this.logger.debug(message, this.enrichMeta(processedMeta));
  }

  http(message: string, meta?: any): void {
    const processedMeta = this.processErrorInMeta({
      event: "http",
      ...meta,
    });
    this.logger.info(message, this.enrichMeta(processedMeta));
  }

  event<TEvent extends LogEvent>(
    event: TEvent,
    meta: EventLogMeta<TEvent, TMeta>,
    level: "info" | "warn" | "error" | "debug" = "info"
  ): void {
    const processedMeta = this.processErrorInMeta(meta);
    const enrichedMeta = this.enrichMeta(processedMeta);

    switch (level) {
      case "warn":
        this.logger.warn(event, enrichedMeta);
        break;
      case "error":
        this.logger.error(event, enrichedMeta);
        break;
      case "debug":
        this.logger.debug(event, enrichedMeta);
        break;
      case "info":
      default:
        this.logger.info(event, enrichedMeta);
        break;
    }
  }
}
