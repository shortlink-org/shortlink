import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { GetLinkByHashUseCase } from "../GetLinkByHashUseCase.js";
import { GetLinkRequest } from "../../dto/GetLinkRequest.js";
import { GetLinkResponse } from "../../dto/GetLinkResponse.js";
import { ILinkRepository } from "../../../domain/repositories/ILinkRepository.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { Link } from "../../../domain/entities/Link.js";
import { LinkNotFoundError } from "../../../domain/exceptions/index.js";

describe("GetLinkByHashUseCase", () => {
  let useCase: GetLinkByHashUseCase;
  let mockLinkRepository: {
    findByHash: ReturnType<typeof vi.fn>;
    save: ReturnType<typeof vi.fn>;
    exists: ReturnType<typeof vi.fn>;
  };

  beforeEach(() => {
    // Создаем мок репозитория
    mockLinkRepository = {
      findByHash: vi.fn(),
      save: vi.fn(),
      exists: vi.fn(),
    } as any;

    // Создаем экземпляр Use Case с моком
    useCase = new GetLinkByHashUseCase(mockLinkRepository as any);
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  describe("execute", () => {
    it("should return link when found", async () => {
      // Arrange
      const hash = new Hash("abc123");
      const link = new Link(
        hash,
        "https://example.com",
        new Date("2024-01-01"),
        new Date("2024-01-02")
      );
      const request = new GetLinkRequest("abc123");

      mockLinkRepository.findByHash.mockResolvedValue(link);

      // Act
      const result = await useCase.execute(request);

      // Assert
      expect(result).toBeInstanceOf(GetLinkResponse);
      expect(result.link).toEqual(link);
      expect(result.link.hash.value).toBe("abc123");
      expect(result.link.url).toBe("https://example.com");
      expect(mockLinkRepository.findByHash).toHaveBeenCalledTimes(1);
      expect(mockLinkRepository.findByHash).toHaveBeenCalledWith(hash);
    });

    it("should throw LinkNotFoundError when link not found", async () => {
      // Arrange
      const request = new GetLinkRequest("nonexistent");

      mockLinkRepository.findByHash.mockResolvedValue(null);

      // Act & Assert
      await expect(useCase.execute(request)).rejects.toThrow(LinkNotFoundError);
      expect(mockLinkRepository.findByHash).toHaveBeenCalledTimes(1);
    });

    it("should handle repository errors", async () => {
      // Arrange
      const request = new GetLinkRequest("error");
      const repositoryError = new Error("Database connection failed");

      mockLinkRepository.findByHash.mockRejectedValue(repositoryError);

      // Act & Assert
      await expect(useCase.execute(request)).rejects.toThrow(
        "Database connection failed"
      );
    });
  });
});

