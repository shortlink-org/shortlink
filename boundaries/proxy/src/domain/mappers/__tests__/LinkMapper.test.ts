import { describe, it, expect } from "vitest";
import { LinkMapper } from "../LinkMapper.js";
import { Link } from "../../entities/Link.js";
import { Hash } from "../../entities/Hash.js";
import { Link as LinkProto, LinkSchema } from "../../../../infrastructure/proto/infrastructure/rpc/link/v1/link_pb.js";
import { create } from "@bufbuild/protobuf";
import { TimestampSchema } from "@bufbuild/protobuf/wkt";

describe("LinkMapper", () => {
  describe("toDomain", () => {
    it("should convert protobuf Link to domain entity", () => {
      // Arrange
      const now = Date.now();
      const seconds = BigInt(Math.floor(now / 1000));
      const nanos = (now % 1000) * 1000000;

      const protoLink = create(LinkSchema, {
        hash: "testhash123",
        url: "https://example.com",
        createdAt: create(TimestampSchema, { seconds, nanos }),
        updatedAt: create(TimestampSchema, { seconds, nanos }),
      });

      // Act
      const domainLink = LinkMapper.toDomain(protoLink);

      // Assert
      expect(domainLink).toBeInstanceOf(Link);
      expect(domainLink.hash.value).toBe("testhash123");
      expect(domainLink.url).toBe("https://example.com");
      expect(domainLink.createdAt).toBeInstanceOf(Date);
      expect(domainLink.updatedAt).toBeInstanceOf(Date);
    });

    it("should use current date when timestamp is missing", () => {
      // Arrange
      const protoLink = create(LinkSchema, {
        hash: "testhash123",
        url: "https://example.com",
      });

      // Act
      const domainLink = LinkMapper.toDomain(protoLink);

      // Assert
      expect(domainLink.createdAt).toBeInstanceOf(Date);
      expect(domainLink.updatedAt).toBeInstanceOf(Date);
    });
  });

  describe("toProto", () => {
    it("should convert domain entity to protobuf", () => {
      // Arrange
      const hash = new Hash("testhash123");
      const createdAt = new Date("2024-01-01");
      const updatedAt = new Date("2024-01-02");
      const link = new Link(hash, "https://example.com", createdAt, updatedAt);

      // Act
      const protoLink = LinkMapper.toProto(link);

      // Assert
      expect(protoLink.hash).toBe("testhash123");
      expect(protoLink.url).toBe("https://example.com");
      expect(protoLink.createdAt).toBeDefined();
      expect(protoLink.updatedAt).toBeDefined();
    });

    it("should handle links without dates", () => {
      // Arrange
      const hash = new Hash("testhash123");
      const link = new Link(hash, "https://example.com");

      // Act
      const protoLink = LinkMapper.toProto(link);

      // Assert
      expect(protoLink.hash).toBe("testhash123");
      expect(protoLink.url).toBe("https://example.com");
    });
  });

  describe("toDomainArray", () => {
    it("should convert array of protobuf Links to domain entities", () => {
      // Arrange
      const protoLinks = [
        create(LinkSchema, { hash: "hash1", url: "https://example1.com" }),
        create(LinkSchema, { hash: "hash2", url: "https://example2.com" }),
      ];

      // Act
      const domainLinks = LinkMapper.toDomainArray(protoLinks);

      // Assert
      expect(domainLinks).toHaveLength(2);
      expect(domainLinks[0].hash.value).toBe("hash1");
      expect(domainLinks[1].hash.value).toBe("hash2");
    });

    it("should handle empty array", () => {
      // Act
      const domainLinks = LinkMapper.toDomainArray([]);

      // Assert
      expect(domainLinks).toHaveLength(0);
    });
  });

  describe("toProtoArray", () => {
    it("should convert array of domain entities to protobuf", () => {
      // Arrange
      const links = [
        new Link(new Hash("hash1"), "https://example1.com"),
        new Link(new Hash("hash2"), "https://example2.com"),
      ];

      // Act
      const protoLinks = LinkMapper.toProtoArray(links);

      // Assert
      expect(protoLinks).toHaveLength(2);
      expect(protoLinks[0].hash).toBe("hash1");
      expect(protoLinks[1].hash).toBe("hash2");
    });

    it("should handle empty array", () => {
      // Act
      const protoLinks = LinkMapper.toProtoArray([]);

      // Assert
      expect(protoLinks).toHaveLength(0);
    });
  });
});

