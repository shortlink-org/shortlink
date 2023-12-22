# 6. Istio

Date: 2023-12-22

## Status

Deprecated

## Context

We long time used Istio as a service mesh. But now we use Cilium as CNI and it has its own service mesh. 
So we don't need Istio anymore.

## Decision

We will remove Istio from the cluster and codebase. But we write this ADR for save history and for future reference.

## Consequences

We write a cookbook for Istio. We will remove it. But we will keep it in the history.

## Cookbook

### Recipe 1: Configuring kind: Sidecar in Istio

**kind: Sidecar** in Istio is a crucial configuration for optimizing the behavior of sidecar proxies within your service mesh. 
It's primarily used to define the accessibility of services to each sidecar, enhancing security and performance.

```yaml
apiVersion: networking.istio.io/v1beta1
kind: Sidecar
metadata:
  name: default
spec:
  egress:
    - hosts:
      - "./*"                   # current namespace
      - "istio-system/*"        # istio-system services
      - "istio-ingress/*"       # istio-ingress services
      - "prometheus-operator/*" # prometheus-operator services
      - "kube-system/*"         # kube-system services
```

> [!TIP]
> 
> **Props**:
> 
>   - **Security**: Restricts sidecar access to essential services, minimizing risk.
>   - **Performance**: Reduces sidecar workload, improving efficiency.
>
> **Cons**:
> 
>   - **Complexity**: Demands precise understanding of service interactions.
>   - **Maintenance**: Requires ongoing updates to align with service changes.

### Recipe 2: Configuring Telemetry in Istio

**Telemetry** in Istio is a crucial configuration for monitoring the behavior of your service mesh.

```yaml
apiVersion: telemetry.istio.io/v1alpha1
kind: Telemetry
metadata:
  name: default
spec:
  # no selector specified, applies to all workloads
  accessLogging:
    - providers:
        - name: envoy
      # By default, this turns on access logging (no need to set `disabled: false`).
      # Unspecified `disabled` will be treated as `disabled: false`, except in
      # cases where a parent configuration has marked as `disabled: true`. In
      # those cases, `disabled: false` must be set explicitly to override.
```
