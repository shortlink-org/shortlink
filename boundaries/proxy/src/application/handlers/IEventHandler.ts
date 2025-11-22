import { DomainEvent } from "../../domain/events/index.js";

/**
 * Базовый интерфейс для обработчиков доменных событий
 * Реализует Notification Pattern для асинхронной реакции на события
 */
export interface IEventHandler<TEvent extends DomainEvent> {
  /**
   * Проверяет, может ли обработчик обработать данное событие
   *
   * @param event - доменное событие
   * @returns true если обработчик может обработать событие
   */
  canHandle(event: DomainEvent): boolean;

  /**
   * Обрабатывает доменное событие
   *
   * @param event - доменное событие для обработки
   * @returns Promise, который разрешается после обработки события
   */
  handle(event: TEvent): Promise<void>;
}

