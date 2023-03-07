# 16. Lifecycle deploy

Date: 2023-02-10

## Status

Accepted

## Context

We need to decide how we will update images in ArgoCD.

## Decision

We use [argocd-image-updater](https://github.com/argoproj-labs/argocd-image-updater#argo-cd-image-updater) for update images in ArgoCD.

## Consequences

![argocd-image-updater](./images/ADR-0016/lifecycle-deploy.png)

+ We can update images in ArgoCD use [argocd-image-updater](https://github.com/argoproj-labs/argocd-image-updater#argo-cd-image-updater).
