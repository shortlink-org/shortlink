# 3. Maintaining Kubernetes

Date: 2023-12-20

## Status

Accepted

## Context

We need to decide how to deploy our services to Kubernetes and maintain our Kubernetes clusters. 
Previously, we used Kubespray for cluster maintenance but needed a more efficient solution. 
We also required a seamless method for deploying services to Kubernetes.

## Decision

### Kubernetes Cluster Maintenance

![Talos](./proof/ADR-0003/talos.png)

We have transitioned from Kubespray to Talos for maintaining our Kubernetes clusters. 
Talos is a modern OS designed specifically for Kubernetes, providing a minimal, immutable, and secure environment. 
We have also integrated the GitLab Agent in our Kubernetes clusters for improved CI pipeline connectivity.

### Deployment of Services

For the deployment of our services to Kubernetes, we will use the Helm package manager. 
This approach ensures efficient management of Kubernetes applications.

### Hardware Configuration

```puml
@startuml k8s_nodes

' Define KubernetesPuml for URL inclusion
!define KubernetesPuml https://raw.githubusercontent.com/dcasati/kubernetes-PlantUML/master/dist

' Include common Kubernetes components
!includeurl KubernetesPuml/kubernetes_Common.puml
!includeurl KubernetesPuml/kubernetes_Context.puml

' Include specific components for nodes
!includeurl KubernetesPuml/OSS/KubernetesNode.puml

' Include C4 model
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

skinparam backgroundColor #FEFEFE
skinparam handwritten false

left to right direction

' Define C4 Context
LAYOUT_TOP_DOWN()

System_Boundary(k8s_cluster, "K8s Cluster") {
    ' Define Nodes
    KubernetesNode(talos_bxi_ihe, "talos-bxi-ihe\nStatus: Ready\nRoles: <none>\nAge: 78d\nVersion: v1.30.3", "#ffcccc")
    KubernetesNode(talos_coj_emd, "talos-coj-emd\nStatus: Ready\nRoles: <none>\nAge: 78d\nVersion: v1.30.3", "#ccffcc")
    KubernetesNode(talos_dks_th8, "talos-dks-th8\nStatus: Ready\nRoles: <none>\nAge: 78d\nVersion: v1.30.3", "#ccccff")
    KubernetesNode(talos_t8v_b87, "talos-t8v-b87\nStatus: Ready\nRoles: control-plane\nAge: 78d\nVersion: v1.30.3", "#ffffcc")
}

@enduml
```

### CI/CD

We will use the GitLab CI/CD to deploy our services to Kubernetes.

### GitOps

We will use ArgoCD to manage our Kubernetes cluster.

- [Introduction to Continuous Delivery and GitOps using Argo CD](https://academy.akuity.io/courses/gitops-argocd-intro)

![argocd-install](./proof/ADR-0003/argocd-install.png)

### Secrets

We use SOPS to manage our secrets.

## Consequences

- **Enhanced Cluster Security and Simplification**: Talos improves security and simplifies Kubernetes operations.
- **Efficient Updates and Rollbacks**: Talos offers reliable, minimal-downtime update and rollback processes.
- **Improved CI Integration**: GitLab Agent enables automated, secure CI connections.
- **GitOps Efficiency**: ArgoCD enhances cluster management via GitOps.
- **Secure Secret Management**: SOPS ensures secure handling of secrets.

## References

1. **Talos Systems**: For detailed information about Talos, visit their official website: [Talos Systems](https://www.talos.dev/).
2. **GitLab Kubernetes Agent**: To learn more about the GitLab Agent for Kubernetes, refer to the official documentation at [GitLab Kubernetes Agent](https://docs.gitlab.com/ee/user/clusters/agent/).
