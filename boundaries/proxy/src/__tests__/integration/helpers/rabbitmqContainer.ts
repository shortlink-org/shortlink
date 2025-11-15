import {
  RabbitMQContainer,
  StartedRabbitMQContainer,
} from "@testcontainers/rabbitmq";

/**
 * Helper для создания и управления RabbitMQ контейнером в тестах
 * Использует Testcontainers для изоляции тестов
 */
export class RabbitMQTestContainer {
  private container: RabbitMQContainer | undefined;
  private startedContainer: StartedRabbitMQContainer | undefined;

  /**
   * Запускает RabbitMQ контейнер
   * @returns URI для подключения к RabbitMQ
   */
  async start(): Promise<string> {
    // Используем стандартный образ RabbitMQ с management plugin
    // Используем более стабильный образ
    this.container = new RabbitMQContainer("rabbitmq:3.13-management");

    // Запускаем контейнер (Testcontainers автоматически настроит wait strategy)
    this.startedContainer = await this.container.start();

    // Дополнительная проверка готовности через rabbitmqctl
    // Ждем, пока RabbitMQ полностью запустится и будет готов принимать соединения
    // Используем await_startup для надежной проверки готовности
    let retries = 10;
    let lastError: Error | undefined;

    while (retries > 0) {
      try {
        const result = await this.startedContainer.exec([
          "rabbitmqctl",
          "await_startup",
        ]);
        if (result.exitCode === 0) {
          // RabbitMQ готов
          break;
        }
        lastError = new Error(
          `RabbitMQ await_startup failed with exit code ${result.exitCode}`
        );
      } catch (error) {
        lastError = error instanceof Error ? error : new Error(String(error));
      }

      retries--;
      if (retries > 0) {
        // Ждем перед следующей попыткой
        await new Promise((resolve) => setTimeout(resolve, 1000));
      }
    }

    // Если await_startup не сработал, пробуем status как финальную проверку
    if (retries === 0) {
      const result = await this.startedContainer.exec([
        "rabbitmqctl",
        "status",
      ]);
      if (result.exitCode !== 0) {
        throw new Error(
          `RabbitMQ status check failed with exit code ${result.exitCode}. Last error: ${lastError?.message}`
        );
      }
    }

    // Возвращаем AMQP URI для подключения
    return this.startedContainer.getAmqpUrl();
  }

  /**
   * Останавливает RabbitMQ контейнер
   */
  async stop(): Promise<void> {
    if (this.startedContainer) {
      await this.startedContainer.stop();
      this.startedContainer = undefined;
    }
    this.container = undefined;
  }

  /**
   * Получает URI для подключения к RabbitMQ
   * @throws Error если контейнер не запущен
   */
  getAmqpUrl(): string {
    if (!this.startedContainer) {
      throw new Error("RabbitMQ container is not started. Call start() first.");
    }
    return this.startedContainer.getAmqpUrl();
  }

  /**
   * Получает HTTP URL для управления RabbitMQ (Management API)
   */
  getHttpUrl(): string {
    if (!this.startedContainer) {
      throw new Error("RabbitMQ container is not started. Call start() first.");
    }
    const host = this.startedContainer.getHost();
    const port = this.startedContainer.getMappedPort(15672);
    return `http://${host}:${port}`;
  }
}

