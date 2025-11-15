import { connect, type ChannelModel, type Channel, type ConsumeMessage } from "amqplib";

/**
 * Helper для подписки на сообщения из RabbitMQ в тестах
 * Позволяет проверять, что сообщения действительно публикуются
 */
export class RabbitMQTestConsumer {
  private connection: ChannelModel | undefined;
  private channel: Channel | undefined;
  private queueName: string | undefined;
  private messages: ConsumeMessage[] = [];

  /**
   * Подключается к RabbitMQ и создает очередь для подписки на exchange
   * @param amqpUrl - URI для подключения к RabbitMQ
   * @param exchangeName - имя exchange для подписки
   * @param exchangeType - тип exchange (fanout, direct, topic)
   */
  async connect(
    amqpUrl: string,
    exchangeName: string,
    exchangeType: "fanout" | "direct" | "topic" = "fanout"
  ): Promise<void> {
    // Создаем connection
    this.connection = await connect(amqpUrl);

    // Создаем channel
    this.channel = await this.connection.createChannel();

    // Создаем exchange (если еще не создан - это безопасно, RabbitMQ идемпотентен)
    await this.channel.assertExchange(exchangeName, exchangeType, {
      durable: true,
    });

    // Создаем временную очередь с автоматическим именем
    const queueResult = await this.channel.assertQueue("", {
      exclusive: true,
    });
    this.queueName = queueResult.queue;

    // Биндим очередь к exchange
    // Для fanout exchange routingKey не важен, но можно передать пустую строку
    await this.channel.bindQueue(this.queueName, exchangeName, "");

    // Подписываемся на сообщения
    await this.channel.consume(
      this.queueName,
      (message) => {
        if (message) {
          this.messages.push(message);
          // Подтверждаем получение сообщения
          this.channel?.ack(message);
        }
      },
      {
        noAck: false,
      }
    );
  }

  /**
   * Получает все полученные сообщения
   */
  getMessages(): ConsumeMessage[] {
    return [...this.messages];
  }

  /**
   * Получает количество полученных сообщений
   */
  getMessageCount(): number {
    return this.messages.length;
  }

  /**
   * Очищает список полученных сообщений
   */
  clearMessages(): void {
    this.messages = [];
  }

  /**
   * Ждет получения указанного количества сообщений
   * @param count - количество сообщений для ожидания
   * @param timeout - таймаут в миллисекундах (по умолчанию 5000)
   */
  async waitForMessages(
    count: number,
    timeout: number = 5000
  ): Promise<ConsumeMessage[]> {
    const startTime = Date.now();

    while (this.messages.length < count && Date.now() - startTime < timeout) {
      await new Promise((resolve) => setTimeout(resolve, 100));
    }

    if (this.messages.length < count) {
      throw new Error(
        `Timeout waiting for messages. Expected ${count}, got ${this.messages.length}`
      );
    }

    return this.messages.slice(0, count);
  }

  /**
   * Закрывает подключение к RabbitMQ
   */
  async disconnect(): Promise<void> {
    if (this.channel) {
      await this.channel.close();
      this.channel = undefined;
    }

    if (this.connection) {
      await this.connection.close();
      this.connection = undefined;
    }

    this.queueName = undefined;
    this.messages = [];
  }
}

