import { ConfigReader } from "./ConfigReader.js";

/**
 * Конфигурация приложения
 * Децентрализованный подход - модуль сам определяет свои дефолты
 */
export class AppConfig {
  /**
   * Порт для HTTP сервера
   */
  readonly port: number;

  /**
   * Окружение (development, production, test)
   */
  readonly environment: string;

  /**
   * Имя сервиса
   */
  readonly serviceName: string;

  constructor() {
    this.port = ConfigReader.number("PORT", 3020);
    this.environment = ConfigReader.string("NODE_ENV", "development");
    this.serviceName = ConfigReader.string("SERVICE_NAME", "proxy-service");
  }
}

