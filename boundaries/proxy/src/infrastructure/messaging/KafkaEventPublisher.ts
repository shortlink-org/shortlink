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
import { toBinary } from "@bufbuild/protobuf";

/**
 * IEventPublisher implementation for publishing events to Kafka
 * Converts domain events to protobuf for serialization
 * Uses IMessageBus abstraction instead of direct Kafka usage
 */
export class KafkaEventPublisher implements IEventPublisher {
  constructor(
    private readonly messageBus: IMessageBus,
    private readonly logger: ILogger
  ) {}

  /**
   * Initializes all necessary topics in advance (not lazy)
   * Must be called after connect() messageBus
   * In Kafka, topics are created automatically on first publish (allowAutoTopicCreation: true)
   * or via KafkaTopic CRD, so this method can be empty
   */
  async initializeExchanges(): Promise<void> {
    // In Kafka, topics are created automatically or via CRD
    // This method is kept for interface compatibility
    this.logger.debug(
      "Kafka topics will be created automatically on first publish"
    );
  }

  async publish(event: DomainEvent): Promise<void> {
    try {
      // Convert domain event to protobuf for serialization
      const protoEvent = this.toProto(event);

      // Get topic name for event type
      const topic = this.getTopicName(event.type);

      // Serialize protobuf object to binary format using toBinary from @bufbuild/protobuf
      const binaryData = toBinary(LinkSchema, protoEvent);
      const messageBuffer = Buffer.from(binaryData);

      // Publish event via Message Bus
      // Trace context is automatically injected by OpenTelemetry KafkaJsInstrumentation
      // In Kafka, exchange corresponds to topic
      await this.messageBus.publish(
        topic,
        undefined, // routingKey not used for Kafka (can be used as partition key)
        messageBuffer,
        {
          persistent: true,
        }
      );

      this.logger.debug("Event published to Kafka", {
        eventType: event.type,
        topic,
        occurredAt: event.occurredAt,
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

      this.logger.error("Failed to publish event to Kafka", {
        error: errorDetails,
        eventType: event.type,
        topic: this.getTopicName(event.type),
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
   * Gets topic name for event type
   * Kafka uses format: shortlink.{domain}.event.{event_name}
   */
  private getTopicName(eventType: string): string {
    const topicMap: Record<string, string> = {
      LinkRedirected: "shortlink.link.event.redirected",
      LinkCreated: "shortlink.link.event.created",
      LinkUpdated: "shortlink.link.event.updated",
      LinkDeleted: "shortlink.link.event.deleted",
    };

    return topicMap[eventType] || `shortlink.event.${eventType.toLowerCase()}`;
  }
}
