# Gateway API HTTPRoute Migration Guide

This guide explains how to migrate from Kubernetes Ingress to Gateway API HTTPRoute for ShortLink services.

## Overview

The ShortLink project now supports both traditional Kubernetes Ingress and the newer Gateway API HTTPRoute resources. HTTPRoute provides better routing capabilities, security features, and is the future direction for ingress in Kubernetes.

## Services with HTTPRoute Support

The following services have HTTPRoute templates available:

- **proxy** - Handles URL shortlink redirection (`/s/*`)
- **ui** - Serves the Next.js frontend (`/next/*`)
- **bff** - Backend for Frontend API (`/api/*`)

## Prerequisites

Before migrating to HTTPRoute, ensure:

1. Gateway API CRDs are installed in your cluster:
   ```bash
   kubectl get crd gateways.gateway.networking.k8s.io
   kubectl get crd httproutes.gateway.networking.k8s.io
   ```

2. A Gateway resource is deployed. Example gateways are available in:
   - `boundaries/common/ops/common/draft/gateway/external-gateway.yaml`
   - `boundaries/common/ops/common/draft/gateway/internal-gateway.yaml`

## Configuration

### Default Configuration (Ingress)

By default, all services use Kubernetes Ingress:

```yaml
# In values.yaml
httpRoute:
  enabled: false  # HTTPRoute disabled by default

ingress:
  enabled: true   # Ingress enabled by default
  ingressClassName: nginx
  # ... other ingress config
```

### Migrating to HTTPRoute

To enable HTTPRoute for a service, update your values.yaml or use Helm set flags:

#### Option 1: Update values.yaml

```yaml
# Enable HTTPRoute
httpRoute:
  enabled: true
  annotations: {}
  parentRefs:
    - name: external-gateway
      namespace: shortlink-common
  hostnames:
    - shortlink.best
  rules:
    - matches:
        - path:
            type: PathPrefix
            value: /s  # For proxy service
      backendRefs:
        - name: shortlink-link-proxy
          port: 3020

# Disable Ingress (optional, both can coexist during migration)
ingress:
  enabled: false
```

#### Option 2: Use Helm CLI

```bash
# Enable HTTPRoute and disable Ingress
helm upgrade proxy boundaries/proxy/ops/proxy \
  --set httpRoute.enabled=true \
  --set ingress.enabled=false
```

## Service-Specific Examples

### Proxy Service

```yaml
httpRoute:
  enabled: true
  parentRefs:
    - name: external-gateway
      namespace: shortlink-common
  hostnames:
    - shortlink.best
  rules:
    - matches:
        - path:
            type: PathPrefix
            value: /s
      backendRefs:
        - name: shortlink-link-proxy
          port: 3020
```

### UI Service

```yaml
httpRoute:
  enabled: true
  parentRefs:
    - name: external-gateway
      namespace: shortlink-common
  hostnames:
    - shortlink.best
  rules:
    - matches:
        - path:
            type: PathPrefix
            value: /next
      backendRefs:
        - name: shortlink-link-ui
          port: 8080
```

### BFF Service

```yaml
httpRoute:
  enabled: true
  parentRefs:
    - name: external-gateway
      namespace: shortlink-common
  hostnames:
    - shortlink.best
  rules:
    - matches:
        - path:
            type: PathPrefix
            value: /api
      backendRefs:
        - name: shortlink-link-bff
          port: 7070
```

## Advanced HTTPRoute Features

HTTPRoute provides advanced routing capabilities beyond traditional Ingress:

### Path Matching Types

```yaml
rules:
  - matches:
      # Exact path match
      - path:
          type: Exact
          value: /api/v1
      # Path prefix match (default)
      - path:
          type: PathPrefix
          value: /api
      # Regular expression match
      - path:
          type: RegularExpression
          value: /api/v[0-9]+
```

### Header-Based Routing

```yaml
rules:
  - matches:
      - path:
          type: PathPrefix
          value: /api
        headers:
          - name: X-Version
            value: v2
    backendRefs:
      - name: api-v2
        port: 8080
```

### Query Parameter Routing

