#!/bin/sh

# Clean up the ArgoCD resources
kubectl -n argocd delete deploy -l app.kubernetes.io/instance=argocd
kubectl -n argocd delete svc -l app.kubernetes.io/instance=argocd
kubectl -n argocd delete statefulset -l app.kubernetes.io/instance=argocd
