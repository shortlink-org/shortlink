import { DomainEvent } from "../../domain/events/index.js";
import { IEventPublisher } from "../../application/use-cases/PublishEventUseCase.js";
import { LinkMapper } from "../../domain/mappers/LinkMapper.js";
import {
  Link as LinkProto,
  LinkSchema,
} from "@buf/shortlink-org_shortlink-link-link.bufbuild_es/infrastructure/rpc/link/v1/link_rpc_pb.js";
import { LinkRedirectedEvent } from "../../domain/events/index.js";
import { ILogger } from "../logging/ILogger.js";
import { IMessageBus } from "../../domain/interfaces/IMessageBus.js";
import {
  MQ_EVENT_LINK_NEW,
  LinkEventTopics,
  EventTypeToTopic,
} from "../../domain/event.js";
import { toBinary } from "@bufbuild/protobuf";
import { context, propagation, trace } from "@opentelemetry/api";
import { RabbitMQMessageBus } from "./RabbitMQMessageBus.js";

/**
 * IEventPublisher implementation for publishing events to AMQP (RabbitMQ)
 * Converts domain events to protobuf for serialization
 * Uses IMessageBus abstraction instead of direct AMQP usage
 */
export class AMQPEventPublisher implements IEventPublisher {
  constructor(
    private readonly messageBus: IMessageBus,
    private readonly logger: ILogger
  ) {}

  /**
   * Initializes all necessary exchanges in advance (not lazy)
   * Must be called after connect() messageBus
   */
  async initializeExchanges(): Promise<void> {
    // Check that messageBus has setupExchange method
    const bus = this.messageBus;

    if (!this.supportsExchangeSetup(bus)) {
      this.logger.warn(
        "MessageBus does not support setupExchange, skipping exchange initialization"
      );
      return;
    }

    // Exchange names following ADR-0002 canonical naming: {service}.{aggregate}.{event}.{version}
    // Use constants from domain layer for type safety
    const exchanges = [
      { name: MQ_EVENT_LINK_NEW, type: "fanout" as const }, // Legacy, keep for backward compatibility
      { name: LinkEventTopics.REDIRECTED, type: "fanout" as const },
      { name: LinkEventTopics.CREATED, type: "fanout" as const },
      { name: LinkEventTopics.UPDATED, type: "fanout" as const },
      { name: LinkEventTopics.DELETED, type: "fanout" as const },
    ];

    for (const exchange of exchanges) {
      try {
        await bus.setupExchange(
          exchange.name,
          exchange.type,
          true // durable
        );
      } catch (error) {
        this.logger.error("Failed to setup exchange", {
          error: error instanceof Error ? error : new Error(String(error)),
          exchange: exchange.name,
        });
        // Continue initializing other exchanges even if one fails
      }
    }
  }

  async publish(event: DomainEvent): Promise<void> {
    try {
      // Convert domain event to protobuf for serialization
      const protoEvent = this.toProto(event);

      // Get exchange name for event type
      const exchange = this.getExchangeName(event.type);

      // Serialize protobuf object to binary format using toBinary from @bufbuild/protobuf
      const binaryData = toBinary(LinkSchema, protoEvent);
      const messageBuffer = Buffer.from(binaryData);

      // Extract current trace context and serialize into message headers
      const headers = this.buildTraceHeaders();
      const traceparentValue =
        headers && typeof headers["traceparent"] === "string"
          ? (headers["traceparent"] as string)
          : undefined;

      // Publish event via Message Bus
      // Exchange is already created in advance via initializeExchanges(), so we don't pass exchangeType
      await this.messageBus.publish(
        exchange,
        undefined, // routingKey is not used for fanout exchange
        messageBuffer,
        {
          persistent: true,
          headers,
        }
      );

      this.logger.debug("Event published to AMQP", {
        eventType: event.type,
        exchange,
        occurredAt: event.occurredAt,
        traceparent: traceparentValue,
      });
    } catch (error) {
      // Enhanced error logging
      const errorDetails =
        error instanceof Error
          ? { message: error.message, stack: error.stack, name: error.name }
          : {
              errorConstructor: error?.constructor?.name || "Unknown",
              errorType: typeof error,
              errorString: String(error),
              errorKeys: Object.keys(error || {}),
              errorJSON: JSON.stringify(
                error,
                Object.getOwnPropertyNames(error || {})
              ),
            };

      this.logger.error("Failed to publish event to AMQP", {
        error: errorDetails,
        eventType: event.type,
        exchange: this.getExchangeName(event.type),
      });
      throw error;
    }
  }

  /**
   * Converts domain event to protobuf for serialization
   */
  private toProto(event: DomainEvent): LinkProto {
    if (event.type === "LinkRedirected") {
      const redirectedEvent = event as LinkRedirectedEvent;
      return LinkMapper.toProto(redirectedEvent.link);
    }

    // Can add conversion for other event types
    throw new Error(`Unsupported event type: ${event.type}`);
  }

  /**
   * Gets exchange name for event type using canonical naming (ADR-0002)
   * Uses constants from domain layer for type safety
   */
  private getExchangeName(eventType: string): string {
    return (
      EventTypeToTopic[eventType] || `link.link.${eventType.toLowerCase()}.v1`
    );
  }

  /**
   * Builds message headers with trace context according to W3C Trace Context
   */
  private buildTraceHeaders(): Record<string, unknown> | undefined {
    const carrier: Record<string, unknown> = {};
    const activeContext = context.active();

    propagation.inject(activeContext, carrier);

    if (!carrier["traceparent"]) {
      const span = trace.getActiveSpan();
      const spanContext = span?.spanContext();

      if (spanContext?.traceId && spanContext?.spanId) {
        const flags = spanContext.traceFlags
          .toString(16)
          .padStart(2, "0")
          .toLowerCase();

        carrier[
          "traceparent"
        ] = `00-${spanContext.traceId}-${spanContext.spanId}-${flags}`;

        if (spanContext.traceState) {
          carrier["tracestate"] = spanContext.traceState.serialize();
        }
      }
    }

    return Object.keys(carrier).length > 0 ? carrier : undefined;
  }

  /**
   * Checks if message bus supports exchange setup
   * Uses safe instanceof check to avoid Awilix proxy activation
   */
  private supportsExchangeSetup(bus: IMessageBus): bus is IMessageBus & {
    setupExchange: (
      exchange: string,
      type?: "direct" | "fanout" | "topic" | "headers",
      durable?: boolean
    ) => Promise<void>;
  } {
    // Check via instanceof to avoid Awilix proxy activation
    // RabbitMQMessageBus is the only class that supports setupExchange
    if (bus instanceof RabbitMQMessageBus) {
      return true;
    }

    // Fallback: check via constructor name (safe for proxy)
    const constructorName = bus.constructor?.name;
    if (constructorName === "RabbitMQMessageBus") {
      return true;
    }

    return false;
  }
}
