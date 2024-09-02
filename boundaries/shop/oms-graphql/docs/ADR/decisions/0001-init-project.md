# 1. Init project

Date: 2024-09-02

## Status

Accepted

## Context

I am working on an OMS GraphQL API Bridge to translate my gRPC API into a GraphQL interface for public API access. 
This is necessary to expose the gRPC functionality to a broader audience that utilizes GraphQL.

## Decision

We will implement the OMS GraphQL API Bridge using `https://tailcall.run/`. 
This tool was chosen for its simplicity, robust feature set, and ease of use in translating gRPC APIs to GraphQL. 
Additionally, Tailcall is written in Rust, ensuring high performance and safety. 
It is also actively developed, which provides confidence in its future support and improvements.

### Alternatives

#### The Guild's GraphQL Mesh

Another alternative considered was `https://the-guild.dev/graphql/mesh/docs/handlers/grpc#use-reflection-instead-of-proto-files`, 
which offers capabilities for integrating gRPC with GraphQL, including the use of reflection instead of proto files. 
However, this tool was not selected due to a known issue (`Invalid value used as weak map key`) that has not yet been resolved, 
making the product unreliable for our use case.

## Consequences

This decision simplifies the process of exposing our gRPC API as a public GraphQL API, 
making it easier to interact with the service. However, this introduces a dependency on `tailcall.run`, 
which will need to be monitored for any updates or changes that might affect our integration. 
