package v1

// Metadata CRUD method IDs
// These are kept for backward compatibility but are no longer used
// Events are now handled via CQRS (go-sdk/cqrs)
var (
	METHOD_ADD    uint32 = 1
	METHOD_GET    uint32 = 2
	METHOD_LIST   uint32 = 3
	METHOD_UPDATE uint32 = 4
	METHOD_DELETE uint32 = 5
)

const (
	MQ_EVENT_CQRS_NEW = "shortlink.metadata.cqrs.new"
)
