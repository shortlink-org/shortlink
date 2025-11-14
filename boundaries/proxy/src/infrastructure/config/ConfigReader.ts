/**
 * Утилита для чтения переменных окружения с дефолтными значениями
 * Поддерживает децентрализованную конфигурацию - каждый модуль сам определяет свои дефолты
 */
export class ConfigReader {
  /**
   * Читает строковую переменную окружения
   *
   * @param key - имя переменной окружения
   * @param defaultValue - значение по умолчанию
   * @returns значение переменной или дефолт
   */
  static string(key: string, defaultValue: string): string {
    return process.env[key] ?? defaultValue;
  }

  /**
   * Читает числовую переменную окружения
   *
   * @param key - имя переменной окружения
   * @param defaultValue - значение по умолчанию
   * @returns значение переменной или дефолт
   */
  static number(key: string, defaultValue: number): number {
    const value = process.env[key];
    if (value === undefined) {
      return defaultValue;
    }
    const parsed = Number(value);
    return isNaN(parsed) ? defaultValue : parsed;
  }

  /**
   * Читает булеву переменную окружения
   * Поддерживает строки: "true", "1", "yes", "on" (case-insensitive)
   *
   * @param key - имя переменной окружения
   * @param defaultValue - значение по умолчанию
   * @returns значение переменной или дефолт
   */
  static boolean(key: string, defaultValue: boolean): boolean {
    const value = process.env[key];
    if (value === undefined) {
      return defaultValue;
    }
    const normalized = value.toLowerCase().trim();
    return (
      normalized === "true" ||
      normalized === "1" ||
      normalized === "yes" ||
      normalized === "on"
    );
  }

  /**
   * Читает переменную окружения с валидацией через функцию
   *
   * @param key - имя переменной окружения
   * @param defaultValue - значение по умолчанию
   * @param validator - функция валидации
   * @returns значение переменной или дефолт
   */
  static withValidation<T>(
    key: string,
    defaultValue: T,
    validator: (value: string) => T | null
  ): T {
    const value = process.env[key];
    if (value === undefined) {
      return defaultValue;
    }
    const validated = validator(value);
    return validated !== null ? validated : defaultValue;
  }
}

