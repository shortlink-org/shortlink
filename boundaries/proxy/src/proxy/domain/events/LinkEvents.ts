import { Link } from "../entities/Link.js";
import { Hash } from "../entities/Hash.js";

/**
 * Базовый интерфейс для доменных событий
 */
export interface DomainEvent {
  readonly type: string;
  readonly occurredAt: Date;
}

/**
 * Событие создания новой ссылки
 */
export interface LinkCreatedEvent extends DomainEvent {
  readonly type: "LinkCreated";
  readonly link: Link;
  readonly hash: Hash;
}

/**
 * Событие редиректа по ссылке
 */
export interface LinkRedirectedEvent extends DomainEvent {
  readonly type: "LinkRedirected";
  readonly hash: Hash;
  readonly link: Link;
  readonly timestamp: Date;
}

/**
 * Событие обновления ссылки
 */
export interface LinkUpdatedEvent extends DomainEvent {
  readonly type: "LinkUpdated";
  readonly hash: Hash;
  readonly oldUrl: string;
  readonly newUrl: string;
}

/**
 * Событие удаления ссылки
 */
export interface LinkDeletedEvent extends DomainEvent {
  readonly type: "LinkDeleted";
  readonly hash: Hash;
}

/**
 * Тип объединения всех событий ссылок
 */
export type LinkEvent =
  | LinkCreatedEvent
  | LinkRedirectedEvent
  | LinkUpdatedEvent
  | LinkDeletedEvent;

/**
 * Фабричные функции для создания событий
 */
export const LinkEvents = {
  created(link: Link): LinkCreatedEvent {
    return {
      type: "LinkCreated",
      occurredAt: new Date(),
      link,
      hash: link.hash,
    };
  },

  redirected(hash: Hash, link: Link): LinkRedirectedEvent {
    return {
      type: "LinkRedirected",
      occurredAt: new Date(),
      hash,
      link,
      timestamp: new Date(),
    };
  },

  updated(hash: Hash, oldUrl: string, newUrl: string): LinkUpdatedEvent {
    return {
      type: "LinkUpdated",
      occurredAt: new Date(),
      hash,
      oldUrl,
      newUrl,
    };
  },

  deleted(hash: Hash): LinkDeletedEvent {
    return {
      type: "LinkDeleted",
      occurredAt: new Date(),
      hash,
    };
  },
};

