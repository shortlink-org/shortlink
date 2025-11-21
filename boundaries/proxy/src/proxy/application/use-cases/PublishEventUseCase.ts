import { Result, ok } from "neverthrow";
import { DomainEvent } from "../../domain/events/index.js";
import { IUseCase } from "./IUseCase.js";

/**
 * Request DTO для PublishEventUseCase
 * Application DTO - чистый, без зависимостей от Express
 */
export interface PublishEventRequest {
  event: DomainEvent;
}

/**
 * Response DTO для PublishEventUseCase
 */
export interface PublishEventResponse {
  success: boolean;
}

/**
 * Интерфейс для публикации событий
 * Будет реализован в Infrastructure Layer (AMQP, Event Bus)
 */
export interface IEventPublisher {
  publish(event: DomainEvent): Promise<void>;
}

/**
 * Use Case для публикации доменных событий
 * Публикует события в message bus (AMQP)
 */
export class PublishEventUseCase
  implements IUseCase<PublishEventRequest, Result<PublishEventResponse, never>>
{
  constructor(
    private readonly eventPublisher: IEventPublisher
  ) {}

  async execute(
    request: PublishEventRequest
  ): Promise<Result<PublishEventResponse, never>> {
    await this.eventPublisher.publish(request.event);

    return ok({ success: true });
  }
}

