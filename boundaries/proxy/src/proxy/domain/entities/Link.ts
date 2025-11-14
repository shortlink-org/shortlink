import { Hash } from "./Hash.js";

/**
 * Доменная сущность ссылки
 * Содержит бизнес-логику и инварианты домена
 */
export class Link {
  constructor(
    public readonly hash: Hash,
    public readonly url: string,
    public readonly createdAt: Date = new Date(),
    public readonly updatedAt: Date = new Date()
  ) {
    this.validate();
  }

  /**
   * Валидация инвариантов домена
   */
  private validate(): void {
    if (!this.url || typeof this.url !== "string" || this.url.trim().length === 0) {
      throw new Error("Link URL cannot be empty");
    }

    // Валидация URL формата (базовая проверка)
    try {
      new URL(this.url);
    } catch {
      throw new Error(`Invalid URL format: "${this.url}"`);
    }
  }

  /**
   * Проверяет, является ли ссылка валидной для редиректа
   */
  isValidForRedirect(): boolean {
    return this.url.length > 0 && this.hash.value.length > 0;
  }

  /**
   * Сравнивает две ссылки на равенство по хешу
   */
  equals(other: Link): boolean {
    return this.hash.equals(other.hash);
  }
}

