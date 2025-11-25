// Legacy event name (kept for backward compatibility)
export const MQ_EVENT_LINK_NEW = "shortlink.link.event.new";

/**
 * Canonical event names following ADR-0002 format: {service}.{aggregate}.{event}.{version}
 * For Link Service events: link.link.{event}.v1
 */
export const LinkEventTopics = {
  /** Link redirected event - canonical name: link.link.redirected.v1 */
  REDIRECTED: "link.link.redirected.v1",
  /** Link created event - canonical name: link.link.created.v1 */
  CREATED: "link.link.created.v1",
  /** Link updated event - canonical name: link.link.updated.v1 */
  UPDATED: "link.link.updated.v1",
  /** Link deleted event - canonical name: link.link.deleted.v1 */
  DELETED: "link.link.deleted.v1",
} as const;

/**
 * Type-safe mapping from event type to canonical topic name
 */
export const EventTypeToTopic: Record<string, string> = {
  LinkRedirected: LinkEventTopics.REDIRECTED,
  LinkCreated: LinkEventTopics.CREATED,
  LinkUpdated: LinkEventTopics.UPDATED,
  LinkDeleted: LinkEventTopics.DELETED,
} as const;
