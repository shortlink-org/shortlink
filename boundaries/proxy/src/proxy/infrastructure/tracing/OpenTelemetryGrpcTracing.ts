import { injectable } from "inversify";
import { trace, Span, SpanStatusCode, context } from "@opentelemetry/api";
import { IGrpcTracing } from "./IGrpcTracing.js";

/**
 * Реализация трейсинга gRPC вызовов через OpenTelemetry Tracing API
 * Создает spans для каждого gRPC вызова для распределенного трейсинга
 */
@injectable()
export class OpenTelemetryGrpcTracing implements IGrpcTracing {
  private readonly tracer = trace.getTracer("proxy-service", "1.0.0");

  startSpan(method: string, operation: string): Span {
    return this.tracer.startSpan(`grpc.${operation}`, {
      kind: 2, // SpanKind.CLIENT
      attributes: {
        "rpc.system": "grpc",
        "rpc.service": "LinkService",
        "rpc.method": method,
        "rpc.grpc.status_code": 0, // Will be updated on completion
      },
    });
  }

  endSpan(span: Span): void {
    span.setStatus({ code: SpanStatusCode.OK });
    span.setAttribute("rpc.grpc.status_code", 0); // OK
    span.end();
  }

  endSpanWithError(span: Span, error: Error): void {
    span.setStatus({
      code: SpanStatusCode.ERROR,
      message: error.message,
    });
    span.recordException(error);
    span.end();
  }
}

