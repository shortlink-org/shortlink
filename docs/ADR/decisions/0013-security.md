# 13. Security

Date: 2023-01-16

## Status

Accepted

## Context

We need to encrypt secret data for deployment.

## Decision

We will use [SOPS](https://github.com/mozilla/sops) to encrypt secret data.

- for **HELM** we will use [helm-secrets](https://github.com/jkroepke/helm-secrets/wiki/Usage) plugin.
- for **ArgoCD** we use - [Argo CD Integration](https://github.com/jkroepke/helm-secrets/blob/main/docs/ArgoCD%20Integration.md).

## Consequences

We will need to use SOPS to decrypt secret data before deploying.

