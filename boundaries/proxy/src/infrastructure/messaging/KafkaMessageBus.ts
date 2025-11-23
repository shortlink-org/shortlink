import { Kafka, type Producer, type KafkaConfig } from "kafkajs";
import {
  IMessageBus,
  MessageBusPublishOptions,
} from "../../domain/interfaces/IMessageBus.js";
import { ILogger } from "../logging/ILogger.js";
import { ConfigReader } from "../config/ConfigReader.js";

/**
 * IMessageBus implementation for Kafka
 * Encapsulates connection and message publishing logic
 * Each module reads its own configuration (decentralized approach)
 */
export class KafkaMessageBus implements IMessageBus {
  private producer: Producer | undefined;
  private isConnecting: boolean = false;
  private connected: boolean = false;
  private kafka: Kafka | undefined;

  // Configuration is read directly by the module
  private readonly enabled: boolean;
  private readonly kafkaUri: string;
  private readonly clientId: string;

  constructor(private readonly logger: ILogger) {
    // Module reads its own configuration
    this.enabled = ConfigReader.boolean("MQ_ENABLED", false);
    this.kafkaUri = ConfigReader.string(
      "MQ_KAFKA_URI",
      "localhost:9092"
    );
    this.clientId = ConfigReader.string(
      "MQ_KAFKA_CLIENT_ID",
      "shortlink-proxy"
    );

    // Log configuration for diagnostics
    this.logger.debug("KafkaMessageBus initialized", {
      enabled: this.enabled,
      uri: this.kafkaUri,
      clientId: this.clientId,
      envMQEnabled: process.env.MQ_ENABLED,
      envMQKafkaURI: process.env.MQ_KAFKA_URI,
    });
  }

  async connect(): Promise<void> {
    if (this.isConnected()) {
      this.logger.debug("Kafka already connected");
      return;
    }

    if (this.isConnecting) {
      this.logger.debug("Kafka connection in progress");
      return;
    }

    if (!this.enabled) {
      this.logger.info("Kafka disabled, skipping connection");
      return;
    }

    try {
      this.isConnecting = true;
      this.logger.info("Connecting to Kafka", {
        uri: this.kafkaUri,
        enabled: this.enabled,
        clientId: this.clientId,
      });

      // Parse URI (can contain multiple brokers separated by comma)
      const brokers = this.kafkaUri.split(",").map((b) => b.trim());

      // Create Kafka client configuration
      const kafkaConfig: KafkaConfig = {
        clientId: this.clientId,
        brokers,
        retry: {
          retries: 8,
          initialRetryTime: 100,
          multiplier: 2,
          maxRetryTime: 30000,
        },
        requestTimeout: 30000,
        connectionTimeout: 3000,
      };

      this.kafka = new Kafka(kafkaConfig);

      // Create producer
      this.producer = this.kafka.producer({
        allowAutoTopicCreation: true,
        idempotent: true,
        maxInFlightRequests: 1,
        transactionTimeout: 30000,
      });

      // Setup event handlers
      this.producer.on("producer.connect", () => {
        this.logger.debug("Kafka producer connected");
      });

      this.producer.on("producer.disconnect", () => {
        this.connected = false;
        this.logger.warn("Kafka producer disconnected");
      });

      this.producer.on("producer.network.request_timeout", (payload) => {
        this.logger.error("Kafka producer network timeout", payload);
      });

      // Connect
      await this.producer.connect();

      // Set connection flag after successful connection
      this.connected = true;
      this.isConnecting = false;
      this.logger.info("Kafka producer connected", {
        uri: this.kafkaUri,
        clientId: this.clientId,
      });
    } catch (error) {
      this.connected = false;
      this.isConnecting = false;
      this.logger.error("Failed to connect to Kafka", error);
      throw error;
    }
  }

  async disconnect(): Promise<void> {
    try {
      this.connected = false;

      if (this.producer) {
        this.logger.info("Disconnecting from Kafka");
        await this.producer.disconnect();
        this.producer = undefined;
      }

      this.kafka = undefined;
    } catch (error) {
      this.connected = false;
      this.logger.error("Error disconnecting from Kafka", error);
      throw error;
    }
  }

  isConnected(): boolean {
    return this.connected && !!this.producer;
  }

  async publish(
    exchange: string,
    routingKey: string | undefined,
    message: Buffer,
    options?: MessageBusPublishOptions
  ): Promise<void> {
    if (!this.enabled) {
      this.logger.debug("Kafka disabled, skipping publish", { exchange });
      return;
    }

    if (!this.isConnected()) {
      throw new Error(
        "Kafka not connected. Call connect() before publishing messages."
      );
    }

    if (!this.producer) {
      throw new Error("Kafka producer not available");
    }

    try {
      // In Kafka, exchange corresponds to topic
      // routingKey can be used as partition key
      const topic = exchange;
      const partitionKey = routingKey || undefined;

      // Convert headers from options
      const headers: Record<string, string> = {};
      if (options?.headers) {
        for (const [key, value] of Object.entries(options.headers)) {
          // Kafka headers must be strings or buffers
          if (typeof value === "string") {
            headers[key] = value;
          } else if (value !== null && value !== undefined) {
            headers[key] = String(value);
          }
        }
      }

      this.logger.debug("Publishing message to Kafka", {
        topic,
        partitionKey: partitionKey || "none",
        messageSize: message.length,
        hasHeaders: Object.keys(headers).length > 0,
      });

      // Publish message
      await this.producer.send({
        topic,
        messages: [
          {
            key: partitionKey,
            value: message,
            headers: Object.keys(headers).length > 0 ? headers : undefined,
          },
        ],
      });

      this.logger.debug("Message published to Kafka", {
        topic,
        partitionKey: partitionKey || "none",
        messageSize: message.length,
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

      this.logger.error("Failed to publish message to Kafka", {
        error: errorDetails,
        exchange,
        routingKey: routingKey || "none",
      });
      throw error;
    }
  }
}

