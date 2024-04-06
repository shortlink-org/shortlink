### ArgoCD configurations

This directory contains the ArgoCD configurations for the cluster.

#### Structure directory

The directory is structured as follows:

```bash
.
├── draft              # Draft configurations
├── infrastructure     # ArgoCD applications for infrastructure components
├── knative            # ArgoCD applications for knative components
├── kustomize          # Kustomize patches for ArgoCD applications
├── shortlink-link     # ArgoCD applications for link boundary components
├── shortlink-billing  # ArgoCD applications for billing boundary components
├── shortlink-shop     # ArgoCD applications for shop boundary components
└── shortlink          # ArgoCD applications for common shortlink components
```
