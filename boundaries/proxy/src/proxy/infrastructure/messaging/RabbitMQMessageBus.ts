import { injectable, inject } from "inversify";
import { connect, Connection, ConfirmChannel } from "amqplib";
import {
  IMessageBus,
  MessageBusPublishOptions,
} from "../../domain/interfaces/IMessageBus.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { ConfigReader } from "../../../infrastructure/config/ConfigReader.js";
import TYPES from "../../../types.js";

/**
 * Реализация IMessageBus для RabbitMQ через AMQP
 * Инкапсулирует логику подключения и публикации сообщений
 * Каждый модуль сам читает свою конфигурацию (децентрализованный подход)
 */
@injectable()
export class RabbitMQMessageBus implements IMessageBus {
  private connection: Connection | undefined;
  private channel: ConfirmChannel | undefined;
  private isConnecting: boolean = false;
  private connected: boolean = false;

  // Конфигурация читается модулем напрямую
  private readonly enabled: boolean;
  private readonly rabbitUri: string;

  constructor(
    @inject(TYPES.INFRASTRUCTURE.Logger) private readonly logger: ILogger
  ) {
    // Модуль сам читает свою конфигурацию
    this.enabled = ConfigReader.boolean("MQ_ENABLED", false);
    this.rabbitUri = ConfigReader.string(
      "MQ_RABBIT_URI",
      "amqp://localhost:5672"
    );

    // Логируем конфигурацию для диагностики
    this.logger.debug("RabbitMQMessageBus initialized", {
      enabled: this.enabled,
      uri: this.rabbitUri,
      envMQEnabled: process.env.MQ_ENABLED,
      envMQRabbitURI: process.env.MQ_RABBIT_URI,
    });
  }

  async connect(): Promise<void> {
    if (this.isConnected()) {
      this.logger.debug("AMQP already connected");
      return;
    }

    if (this.isConnecting) {
      this.logger.debug("AMQP connection in progress");
      return;
    }

    if (!this.enabled) {
      this.logger.info("AMQP disabled, skipping connection");
      return;
    }

    try {
      this.isConnecting = true;
      this.logger.info("Connecting to AMQP", {
        uri: this.rabbitUri,
        enabled: this.enabled,
      });

      // Создаем connection
      this.connection = await connect(this.rabbitUri);

      // Настраиваем обработчики событий подключения
      this.connection.on("error", (err: Error) => {
        this.connected = false;
        this.logger.error("CONNECTION ERROR RAW", {
          raw: err,
          name: err?.name,
          message: err?.message,
          stack: err?.stack,
        });
        this.logger.error("AMQP connection error", err);
        this.isConnecting = false;
      });

      this.connection.on("close", () => {
        this.connected = false;
        this.channel = undefined;
        this.logger.warn("AMQP connection closed");
        this.isConnecting = false;
      });

      // Создаем confirm channel для подтверждения отправки сообщений
      this.channel = await this.connection.createConfirmChannel();

      this.channel.on("error", (err: Error) => {
        this.connected = false;
        this.logger.error("CHANNEL ERROR RAW", {
          raw: err,
          name: err?.name,
          message: err?.message,
          stack: err?.stack,
        });
        this.logger.error("AMQP channel error", err);
      });

      this.channel.on("close", () => {
        this.connected = false;
        this.channel = undefined;
        this.logger.warn("AMQP channel closed");
      });

      // Устанавливаем флаг подключения после успешного создания channel
      this.connected = true;
      this.isConnecting = false;
      this.logger.info("AMQP connection and channel initialized", {
        uri: this.rabbitUri,
      });
    } catch (error) {
      this.connected = false;
      this.isConnecting = false;
      this.logger.error("Failed to connect to AMQP", error);
      throw error;
    }
  }

  async disconnect(): Promise<void> {
    try {
      this.connected = false;

      if (this.channel) {
        this.logger.debug("Closing AMQP channel");
        await this.channel.close();
        this.channel = undefined;
      }

      if (this.connection) {
        this.logger.info("Disconnecting from AMQP");
        await this.connection.close();
        this.connection = undefined;
      }
    } catch (error) {
      this.connected = false;
      this.logger.error("Error disconnecting from AMQP", error);
      throw error;
    }
  }

