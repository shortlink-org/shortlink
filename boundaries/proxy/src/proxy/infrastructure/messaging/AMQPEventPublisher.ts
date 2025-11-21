import { DomainEvent } from "../../domain/events/index.js";
import { IEventPublisher } from "../../application/use-cases/PublishEventUseCase.js";
import { LinkMapper } from "../../domain/mappers/LinkMapper.js";
import {
  Link as LinkProto,
  LinkSchema,
} from "../../../proto/infrastructure/rpc/link/v1/link_pb.js";
import { LinkRedirectedEvent } from "../../domain/events/index.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { IMessageBus } from "../../domain/interfaces/IMessageBus.js";
import { MQ_EVENT_LINK_NEW } from "../../domain/event.js";
import { toBinary } from "@bufbuild/protobuf";
import { context, propagation, trace } from "@opentelemetry/api";

/**
 * Реализация IEventPublisher для публикации событий в AMQP (RabbitMQ)
 * Преобразует доменные события в protobuf для сериализации
 * Использует IMessageBus абстракцию вместо прямого использования AMQP
 */
export class AMQPEventPublisher implements IEventPublisher {
  constructor(
    private readonly messageBus: IMessageBus,
    private readonly logger: ILogger
  ) {}

  /**
   * Инициализирует все необходимые exchange заранее (не лениво)
   * Должен быть вызван после connect() messageBus
   */
  async initializeExchanges(): Promise<void> {
    // Проверяем, что messageBus имеет метод setupExchange
    const bus = this.messageBus;

    if (!this.supportsExchangeSetup(bus)) {
      this.logger.warn(
        "MessageBus does not support setupExchange, skipping exchange initialization"
      );
      return;
    }

    // Exchange имена согласно ADR-0002: shortlink.{domain}.event.{event_name}
    const exchanges = [
      { name: MQ_EVENT_LINK_NEW, type: "fanout" as const },
      { name: "shortlink.link.event.redirected", type: "fanout" as const },
      { name: "shortlink.link.event.created", type: "fanout" as const },
      { name: "shortlink.link.event.updated", type: "fanout" as const },
      { name: "shortlink.link.event.deleted", type: "fanout" as const },
    ];

    for (const exchange of exchanges) {
      try {
        await bus.setupExchange(
          exchange.name,
          exchange.type,
          true // durable
        );
      } catch (error) {
        this.logger.error("Failed to setup exchange", error, {
          exchange: exchange.name,
        });
        // Продолжаем инициализацию других exchange даже если один не удался
      }
    }
  }

  async publish(event: DomainEvent): Promise<void> {
    try {
      // Преобразуем доменное событие в protobuf для сериализации
      const protoEvent = this.toProto(event);

      // Получаем имя exchange для типа события
      const exchange = this.getExchangeName(event.type);

      // Сериализуем protobuf объект в бинарный формат используя toBinary из @bufbuild/protobuf
      const binaryData = toBinary(LinkSchema, protoEvent);
      const messageBuffer = Buffer.from(binaryData);

      // Извлекаем текущий trace context и сериализуем в заголовки сообщения
      const headers = this.buildTraceHeaders();
      const traceparentValue =
        headers && typeof headers["traceparent"] === "string"
          ? (headers["traceparent"] as string)
          : undefined;

      // Публикуем событие через Message Bus
      // Exchange уже создан заранее через initializeExchanges(), поэтому не передаем exchangeType
      await this.messageBus.publish(
        exchange,
        undefined, // routingKey не используется для fanout exchange
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
      // Улучшенное логирование ошибок
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

      this.logger.error("Failed to publish event to AMQP", errorDetails, {
        eventType: event.type,
        exchange: this.getExchangeName(event.type),
      });
      throw error;
    }
  }

  /**
   * Преобразует доменное событие в protobuf для сериализации
   */
  private toProto(event: DomainEvent): LinkProto {
    if (event.type === "LinkRedirected") {
      const redirectedEvent = event as LinkRedirectedEvent;
      return LinkMapper.toProto(redirectedEvent.link);
    }

    // Для других типов событий можно добавить преобразование
    throw new Error(`Unsupported event type: ${event.type}`);
  }

  /**
   * Получает имя exchange для типа события
   */
  private getExchangeName(eventType: string): string {
    const exchangeMap: Record<string, string> = {
      LinkRedirected: "shortlink.link.event.redirected",
      LinkCreated: "shortlink.link.event.created",
      LinkUpdated: "shortlink.link.event.updated",
      LinkDeleted: "shortlink.link.event.deleted",
    };

    return (
      exchangeMap[eventType] || `shortlink.event.${eventType.toLowerCase()}`
    );
  }

  /**
   * Формирует заголовки сообщения с trace context согласно W3C Trace Context
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
   * Проверяет, поддерживает ли message bus предварительную настройку exchange
   */
  private supportsExchangeSetup(bus: IMessageBus): bus is IMessageBus & {
    setupExchange: (
      exchange: string,
      type?: "direct" | "fanout" | "topic" | "headers",
      durable?: boolean
    ) => Promise<void>;
  } {
    return (
      typeof (bus as { setupExchange?: unknown }).setupExchange === "function"
    );
  }
}
