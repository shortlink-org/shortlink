import {
  describe,
  it,
  expect,
  beforeAll,
  afterAll,
  beforeEach,
  afterEach,
  vi,
} from "vitest";
import type { FastifyInstance } from "fastify";
import { createTestServer } from "../helpers/testServer.js";
import { RabbitMQTestContainer } from "../helpers/rabbitmqContainer.js";
import { RabbitMQTestConsumer } from "../helpers/rabbitmqConsumer.js";
import { LinkApplicationService } from "../../../application/services/LinkApplicationService.js";
import { GetLinkByHashUseCase } from "../../../application/use-cases/GetLinkByHashUseCase.js";
import { PublishEventUseCase } from "../../../application/use-cases/PublishEventUseCase.js";
import { LinkServiceRepository } from "../../../infrastructure/repositories/LinkServiceRepository.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import { LinkEventTopics } from "../../../domain/event.js";
import type { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { IMessageBus } from "../../../domain/interfaces/IMessageBus.js";
import { RabbitMQMessageBus } from "../../../infrastructure/messaging/RabbitMQMessageBus.js";
import { AMQPEventPublisher } from "../../../infrastructure/messaging/AMQPEventPublisher.js";
import type { IEventPublisher } from "../../../application/use-cases/PublishEventUseCase.js";
import {
  UseCasePipeline,
  LoggingInterceptor,
  MetricsInterceptor,
} from "../../../application/pipeline/index.js";
import { WinstonLogger } from "../../../infrastructure/logging/WinstonLogger.js";
import { LinkSchema } from "@buf/shortlink-org_shortlink-link-link.bufbuild_es/infrastructure/rpc/link/v1/link_rpc_pb.js";
import { fromBinary } from "@bufbuild/protobuf";
import type { ILinkCache } from "../../../infrastructure/cache/RedisLinkCache.js";
import type { ILinkServiceAdapter } from "../../../infrastructure/adapters/ILinkServiceAdapter.js";

/**
 * End-to-end интеграционные тесты с реальным RabbitMQ через Testcontainers
 * Тестируют полный flow: Controller → Application Service → Use Cases → Message Bus → RabbitMQ
 *
 * Требования:
 * - Docker должен быть запущен
 * - Тесты используют Testcontainers для изоляции
 */
describe("Redirect Flow E2E with RabbitMQ (Testcontainers)", () => {
  let app: FastifyInstance;
  let rabbitMQContainer: RabbitMQTestContainer;
  let rabbitMQConsumer: RabbitMQTestConsumer;
  let getLinkByHashMock: ReturnType<
    typeof vi.fn<(hash: Hash) => Promise<Link | null>>
  >;
  let cacheGetMock: ReturnType<
    typeof vi.fn<(hash: Hash) => Promise<Link | null | undefined>>
  >;
  let cacheSetPositiveMock: ReturnType<
    typeof vi.fn<(hash: Hash, link: Link) => Promise<void>>
  >;
  let cacheSetNegativeMock: ReturnType<
    typeof vi.fn<(hash: Hash) => Promise<void>>
  >;
  let cacheClearMock: ReturnType<typeof vi.fn<(hash: Hash) => Promise<void>>>;
  let mockLinkServiceAdapter: ILinkServiceAdapter;
  let mockLinkCache: ILinkCache;
  let messageBus: IMessageBus;
  let amqpUrl: string;

  beforeAll(async () => {
    // Запускаем RabbitMQ контейнер один раз для всех тестов
    rabbitMQContainer = new RabbitMQTestContainer();
    amqpUrl = await rabbitMQContainer.start();

    console.log(`[Test] RabbitMQ started at: ${amqpUrl}`);
    console.log(`[Test] RabbitMQ URL details:`, {
      url: amqpUrl,
      host: new URL(amqpUrl).hostname,
      port: new URL(amqpUrl).port,
      protocol: new URL(amqpUrl).protocol,
    });
  }, 120000); // Увеличиваем таймаут для запуска контейнера (2 минуты)

  afterAll(async () => {
    // Останавливаем RabbitMQ контейнер
    await rabbitMQContainer.stop();
    console.log("[Test] RabbitMQ stopped");
  }, 30000);

  beforeEach(async () => {
    // Мокаем LinkServiceAdapter
    getLinkByHashMock = vi.fn<(hash: Hash) => Promise<Link | null>>();
    mockLinkServiceAdapter = {
      getLinkByHash: getLinkByHashMock,
    } as any;

    cacheGetMock = vi
      .fn<(hash: Hash) => Promise<Link | null | undefined>>()
      .mockResolvedValue(undefined);
    cacheSetPositiveMock = vi
      .fn<(hash: Hash, link: Link) => Promise<void>>()
      .mockResolvedValue(undefined);
    cacheSetNegativeMock = vi
      .fn<(hash: Hash) => Promise<void>>()
      .mockResolvedValue(undefined);
    cacheClearMock = vi
      .fn<(hash: Hash) => Promise<void>>()
      .mockResolvedValue(undefined);
    mockLinkCache = {
      get: cacheGetMock,
      setPositive: cacheSetPositiveMock,
      setNegative: cacheSetNegativeMock,
      clear: cacheClearMock,
    } as any;

    // Создаем реальный Logger
    const logger = new WinstonLogger();

    // Создаем реальный RabbitMQMessageBus с URL из Testcontainers
    // Устанавливаем переменные окружения для конфигурации
    process.env.MQ_ENABLED = "true";
    process.env.MQ_RABBIT_URI = amqpUrl;

    messageBus = new RabbitMQMessageBus(logger);
    await messageBus.connect();

    // Создаем реальный AMQPEventPublisher
    const eventPublisher = new AMQPEventPublisher(messageBus, logger);
    // Инициализируем exchange заранее (не лениво)
    await eventPublisher.initializeExchanges();

    // Создаем Fastify сервер с реальными зависимостями, но мокированными адаптерами
    app = await createTestServer({
      linkServiceAdapter: mockLinkServiceAdapter,
      linkCache: mockLinkCache,
      eventPublisher,
      messageBus,
      logger,
      // Остальные зависимости будут взяты из контейнера
    } as any);

    // Создаем consumer для подписки на события
    rabbitMQConsumer = new RabbitMQTestConsumer();
    await rabbitMQConsumer.connect(
      amqpUrl,
      LinkEventTopics.REDIRECTED,
      "fanout"
    );
  }, 30000); // Увеличиваем таймаут для подключения к RabbitMQ

  afterEach(async () => {
    // Закрываем Fastify сервер
    await app.close();

    // Очищаем consumer
    await rabbitMQConsumer.disconnect();

    // Отключаемся от Message Bus
    if (messageBus.isConnected()) {
      await messageBus.disconnect();
    }

    // Очищаем переменные окружения
    delete process.env.MQ_ENABLED;
    delete process.env.MQ_RABBIT_URI;

    vi.clearAllMocks();
  });

  describe("Full redirect flow with RabbitMQ", () => {
    it("should publish LinkRedirected event to RabbitMQ on successful redirect", async () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");

      getLinkByHashMock.mockResolvedValue(link);
      rabbitMQConsumer.clearMessages();

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/testhash123",
      });

      // Assert - проверяем редирект
      expect(response.statusCode).toBe(301);
      expect(response.headers.location).toBe("https://example.com");

      // Ждем получения сообщения из RabbitMQ (увеличиваем таймаут для надежности)
      const messages = await rabbitMQConsumer.waitForMessages(1, 10000);
      expect(messages.length).toBe(1);

      // Проверяем содержимое сообщения
      const message = messages[0];
      // amqplib ConsumeMessage имеет свойство content (Buffer)
      const messageContent = message.content;
      expect(messageContent).toBeDefined();

      // Декодируем protobuf сообщение
      const linkProto = fromBinary(LinkSchema, messageContent);
      expect(linkProto.hash).toBe("testhash123");
      expect(linkProto.url).toBe("https://example.com");
    });

    it("should not publish event when link is not found", async () => {
      // Arrange
      getLinkByHashMock.mockResolvedValue(null);
      rabbitMQConsumer.clearMessages();

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/nonexistent",
      });

      // Assert
      expect(response.statusCode).toBe(404);

      // Ждем немного, чтобы убедиться, что сообщение не было отправлено
      await new Promise((resolve) => setTimeout(resolve, 1000));
      expect(rabbitMQConsumer.getMessageCount()).toBe(0);
    });

    it("should handle multiple redirects and publish multiple events", async () => {
      // Arrange
      const hash1 = new Hash("hash1");
      const link1 = new Link(hash1, "https://example1.com");
      const hash2 = new Hash("hash2");
      const link2 = new Link(hash2, "https://example2.com");

      getLinkByHashMock
        .mockResolvedValueOnce(link1)
        .mockResolvedValueOnce(link2);

      rabbitMQConsumer.clearMessages();

      // Act
      const response1 = await app.inject({
        method: "GET",
        url: "/s/hash1",
      });
      const response2 = await app.inject({
        method: "GET",
        url: "/s/hash2",
      });

      // Assert
      expect(response1.statusCode).toBe(301);
      expect(response2.statusCode).toBe(301);

      // Ждем получения обоих сообщений (увеличиваем таймаут для надежности)
      const messages = await rabbitMQConsumer.waitForMessages(2, 10000);
      expect(messages.length).toBe(2);

      // Проверяем содержимое сообщений
      const linkProto1 = fromBinary(LinkSchema, messages[0].content);
      const linkProto2 = fromBinary(LinkSchema, messages[1].content);

      expect(linkProto1.hash).toBe("hash1");
      expect(linkProto1.url).toBe("https://example1.com");
      expect(linkProto2.hash).toBe("hash2");
      expect(linkProto2.url).toBe("https://example2.com");
    });

    it("should continue redirect even if RabbitMQ is unavailable", async () => {
      // Arrange
      const hash = new Hash("testhash");
      const link = new Link(hash, "https://example.com");

      getLinkByHashMock.mockResolvedValue(link);

      // Отключаемся от RabbitMQ перед запросом
      await messageBus.disconnect();
      rabbitMQConsumer.clearMessages();

      // Act
      const response = await app.inject({
        method: "GET",
        url: "/s/testhash",
      });

      // Assert - редирект должен выполниться, даже если RabbitMQ недоступен
      expect(response.statusCode).toBe(301);
      expect(response.headers.location).toBe("https://example.com");

      // Сообщение не должно быть отправлено (RabbitMQ недоступен)
      await new Promise((resolve) => setTimeout(resolve, 1000));
      expect(rabbitMQConsumer.getMessageCount()).toBe(0);
    });
  });
});
