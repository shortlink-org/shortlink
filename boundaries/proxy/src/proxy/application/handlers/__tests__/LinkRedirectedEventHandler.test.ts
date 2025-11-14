import { describe, it, expect } from "vitest";
import { LinkRedirectedEventHandler } from "../LinkRedirectedEventHandler.js";
import { LinkRedirectedEvent, DomainEvent } from "../../../domain/events/index.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";

describe("LinkRedirectedEventHandler", () => {
  let handler: LinkRedirectedEventHandler;

  beforeEach(() => {
    handler = new LinkRedirectedEventHandler();
  });

  describe("canHandle", () => {
    it("should return true for LinkRedirected event", () => {
      // Arrange
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        hash: new Hash("testhash123"),
        link: new Link(new Hash("testhash123"), "https://example.com"),
        occurredAt: new Date(),
      };

      // Act
      const result = handler.canHandle(event);

      // Assert
      expect(result).toBe(true);
    });

    it("should return false for other event types", () => {
      // Arrange
      const event: DomainEvent = {
        type: "OtherEvent",
        occurredAt: new Date(),
      } as DomainEvent;

      // Act
      const result = handler.canHandle(event);

      // Assert
      expect(result).toBe(false);
    });
  });

  describe("handle", () => {
    it("should handle event without throwing", async () => {
      // Arrange
      const event: LinkRedirectedEvent = {
        type: "LinkRedirected",
        hash: new Hash("testhash123"),
        link: new Link(new Hash("testhash123"), "https://example.com"),
        occurredAt: new Date(),
      };

      // Act & Assert
      await expect(handler.handle(event)).resolves.toBeUndefined();
    });
  });
});

