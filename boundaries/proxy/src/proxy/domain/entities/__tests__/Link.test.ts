import { Link } from "../Link.js";
import { Hash } from "../Hash.js";

describe("Link Entity", () => {
  const validHash = new Hash("abc123");
  const validUrl = "https://example.com";
  const createdAt = new Date("2024-01-01T00:00:00Z");
  const updatedAt = new Date("2024-01-02T00:00:00Z");

  describe("constructor", () => {
    it("should create Link with valid URL", () => {
      const link = new Link(validHash, validUrl);
      expect(link.hash).toEqual(validHash);
      expect(link.url).toBe(validUrl);
      expect(link.createdAt).toBeInstanceOf(Date);
      expect(link.updatedAt).toBeInstanceOf(Date);
    });

    it("should create Link with custom dates", () => {
      const link = new Link(validHash, validUrl, createdAt, updatedAt);
      expect(link.createdAt).toEqual(createdAt);
      expect(link.updatedAt).toEqual(updatedAt);
    });

    it("should use current date if dates not provided", () => {
      const beforeCreation = new Date();
      const link = new Link(validHash, validUrl);
      const afterCreation = new Date();

      expect(link.createdAt.getTime()).toBeGreaterThanOrEqual(
        beforeCreation.getTime()
      );
      expect(link.createdAt.getTime()).toBeLessThanOrEqual(
        afterCreation.getTime()
      );
      expect(link.updatedAt.getTime()).toBeGreaterThanOrEqual(
        beforeCreation.getTime()
      );
      expect(link.updatedAt.getTime()).toBeLessThanOrEqual(
        afterCreation.getTime()
      );
    });

    it("should throw Error for invalid URL", () => {
      expect(() => new Link(validHash, "not-a-url")).toThrow();
      expect(() => new Link(validHash, "")).toThrow();
      expect(() => new Link(validHash, "ftp://example.com")).not.toThrow(); // FTP is valid URL
    });

    it("should accept HTTP URLs", () => {
      const link = new Link(validHash, "http://example.com");
      expect(link.url).toBe("http://example.com");
    });

    it("should accept HTTPS URLs", () => {
      const link = new Link(validHash, "https://example.com");
      expect(link.url).toBe("https://example.com");
    });

    it("should accept URLs with paths", () => {
      const link = new Link(
        validHash,
        "https://example.com/path/to/resource"
      );
      expect(link.url).toBe("https://example.com/path/to/resource");
    });

    it("should accept URLs with query parameters", () => {
      const link = new Link(
        validHash,
        "https://example.com?param=value&other=123"
      );
      expect(link.url).toBe("https://example.com?param=value&other=123");
    });
  });

  describe("URL validation", () => {
    it("should accept valid URLs", () => {
      expect(() => new Link(validHash, "https://example.com")).not.toThrow();
      expect(() => new Link(validHash, "http://example.com")).not.toThrow();
      expect(() => new Link(validHash, "https://example.com/path")).not.toThrow();
      expect(() => new Link(validHash, "https://example.com?query=value")).not.toThrow();
    });

    it("should reject invalid URLs", () => {
      expect(() => new Link(validHash, "")).toThrow();
      expect(() => new Link(validHash, "not-a-url")).toThrow();
      expect(() => new Link(validHash, "example.com")).toThrow(); // Missing protocol
    });
  });

  describe("isValidForRedirect", () => {
    it("should return true for valid link", () => {
      const link = new Link(validHash, validUrl);
      expect(link.isValidForRedirect()).toBe(true);
    });

    it("should return true for HTTP link", () => {
      const link = new Link(validHash, "http://example.com");
      expect(link.isValidForRedirect()).toBe(true);
    });
  });

  describe("equals", () => {
    it("should return true for links with same hash", () => {
      const link1 = new Link(validHash, validUrl);
      const link2 = new Link(validHash, "https://different.com");
      expect(link1.equals(link2)).toBe(true);
    });

    it("should return false for links with different hash", () => {
      const hash1 = new Hash("abc123");
      const hash2 = new Hash("def456");
      const link1 = new Link(hash1, validUrl);
      const link2 = new Link(hash2, validUrl);
      expect(link1.equals(link2)).toBe(false);
    });
  });
});

