import { InvalidHashError } from "../exceptions/InvalidHashError.js";

/**
 * Value Object для хеша ссылки
 * Инкапсулирует валидацию формата хеша
 */
export class Hash {
  private static readonly VALID_HASH_PATTERN = /^[a-zA-Z0-9]+$/;

  constructor(public readonly value: string) {
    if (!this.isValid(value)) {
      throw new InvalidHashError(
        value,
        "Hash must contain only alphanumeric characters"
      );
    }
  }

  /**
   * Проверяет валидность формата хеша
   */
  private isValid(value: string): boolean {
    return (
      typeof value === "string" &&
      value.length > 0 &&
      Hash.VALID_HASH_PATTERN.test(value)
    );
  }

  /**
   * Сравнивает два хеша на равенство
   */
  equals(other: Hash): boolean {
    return this.value === other.value;
  }

  /**
   * Возвращает строковое представление хеша
   */
  toString(): string {
    return this.value;
  }
}

