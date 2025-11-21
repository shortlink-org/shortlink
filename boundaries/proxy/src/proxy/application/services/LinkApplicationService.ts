import { metrics, trace, context } from "@opentelemetry/api";
import type { Counter, Histogram, Meter } from "@opentelemetry/api";
import { Result, ok, err } from "neverthrow";
import { Hash } from "../../domain/entities/Hash.js";
import { Link } from "../../domain/entities/Link.js";
import { LinkNotFoundError } from "../../domain/exceptions/index.js";
import { LinkRedirectedEvent, LinkEvents } from "../../domain/events/index.js";
import { GetLinkByHashUseCase } from "../use-cases/GetLinkByHashUseCase.js";
import {
  PublishEventUseCase,
  PublishEventRequest,
} from "../use-cases/PublishEventUseCase.js";
import {
  GetLinkRequest,
  GetLinkResponse,
} from "../dto/index.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import {
  UseCasePipeline,
  LoggingInterceptor,
  MetricsInterceptor,
} from "../pipeline/index.js";

/**
 * Request DTO для handleRedirect
 * Application DTO - чистый, без зависимостей от Express
 */
export interface HandleRedirectRequest {
  hash: Hash;
}

/**
 * Response DTO для handleRedirect
 */
export interface HandleRedirectResponse {
  link: Link;
}

/**
 * Application Service для оркестрации Use Cases работы со ссылками
 * Координирует выполнение нескольких Use Cases для сложных бизнес-операций
 */
export class LinkApplicationService {
  private readonly pipeline: UseCasePipeline;
  private readonly interceptors: Array<
    LoggingInterceptor | MetricsInterceptor
  >;
  private readonly meter: Meter;
  private readonly redirectCounter: Counter;
  private readonly redirectLatency: Histogram;

  constructor(
    private readonly getLinkByHashUseCase: GetLinkByHashUseCase,
    private readonly publishEventUseCase: PublishEventUseCase,
    private readonly logger: ILogger,
    private readonly useCasePipeline: UseCasePipeline,
    private readonly loggingInterceptor: LoggingInterceptor,
    private readonly metricsInterceptor: MetricsInterceptor
  ) {
    this.pipeline = useCasePipeline;
    this.interceptors = [loggingInterceptor, metricsInterceptor];
    this.meter = metrics.getMeter("proxy-service", "1.0.0");
    this.redirectCounter = this.meter.createCounter("redirects_total", {
      description: "Total number of redirect attempts",
      unit: "1",
    });
    this.redirectLatency = this.meter.createHistogram(
      "redirect_latency_seconds",
      {
        description: "Redirect handling latency in seconds",
        unit: "s",
      }
    );
  }

  /**
   * Получает ссылку по хешу
   * Простой фасад для GetLinkByHashUseCase
   */
  getByHash(hash: string): Promise<GetLinkResponse> {
    return this.getLinkByHashUseCase.execute({ hash });
  }

  /**
   * Обрабатывает редирект по ссылке
   * Оркестрирует несколько Use Cases:
   * 1. Получает ссылку по хешу
   * 2. Публикует событие редиректа
   * 
   * Статистика собирается через eBPF, не требует записи в БД
   *
   * @param request - запрос на редирект
   * @returns Result с ссылкой или ошибкой
   */
  async handleRedirect(
    request: HandleRedirectRequest
  ): Promise<Result<HandleRedirectResponse, LinkNotFoundError>> {
    const activeSpan = trace.getSpan(context.active());
    if (activeSpan) {
      activeSpan.setAttribute("url.hash", request.hash.value);
    }

    const startedAt = process.hrtime.bigint();
    let outcome: "success" | "not_found" | "error" = "error";
    let linkFound = false;
    let eventPublished = false;

    try {
      // 1. Получаем ссылку по хешу
      const linkResponse = await this.getByHash(request.hash.value);
      const { link } = linkResponse;
      linkFound = true;

      // 2. Публикуем событие редиректа (асинхронно, не блокируем основной поток)
      const event: LinkRedirectedEvent = LinkEvents.redirected(request.hash, link);
      try {
        await this.pipeline.execute(
          this.publishEventUseCase,
          { event },
          this.interceptors
        );
        eventPublished = true;
      } catch (error) {
        // Если публикация события не удалась, логируем, но не прерываем выполнение
        this.logger.warn("Failed to publish redirect event", {
          hash: request.hash.value,
          error,
        });
      }

      outcome = "success";
      return ok({ link });
    } catch (error) {
      // Обрабатываем ошибки (например, LinkNotFoundError)
      if (error instanceof LinkNotFoundError) {
        outcome = "not_found";
        return err(error);
      }
      // Неожиданные ошибки пробрасываем дальше
      throw error;
    } finally {
      const durationNs = process.hrtime.bigint() - startedAt;
      const durationSeconds = Number(durationNs) / 1_000_000_000;

      this.redirectCounter.add(1, { outcome });
      this.redirectLatency.record(durationSeconds, { outcome });

      if (activeSpan) {
        activeSpan.setAttribute("link.found", linkFound);
        activeSpan.setAttribute("mq.event_published", eventPublished);
        activeSpan.setAttribute("url.hash", request.hash.value);
      }
    }
  }
}

