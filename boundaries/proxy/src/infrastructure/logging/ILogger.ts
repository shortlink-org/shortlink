/**
 * Интерфейс логгера для абстракции от конкретной реализации
 * Позволяет легко заменить Winston на другую библиотеку логирования
 */
export interface ILogger {
  /**
   * Логирует информационное сообщение
   */
  info(message: string, meta?: any): void;

  /**
   * Логирует предупреждение
   */
  warn(message: string, meta?: any): void;

  /**
   * Логирует ошибку
   */
  error(message: string, error?: any, meta?: any): void;

  /**
   * Логирует отладочное сообщение
   */
  debug(message: string, meta?: any): void;

  /**
   * Логирует HTTP запрос (специальный метод для HTTP логирования)
   */
  http(message: string): void;
}

