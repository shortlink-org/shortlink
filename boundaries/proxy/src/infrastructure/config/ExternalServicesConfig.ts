import { ConfigReader } from "./ConfigReader.js";

/**
 * Конфигурация внешних сервисов
 * Децентрализованный подход - модуль сам определяет свои дефолты
 */
export class ExternalServicesConfig {
  /**
   * URL Link Service (gRPC)
   */
  readonly linkServiceGrpcUrl: string;

  /**
   * Таймаут запросов к внешним сервисам в миллисекундах
   */
  readonly requestTimeout: number;

  /**
   * Количество повторных попыток при ошибке
   */
  readonly retryCount: number;

  constructor() {
    this.linkServiceGrpcUrl = ConfigReader.string(
      "LINK_SERVICE_GRPC_URL",
      "link-service:50051"
    );
    this.requestTimeout = ConfigReader.number("EXTERNAL_SERVICE_TIMEOUT", 5000);
    this.retryCount = ConfigReader.number("EXTERNAL_SERVICE_RETRY_COUNT", 3);
  }
}

