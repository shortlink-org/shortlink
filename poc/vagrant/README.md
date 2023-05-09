### Vagrant && K8S

This POC is about how to use Vagrant for local K8S cluster.

1. We use kubeadm for installation K8S cluster.
2. We use Vagrant for running VMs.
3. We use Ansible for provisioning VMs.

```mermaid
graph LR

  A[Vagrant] --> B[VMs]
  B --> C[Kubeadm]
  C --> D[K8S Cluster]
  A --> E[Ansible]
  E --> B

  subgraph "Local K8S Cluster"
    D
  end

  classDef title fill:#f9d2d2,stroke:#333,stroke-width:2px;
  class A title
```

