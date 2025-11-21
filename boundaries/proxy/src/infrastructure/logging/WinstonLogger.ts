import winston from "winston";
import { trace } from "@opentelemetry/api";
import { ILogger } from "./ILogger.js";

/**
 * Реализация ILogger на основе Winston
 * Инкапсулирует Winston логику в infrastructure слое
 * Автоматически добавляет trace_id и span_id из OpenTelemetry context
 */
export class WinstonLogger implements ILogger {
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
          format: winston.format.combine(
            winston.format.colorize(),
            winston.format.simple()
          ),
        }),
      ],
    });
  }

  /**
   * Извлекает trace_id и span_id из текущего OpenTelemetry context
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
   * Объединяет метаданные с trace context
   */
  private enrichMeta(meta?: any): any {
    const traceContext = this.getTraceContext();
    return { ...traceContext, ...meta };
  }

  info(message: string, meta?: any): void {
    this.logger.info(message, this.enrichMeta(meta));
  }

  warn(message: string, meta?: any): void {
    this.logger.warn(message, this.enrichMeta(meta));
  }

  error(message: string, error?: any, meta?: any): void {
    this.logger.error(message, this.enrichMeta({ error, ...meta }));
  }

  debug(message: string, meta?: any): void {
    this.logger.debug(message, this.enrichMeta(meta));
  }

  http(message: string): void {
    this.logger.info(message, this.enrichMeta({ severity: "http" }));
  }
}