```yaml
rules:
  - matches:
      - path:
          type: PathPrefix
          value: /api
        queryParams:
          - name: version
            value: beta
    backendRefs:
      - name: api-beta
        port: 8080
```

### Traffic Splitting

```yaml
rules:
  - matches:
      - path:
          type: PathPrefix
          value: /api
    backendRefs:
      - name: api-stable
        port: 8080
        weight: 90
      - name: api-canary
        port: 8080
        weight: 10
```

### Request/Response Filters

```yaml
rules:
  - matches:
      - path:
          type: PathPrefix
          value: /api
    filters:
      - type: RequestHeaderModifier
        requestHeaderModifier:
          add:
            - name: X-Custom-Header
              value: custom-value
      - type: RequestRedirect
        requestRedirect:
          scheme: https
          statusCode: 301
    backendRefs:
      - name: api-service
        port: 8080
```

## Migration Strategy

### Phased Approach

1. **Test Phase**: Enable HTTPRoute alongside Ingress
   ```yaml
   httpRoute:
     enabled: true
   ingress:
     enabled: true  # Keep both running
   ```

2. **Validation Phase**: Verify HTTPRoute works correctly
   ```bash
   kubectl get httproute -n your-namespace
   kubectl describe httproute your-service
   ```

3. **Cutover Phase**: Disable Ingress once HTTPRoute is validated
   ```yaml
   httpRoute:
     enabled: true
   ingress:
     enabled: false
   ```

### Rollback Plan

If issues occur with HTTPRoute, simply re-enable Ingress:

```bash
helm upgrade your-service path/to/chart \
  --set httpRoute.enabled=false \
  --set ingress.enabled=true
```

## Gateway Configuration

### External Gateway Example

For public-facing services (proxy, ui, bff):

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: external-gateway
  namespace: shortlink-common
spec:
  gatewayClassName: istio  # or nginx, cilium, etc.
  listeners:
    - name: https
      hostname: "shortlink.best"
      port: 443
      protocol: HTTPS
      tls:
        certificateRefs:
          - name: shortlink-best-tls
            kind: Secret
      allowedRoutes:
        namespaces:
          from: All  # Allow routes from all namespaces
```

### Internal Gateway Example

For internal gRPC services (link, metadata):

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: internal-gateway
  namespace: shortlink-common
spec:
  gatewayClassName: istio
  listeners:
    - name: grpc
      port: 50051
      hostname: "*.shortlink"
      protocol: HTTPS
      tls:
        certificateRefs:
          - name: self-signed
      allowedRoutes:
        namespaces:
          from: Same
```

## Troubleshooting

### HTTPRoute not working

1. Check Gateway status:
   ```bash
   kubectl get gateway -n shortlink-common
   kubectl describe gateway external-gateway -n shortlink-common
   ```

2. Check HTTPRoute status:
   ```bash
   kubectl get httproute -n your-namespace
   kubectl describe httproute your-service -n your-namespace
   ```

3. Verify parentRefs match your Gateway:
   ```bash
   # The Gateway name and namespace must match
   kubectl get gateway external-gateway -n shortlink-common
   ```

### Service not reachable

1. Check if the Service exists and has endpoints:
   ```bash
   kubectl get svc your-service
   kubectl get endpoints your-service
   ```

2. Verify the backend service name and port match:
   ```bash
   kubectl get svc your-service -o yaml | grep -A 5 ports
   ```

## Benefits of HTTPRoute

- **More expressive routing**: Header, query parameter, and method-based routing
- **Traffic splitting**: Native support for canary deployments and A/B testing
- **Better security**: Integration with policy resources
- **Future-proof**: Gateway API is the standard for Kubernetes ingress
- **Vendor-neutral**: Works with multiple gateway implementations (Istio, Nginx, Cilium, etc.)

## References

- [Gateway API Documentation](https://gateway-api.sigs.k8s.io/)
- [HTTPRoute Reference](https://gateway-api.sigs.k8s.io/api-types/httproute/)
- [Gateway API Migration Guide](https://gateway-api.sigs.k8s.io/guides/migrating-from-ingress/)
