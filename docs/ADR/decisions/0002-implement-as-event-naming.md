# 2. Implement as event naming

Date: 2022-08-24

## Status

Accepted

## Context

We use event-driven architecture is a software architecture paradigm promoting 
the production, detection, consumption of, and reaction to event.

- Need set Event naming

## Decision

We will stick to the following format of naming system events: **shortlink.link.event.new**
  
Where:
- **shortlink** is the name of the all system
- **link** is the domain name of the system
- **event** event trait
- **new** is the name of the event

Example:

```go
// MQ events
const (
	// MQ_EVENT_LINK_NEW - subscribe on request created link
	MQ_EVENT_LINK_NEW = "shortlink.link.event.new"

	// MQ_EVENT_LINK_CREATED - subscribe on created link
	MQ_EVENT_LINK_CREATED = "shortlink.link.event.created"
)
```

## Consequences

We get a single naming convention for events in the system.
