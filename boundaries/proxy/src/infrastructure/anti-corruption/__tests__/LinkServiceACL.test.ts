import { describe, it, expect, beforeEach } from "vitest";
import { LinkServiceACL } from "../LinkServiceACL.js";
import { Link } from "../../../domain/entities/Link.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link as LinkProto, LinkSchema } from "../../proto/infrastructure/rpc/link/v1/link_pb.js";
import { create } from "@bufbuild/protobuf";
import { TimestampSchema } from "@bufbuild/protobuf/wkt";

describe("LinkServiceACL", () => {
  let acl: LinkServiceACL;

  beforeEach(() => {
    acl = new LinkServiceACL();
  });

  describe("toDomainEntityFromProto", () => {
    it("should convert protobuf Link to domain entity", () => {
      // Arrange
      const protoLink = create(LinkSchema, {
        hash: "testhash123",
        url: "https://example.com",
        createdAt: create(TimestampSchema, {
          seconds: BigInt(Math.floor(Date.now() / 1000)),
          nanos: 0,
        }),
        updatedAt: create(TimestampSchema, {
          seconds: BigInt(Math.floor(Date.now() / 1000)),
          nanos: 0,
        }),
      });

      // Act
      const domainLink = acl.toDomainEntityFromProto(protoLink);

      // Assert
      expect(domainLink).toBeInstanceOf(Link);
      expect(domainLink.hash.value).toBe("testhash123");
      expect(domainLink.url).toBe("https://example.com");
      expect(domainLink.createdAt).toBeInstanceOf(Date);
      expect(domainLink.updatedAt).toBeInstanceOf(Date);
    });
  });

  describe("toDomainEntityFromHttp", () => {
    it("should convert HTTP JSON response to domain entity", () => {
      // Arrange
      const httpResponse = {
        hash: "testhash123",
        url: "https://example.com",
        createdAt: "2024-01-01T00:00:00Z",
        updatedAt: "2024-01-02T00:00:00Z",
      };

      // Act
      const domainLink = acl.toDomainEntityFromHttp(httpResponse);

      // Assert
      expect(domainLink).toBeInstanceOf(Link);
      expect(domainLink.hash.value).toBe("testhash123");
      expect(domainLink.url).toBe("https://example.com");
      expect(domainLink.createdAt).toBeInstanceOf(Date);
      expect(domainLink.updatedAt).toBeInstanceOf(Date);
    });

    it("should handle Date objects in HTTP response", () => {
      // Arrange
      const createdAt = new Date("2024-01-01");
      const updatedAt = new Date("2024-01-02");
      const httpResponse = {
        hash: "testhash123",
        url: "https://example.com",
        createdAt,
        updatedAt,
      };

      // Act
      const domainLink = acl.toDomainEntityFromHttp(httpResponse);

      // Assert
      expect(domainLink.createdAt).toBe(createdAt);
      expect(domainLink.updatedAt).toBe(updatedAt);
    });

    it("should use current date when dates are missing", () => {
      // Arrange
      const httpResponse = {
        hash: "testhash123",
        url: "https://example.com",
      };

      // Act
      const domainLink = acl.toDomainEntityFromHttp(httpResponse);

      // Assert
      expect(domainLink.createdAt).toBeInstanceOf(Date);
      expect(domainLink.updatedAt).toBeInstanceOf(Date);
      expect(domainLink.createdAt.getTime()).toBeGreaterThan(0);
      expect(domainLink.updatedAt.getTime()).toBeGreaterThan(0);
    });
  });

  describe("toProto", () => {
    it("should convert domain entity to protobuf", () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");

      // Act
      const protoLink = acl.toProto(link);

      // Assert
      expect(protoLink.hash).toBe("testhash123");
      expect(protoLink.url).toBe("https://example.com");
    });
  });
});