  isConnected(): boolean {
    return this.connected && !!this.connection && !!this.channel;
  }

  /**
   * Настраивает exchange заранее (не лениво)
   * Должен быть вызван после connect() и перед publish()
   *
   * @param exchange - имя exchange
   * @param type - тип exchange (direct, fanout, topic, headers)
   * @param durable - должен ли exchange быть durable
   */
  async setupExchange(
    exchange: string,
    type: "direct" | "fanout" | "topic" | "headers" = "fanout",
    durable: boolean = true
  ): Promise<void> {
    if (!this.enabled) {
      this.logger.debug("AMQP disabled, skipping exchange setup", { exchange });
      return;
    }

    if (!this.isConnected()) {
      throw new Error(
        "AMQP not connected. Call connect() before setting up exchanges."
      );
    }

    if (!this.channel) {
      throw new Error("AMQP channel not available");
    }

    try {
      await this.channel.assertExchange(exchange, type, { durable });
      this.logger.debug("Exchange setup completed", {
        exchange,
        type,
        durable,
      });
    } catch (error) {
      this.logger.error("Failed to setup exchange", error, {
        exchange,
        type,
        durable,
      });
      throw error;
    }
  }

  async publish(
    exchange: string,
    routingKey: string | undefined,
    message: Buffer,
    options?: MessageBusPublishOptions
  ): Promise<void> {
    if (!this.enabled) {
      this.logger.debug("AMQP disabled, skipping publish", { exchange });
      return;
    }

    if (!this.isConnected()) {
      throw new Error(
        "AMQP not connected. Call connect() before publishing messages."
      );
    }

    if (!this.channel) {
      throw new Error("AMQP channel not available");
    }

    try {
      // Публикуем сообщение (exchange должен быть уже создан через setupExchange)
      const persistent = options?.persistent !== false; // По умолчанию persistent
      const headers = options?.headers;
      const hasTraceContext =
        !!headers &&
        (Object.prototype.hasOwnProperty.call(headers, "traceparent") ||
          Object.prototype.hasOwnProperty.call(headers, "tracestate"));

      this.logger.debug("Publishing message to AMQP", {
        exchange,
        routingKey: routingKey || "none",
        messageSize: message.length,
        hasTraceContext,
      });

      const published = this.channel.publish(
        exchange,
        routingKey || "",
        message,
        {
          persistent,
          headers,
        }
      );

      if (!published) {
        throw new Error(
          "Failed to publish message: channel buffer is full or connection closed"
        );
      }

      // Ждем подтверждения отправки
      // waitForConfirms() выбрасывает ошибку, если сообщение не было подтверждено брокером
      // (например, если exchange не существует)
      try {
        await this.channel.waitForConfirms();
      } catch (confirmError) {
        // Если ошибка подтверждения, логируем детально
        const confirmErrorDetails =
          confirmError instanceof Error
            ? {
                message: confirmError.message,
                stack: confirmError.stack,
                name: confirmError.name,
              }
            : {
                errorConstructor: confirmError?.constructor?.name || "Unknown",
                errorType: typeof confirmError,
                errorString: String(confirmError),
                errorKeys: Object.keys(confirmError || {}),
                errorJSON: JSON.stringify(
                  confirmError,
                  Object.getOwnPropertyNames(confirmError || {})
                ),
              };

        this.logger.error(
          "Failed to confirm message publication",
          confirmErrorDetails,
          {
            exchange,
            routingKey: routingKey || "none",
          }
        );
        throw confirmError;
      }

      this.logger.debug("Message published to AMQP", {
        exchange,
        routingKey: routingKey || "none",
        messageSize: message.length,
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

      this.logger.error("Failed to publish message to AMQP", errorDetails, {
        exchange,
        routingKey: routingKey || "none",
      });
      throw error;
    }
  }
}
