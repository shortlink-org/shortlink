# 2. Use Dapr

Date: 2023-07-03

## Status

Accepted

## Context

As we plan to potentially shift to a microservices architecture for our ShortLink Merch Service, 
we need a platform that simplifies the development and management of our services, 
and facilitates communication between them.

## Decision

We propose to adopt Dapr (Distributed Application Runtime), an event-driven, portable runtime that simplifies 
building microservices. It provides a range of capabilities like state management, service-to-service invocation, 
and pub/sub messaging via a consistent, language-agnostic API, which will allow us to build resilient, scalable services.

## Consequences

Using Dapr will simplify the development of our microservices, enabling us to focus more on business logic rather 
than infrastructure concerns. It will also provide us with a scalable, resilient foundation for our services.
