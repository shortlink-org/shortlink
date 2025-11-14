import { DomainEvent } from "../../domain/events/index.js";
import { IEventHandler } from "./IEventHandler.js";

/**
 * Диспетчер событий для реализации Notification Pattern
 * Управляет reaction chain - цепочкой обработчиков для одного события
 */
export class EventDispatcher {
  private handlers: Map<string, IEventHandler<DomainEvent>[]> = new Map();

  /**
   * Регистрирует обработчик для определенного типа события
   *
   * @param eventType - тип события
   * @param handler - обработчик события
   */
  register<TEvent extends DomainEvent>(
    eventType: string,
    handler: IEventHandler<TEvent>
  ): void {
    if (!this.handlers.has(eventType)) {
      this.handlers.set(eventType, []);
    }
    this.handlers.get(eventType)!.push(handler as IEventHandler<DomainEvent>);
  }

  /**
   * Диспетчеризует событие всем зарегистрированным обработчикам
   * Выполняет все обработчики асинхронно (reaction chain)
   *
   * @param event - доменное событие для диспетчеризации
   */
  async dispatch<TEvent extends DomainEvent>(
    event: TEvent & { type: string }
  ): Promise<void> {
    const handlers = this.handlers.get(event.type) || [];

    // Выполняем все обработчики асинхронно (reaction chain)
    await Promise.all(
      handlers
        .filter((h) => h.canHandle(event))
        .map((h) => h.handle(event))
    );
  }
}

