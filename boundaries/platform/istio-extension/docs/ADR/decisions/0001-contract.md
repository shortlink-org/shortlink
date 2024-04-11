# 1. Contract for Istio extension.

Date: 2024-04-11

## Status

Accepted

## Context

We use Istio CNI for service mesh. And we want to write plugins for Istio that are written in Go and compiled to WebAssembly
for extending the functionality of the service mesh.

## Decision

For successful integration of the plugin into the service mesh, the plugin must implement the following contract:

1. Add header to response `injected-by: <plugin-name>` if we change the response.

## Consequences

We can easily integrate the plugin into the service mesh.
