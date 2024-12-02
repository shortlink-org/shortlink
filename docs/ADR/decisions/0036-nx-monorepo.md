# 36. Nx, monorepo

Date: 2023-04-27

## Status

Deprecated

## Context

Our project consists of multiple applications and libraries, which are interdependent and need to be managed efficiently. 
Monorepos allow for easier maintenance, cross-package code sharing, and improved developer productivity.

## Decision

We will set up a monorepo using Nx to streamline development and manage our applications and libraries more efficiently.

### Alternatives

- Use a multi-repo approach
- Use turborepo
- Use NX

#### Turborepo

Turborepo is a tool that allows you to manage multiple repos as a single repo. 
It is similar to Nx, but it is not as mature and does not have as many features.

#### Benchmarking Nx, Turbo, and Lerna

> [link to benchmarking repo](https://github.com/vsavkin/large-monorepo#benchmarking-nx-turbo-and-lerna)

![Benchmarking Nx, Turbo, and Lerna](https://raw.githubusercontent.com/vsavkin/large-monorepo/main/readme-assets/turbo-nx-perf.gif)

## Consequences

By initializing an Nx monorepo, we can expect the following benefits:

- Improved developer productivity
- Easier maintenance
- Cross-package code sharing
- Improved CI/CD

However, there are some potential drawbacks:

- Increased complexity for developers unfamiliar with monorepos or Nx.

## Deprecated

This ADR is deprecated because we decided not to use Nx for our project.
We grouped our applications and libraries by bounded contexts.

We have problems with Nx:
- Nx is bad controlled dependencies between packages
- All packages are in one repository, which makes it difficult to manage
