import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { PublishEventUseCase, PublishEventRequest, IEventPublisher } from "../PublishEventUseCase.js";
import { DomainEvent, LinkRedirectedEvent } from "../../../domain/events/index.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";

describe("PublishEventUseCase", () => {
  let useCase: PublishEventUseCase;
  let mockEventPublisher: {
    publish: ReturnType<typeof vi.fn>;
  };

  beforeEach(() => {
    // Создаем мок Event Publisher
    mockEventPublisher = {
      publish: vi.fn().mockResolvedValue(undefined),
    } as any;

    // Создаем экземпляр Use Case с моком
    useCase = new PublishEventUseCase(mockEventPublisher as any);
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  describe("execute", () => {
    it("should publish event successfully", async () => {
      // Arrange
      const hash = new Hash("abc123");
      const link = new Link(
        hash,
        "https://example.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        occurredAt: new Date("2024-01-01T12:00:00Z"),
        hash,
        link,
        timestamp: new Date("2024-01-01T12:00:00Z"),
      };
      const request: PublishEventRequest = { event };

      // Act
      const result = await useCase.execute(request);

      // Assert
      expect(result.isOk()).toBe(true);
      if (result.isOk()) {
        expect(result.value.success).toBe(true);
      }
      expect(mockEventPublisher.publish).toHaveBeenCalledWith(event);
      expect(mockEventPublisher.publish).toHaveBeenCalledTimes(1);
    });

    it("should handle different event types", async () => {
      // Arrange
      const hash = new Hash("test123");
      const link = new Link(
        hash,
        "https://test.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        occurredAt: new Date("2024-01-01T12:00:00Z"),
        hash,
        link,
        timestamp: new Date("2024-01-01T12:00:00Z"),
      };
      const request: PublishEventRequest = { event };

      // Act
      const result = await useCase.execute(request);

      // Assert
      expect(result.isOk()).toBe(true);
      if (result.isOk()) {
        expect(result.value.success).toBe(true);
      }
      expect(mockEventPublisher.publish).toHaveBeenCalledWith(event);
    });

    it("should handle publisher errors", async () => {
      // Arrange
      const hash = new Hash("error");
      const link = new Link(
        hash,
        "https://error.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        occurredAt: new Date("2024-01-01T12:00:00Z"),
        hash,
        link,
        timestamp: new Date("2024-01-01T12:00:00Z"),
      };
      const request: PublishEventRequest = { event };
      const publisherError = new Error("AMQP connection failed");

      mockEventPublisher.publish.mockRejectedValue(publisherError);

      // Act & Assert
      await expect(useCase.execute(request)).rejects.toThrow(
        "AMQP connection failed"
      );
      expect(mockEventPublisher.publish).toHaveBeenCalledWith(event);
    });

    it("should handle async publisher operations", async () => {
      // Arrange
      const hash = new Hash("async");
      const link = new Link(
        hash,
        "https://async.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        occurredAt: new Date("2024-01-01T12:00:00Z"),
        hash,
        link,
        timestamp: new Date("2024-01-01T12:00:00Z"),
      };
      const request: PublishEventRequest = { event };

      // Симулируем асинхронную операцию
      let publishCalled = false;
      mockEventPublisher.publish.mockImplementation(async () => {
        publishCalled = true;
        await new Promise((resolve) => setTimeout(resolve, 10));
      });

      // Act
      const result = await useCase.execute(request);

      // Assert
      expect(result.isOk()).toBe(true);
      expect(publishCalled).toBe(true);
      if (result.isOk()) {
        expect(result.value.success).toBe(true);
      }
    });
  });
});

