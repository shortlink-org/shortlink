import { describe, it, expect, beforeEach, vi } from "vitest";
import { EventDispatcher } from "../EventDispatcher.js";
import { IEventHandler } from "../IEventHandler.js";
import { DomainEvent } from "../../../domain/events/index.js";
import { LinkRedirectedEvent } from "../../../domain/events/index.js";

describe("EventDispatcher", () => {
  let dispatcher: EventDispatcher;
  let mockHandler1: IEventHandler<DomainEvent>;
  let mockHandler2: IEventHandler<DomainEvent>;

  beforeEach(() => {
    dispatcher = new EventDispatcher();
    mockHandler1 = {
      canHandle: vi.fn().mockReturnValue(true),
      handle: vi.fn().mockResolvedValue(undefined),
    } as unknown as IEventHandler<DomainEvent>;

    mockHandler2 = {
      canHandle: vi.fn().mockReturnValue(true),
      handle: vi.fn().mockResolvedValue(undefined),
    } as unknown as IEventHandler<DomainEvent>;
  });

  describe("register", () => {
    it("should register handler for event type", () => {
      // Act
      dispatcher.register("LinkRedirected", mockHandler1);

      // Assert
      expect(dispatcher["handlers"].has("LinkRedirected")).toBe(true);
    });

    it("should register multiple handlers for same event type", () => {
      // Act
      dispatcher.register("LinkRedirected", mockHandler1);
      dispatcher.register("LinkRedirected", mockHandler2);

      // Assert
      const handlers = dispatcher["handlers"].get("LinkRedirected");
      expect(handlers).toHaveLength(2);
      expect(handlers).toContain(mockHandler1);
      expect(handlers).toContain(mockHandler2);
    });
  });

  describe("dispatch", () => {
    it("should dispatch event to registered handlers", async () => {
      // Arrange
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        hash: { value: "test-hash" } as any,
        link: {} as any,
        occurredAt: new Date(),
        timestamp: new Date(),
      };

      dispatcher.register("LinkRedirected", mockHandler1);
      dispatcher.register("LinkRedirected", mockHandler2);

      // Act
      await dispatcher.dispatch(event);

      // Assert
      expect(mockHandler1.canHandle).toHaveBeenCalledWith(event);
      expect(mockHandler1.handle).toHaveBeenCalledWith(event);
      expect(mockHandler2.canHandle).toHaveBeenCalledWith(event);
      expect(mockHandler2.handle).toHaveBeenCalledWith(event);
    });

    it("should not dispatch to handlers that cannot handle event", async () => {
      // Arrange
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        hash: { value: "test-hash" } as any,
        link: {} as any,
        occurredAt: new Date(),
        timestamp: new Date(),
      };

      vi.mocked(mockHandler1.canHandle).mockReturnValue(false);
      dispatcher.register("LinkRedirected", mockHandler1);

      // Act
      await dispatcher.dispatch(event);

      // Assert
      expect(mockHandler1.canHandle).toHaveBeenCalledWith(event);
      expect(mockHandler1.handle).not.toHaveBeenCalled();
    });

    it("should handle events with no registered handlers", async () => {
      // Arrange
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        hash: { value: "test-hash" } as any,
        link: {} as any,
        occurredAt: new Date(),
        timestamp: new Date(),
      };

      // Act & Assert (should not throw)
      await expect(dispatcher.dispatch(event)).resolves.toBeUndefined();
    });

    it("should execute all handlers in parallel", async () => {
      // Arrange
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        hash: { value: "test-hash" } as any,
        link: {} as any,
        occurredAt: new Date(),
        timestamp: new Date(),
      };

      let handler1Start: number | null = null;
      let handler2Start: number | null = null;

      const asyncHandler1 = {
        canHandle: () => true,
        handle: async () => {
          handler1Start = Date.now();
          await new Promise((resolve) => setTimeout(resolve, 10));
        },
      };

      const asyncHandler2 = {
        canHandle: () => true,
        handle: async () => {
          handler2Start = Date.now();
          await new Promise((resolve) => setTimeout(resolve, 10));
        },
      };

      dispatcher.register("LinkRedirected", asyncHandler1);
      dispatcher.register("LinkRedirected", asyncHandler2);

      // Act
      await dispatcher.dispatch(event);

      // Assert - handlers should start at approximately the same time (parallel execution)
      expect(handler1Start).not.toBeNull();
      expect(handler2Start).not.toBeNull();
      if (handler1Start && handler2Start) {
        const timeDiff = Math.abs(handler1Start - handler2Start);
        expect(timeDiff).toBeLessThan(5); // Should start within 5ms of each other
      }
    });
  });
});

