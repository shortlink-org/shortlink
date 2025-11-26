package domain

// Link event topic names that metadata service consumes
// Following ADR-0002 canonical naming format: {service}.{aggregate}.{event}.{version}

const (
	// LinkCreatedTopic is the canonical topic name for LinkCreated events from Link Service
	// Format: link.link.created.v1
	// Metadata service subscribes to this topic to extract metadata for new links
	LinkCreatedTopic = "link.link.created.v1"

	// MetadataExtractedTopic is the canonical topic name for MetadataExtracted events
	// Format: metadata.metadata.extracted.v1
	// Published when metadata is successfully extracted from a URL
	MetadataExtractedTopic = "metadata.metadata.extracted.v1"
)
