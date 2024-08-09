# 32. ArgoCD tips [Cookbook]

Date: 2024-08-09

## Status

Accepted

## Cookbook

### Recipe 1: Troubleshooting ArgoCD Applications

```
kubectl -n argocd patch app temporal -p '{"metadata": {"finalizers": null}}' --type merge
```
