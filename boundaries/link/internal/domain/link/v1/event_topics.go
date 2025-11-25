package v1

// Link event topic names following ADR-0002 canonical naming format:
// {service}.{aggregate}.{event}.{version}
// For Link Service: link.link.{event}.v1

const (
	// LinkCreatedTopic is the canonical topic name for LinkCreated events
	// Format: link.link.created.v1
	LinkCreatedTopic = "link.link.created.v1"

	// LinkUpdatedTopic is the canonical topic name for LinkUpdated events
	// Format: link.link.updated.v1
	LinkUpdatedTopic = "link.link.updated.v1"

	// LinkDeletedTopic is the canonical topic name for LinkDeleted events
	// Format: link.link.deleted.v1
	LinkDeletedTopic = "link.link.deleted.v1"
)

