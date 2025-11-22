import { Hash } from "../Hash.js";
import { InvalidHashError } from "../../exceptions/InvalidHashError.js";

describe("Hash Value Object", () => {
  describe("constructor", () => {
    it("should create Hash with valid alphanumeric string", () => {
      const hash = new Hash("abc123");
      expect(hash.value).toBe("abc123");
    });

    it("should create Hash with uppercase letters", () => {
      const hash = new Hash("ABC123");
      expect(hash.value).toBe("ABC123");
    });

    it("should create Hash with mixed case", () => {
      const hash = new Hash("AbC123");
      expect(hash.value).toBe("AbC123");
    });

    it("should create Hash with only numbers", () => {
      const hash = new Hash("123456");
      expect(hash.value).toBe("123456");
    });

    it("should create Hash with only letters", () => {
      const hash = new Hash("abcdef");
      expect(hash.value).toBe("abcdef");
    });

    it("should throw InvalidHashError for empty string", () => {
      expect(() => new Hash("")).toThrow(InvalidHashError);
    });

    it("should throw InvalidHashError for string with spaces", () => {
      expect(() => new Hash("abc 123")).toThrow(InvalidHashError);
    });

    it("should throw InvalidHashError for string with special characters", () => {
      expect(() => new Hash("abc-123")).toThrow(InvalidHashError);
      expect(() => new Hash("abc_123")).toThrow(InvalidHashError);
      expect(() => new Hash("abc.123")).toThrow(InvalidHashError);
      expect(() => new Hash("abc@123")).toThrow(InvalidHashError);
    });

    it("should throw InvalidHashError for string with unicode characters", () => {
      expect(() => new Hash("abc123абв")).toThrow(InvalidHashError);
    });
  });

  describe("validation", () => {
    it("should accept valid hash formats", () => {
      expect(() => new Hash("abc123")).not.toThrow();
      expect(() => new Hash("ABC123")).not.toThrow();
      expect(() => new Hash("123456")).not.toThrow();
      expect(() => new Hash("abcdef")).not.toThrow();
    });

    it("should reject invalid hash formats", () => {
      expect(() => new Hash("")).toThrow(InvalidHashError);
      expect(() => new Hash("abc-123")).toThrow(InvalidHashError);
      expect(() => new Hash("abc 123")).toThrow(InvalidHashError);
      expect(() => new Hash("abc@123")).toThrow(InvalidHashError);
    });
  });

  describe("equals", () => {
    it("should return true for equal hashes", () => {
      const hash1 = new Hash("abc123");
      const hash2 = new Hash("abc123");
      expect(hash1.equals(hash2)).toBe(true);
    });

    it("should return false for different hashes", () => {
      const hash1 = new Hash("abc123");
      const hash2 = new Hash("def456");
      expect(hash1.equals(hash2)).toBe(false);
    });

    it("should be case-sensitive", () => {
      const hash1 = new Hash("abc123");
      const hash2 = new Hash("ABC123");
      expect(hash1.equals(hash2)).toBe(false);
    });
  });

  describe("toString", () => {
    it("should return hash value as string", () => {
      const hash = new Hash("abc123");
      expect(hash.toString()).toBe("abc123");
    });
  });
});

