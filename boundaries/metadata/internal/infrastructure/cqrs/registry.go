package cqrs

import (
	"fmt"

	linkpb "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/domain/link/v1"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/shortlink-org/go-sdk/cqrs/bus"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"

	metadatapb "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
)

// NewEventRegistry creates and configures a CQRS registry for metadata service
// Registers events from Link Service that metadata service consumes
func NewEventRegistry() (*bus.TypeRegistry, error) {
	registry := bus.NewTypeRegistry()

	// Register Link Service events that metadata service consumes
	// LinkCreated - metadata service processes new links to extract metadata
	err := registry.RegisterEvent(&linkpb.LinkCreated{})
	if err != nil {
		return nil, fmt.Errorf("register link created event: %w", err)
	}

	// Register metadata service's own events
	// MetadataExtracted - published when metadata is successfully extracted from a URL
	err = registry.RegisterEvent(&metadatapb.MetadataExtracted{})
	if err != nil {
		return nil, fmt.Errorf("register metadata extracted event: %w", err)
	}

	return registry, nil
}

// NewShortlinkNamer creates a singleton namer for "metadata" service.
func NewShortlinkNamer() cqrsmessage.Namer { //nolint:ireturn // CQRS components expect the interface type
	return cqrsmessage.NewShortlinkNamer("metadata")
}

// NewProtoMarshaler creates ProtoMarshaler using the provided namer
// The namer must be the same instance used by EventBus and CommandBus
func NewProtoMarshaler(namer cqrsmessage.Namer) *cqrsmessage.ProtoMarshaler {
	return cqrsmessage.NewProtoMarshaler(namer)
}

// NewEventBus creates EventBus with publisher, marshaler and namer
// All three components must use the same namer instance for consistency
func NewEventBus(
	publisher message.Publisher,
	marshaler cqrsmessage.Marshaler,
	namer cqrsmessage.Namer,
) (*bus.EventBus, error) {
	// Create EventBus
	eventBus := bus.NewEventBus(publisher, marshaler, namer)

	return eventBus, nil
}

// NewCommandBus creates CommandBus with publisher, marshaler and namer
// All three components must use the same namer instance for consistency
func NewCommandBus(
	publisher message.Publisher,
	marshaler cqrsmessage.Marshaler,
	namer cqrsmessage.Namer,
) (*bus.CommandBus, error) {
	// Create CommandBus
	commandBus := bus.NewCommandBus(publisher, marshaler, namer)

	return commandBus, nil
}
