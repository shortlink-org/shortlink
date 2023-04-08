# 21. microservice structure

Date: 2023-04-08

## Status

Accepted

## Context

We want to have a clear structure for our microservices.

## Decision

We will use the following structure:

```
├── ops/dockerfile/{{serviceName}}.Dockerfile
├── ops/docker-compose/application/{{serviceName}}/{{serviceName}}.yml
├── ops/Helm/{{serviceName}}/Chart.yaml
├── ops/gitlab/workflows/matrix_build_base.yml
├── ops/gitlab/workflows/matrix_build_helm.yml
├── ops/argocd/shortlink/{{serviceName}}/application.yaml
├── docs/c4/containers/services
├── internal/services/{{serviceName}}
│   ├── application/
│   ├── cmd/
│   ├── di/
│   ├── domain/
│   ├── infrastructure/
│   ├── tests/
│   └── README.md
└── README.md
```

### Project README

The project README should contain the following sections:

- Project description
- C4 container diagram
- C4 component diagram
- C4 use case diagram

## Consequences

We have a clear structure for our microservices.
