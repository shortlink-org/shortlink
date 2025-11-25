package cqrs

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/shortlink-org/go-sdk/cqrs/bus"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"
	linkpb "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// NewEventRegistry creates and configures a CQRS registry for link service
// Registers all events and commands following ADR-0002 canonical naming
func NewEventRegistry() (*bus.TypeRegistry, error) {
	registry := bus.NewTypeRegistry()

	// Register events
	if err := registry.RegisterEvent(&linkpb.LinkCreated{}); err != nil {
		return nil, err
	}
	if err := registry.RegisterEvent(&linkpb.LinkUpdated{}); err != nil {
		return nil, err
	}
	if err := registry.RegisterEvent(&linkpb.LinkDeleted{}); err != nil {
		return nil, err
	}

	// Register commands
	if err := registry.RegisterCommand(&linkpb.CreateLink{}); err != nil {
		return nil, err
	}
	if err := registry.RegisterCommand(&linkpb.UpdateLink{}); err != nil {
		return nil, err
	}
	if err := registry.RegisterCommand(&linkpb.DeleteLink{}); err != nil {
		return nil, err
	}

	return registry, nil
}

// NewShortlinkNamer creates a singleton namer for "link" service
// This namer is used consistently across EventBus, CommandBus, and ProtoMarshaler
// to ensure stable canonical naming (e.g., link.link.created.v1)
func NewShortlinkNamer() cqrsmessage.Namer {
	return cqrsmessage.NewShortlinkNamer("link")
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

