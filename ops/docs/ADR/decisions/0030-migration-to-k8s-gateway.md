# 30. Migration to K8S Gateway

Date: 2023-10-30

## Status

Accepted

## Context

The migration to Kubernetes Gateway is motivated by the need to evolve Kubernetes service networking through expressive, 
extensible, and role-oriented interfaces as provided by the Gateway API. The Gateway, a core component of the Gateway API, 
represents a logical set of capabilities provided by the system for managing ingress traffic and potentially, 
service mesh configurations under the GAMMA initiative.

Integration with cert-manager requires enabling the `featureGateways` as per the [cert-manager documentation](https://cert-manager.io/docs/usage/gateway/).

The migration also entails the continuation of using VirtualService for managing gRPC traffic for internal communications between services. 
Additionally, some services from an external company will continue using ingress, necessitating the support 
for both implementations to ensure seamless traffic management.

## Decision


![API Model](https://gateway-api.sigs.k8s.io/images/api-model.png)

The decision is to migrate to the K8s Gateway to leverage its role-oriented, portable, expressive, and extensible API resources, 
while also maintaining the existing ingress and VirtualService configurations where necessary.

## Consequences

This migration will facilitate better management of ingress traffic and service networking within Kubernetes. 
It may also pave the way for more advanced service mesh configurations once the GAMMA initiative reaches a stable release. 
However, as GAMMA is still experimental, there's a risk associated with its use in production environments.

Maintaining both the new Gateway and the existing ingress configurations may result in increased complexity and require 
careful coordination to ensure that traffic routing and service interactions function as expected.

### References

1. [Kubernetes Gateway API Documentation](https://gateway-api.sigs.k8s.io/)
2. [GAMMA Initiative](https://gateway-api.sigs.k8s.io/concepts/gamma/)
3. [Cert-manager Gateway Documentation](https://cert-manager.io/docs/usage/gateway/)
