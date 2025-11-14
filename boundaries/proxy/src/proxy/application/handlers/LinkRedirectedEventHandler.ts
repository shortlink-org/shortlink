import { injectable } from "inversify";
import { DomainEvent, LinkRedirectedEvent } from "../../domain/events/index.js";
import { IEventHandler } from "./IEventHandler.js";

/**
 * Обработчик события редиректа ссылки
 * Реализует Notification Pattern - реакция на доменное событие
 * Выполняется асинхронно после публикации события
 * 
 * Статистика собирается через eBPF, поэтому обработчик может быть пустым
 * или использоваться для других реакций на событие редиректа
 */
@injectable()
export class LinkRedirectedEventHandler
  implements IEventHandler<LinkRedirectedEvent>
{
  canHandle(event: DomainEvent): boolean {
    return event.type === "LinkRedirected";
  }

  async handle(event: LinkRedirectedEvent): Promise<void> {
    // Статистика собирается через eBPF, не требует обработки здесь
    // Этот обработчик может быть использован для других реакций на событие редиректа
    // Например, отправка уведомлений, обновление кэша и т.д.
  }
}

