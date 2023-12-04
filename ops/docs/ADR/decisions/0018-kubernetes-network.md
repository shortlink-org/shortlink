# 18. Kubernetes network

Date: 2023-03-07

## Status

Proposed

## Context

As we plan the network configuration for our Kubernetes cluster, we need to consider the various network plugins 
available and select the one that best meets our needs.

Right now, we are using the default network plugin, which is Calico. However, we have identified several CNIs that can 
provide better performance and security, including Cilium, Weave, and Flannel.

## Decision

After evaluating several network plugins, we have decided to use Cilium as our network plugin. 
We made this decision based on the following factors:

- **eBPF**: Cilium uses eBPF (extended Berkeley Packet Filter), which provides efficient and flexible packet filtering and 
  allows us to perform complex network operations.
- **L7** policy: Cilium includes a Layer 7 policy engine, which allows us to enforce 
  network policies on the application layer.
- **Observability**: Cilium provides a comprehensive set of observability tools, including 
  network metrics, network traces, and network policy enforcement reports.
- **Security**: Cilium provides a set of security features, including encryption, 
  identity-based access control, and network segmentation.
- **Ease of use**: Cilium is easy to install and configure, and it integrates well with 
  Kubernetes.
- **Community**: Cilium is an open source project with a large and active community.
- **Support**: Cilium is supported by the CNCF and has a large number of contributors.

### Docs

- [Cilium](https://cilium.io/)
- [Cilium for Kubernetes networking: Why we use it and why we love it](https://blog.palark.com/why-cilium-for-kubernetes-networking/)

## Consequences

By selecting Cilium as our network plugin, we expect to achieve the following benefits:

- Improved network performance and efficiency, thanks to eBPF
![ebpf host routing diagram](./proof/ADR-0018/ebpf-host-routing-diagram.png)
- Better security and policy enforcement, thanks to the L7 policy engine
![cilium policy enforcement](https://docs.cilium.io/en/latest/_images/cilium_bpf_endpoint.svg)
- Simplified network management and troubleshooting
- Improved observability and monitoring

### References

- [Migrating from MetaLB to Cilium](https://blog.stonegarden.dev/articles/2023/12/migrating-from-metallb-to-cilium/)
