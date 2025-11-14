import { IUseCase } from "../use-cases/IUseCase.js";
import { Result } from "neverthrow";

/**
 * Контекст выполнения Use Case для интерцепторов
 * Содержит информацию о запросе, ответе и метаданные
 */
export interface UseCaseExecutionContext<TRequest, TResponse> {
  /**
   * Название Use Case
   */
  useCaseName: string;

  /**
   * Входные данные запроса
   */
  request: TRequest;

  /**
   * Результат выполнения Use Case (заполняется после выполнения)
   */
  response?: Result<TResponse, any>;

  /**
   * Ошибка выполнения (если произошла)
   */
  error?: Error;

  /**
   * Время начала выполнения
   */
  startTime: number;

  /**
   * Время окончания выполнения (заполняется после выполнения)
   */
  endTime?: number;

  /**
   * Длительность выполнения в миллисекундах
   */
  duration?: number;

  /**
   * Метаданные для передачи между интерцепторами
   */
  metadata: Map<string, unknown>;
}

/**
 * Интерфейс для интерцепторов Use Case Pipeline
 * Позволяет добавлять cross-cutting concerns без изменения Use Cases
 */
export interface IUseCaseInterceptor<TRequest = any, TResponse = any> {
  /**
   * Вызывается перед выполнением Use Case
   * Можно модифицировать запрос или выполнить предварительные действия
   *
   * @param context - Контекст выполнения Use Case
   * @returns Promise с модифицированным запросом или исходным запросом
   */
  before(
    context: UseCaseExecutionContext<TRequest, TResponse>
  ): Promise<TRequest> | TRequest;

  /**
   * Вызывается после успешного выполнения Use Case
   * Можно обработать результат или выполнить пост-обработку
   *
   * @param context - Контекст выполнения Use Case с результатом
   * @returns Promise<void>
   */
  after(
    context: UseCaseExecutionContext<TRequest, TResponse>
  ): Promise<void> | void;

  /**
   * Вызывается при ошибке выполнения Use Case
   * Можно обработать ошибку или выполнить cleanup
   *
   * @param context - Контекст выполнения Use Case с ошибкой
   * @returns Promise<void>
   */
  onError(
    context: UseCaseExecutionContext<TRequest, TResponse>
  ): Promise<void> | void;

  /**
   * Вызывается всегда после выполнения Use Case (success или error)
   * Используется для cleanup и финальных действий
   *
   * @param context - Контекст выполнения Use Case
   * @returns Promise<void>
   */
  finally?(
    context: UseCaseExecutionContext<TRequest, TResponse>
  ): Promise<void> | void;
}

