/**
 * Абстракция Message Bus для публикации сообщений
 * Domain Layer интерфейс - не зависит от деталей реализации (AMQP, Kafka, etc.)
 */
export interface IMessageBus {
  /**
   * Публикует сообщение в указанный exchange/routing key
   *
   * @param exchange - имя exchange (например, "shortlink.link.event.redirected")
   * @param routingKey - routing key для сообщения (опционально)
   * @param message - данные сообщения в виде Buffer
   * @param options - дополнительные опции (durable, persistent и т.д.)
   */
  publish(
    exchange: string,
    routingKey: string | undefined,
    message: Buffer,
    options?: MessageBusPublishOptions
  ): Promise<void>;

  /**
   * Инициализирует подключение к message bus
   * Должен быть вызван перед использованием publish
   */
  connect(): Promise<void>;

  /**
   * Закрывает подключение к message bus
   */
  disconnect(): Promise<void>;

  /**
   * Проверяет, подключен ли message bus
   */
  isConnected(): boolean;
}

/**
 * Опции для публикации сообщений
 */
export interface MessageBusPublishOptions {
  /**
   * Должен ли exchange быть durable (пережить перезапуск брокера)
   */
  durable?: boolean;

  /**
   * Должно ли сообщение быть persistent (сохраняться на диск)
   */
  persistent?: boolean;

  /**
   * Пользовательские заголовки сообщения (например, traceparent, tracestate)
   */
  headers?: Record<string, unknown>;

  /**
   * Тип exchange (direct, fanout, topic, headers)
   */
  exchangeType?: "direct" | "fanout" | "topic" | "headers";
}
